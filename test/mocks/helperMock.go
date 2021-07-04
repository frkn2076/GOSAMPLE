package mocks

import (
	"os"

	"app/GoSample/config/session"
	"app/GoSample/infra/constant"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

type MockHelper struct{}

func (e MockHelper) BindRequest(context *gin.Context, model interface{}) bool {
	return true
}

func (e MockHelper) AddToSession(context *gin.Context, key string, value string) bool {
	return true
}

func (e MockHelper) GetSession(context *gin.Context) (ses *sessions.Session, err error) {
	sessionCookieName := os.Getenv("SessionCookieName")
	ses, err = session.Store.Get(context.Request, sessionCookieName)
	if err != nil {
		panic(err)
	}
	return
}

func (e MockHelper) HashPassword(context *gin.Context, password string) (string, bool) {
	 return "HashedPassword", true
}

func (e MockHelper) CheckPasswordHash(password, hash string) bool {
	return true
}

func (e MockHelper) GenerateToken(userId uint) (string, bool) {
	return "DummyToken", true
}

func (e MockHelper) StringToUint(input string) (uint, bool){
	return 1, true
}


type MockRequestFailHelper struct{}

func (e MockRequestFailHelper) BindRequest(context *gin.Context, model interface{}) bool {
	return false
}

func (e MockRequestFailHelper) AddToSession(context *gin.Context, key string, value string) bool {
	return true
}

func (e MockRequestFailHelper) GetSession(context *gin.Context) (ses *sessions.Session, err error) {
	sessionCookieName := os.Getenv("SessionCookieName")
	ses, err = session.Store.Get(context.Request, sessionCookieName)
	if err != nil {
		panic(err)
	}
	return
}

func (e MockRequestFailHelper) HashPassword(context *gin.Context, password string) (string, bool) {
	 return "HashedPassword", true
}

func (e MockRequestFailHelper) CheckPasswordHash(password, hash string) bool {
	return true
}

func (e MockRequestFailHelper) GenerateToken(userId uint) (string, bool) {
	return "DummyToken", true
}

func (e MockRequestFailHelper) StringToUint(input string) (uint, bool){
	return 1, true
}


type MockHashPasswordFailHelper struct{}

func (e MockHashPasswordFailHelper) BindRequest(context *gin.Context, model interface{}) bool {
	return true
}

func (e MockHashPasswordFailHelper) AddToSession(context *gin.Context, key string, value string) bool {
	return true
}

func (e MockHashPasswordFailHelper) GetSession(context *gin.Context) (ses *sessions.Session, err error) {
	sessionCookieName := os.Getenv("SessionCookieName")
	ses, err = session.Store.Get(context.Request, sessionCookieName)
	if err != nil {
		panic(err)
	}
	return
}

func (e MockHashPasswordFailHelper) HashPassword(context *gin.Context, password string) (string, bool) {
	 return constant.EmptyString, false
}

func (e MockHashPasswordFailHelper) CheckPasswordHash(password, hash string) bool {
	return true
}

func (e MockHashPasswordFailHelper) GenerateToken(userId uint) (string, bool) {
	return "DummyToken", true
}
func (e MockHashPasswordFailHelper) StringToUint(input string) (uint, bool){
	return 1, true
}


type MockAddToSessionFailHelper struct{}

func (e MockAddToSessionFailHelper) BindRequest(context *gin.Context, model interface{}) bool {
	return true
}

func (e MockAddToSessionFailHelper) AddToSession(context *gin.Context, key string, value string) bool {
	return false
}

func (e MockAddToSessionFailHelper) GetSession(context *gin.Context) (ses *sessions.Session, err error) {
	sessionCookieName := os.Getenv("SessionCookieName")
	ses, err = session.Store.Get(context.Request, sessionCookieName)
	if err != nil {
		panic(err)
	}
	return
}

func (e MockAddToSessionFailHelper) HashPassword(context *gin.Context, password string) (string, bool) {
	 return "HashedPassword", true
}

func (e MockAddToSessionFailHelper) CheckPasswordHash(password, hash string) bool {
	return true
}

func (e MockAddToSessionFailHelper) GenerateToken(userId uint) (string, bool) {
	return "DummyToken", true
}

func (e MockAddToSessionFailHelper) StringToUint(input string) (uint, bool){
	return 1, true
}


type MockGenerateTokenFailHelper struct{}

func (e MockGenerateTokenFailHelper) BindRequest(context *gin.Context, model interface{}) bool {
	return true
}

func (e MockGenerateTokenFailHelper) AddToSession(context *gin.Context, key string, value string) bool {
	return true
}

func (e MockGenerateTokenFailHelper) GetSession(context *gin.Context) (ses *sessions.Session, err error) {
	sessionCookieName := os.Getenv("SessionCookieName")
	ses, err = session.Store.Get(context.Request, sessionCookieName)
	if err != nil {
		panic(err)
	}
	return
}

func (e MockGenerateTokenFailHelper) HashPassword(context *gin.Context, password string) (string, bool) {
	 return "HashedPassword", true
}

func (e MockGenerateTokenFailHelper) CheckPasswordHash(password, hash string) bool {
	return true
}

func (e MockGenerateTokenFailHelper) GenerateToken(userId uint) (string, bool) {
	return constant.EmptyString, false
}

func (e MockGenerateTokenFailHelper) StringToUint(input string) (uint, bool){
	return 1, true
}


type MockCheckPasswordFailHelper struct{}

func (e MockCheckPasswordFailHelper) BindRequest(context *gin.Context, model interface{}) bool {
	return true
}

func (e MockCheckPasswordFailHelper) AddToSession(context *gin.Context, key string, value string) bool {
	return true
}

func (e MockCheckPasswordFailHelper) GetSession(context *gin.Context) (ses *sessions.Session, err error) {
	sessionCookieName := os.Getenv("SessionCookieName")
	ses, err = session.Store.Get(context.Request, sessionCookieName)
	if err != nil {
		panic(err)
	}
	return
}

func (e MockCheckPasswordFailHelper) HashPassword(context *gin.Context, password string) (string, bool) {
	 return "HashedPassword", true
}

func (e MockCheckPasswordFailHelper) CheckPasswordHash(password, hash string) bool {
	return false
}

func (e MockCheckPasswordFailHelper) GenerateToken(userId uint) (string, bool) {
	return "DummyToken", true
}

func (e MockCheckPasswordFailHelper) StringToUint(input string) (uint, bool){
	return 1, true
}

type MockStringToUintFailHelper struct{}

func (e MockStringToUintFailHelper) BindRequest(context *gin.Context, model interface{}) bool {
	return true
}

func (e MockStringToUintFailHelper) AddToSession(context *gin.Context, key string, value string) bool {
	return true
}

func (e MockStringToUintFailHelper) GetSession(context *gin.Context) (ses *sessions.Session, err error) {
	sessionCookieName := os.Getenv("SessionCookieName")
	ses, err = session.Store.Get(context.Request, sessionCookieName)
	if err != nil {
		panic(err)
	}
	return
}

func (e MockStringToUintFailHelper) HashPassword(context *gin.Context, password string) (string, bool) {
	 return "HashedPassword", true
}

func (e MockStringToUintFailHelper) CheckPasswordHash(password, hash string) bool {
	return true
}

func (e MockStringToUintFailHelper) GenerateToken(userId uint) (string, bool) {
	return "DummyToken", true
}

func (e MockStringToUintFailHelper) StringToUint(input string) (uint, bool){
	return 0, false
}