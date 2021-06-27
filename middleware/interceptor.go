package middleware

import (
	"bytes"
	"encoding/json"
	"time"
	"io/ioutil"
	"fmt"
	"context"

	"app/GoSample/controllers/models/response"
	"app/GoSample/db"
	"app/GoSample/infra/constant"
	"app/GoSample/infra/resource"
	"app/GoSample/logger"

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

		now := time.Now()
		current := now.Format("2006-01-02 15:04:05")
		requestBodyBytes, _ := ioutil.ReadAll(c.Request.Body)
		var clientIP string
		if c.ClientIP() == "::1" {
			clientIP = "localhost"
		} else {
			clientIP = c.ClientIP()
		}

		logRecord["RequestURI"] = c.Request.RequestURI
		logRecord["ClientIP"] = clientIP
		logRecord["RequestedAt"] = current

		var language string
		var requestBody map[string]interface{}
		if err := json.Unmarshal(requestBodyBytes, &requestBody); err != nil {
			logger.ErrorLog("An error occured while request body unmarshal:", requestBodyBytes, " - And selected 'EN' as default language")
			language = "EN"
		} else {
			language = fmt.Sprintf("%v", requestBody["Language"])
		}

		logRecord["RequestBody"] = requestBody

		c.Request.Body = ioutil.NopCloser(bytes.NewReader(requestBodyBytes))

		bodyLogWriter := &bodyLogWriter{body: bytes.NewBufferString(constant.EmptyString), ResponseWriter: c.Writer}
		c.Writer = bodyLogWriter

		c.Next() // < the rest of handlers in the chain are executed here!

		if len(c.Errors) > 0 {
			errorMessageKey := c.Errors[0].Error()
			errorMessage := resource.GetResource(errorMessageKey, language)
			responseBody := &response.BaseResponse{IsSuccess: false, ErrorMessage: errorMessage} 
			logRecord["ResponseBody"] = responseBody
			insertLogRecord(logRecord)
			c.JSON(200, responseBody)
		} else {
			var responseBody map[string]interface{}
			if err := json.Unmarshal(bodyLogWriter.body.Bytes(), &responseBody); err != nil {
				logger.ErrorLog("An error occured while response body unmarshal", bodyLogWriter.body.Bytes())
				logRecord["ResponseBody"] = bodyLogWriter.body.String()
			} else {
				logRecord["ResponseBody"] = responseBody
			}
			insertLogRecord(logRecord)
		}
	}
}

func insertLogRecord(record map[string]interface{}) {
	collection := db.MongoDB.Collection("RequestReponseLogs")
	_, err := collection.InsertOne(context.Background(), record)
	if err != nil {
		logger.ErrorLog("An error occured while inserting log record to mongo db - Error:", err.Error(), "- LogRecord:", record)
	}
}

