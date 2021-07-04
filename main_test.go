package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"app/GoSample/config/environments"
	"app/GoSample/config/session"
	"app/GoSample/controllers"
	"app/GoSample/controllers/helper"
	req "app/GoSample/controllers/models/request"
	res "app/GoSample/controllers/models/response"
	"app/GoSample/db"
	"app/GoSample/db/repo"
	"app/GoSample/infra/constant"
	"app/GoSample/logger"
	"app/GoSample/middleware"
	"app/GoSample/test/mocks"
	"app/GoSample/db/entities"

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

func TestAddItem(t *testing.T) {
	addItemRequestFailCase(t)
	addItemStringToUintFailCase(t)
	addItemCreateTodoFailCase(t)
	addItemSuccessfullyCase(t)
}

func TestGetAllItems(t *testing.T) {
	getAllGettingAllFailCase(t)
	getAllSuccessfullyCase(t)
}

func TestUpdateItem(t *testing.T){
	updateItemRequestFailCase(t)
	updateItemStringToUintFailCase(t)
	updateItemUpdateTodoFailCase(t)
	updateItemSuccessfullyCase(t)
}

func TestDeleteItem(t *testing.T){
	deleteItemStringToUintFailCase(t)
	deleteItemSuccessfullyCase(t)
}

func TestGetAllReports(t *testing.T){
	getAllReportsSuccessfullyCase(t)
}

func registerRequestFailCase(t *testing.T) {
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

	request := req.AccountRequest{BaseRequest: req.BaseRequest{Language: "TR"}, UserName: "furkan", Password: "12345"}
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

func registerHashPasswordFailCase(t *testing.T) {
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

	request := req.AccountRequest{BaseRequest: req.BaseRequest{Language: "TR"}, UserName: "furkan", Password: "12345"}
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

func registerUserExistCase(t *testing.T) {
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

	request := req.AccountRequest{BaseRequest: req.BaseRequest{Language: "TR"}, UserName: "furkan", Password: "12345"}
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

func registerAccountCreateFailCase(t *testing.T) {
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

	request := req.AccountRequest{BaseRequest: req.BaseRequest{Language: "TR"}, UserName: "furkan", Password: "12345"}
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

func registerAddToSessionFailCase(t *testing.T) {
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

	request := req.AccountRequest{BaseRequest: req.BaseRequest{Language: "TR"}, UserName: "furkan", Password: "12345"}
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

func registerGenerateTokenFailCase(t *testing.T) {
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

	request := req.AccountRequest{BaseRequest: req.BaseRequest{Language: "TR"}, UserName: "furkan", Password: "12345"}
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

func registerSuccessfullyCase(t *testing.T) {
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

	request := req.AccountRequest{BaseRequest: req.BaseRequest{Language: "TR"}, UserName: "furkan", Password: "12345"}
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

func loginRequestFailCase(t *testing.T) {
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

	request := req.AccountRequest{BaseRequest: req.BaseRequest{Language: "TR"}, UserName: "furkan", Password: "12345"}
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

func loginUserNameNotFoundCase(t *testing.T) {
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

	request := req.AccountRequest{BaseRequest: req.BaseRequest{Language: "TR"}, UserName: "furkan", Password: "12345"}
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

func loginCheckPasswordFailCase(t *testing.T) {
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

	request := req.AccountRequest{BaseRequest: req.BaseRequest{Language: "TR"}, UserName: "furkan", Password: "12345"}
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

func loginAddToSessionFailCase(t *testing.T) {
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

	request := req.AccountRequest{BaseRequest: req.BaseRequest{Language: "TR"}, UserName: "furkan", Password: "12345"}
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

func loginGenerateTokenFailCase(t *testing.T) {
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

	request := req.AccountRequest{BaseRequest: req.BaseRequest{Language: "TR"}, UserName: "furkan", Password: "12345"}
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

func loginSuccessfullyCase(t *testing.T) {
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

	request := req.AccountRequest{BaseRequest: req.BaseRequest{Language: "TR"}, UserName: "furkan", Password: "12345"}
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

func addItemRequestFailCase(t *testing.T) {
	// Arrange
	var mockTodoRepo repo.ITodoRepo = mocks.MockTodoRepo{}
	var mockRepo repo.IRepo = mocks.MockRepo{}
	var mockHelper helper.IHelper = mocks.MockRequestFailHelper{}
	todoController := controllers.TodoController{TodoRepo: mockTodoRepo, Repo: mockRepo, Helper: mockHelper}

	responseWriter := httptest.NewRecorder()
	context, router := gin.CreateTestContext(responseWriter)

	var mockLocalizationRepo repo.ILocalizationRepo = mocks.MockLocalizationRepo{}
	var mockMongoOperator db.IMongo = mocks.MockMongoOperator{}
	serviceLogAndErrorMiddleware := middleware.ServiceLogAndErrorMiddleware{LocalizationRepo: mockLocalizationRepo, MongoOperator: mockMongoOperator}
	router.Use(serviceLogAndErrorMiddleware.ServiceLogAndErrorHandler())
	router.POST("add", todoController.AddItem)

	request := req.TodoRequest{BaseRequest: req.BaseRequest{Language: "TR"},
		Name: "DummyName1", Description: "DummyDescription1", Deadline: time.Now(), IsCompleted: true}
	requestBytes, _ := json.Marshal(request)

	context.Request, _ = http.NewRequest(http.MethodPost, "/add", bytes.NewBuffer(requestBytes))

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

func addItemStringToUintFailCase(t *testing.T) {
	// Arrange
	var mockTodoRepo repo.ITodoRepo = mocks.MockTodoRepo{}
	var mockRepo repo.IRepo = mocks.MockRepo{}
	var mockHelper helper.IHelper = mocks.MockStringToUintFailHelper{}
	todoController := controllers.TodoController{TodoRepo: mockTodoRepo, Repo: mockRepo, Helper: mockHelper}

	responseWriter := httptest.NewRecorder()
	context, router := gin.CreateTestContext(responseWriter)

	var mockLocalizationRepo repo.ILocalizationRepo = mocks.MockLocalizationRepo{}
	var mockMongoOperator db.IMongo = mocks.MockMongoOperator{}
	serviceLogAndErrorMiddleware := middleware.ServiceLogAndErrorMiddleware{LocalizationRepo: mockLocalizationRepo, MongoOperator: mockMongoOperator}
	router.Use(serviceLogAndErrorMiddleware.ServiceLogAndErrorHandler())
	router.POST("add", todoController.AddItem)

	request := req.TodoRequest{BaseRequest: req.BaseRequest{Language: "TR"},
		Name: "DummyName1", Description: "DummyDescription1", Deadline: time.Now(), IsCompleted: true}
	requestBytes, _ := json.Marshal(request)

	context.Request, _ = http.NewRequest(http.MethodPost, "/add", bytes.NewBuffer(requestBytes))

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

func addItemCreateTodoFailCase(t *testing.T) {
	// Arrange
	var mockTodoRepo repo.ITodoRepo = mocks.MockCreateTodoFailTodoRepo{}
	var mockRepo repo.IRepo = mocks.MockRepo{}
	var mockHelper helper.IHelper = mocks.MockHelper{}
	todoController := controllers.TodoController{TodoRepo: mockTodoRepo, Repo: mockRepo, Helper: mockHelper}

	responseWriter := httptest.NewRecorder()
	context, router := gin.CreateTestContext(responseWriter)

	var mockLocalizationRepo repo.ILocalizationRepo = mocks.MockLocalizationRepo{}
	var mockMongoOperator db.IMongo = mocks.MockMongoOperator{}
	serviceLogAndErrorMiddleware := middleware.ServiceLogAndErrorMiddleware{LocalizationRepo: mockLocalizationRepo, MongoOperator: mockMongoOperator}
	router.Use(serviceLogAndErrorMiddleware.ServiceLogAndErrorHandler())
	router.POST("add", todoController.AddItem)

	request := req.TodoRequest{BaseRequest: req.BaseRequest{Language: "TR"},
		Name: "DummyName1", Description: "DummyDescription1", Deadline: time.Now(), IsCompleted: true}
	requestBytes, _ := json.Marshal(request)

	context.Request, _ = http.NewRequest(http.MethodPost, "/add", bytes.NewBuffer(requestBytes))

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

func addItemSuccessfullyCase(t *testing.T) {
	// Arrange
	var mockTodoRepo repo.ITodoRepo = mocks.MockTodoRepo{}
	var mockRepo repo.IRepo = mocks.MockRepo{}
	var mockHelper helper.IHelper = mocks.MockHelper{}
	todoController := controllers.TodoController{TodoRepo: mockTodoRepo, Repo: mockRepo, Helper: mockHelper}

	responseWriter := httptest.NewRecorder()
	context, router := gin.CreateTestContext(responseWriter)

	var mockLocalizationRepo repo.ILocalizationRepo = mocks.MockLocalizationRepo{}
	var mockMongoOperator db.IMongo = mocks.MockMongoOperator{}
	serviceLogAndErrorMiddleware := middleware.ServiceLogAndErrorMiddleware{LocalizationRepo: mockLocalizationRepo, MongoOperator: mockMongoOperator}
	router.Use(serviceLogAndErrorMiddleware.ServiceLogAndErrorHandler())
	router.POST("add", todoController.AddItem)

	request := req.TodoRequest{BaseRequest: req.BaseRequest{Language: "TR"},
		Name: "DummyName1", Description: "DummyDescription1", Deadline: time.Now(), IsCompleted: true}
	requestBytes, _ := json.Marshal(request)

	context.Request, _ = http.NewRequest(http.MethodPost, "/add", bytes.NewBuffer(requestBytes))

	// Act
	router.ServeHTTP(responseWriter, context.Request)

	// Assert
	var responseBody res.BaseResponse
	if err := json.Unmarshal(responseWriter.Body.Bytes(), &responseBody); err != nil {
		panic(err)
	}

	assert.Equal(t, 200, responseWriter.Code)
	assert.True(t, responseBody.IsSuccess)
}

func getAllGettingAllFailCase(t *testing.T) {
	// Arrange
	var mockTodoRepo repo.ITodoRepo = mocks.MockGetAllFailTodoRepo{}
	var mockRepo repo.IRepo = mocks.MockRepo{}
	var mockHelper helper.IHelper = mocks.MockHelper{}
	todoController := controllers.TodoController{TodoRepo: mockTodoRepo, Repo: mockRepo, Helper: mockHelper}

	responseWriter := httptest.NewRecorder()
	context, router := gin.CreateTestContext(responseWriter)

	var mockLocalizationRepo repo.ILocalizationRepo = mocks.MockLocalizationRepo{}
	var mockMongoOperator db.IMongo = mocks.MockMongoOperator{}
	serviceLogAndErrorMiddleware := middleware.ServiceLogAndErrorMiddleware{LocalizationRepo: mockLocalizationRepo, MongoOperator: mockMongoOperator}
	router.Use(serviceLogAndErrorMiddleware.ServiceLogAndErrorHandler())
	router.GET("getall", todoController.GetAllItems)

	context.Request, _ = http.NewRequest(http.MethodGet, "/getall", bytes.NewBuffer([]byte(constant.EmptyString)))

	// Act
	router.ServeHTTP(responseWriter, context.Request)

	// Assert
	var responseBody res.BaseResponse
	if err := json.Unmarshal(responseWriter.Body.Bytes(), &responseBody); err != nil {
		panic(err)
	}

	assert.Equal(t, 200, responseWriter.Code)
	assert.False(t, responseBody.IsSuccess)
}

func getAllSuccessfullyCase(t *testing.T) {
	// Arrange
	expectedTodos := []entities.Todo{
		entities.Todo{ UserId: 1, Name: "DummyName1", Description: "DummyDescription1", Deadline: time.Date(1999, time.January, 03, 0, 0, 0, 0, time.UTC), IsCompleted: true },
		entities.Todo{ UserId: 1, Name: "DummyName2", Description: "DummyDescription2", Deadline: time.Date(1999, time.January, 03, 0, 0, 0, 0, time.UTC), IsCompleted: false },
	}

	var mockTodoRepo repo.ITodoRepo = mocks.MockTodoRepo{}
	var mockRepo repo.IRepo = mocks.MockRepo{}
	var mockHelper helper.IHelper = mocks.MockHelper{}
	todoController := controllers.TodoController{TodoRepo: mockTodoRepo, Repo: mockRepo, Helper: mockHelper}

	responseWriter := httptest.NewRecorder()
	context, router := gin.CreateTestContext(responseWriter)

	var mockLocalizationRepo repo.ILocalizationRepo = mocks.MockLocalizationRepo{}
	var mockMongoOperator db.IMongo = mocks.MockMongoOperator{}
	serviceLogAndErrorMiddleware := middleware.ServiceLogAndErrorMiddleware{LocalizationRepo: mockLocalizationRepo, MongoOperator: mockMongoOperator}
	router.Use(serviceLogAndErrorMiddleware.ServiceLogAndErrorHandler())
	router.GET("getall", todoController.GetAllItems)

	context.Request, _ = http.NewRequest(http.MethodGet, "/getall", bytes.NewBuffer([]byte(constant.EmptyString)))

	// Act
	router.ServeHTTP(responseWriter, context.Request)

	// Assert
	var responseBody res.TodosResponse
	if err := json.Unmarshal(responseWriter.Body.Bytes(), &responseBody); err != nil {
		panic(err)
	}

	assert.Equal(t, 200, responseWriter.Code)
	assert.True(t, responseBody.BaseResponse.IsSuccess)
	assert.Equal(t, expectedTodos[0].Name, responseBody.Todos[0].Name)
	assert.Equal(t, expectedTodos[0].Description, responseBody.Todos[0].Description)
	assert.Equal(t, expectedTodos[0].Deadline, responseBody.Todos[0].Deadline)
	assert.Equal(t, expectedTodos[0].IsCompleted, responseBody.Todos[0].IsCompleted)
	assert.Equal(t, expectedTodos[1].Name, responseBody.Todos[1].Name)
	assert.Equal(t, expectedTodos[1].Description, responseBody.Todos[1].Description)
	assert.Equal(t, expectedTodos[1].Deadline, responseBody.Todos[1].Deadline)
	assert.Equal(t, expectedTodos[1].IsCompleted, responseBody.Todos[1].IsCompleted)
}

func updateItemRequestFailCase(t *testing.T) {
	// Arrange
	var mockTodoRepo repo.ITodoRepo = mocks.MockTodoRepo{}
	var mockRepo repo.IRepo = mocks.MockRepo{}
	var mockHelper helper.IHelper = mocks.MockRequestFailHelper{}
	todoController := controllers.TodoController{TodoRepo: mockTodoRepo, Repo: mockRepo, Helper: mockHelper}

	responseWriter := httptest.NewRecorder()
	context, router := gin.CreateTestContext(responseWriter)

	var mockLocalizationRepo repo.ILocalizationRepo = mocks.MockLocalizationRepo{}
	var mockMongoOperator db.IMongo = mocks.MockMongoOperator{}
	serviceLogAndErrorMiddleware := middleware.ServiceLogAndErrorMiddleware{LocalizationRepo: mockLocalizationRepo, MongoOperator: mockMongoOperator}
	router.Use(serviceLogAndErrorMiddleware.ServiceLogAndErrorHandler())
	router.POST("update", todoController.UpdateItem)

	request := req.TodoRequest{BaseRequest: req.BaseRequest{Language: "TR"},
		Name: "DummyName1", Description: "DummyDescription1", Deadline: time.Now(), IsCompleted: true}
	requestBytes, _ := json.Marshal(request)

	context.Request, _ = http.NewRequest(http.MethodPost, "/update", bytes.NewBuffer(requestBytes))

	// Act
	router.ServeHTTP(responseWriter, context.Request)

	// Assert
	var responseBody res.BaseResponse
	if err := json.Unmarshal(responseWriter.Body.Bytes(), &responseBody); err != nil {
		panic(err)
	}

	assert.Equal(t, 200, responseWriter.Code)
	assert.False(t, responseBody.IsSuccess)
}

func updateItemStringToUintFailCase(t *testing.T) {
	// Arrange
	var mockTodoRepo repo.ITodoRepo = mocks.MockTodoRepo{}
	var mockRepo repo.IRepo = mocks.MockRepo{}
	var mockHelper helper.IHelper = mocks.MockStringToUintFailHelper{}
	todoController := controllers.TodoController{TodoRepo: mockTodoRepo, Repo: mockRepo, Helper: mockHelper}

	responseWriter := httptest.NewRecorder()
	context, router := gin.CreateTestContext(responseWriter)

	var mockLocalizationRepo repo.ILocalizationRepo = mocks.MockLocalizationRepo{}
	var mockMongoOperator db.IMongo = mocks.MockMongoOperator{}
	serviceLogAndErrorMiddleware := middleware.ServiceLogAndErrorMiddleware{LocalizationRepo: mockLocalizationRepo, MongoOperator: mockMongoOperator}
	router.Use(serviceLogAndErrorMiddleware.ServiceLogAndErrorHandler())
	router.POST("update", todoController.UpdateItem)

	request := req.TodoRequest{BaseRequest: req.BaseRequest{Language: "TR"},
		Name: "DummyName1", Description: "DummyDescription1", Deadline: time.Now(), IsCompleted: true}
	requestBytes, _ := json.Marshal(request)

	context.Request, _ = http.NewRequest(http.MethodPost, "/update", bytes.NewBuffer(requestBytes))

	// Act
	router.ServeHTTP(responseWriter, context.Request)

	// Assert
	var responseBody res.BaseResponse
	if err := json.Unmarshal(responseWriter.Body.Bytes(), &responseBody); err != nil {
		panic(err)
	}

	assert.Equal(t, 200, responseWriter.Code)
	assert.False(t, responseBody.IsSuccess)
}

func updateItemUpdateTodoFailCase(t *testing.T) {
	// Arrange
	var mockTodoRepo repo.ITodoRepo = mocks.MockUpdateFailTodoRepo{}
	var mockRepo repo.IRepo = mocks.MockRepo{}
	var mockHelper helper.IHelper = mocks.MockHelper{}
	todoController := controllers.TodoController{TodoRepo: mockTodoRepo, Repo: mockRepo, Helper: mockHelper}

	responseWriter := httptest.NewRecorder()
	context, router := gin.CreateTestContext(responseWriter)

	var mockLocalizationRepo repo.ILocalizationRepo = mocks.MockLocalizationRepo{}
	var mockMongoOperator db.IMongo = mocks.MockMongoOperator{}
	serviceLogAndErrorMiddleware := middleware.ServiceLogAndErrorMiddleware{LocalizationRepo: mockLocalizationRepo, MongoOperator: mockMongoOperator}
	router.Use(serviceLogAndErrorMiddleware.ServiceLogAndErrorHandler())
	router.POST("update", todoController.UpdateItem)

	request := req.TodoRequest{BaseRequest: req.BaseRequest{Language: "TR"},
		Name: "DummyName1", Description: "DummyDescription1", Deadline: time.Now(), IsCompleted: true}
	requestBytes, _ := json.Marshal(request)

	context.Request, _ = http.NewRequest(http.MethodPost, "/update", bytes.NewBuffer(requestBytes))

	// Act
	router.ServeHTTP(responseWriter, context.Request)

	// Assert
	var responseBody res.BaseResponse
	if err := json.Unmarshal(responseWriter.Body.Bytes(), &responseBody); err != nil {
		panic(err)
	}

	assert.Equal(t, 200, responseWriter.Code)
	assert.False(t, responseBody.IsSuccess)
}

func updateItemSuccessfullyCase(t *testing.T) {
	// Arrange
	var mockTodoRepo repo.ITodoRepo = mocks.MockTodoRepo{}
	var mockRepo repo.IRepo = mocks.MockRepo{}
	var mockHelper helper.IHelper = mocks.MockHelper{}
	todoController := controllers.TodoController{TodoRepo: mockTodoRepo, Repo: mockRepo, Helper: mockHelper}

	responseWriter := httptest.NewRecorder()
	context, router := gin.CreateTestContext(responseWriter)

	var mockLocalizationRepo repo.ILocalizationRepo = mocks.MockLocalizationRepo{}
	var mockMongoOperator db.IMongo = mocks.MockMongoOperator{}
	serviceLogAndErrorMiddleware := middleware.ServiceLogAndErrorMiddleware{LocalizationRepo: mockLocalizationRepo, MongoOperator: mockMongoOperator}
	router.Use(serviceLogAndErrorMiddleware.ServiceLogAndErrorHandler())
	router.POST("update", todoController.UpdateItem)

	request := req.TodoRequest{BaseRequest: req.BaseRequest{Language: "TR"},
		Name: "DummyName1", Description: "DummyDescription1", Deadline: time.Now(), IsCompleted: true}
	requestBytes, _ := json.Marshal(request)

	context.Request, _ = http.NewRequest(http.MethodPost, "/update", bytes.NewBuffer(requestBytes))

	// Act
	router.ServeHTTP(responseWriter, context.Request)

	// Assert
	var responseBody res.BaseResponse
	if err := json.Unmarshal(responseWriter.Body.Bytes(), &responseBody); err != nil {
		panic(err)
	}

	assert.Equal(t, 200, responseWriter.Code)
	assert.True(t, responseBody.IsSuccess)
}

func deleteItemStringToUintFailCase(t *testing.T) {
	// Arrange
	var mockTodoRepo repo.ITodoRepo = mocks.MockTodoRepo{}
	var mockRepo repo.IRepo = mocks.MockRepo{}
	var mockHelper helper.IHelper = mocks.MockStringToUintFailHelper{}
	todoController := controllers.TodoController{TodoRepo: mockTodoRepo, Repo: mockRepo, Helper: mockHelper}

	responseWriter := httptest.NewRecorder()
	context, router := gin.CreateTestContext(responseWriter)

	var mockLocalizationRepo repo.ILocalizationRepo = mocks.MockLocalizationRepo{}
	var mockMongoOperator db.IMongo = mocks.MockMongoOperator{}
	serviceLogAndErrorMiddleware := middleware.ServiceLogAndErrorMiddleware{LocalizationRepo: mockLocalizationRepo, MongoOperator: mockMongoOperator}
	router.Use(serviceLogAndErrorMiddleware.ServiceLogAndErrorHandler())
	router.GET("delete/:todoId", todoController.DeleteItem)

	context.Request, _ = http.NewRequest(http.MethodGet, "/delete/1", bytes.NewBuffer([]byte(constant.EmptyString)))

	// Act
	router.ServeHTTP(responseWriter, context.Request)

	// Assert
	var responseBody res.BaseResponse
	if err := json.Unmarshal(responseWriter.Body.Bytes(), &responseBody); err != nil {
		panic(err)
	}

	assert.Equal(t, 200, responseWriter.Code)
	assert.False(t, responseBody.IsSuccess)
}

func deleteItemSuccessfullyCase(t *testing.T) {
	// Arrange
	var mockTodoRepo repo.ITodoRepo = mocks.MockTodoRepo{}
	var mockRepo repo.IRepo = mocks.MockRepo{}
	var mockHelper helper.IHelper = mocks.MockHelper{}
	todoController := controllers.TodoController{TodoRepo: mockTodoRepo, Repo: mockRepo, Helper: mockHelper}

	responseWriter := httptest.NewRecorder()
	context, router := gin.CreateTestContext(responseWriter)

	var mockLocalizationRepo repo.ILocalizationRepo = mocks.MockLocalizationRepo{}
	var mockMongoOperator db.IMongo = mocks.MockMongoOperator{}
	serviceLogAndErrorMiddleware := middleware.ServiceLogAndErrorMiddleware{LocalizationRepo: mockLocalizationRepo, MongoOperator: mockMongoOperator}
	router.Use(serviceLogAndErrorMiddleware.ServiceLogAndErrorHandler())
	router.GET("delete/:todoId", todoController.DeleteItem)

	context.Request, _ = http.NewRequest(http.MethodGet, "/delete/1", bytes.NewBuffer([]byte(constant.EmptyString)))

	// Act
	router.ServeHTTP(responseWriter, context.Request)

	// Assert
	var responseBody res.BaseResponse
	if err := json.Unmarshal(responseWriter.Body.Bytes(), &responseBody); err != nil {
		panic(err)
	}

	assert.Equal(t, 200, responseWriter.Code)
	assert.True(t, responseBody.IsSuccess)
}

func getAllReportsSuccessfullyCase(t *testing.T) {
	// Arrange
	heartBeatController := controllers.HeartBeatController{}

	responseWriter := httptest.NewRecorder()
	context, router := gin.CreateTestContext(responseWriter)

	var mockLocalizationRepo repo.ILocalizationRepo = mocks.MockLocalizationRepo{}
	var mockMongoOperator db.IMongo = mocks.MockMongoOperator{}
	serviceLogAndErrorMiddleware := middleware.ServiceLogAndErrorMiddleware{LocalizationRepo: mockLocalizationRepo, MongoOperator: mockMongoOperator}
	router.Use(serviceLogAndErrorMiddleware.ServiceLogAndErrorHandler())
	router.GET("reports", heartBeatController.GetAllReports)

	context.Request, _ = http.NewRequest(http.MethodGet, "/reports", bytes.NewBuffer([]byte(constant.EmptyString)))

	// Act
	router.ServeHTTP(responseWriter, context.Request)

	// Assert
	assert.Equal(t, 200, responseWriter.Code)
	assert.NotNil(t, responseWriter.Body)
}