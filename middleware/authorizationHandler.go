package middleware

import (
	// "bytes"
	// "encoding/json"
	// "io/ioutil"
	// "fmt"
	// "context"
	"strings"

	// "app/GoSample/controllers/models/response"
	// "app/GoSample/db"
	// "app/GoSample/infra/resource"
	"app/GoSample/infra/auth"
	"app/GoSample/infra/customeError"
	"app/GoSample/logger"
	"app/GoSample/infra/constant"

	"github.com/gin-gonic/gin"
)



func AuthorizationHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("Authorization")
		if clientToken == constant.EmptyString {
			logger.ErrorLog("Authorization header is missing")
			c.Error(customeError.SomethingWentWrong)
			return
		} else {
			clientToken = strings.TrimPrefix(clientToken, "Bearer ")
			if clientToken == constant.EmptyString {
				logger.ErrorLog("Authorization header is missing")
				c.Error(customeError.SomethingWentWrong)
				return
			}
		}

		claims, err := auth.JWT.ValidateToken(clientToken)
		if err != nil {
			c.Error(customeError.SomethingWentWrong)
			return
		}

		c.Set("userId", claims.UserId)
		// value := c.GetString("userId")

		c.Next()
	}
}

