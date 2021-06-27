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
func BindRequest(context *gin.Context, model interface{}) {
	if err := context.ShouldBindJSON(&model); err != nil {
		logger.ErrorLog("Invalid request - helper.go - Model:", model, "- Error:", err.Error())
		context.Error(customeError.SomethingWentWrong)
		return
	}
}

//#endregion

//#region Session
func AddToSession(context *gin.Context, key string, value string) {
	session := GetSession(context)
	session.Values[key] = value
	if err := session.Save(context.Request, context.Writer); err != nil {
		logger.ErrorLog("An error occured while saving session - helper.go - Error:", err.Error())
		context.Error(customeError.SomethingWentWrong)
		return
	}
}

func GetSession(context *gin.Context) *sessions.Session {
	sessionCookieName := os.Getenv("SessionCookieName")
	session, err := session.Store.Get(context.Request, sessionCookieName)
	if err != nil {
		logger.ErrorLog("An error occured while getting session - helper.go - Error:", err.Error())
		context.Error(customeError.SomethingWentWrong)
		return nil
	}
	return session
}

//#endregion

//#region Hashing
func HashPassword(context *gin.Context, password string) string {
	passwordBytes := []byte(password)
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		logger.ErrorLog("An error occured while generating crypted password - Register - Error:", err.Error())
		context.Error(customeError.SomethingWentWrong)
		return constant.EmptyString
	}
	hashedPassword := string(hashedPasswordBytes)
	return hashedPassword
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

//#endregion
