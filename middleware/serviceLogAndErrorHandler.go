package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"app/GoSample/controllers/models/response"
	"app/GoSample/logger"
	"app/GoSample/db"
	"app/GoSample/db/repo"
	"app/GoSample/infra/constant"
	"app/GoSample/infra/resource"

	"github.com/gin-gonic/gin"
)

type ServiceLogAndErrorMiddleware struct{
	MongoOperator db.IMongo
	LocalizationRepo repo.ILocalizationRepo
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (middleware *ServiceLogAndErrorMiddleware) ServiceLogAndErrorHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		logRecord := map[string]interface{}{}
		now := time.Now()
		current := now.Format("2006-01-02 15:04:05")
		requestBodyBytes, _ := ioutil.ReadAll(context.Request.Body)
		var clientIP string
		if context.ClientIP() == "::1" {
			clientIP = "localhost"
		} else {
			clientIP = context.ClientIP()
		}

		logRecord["RequestURI"] = context.Request.RequestURI
		logRecord["ClientIP"] = clientIP
		logRecord["RequestedAt"] = current

		var language string
		var requestBody map[string]interface{}
		if err := json.Unmarshal(requestBodyBytes, &requestBody); err != nil {
			logger.ErrorLog("An error occured while request body unmarshal:", requestBodyBytes, " - And selected 'EN' as default language")
			language = "EN"
		} else if requestBody["Language"] != nil {
			language = fmt.Sprintf("%v", requestBody["Language"])
		} else {
			language = "EN"
		}

		logRecord["RequestBody"] = requestBody

		context.Request.Body = ioutil.NopCloser(bytes.NewReader(requestBodyBytes))

		bodyLogWriter := &bodyLogWriter{body: bytes.NewBufferString(constant.EmptyString), ResponseWriter: context.Writer}
		context.Writer = bodyLogWriter

		// check whether error comes from pre middleware
		if len(context.Errors) == 0 {
			context.Next() 
		}

		if len(context.Errors) > 0 {
			errorMessageKey := context.Errors[0].Error()
			resourcer := resource.Resource{LocalizationRepo: middleware.LocalizationRepo}
			errorMessage := resourcer.GetResource(errorMessageKey, language)
			responseBody := &response.BaseResponse{IsSuccess: false, ErrorMessage: errorMessage} 
			logRecord["ResponseBody"] = responseBody
			middleware.MongoOperator.InsertRecord("RequestReponseLogs", logRecord)
			context.JSON(200, responseBody)
		} else {
			var responseBody map[string]interface{}
			if err := json.Unmarshal(bodyLogWriter.body.Bytes(), &responseBody); err != nil {
				logger.ErrorLog("An error occured while response body unmarshal", bodyLogWriter.body.Bytes())
				logRecord["ResponseBody"] = bodyLogWriter.body.String()
			} else {
				logRecord["ResponseBody"] = responseBody
			}
			middleware.MongoOperator.InsertRecord("RequestReponseLogs", logRecord)
		}
	}
}

