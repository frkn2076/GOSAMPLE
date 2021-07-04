package helper

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

type IHelper interface {
	BindRequest(context *gin.Context, model interface{}) bool
	AddToSession(context *gin.Context, key string, value string) bool
	GetSession(context *gin.Context) (ses *sessions.Session, err error)
	HashPassword(context *gin.Context, password string) (string, bool)
	CheckPasswordHash(password, hash string) bool
	GenerateToken(userId uint) (string, bool)
	StringToUint(input string) (uint, bool)
}