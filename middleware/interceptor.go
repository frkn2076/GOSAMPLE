package middleware

import (
	"bytes"
	"io/ioutil"

	// "net/http"
	"encoding/json"
	"fmt"
	"unsafe"

	// "app/GoSample/controllers/models/request"
	"app/GoSample/controllers/models/response"
	"app/GoSample/db/nosql"
	"app/GoSample/infra/constant"
	"app/GoSample/infra/customeError"
	"app/GoSample/infra/resource"

	// "app/GoSample/logger"

	"github.com/gin-gonic/gin"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func ServiceLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		logRecord := map[string]interface{}{}

		body, _ := ioutil.ReadAll(c.Request.Body)
		var clientIP string
		if c.ClientIP() == "::1" {
			clientIP = "localhost"
		} else {
			clientIP = c.ClientIP()
		}

		logRecord["RequestURI"] = c.Request.RequestURI
		logRecord["ClientIP"] = clientIP
		logRecord["RequestBody"] = string(body)

		c.Request.Body = ioutil.NopCloser(bytes.NewReader(body))

		bodyLogWriter := &bodyLogWriter{body: bytes.NewBufferString(constant.EmptyString), ResponseWriter: c.Writer}
		c.Writer = bodyLogWriter

		c.Next() // < the rest of handlers in the chain are executed here!

		errorHandler(c, logRecord)

		responseBody := bodyLogWriter.body.String()
		logRecord["ResponseBody"] = string(responseBody)

		nosql.InsertLogRecord(logRecord)
	}
}

func errorHandler(c *gin.Context, logRecord map[string]interface{}) {
	requestBody := logRecord["RequestBody"]
	requestBodyBytes := []byte(fmt.Sprintf("%v", requestBody))
	var language string

	var request map[string]string
	if err := json.Unmarshal(requestBodyBytes, &request); err != nil {
		language = request["Language"]
	} else {
		language = "EN"
	}

	if len(c.Errors) > 0 {
		customeError := *(*customeError.CustomeError)(unsafe.Pointer(&c.Errors[0]))
		errorCode := customeError.ErrorCode
		errorMessage := resource.GetResource(customeError.ErrorMessage, language)
		c.JSON(500, &response.BaseResponse{IsSuccess: false, ErrorCode: string(errorCode), ErrorMessage: errorMessage})
	}
	

	// var request request.BaseRequest
	// if err := c.Bind(&request); err != nil {
	// 	logger.ErrorLog("Invalid request - Register - interceptor.go - Error:", err.Error())
	// 	defaultErrorMessage := resource.GetResource(c.Errors[0].Error(), "EN")
	// 	response := &response.BaseResponse{IsSuccess: false, Message: defaultErrorMessage}
	// 	logRecord["ResponseBody"] = response
	// 	c.Writer = &bodyLogWriter{body: bytes.NewBufferString("selam"), ResponseWriter: c.Writer}
	// 	// c.AbortWithStatusJSON(500, response)
	// }

	// if len(c.Errors) > 0 {
	// 	errorMessage := resource.GetResource(c.Errors[0].Error(), language)
	// 	response := &response.BaseResponse{IsSuccess: false, Message: errorMessage}
	// 	logRecord["ResponseBody"] = response
	// 	c.AbortWithStatusJSON(500, response)
	// }	else if c.Writer.Status() < http.StatusOK || c.Writer.Status() > http.StatusIMUsed {
	// 	commonErrorMessage := resource.GetResource("SomethingWentWrong-Err", language)
	// 	logger.ServiceLog("Response: ", c.Request.RequestURI, "- Error Message:", commonErrorMessage)
	// 	c.AbortWithStatusJSON(c.Writer.Status(), &response.BaseResponse{IsSuccess: false, Message: commonErrorMessage})
	// 	return
	// }
}
