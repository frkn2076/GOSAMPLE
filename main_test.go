package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"app/GoSample/db/repo"
	"app/GoSample/db"
	"app/GoSample/config/environments"
	"app/GoSample/config/session"
	"app/GoSample/controllers"
	req "app/GoSample/controllers/models/request"
	res "app/GoSample/controllers/models/response"
	"app/GoSample/controllers/helper"
	"app/GoSample/logger"
	"app/GoSample/test/mocks"
	"app/GoSample/middleware"

	// "gorm.io/gorm"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSetup(t *testing.T) {
	environments.Load()
	session.Start()
	logger.Start()
	gin.SetMode(gin.TestMode)
}

func TestRegister(t *testing.T) {
	registerRequestFailCase(t)
	registerHashPasswordFailCase(t)
	registerUserExistCase(t)
	registerAccountCreateFailCase(t)
	registerAddToSessionFailCase(t)
	registerGenerateTokenFailCase(t)
	registerSuccessfullyCase(t)
}

func TestLogin(t *testing.T) {
	loginRequestFailCase(t)
	loginUserNameNotFoundCase(t)
	loginCheckPasswordFailCase(t)
	loginAddToSessionFailCase(t)
	loginGenerateTokenFailCase(t)
	loginSuccessfullyCase(t)
}

func registerRequestFailCase(t *testing.T){
	// Arrange
	var mockAccountRepo repo.IAccountrepo = mocks.MockNotRegisteredUserAccountRepo{}
	var mockRepo repo.IRepo = mocks.MockRepo{}
	var mockHelper helper.IHelper = mocks.MockRequestFailHelper{}
	accountController := controllers.AccountController{AccountRepo: mockAccountRepo, Repo: mockRepo, Helper: mockHelper}

	responseWriter := httptest.NewRecorder()
	context, router := gin.CreateTestContext(responseWriter)

	var mockLocalizationRepo repo.ILocalizationRepo = mocks.MockLocalizationRepo{}
	var mockMongoOperator db.IMongo = mocks.MockMongoOperator{}
	serviceLogAndErrorMiddleware := middleware.ServiceLogAndErrorMiddleware{LocalizationRepo: mockLocalizationRepo, MongoOperator: mockMongoOperator}
	router.Use(serviceLogAndErrorMiddleware.ServiceLogAndErrorHandler())
	router.POST("register", accountController.Register)

	request := req.AccountRequest{BaseRequest: req.BaseRequest{Version: "0.0.1", Language: "TR"}, UserName: "furkan", Password: "12345"}
	requestBytes, _ := json.Marshal(request)

	context.Request, _ = http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(requestBytes))

	// Act
	router.ServeHTTP(responseWriter, context.Request)

	// Assert
	var responseBody res.AccountResponse
	if err := json.Unmarshal(responseWriter.Body.Bytes(), &responseBody); err != nil {
		panic(err)
	}

	assert.Equal(t, 200, responseWriter.Code)
	assert.False(t, responseBody.BaseResponse.IsSuccess)
	assert.Equal(t, responseBody.BaseResponse.ErrorMessage, "DummyMessage")
}

func registerHashPasswordFailCase(t *testing.T){
	// Arrange
	var mockAccountRepo repo.IAccountrepo = mocks.MockNotRegisteredUserAccountRepo{}
	var mockRepo repo.IRepo = mocks.MockRepo{}
	var mockHelper helper.IHelper = mocks.MockHashPasswordFailHelper{}
	accountController := controllers.AccountController{AccountRepo: mockAccountRepo, Repo: mockRepo, Helper: mockHelper}

	responseWriter := httptest.NewRecorder()
	context, router := gin.CreateTestContext(responseWriter)

	var mockLocalizationRepo repo.ILocalizationRepo = mocks.MockLocalizationRepo{}
	var mockMongoOperator db.IMongo = mocks.MockMongoOperator{}
	serviceLogAndErrorMiddleware := middleware.ServiceLogAndErrorMiddleware{LocalizationRepo: mockLocalizationRepo, MongoOperator: mockMongoOperator}
	router.Use(serviceLogAndErrorMiddleware.ServiceLogAndErrorHandler())
	router.POST("register", accountController.Register)

	request := req.AccountRequest{BaseRequest: req.BaseRequest{Version: "0.0.1", Language: "TR"}, UserName: "furkan", Password: "12345"}
	requestBytes, _ := json.Marshal(request)

	context.Request, _ = http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(requestBytes))

	// Act
	router.ServeHTTP(responseWriter, context.Request)

	// Assert
	var responseBody res.AccountResponse
	if err := json.Unmarshal(responseWriter.Body.Bytes(), &responseBody); err != nil {
		panic(err)
	}

	assert.Equal(t, 200, responseWriter.Code)
	assert.False(t, responseBody.BaseResponse.IsSuccess)
	assert.Equal(t, responseBody.BaseResponse.ErrorMessage, "DummyMessage")
}

func registerUserExistCase(t *testing.T){
	// Arrange
	var mockAccountRepo repo.IAccountrepo = mocks.MockRegisteredUserAccountRepo{}
	var mockRepo repo.IRepo = mocks.MockRepo{}
	var mockHelper helper.IHelper = mocks.MockHelper{}
	accountController := controllers.AccountController{AccountRepo: mockAccountRepo, Repo: mockRepo, Helper: mockHelper}

	responseWriter := httptest.NewRecorder()
	context, router := gin.CreateTestContext(responseWriter)

	var mockLocalizationRepo repo.ILocalizationRepo = mocks.MockLocalizationRepo{}
	var mockMongoOperator db.IMongo = mocks.MockMongoOperator{}
	serviceLogAndErrorMiddleware := middleware.ServiceLogAndErrorMiddleware{LocalizationRepo: mockLocalizationRepo, MongoOperator: mockMongoOperator}
	router.Use(serviceLogAndErrorMiddleware.ServiceLogAndErrorHandler())
	router.POST("register", accountController.Register)

	request := req.AccountRequest{BaseRequest: req.BaseRequest{Version: "0.0.1", Language: "TR"}, UserName: "furkan", Password: "12345"}
	requestBytes, _ := json.Marshal(request)

	context.Request, _ = http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(requestBytes))

	// Act
	router.ServeHTTP(responseWriter, context.Request)

	// Assert
	var responseBody res.AccountResponse
	if err := json.Unmarshal(responseWriter.Body.Bytes(), &responseBody); err != nil {
		panic(err)
	}

	assert.Equal(t, 200, responseWriter.Code)
	assert.False(t, responseBody.BaseResponse.IsSuccess)
	assert.Equal(t, responseBody.BaseResponse.ErrorMessage, "DummyMessage")
}

func registerAccountCreateFailCase(t *testing.T){
	// Arrange
	var mockAccountRepo repo.IAccountrepo = mocks.MockAccountCreateFailAccountRepo{}
	var mockRepo repo.IRepo = mocks.MockRepo{}
	var mockHelper helper.IHelper = mocks.MockHelper{}
	accountController := controllers.AccountController{AccountRepo: mockAccountRepo, Repo: mockRepo, Helper: mockHelper}

	responseWriter := httptest.NewRecorder()
	context, router := gin.CreateTestContext(responseWriter)

	var mockLocalizationRepo repo.ILocalizationRepo = mocks.MockLocalizationRepo{}
	var mockMongoOperator db.IMongo = mocks.MockMongoOperator{}
	serviceLogAndErrorMiddleware := middleware.ServiceLogAndErrorMiddleware{LocalizationRepo: mockLocalizationRepo, MongoOperator: mockMongoOperator}
	router.Use(serviceLogAndErrorMiddleware.ServiceLogAndErrorHandler())
	router.POST("register", accountController.Register)

	request := req.AccountRequest{BaseRequest: req.BaseRequest{Version: "0.0.1", Language: "TR"}, UserName: "furkan", Password: "12345"}
	requestBytes, _ := json.Marshal(request)

	context.Request, _ = http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(requestBytes))

	// Act
	router.ServeHTTP(responseWriter, context.Request)

	// Assert
	var responseBody res.AccountResponse
	if err := json.Unmarshal(responseWriter.Body.Bytes(), &responseBody); err != nil {
		panic(err)
	}

	assert.Equal(t, 200, responseWriter.Code)
	assert.False(t, responseBody.BaseResponse.IsSuccess)
	assert.Equal(t, responseBody.BaseResponse.ErrorMessage, "DummyMessage")
}

func registerAddToSessionFailCase(t *testing.T){
	// Arrange
	var mockAccountRepo repo.IAccountrepo = mocks.MockNotRegisteredUserAccountRepo{}
	var mockRepo repo.IRepo = mocks.MockRepo{}
	var mockHelper helper.IHelper = mocks.MockAddToSessionFailHelper{}
	accountController := controllers.AccountController{AccountRepo: mockAccountRepo, Repo: mockRepo, Helper: mockHelper}

	responseWriter := httptest.NewRecorder()
	context, router := gin.CreateTestContext(responseWriter)

	var mockLocalizationRepo repo.ILocalizationRepo = mocks.MockLocalizationRepo{}
	var mockMongoOperator db.IMongo = mocks.MockMongoOperator{}
	serviceLogAndErrorMiddleware := middleware.ServiceLogAndErrorMiddleware{LocalizationRepo: mockLocalizationRepo, MongoOperator: mockMongoOperator}
	router.Use(serviceLogAndErrorMiddleware.ServiceLogAndErrorHandler())
	router.POST("register", accountController.Register)

	request := req.AccountRequest{BaseRequest: req.BaseRequest{Version: "0.0.1", Language: "TR"}, UserName: "furkan", Password: "12345"}
	requestBytes, _ := json.Marshal(request)

	context.Request, _ = http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(requestBytes))

	// Act
	router.ServeHTTP(responseWriter, context.Request)

	// Assert
	var responseBody res.AccountResponse
	if err := json.Unmarshal(responseWriter.Body.Bytes(), &responseBody); err != nil {
		panic(err)
	}

	assert.Equal(t, 200, responseWriter.Code)
	assert.False(t, responseBody.BaseResponse.IsSuccess)
	assert.Equal(t, responseBody.BaseResponse.ErrorMessage, "DummyMessage")
}

func registerGenerateTokenFailCase(t *testing.T){
	// Arrange
	var mockAccountRepo repo.IAccountrepo = mocks.MockNotRegisteredUserAccountRepo{}
	var mockRepo repo.IRepo = mocks.MockRepo{}
	var mockHelper helper.IHelper = mocks.MockGenerateTokenFailHelper{}
	accountController := controllers.AccountController{AccountRepo: mockAccountRepo, Repo: mockRepo, Helper: mockHelper}

	responseWriter := httptest.NewRecorder()
	context, router := gin.CreateTestContext(responseWriter)

	var mockLocalizationRepo repo.ILocalizationRepo = mocks.MockLocalizationRepo{}
	var mockMongoOperator db.IMongo = mocks.MockMongoOperator{}
	serviceLogAndErrorMiddleware := middleware.ServiceLogAndErrorMiddleware{LocalizationRepo: mockLocalizationRepo, MongoOperator: mockMongoOperator}
	router.Use(serviceLogAndErrorMiddleware.ServiceLogAndErrorHandler())
	router.POST("register", accountController.Register)

	request := req.AccountRequest{BaseRequest: req.BaseRequest{Version: "0.0.1", Language: "TR"}, UserName: "furkan", Password: "12345"}
	requestBytes, _ := json.Marshal(request)

	context.Request, _ = http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(requestBytes))

	// Act
	router.ServeHTTP(responseWriter, context.Request)

	// Assert
	var responseBody res.AccountResponse
	if err := json.Unmarshal(responseWriter.Body.Bytes(), &responseBody); err != nil {
		panic(err)
	}

	assert.Equal(t, 200, responseWriter.Code)
	assert.False(t, responseBody.BaseResponse.IsSuccess)
	assert.Equal(t, responseBody.BaseResponse.ErrorMessage, "DummyMessage")
}

func registerSuccessfullyCase(t *testing.T){
	// Arrange
	var mockAccountRepo repo.IAccountrepo = mocks.MockNotRegisteredUserAccountRepo{}
	var mockRepo repo.IRepo = mocks.MockRepo{}
	var mockHelper helper.IHelper = mocks.MockHelper{}
	accountController := controllers.AccountController{AccountRepo: mockAccountRepo, Repo: mockRepo, Helper: mockHelper}

	responseWriter := httptest.NewRecorder()
	context, router := gin.CreateTestContext(responseWriter)

	var mockLocalizationRepo repo.ILocalizationRepo = mocks.MockLocalizationRepo{}
	var mockMongoOperator db.IMongo = mocks.MockMongoOperator{}
	serviceLogAndErrorMiddleware := middleware.ServiceLogAndErrorMiddleware{LocalizationRepo: mockLocalizationRepo, MongoOperator: mockMongoOperator}
	router.Use(serviceLogAndErrorMiddleware.ServiceLogAndErrorHandler())
	router.POST("register", accountController.Register)

	request := req.AccountRequest{BaseRequest: req.BaseRequest{Version: "0.0.1", Language: "TR"}, UserName: "furkan", Password: "12345"}
	requestBytes, _ := json.Marshal(request)

	context.Request, _ = http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(requestBytes))

	// Act
	router.ServeHTTP(responseWriter, context.Request)

	// Assert
	var responseBody res.AccountResponse
	if err := json.Unmarshal(responseWriter.Body.Bytes(), &responseBody); err != nil {
		panic(err)
	}

	assert.Equal(t, 200, responseWriter.Code)
	assert.True(t, responseBody.BaseResponse.IsSuccess)
	assert.Equal(t, responseBody.Token, "DummyToken")
}

func loginRequestFailCase(t *testing.T){
	// Arrange
	var mockAccountRepo repo.IAccountrepo = mocks.MockRegisteredUserAccountRepo{}
	var mockRepo repo.IRepo = mocks.MockRepo{}
	var mockHelper helper.IHelper = mocks.MockRequestFailHelper{}
	accountController := controllers.AccountController{AccountRepo: mockAccountRepo, Repo: mockRepo, Helper: mockHelper}

	responseWriter := httptest.NewRecorder()
	context, router := gin.CreateTestContext(responseWriter)

	var mockLocalizationRepo repo.ILocalizationRepo = mocks.MockLocalizationRepo{}
	var mockMongoOperator db.IMongo = mocks.MockMongoOperator{}
	serviceLogAndErrorMiddleware := middleware.ServiceLogAndErrorMiddleware{LocalizationRepo: mockLocalizationRepo, MongoOperator: mockMongoOperator}
	router.Use(serviceLogAndErrorMiddleware.ServiceLogAndErrorHandler())
	router.POST("login", accountController.Login)

	request := req.AccountRequest{BaseRequest: req.BaseRequest{Version: "0.0.1", Language: "TR"}, UserName: "furkan", Password: "12345"}
	requestBytes, _ := json.Marshal(request)

	context.Request, _ = http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(requestBytes))

	// Act
	router.ServeHTTP(responseWriter, context.Request)

	// Assert
	var responseBody res.AccountResponse
	if err := json.Unmarshal(responseWriter.Body.Bytes(), &responseBody); err != nil {
		panic(err)
	}

	assert.Equal(t, 200, responseWriter.Code)
	assert.False(t, responseBody.BaseResponse.IsSuccess)
	assert.Equal(t, responseBody.BaseResponse.ErrorMessage, "DummyMessage")
}

func loginUserNameNotFoundCase(t *testing.T){
	// Arrange
	var mockAccountRepo repo.IAccountrepo = mocks.MockNotRegisteredUserAccountRepo{}
	var mockRepo repo.IRepo = mocks.MockRepo{}
	var mockHelper helper.IHelper = mocks.MockHelper{}
	accountController := controllers.AccountController{AccountRepo: mockAccountRepo, Repo: mockRepo, Helper: mockHelper}

	responseWriter := httptest.NewRecorder()
	context, router := gin.CreateTestContext(responseWriter)

	var mockLocalizationRepo repo.ILocalizationRepo = mocks.MockLocalizationRepo{}
	var mockMongoOperator db.IMongo = mocks.MockMongoOperator{}
	serviceLogAndErrorMiddleware := middleware.ServiceLogAndErrorMiddleware{LocalizationRepo: mockLocalizationRepo, MongoOperator: mockMongoOperator}
	router.Use(serviceLogAndErrorMiddleware.ServiceLogAndErrorHandler())
	router.POST("login", accountController.Login)

	request := req.AccountRequest{BaseRequest: req.BaseRequest{Version: "0.0.1", Language: "TR"}, UserName: "furkan", Password: "12345"}
	requestBytes, _ := json.Marshal(request)

	context.Request, _ = http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(requestBytes))

	// Act
	router.ServeHTTP(responseWriter, context.Request)

	// Assert
	var responseBody res.AccountResponse
	if err := json.Unmarshal(responseWriter.Body.Bytes(), &responseBody); err != nil {
		panic(err)
	}

	assert.Equal(t, 200, responseWriter.Code)
	assert.False(t, responseBody.BaseResponse.IsSuccess)
	assert.Equal(t, responseBody.BaseResponse.ErrorMessage, "DummyMessage")
}

func loginCheckPasswordFailCase(t *testing.T){
	// Arrange
	var mockAccountRepo repo.IAccountrepo = mocks.MockRegisteredUserAccountRepo{}
	var mockRepo repo.IRepo = mocks.MockRepo{}
	var mockHelper helper.IHelper = mocks.MockCheckPasswordFailHelper{}
	accountController := controllers.AccountController{AccountRepo: mockAccountRepo, Repo: mockRepo, Helper: mockHelper}

	responseWriter := httptest.NewRecorder()
	context, router := gin.CreateTestContext(responseWriter)

	var mockLocalizationRepo repo.ILocalizationRepo = mocks.MockLocalizationRepo{}
	var mockMongoOperator db.IMongo = mocks.MockMongoOperator{}
	serviceLogAndErrorMiddleware := middleware.ServiceLogAndErrorMiddleware{LocalizationRepo: mockLocalizationRepo, MongoOperator: mockMongoOperator}
	router.Use(serviceLogAndErrorMiddleware.ServiceLogAndErrorHandler())
	router.POST("login", accountController.Login)

	request := req.AccountRequest{BaseRequest: req.BaseRequest{Version: "0.0.1", Language: "TR"}, UserName: "furkan", Password: "12345"}
	requestBytes, _ := json.Marshal(request)

	context.Request, _ = http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(requestBytes))

	// Act
	router.ServeHTTP(responseWriter, context.Request)

	// Assert
	var responseBody res.AccountResponse
	if err := json.Unmarshal(responseWriter.Body.Bytes(), &responseBody); err != nil {
		panic(err)
	}

	assert.Equal(t, 200, responseWriter.Code)
	assert.False(t, responseBody.BaseResponse.IsSuccess)
	assert.Equal(t, responseBody.BaseResponse.ErrorMessage, "DummyMessage")
}

func loginAddToSessionFailCase(t *testing.T){
	// Arrange
	var mockAccountRepo repo.IAccountrepo = mocks.MockNotRegisteredUserAccountRepo{}
	var mockRepo repo.IRepo = mocks.MockRepo{}
	var mockHelper helper.IHelper = mocks.MockAddToSessionFailHelper{}
	accountController := controllers.AccountController{AccountRepo: mockAccountRepo, Repo: mockRepo, Helper: mockHelper}

	responseWriter := httptest.NewRecorder()
	context, router := gin.CreateTestContext(responseWriter)

	var mockLocalizationRepo repo.ILocalizationRepo = mocks.MockLocalizationRepo{}
	var mockMongoOperator db.IMongo = mocks.MockMongoOperator{}
	serviceLogAndErrorMiddleware := middleware.ServiceLogAndErrorMiddleware{LocalizationRepo: mockLocalizationRepo, MongoOperator: mockMongoOperator}
	router.Use(serviceLogAndErrorMiddleware.ServiceLogAndErrorHandler())
	router.POST("login", accountController.Login)

	request := req.AccountRequest{BaseRequest: req.BaseRequest{Version: "0.0.1", Language: "TR"}, UserName: "furkan", Password: "12345"}
	requestBytes, _ := json.Marshal(request)

	context.Request, _ = http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(requestBytes))

	// Act
	router.ServeHTTP(responseWriter, context.Request)

	// Assert
	var responseBody res.AccountResponse
	if err := json.Unmarshal(responseWriter.Body.Bytes(), &responseBody); err != nil {
		panic(err)
	}

	assert.Equal(t, 200, responseWriter.Code)
	assert.False(t, responseBody.BaseResponse.IsSuccess)
	assert.Equal(t, responseBody.BaseResponse.ErrorMessage, "DummyMessage")
}

func loginGenerateTokenFailCase(t *testing.T){
	// Arrange
	var mockAccountRepo repo.IAccountrepo = mocks.MockNotRegisteredUserAccountRepo{}
	var mockRepo repo.IRepo = mocks.MockRepo{}
	var mockHelper helper.IHelper = mocks.MockGenerateTokenFailHelper{}
	accountController := controllers.AccountController{AccountRepo: mockAccountRepo, Repo: mockRepo, Helper: mockHelper}

	responseWriter := httptest.NewRecorder()
	context, router := gin.CreateTestContext(responseWriter)

	var mockLocalizationRepo repo.ILocalizationRepo = mocks.MockLocalizationRepo{}
	var mockMongoOperator db.IMongo = mocks.MockMongoOperator{}
	serviceLogAndErrorMiddleware := middleware.ServiceLogAndErrorMiddleware{LocalizationRepo: mockLocalizationRepo, MongoOperator: mockMongoOperator}
	router.Use(serviceLogAndErrorMiddleware.ServiceLogAndErrorHandler())
	router.POST("login", accountController.Login)

	request := req.AccountRequest{BaseRequest: req.BaseRequest{Version: "0.0.1", Language: "TR"}, UserName: "furkan", Password: "12345"}
	requestBytes, _ := json.Marshal(request)

	context.Request, _ = http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(requestBytes))

	// Act
	router.ServeHTTP(responseWriter, context.Request)

	// Assert
	var responseBody res.AccountResponse
	if err := json.Unmarshal(responseWriter.Body.Bytes(), &responseBody); err != nil {
		panic(err)
	}

	assert.Equal(t, 200, responseWriter.Code)
	assert.False(t, responseBody.BaseResponse.IsSuccess)
	assert.Equal(t, responseBody.BaseResponse.ErrorMessage, "DummyMessage")
}

func loginSuccessfullyCase(t *testing.T){
	// Arrange
	var mockAccountRepo repo.IAccountrepo = mocks.MockRegisteredUserAccountRepo{}
	var mockRepo repo.IRepo = mocks.MockRepo{}
	var mockHelper helper.IHelper = mocks.MockHelper{}
	accountController := controllers.AccountController{AccountRepo: mockAccountRepo, Repo: mockRepo, Helper: mockHelper}

	responseWriter := httptest.NewRecorder()
	context, router := gin.CreateTestContext(responseWriter)

	var mockLocalizationRepo repo.ILocalizationRepo = mocks.MockLocalizationRepo{}
	var mockMongoOperator db.IMongo = mocks.MockMongoOperator{}
	serviceLogAndErrorMiddleware := middleware.ServiceLogAndErrorMiddleware{LocalizationRepo: mockLocalizationRepo, MongoOperator: mockMongoOperator}
	router.Use(serviceLogAndErrorMiddleware.ServiceLogAndErrorHandler())
	router.POST("login", accountController.Login)

	request := req.AccountRequest{BaseRequest: req.BaseRequest{Version: "0.0.1", Language: "TR"}, UserName: "furkan", Password: "12345"}
	requestBytes, _ := json.Marshal(request)

	context.Request, _ = http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(requestBytes))

	// Act
	router.ServeHTTP(responseWriter, context.Request)

	// Assert
	var responseBody res.AccountResponse
	if err := json.Unmarshal(responseWriter.Body.Bytes(), &responseBody); err != nil {
		panic(err)
	}

	assert.Equal(t, 200, responseWriter.Code)
	assert.True(t, responseBody.BaseResponse.IsSuccess)
	assert.Equal(t, responseBody.Token, "DummyToken")
}