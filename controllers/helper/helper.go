package helper

import (
	"os"

	"app/GoSample/config/session"
	"app/GoSample/infra/auth"
	"app/GoSample/infra/constant"
	"app/GoSample/logger"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

var HelperInstance *Helper

func init() {
	HelperInstance = new(Helper)
}

type Helper struct{}

func (h *Helper) BindRequest(context *gin.Context, model interface{}) bool {
	if err := context.ShouldBindJSON(&model); err != nil {
		logger.ErrorLog("Invalid request - helper.go - Model:", model, "- Error:", err.Error())
		return false
	}
	return true
}

func (h *Helper) AddToSession(context *gin.Context, key string, value string) bool {
	session, err := h.GetSession(context)
	if err != nil {
		return false
	}
	session.Values[key] = value
	if err := session.Save(context.Request, context.Writer); err != nil {
		logger.ErrorLog("An error occured while saving session - helper.go - Error:", err.Error())
		return false
	}
	return true
}

func (h *Helper) GetSession(context *gin.Context) (ses *sessions.Session, err error) {
	sessionCookieName := os.Getenv("SessionCookieName")
	ses, err = session.Store.Get(context.Request, sessionCookieName)
	if err != nil {
		logger.ErrorLog("An error occured while getting session - helper.go - Error:", err.Error())
	}
	return
}

func (h *Helper) HashPassword(context *gin.Context, password string) (string, bool) {
	passwordBytes := []byte(password)
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		logger.ErrorLog("An error occured while generating crypted password - Register - Error:", err.Error())
		return constant.EmptyString, false
	}
	hashedPassword := string(hashedPasswordBytes)
	return hashedPassword, true
}

func (h *Helper) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (h *Helper) GenerateToken(userId uint) (string, bool) {
	token, err := auth.JWT.GenerateToken(userId)
	if err != nil {
		return constant.EmptyString, false
	}
	return token, true
}
