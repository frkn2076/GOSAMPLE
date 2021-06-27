package helper

import (
	"os"

	"app/GoSample/config/session"
	"app/GoSample/infra/constant"
	"app/GoSample/infra/customeError"
	"app/GoSample/logger"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

//#region Request/Response
func BindRequest(c *gin.Context, model interface{}) {
	if err := c.ShouldBindJSON(&model); err != nil {
		logger.ErrorLog("Invalid request - helper.go - Model:", model, "- Error:", err.Error())
		c.Error(customeError.SomethingWentWrong)
		return
	}
}

//#endregion

//#region Session
func AddToSession(c *gin.Context, key string, value string) {
	session := GetSession(c)
	session.Values[key] = value
	if err := session.Save(c.Request, c.Writer); err != nil {
		logger.ErrorLog("An error occured while saving session - helper.go - Error:", err.Error())
		c.Error(customeError.SomethingWentWrong)
		return
	}
}

func GetSession(c *gin.Context) *sessions.Session {
	sessionCookieName := os.Getenv("SessionCookieName")
	session, err := session.Store.Get(c.Request, sessionCookieName)
	if err != nil {
		logger.ErrorLog("An error occured while getting session - helper.go - Error:", err.Error())
		c.Error(customeError.SomethingWentWrong)
		return nil
	}
	return session
}

//#endregion

//#region Hashing
func HashPassword(c *gin.Context, password string) string {
	passwordBytes := []byte(password)
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		logger.ErrorLog("An error occured while generating crypted password - Register - Error:", err.Error())
		c.Error(customeError.SomethingWentWrong)
		return constant.EmptyString
	}
	hashedPassword := string(hashedPasswordBytes)
	return hashedPassword
}

//#endregion
