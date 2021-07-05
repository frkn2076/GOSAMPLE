package router

import (
	"app/GoSample/controllers"
	"app/GoSample/middleware"
	"app/GoSample/db"
	"app/GoSample/db/repo"
	"app/GoSample/controllers/helper"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	
	accountController := controllers.AccountController{AccountRepo: repo.Account, Repo: repo.Repo, Helper: helper.HelperInstance}
	todoController := controllers.TodoController{TodoRepo: repo.Todo, Repo: repo.Repo, Helper: helper.HelperInstance}
	heartBeatController := controllers.HeartBeatController{}

	corsMiddleware := middleware.CORSMiddleware{}
	authorizationMiddleware := middleware.AuthorizationMiddleware{}
	serviceLogAndErrorMiddleware := middleware.ServiceLogAndErrorMiddleware{LocalizationRepo: repo.Localization, MongoOperator: db.Mongo}
	
	router.Use(corsMiddleware.CORSHandler())

	accountRoute := router.Group("/account").Use(serviceLogAndErrorMiddleware.ServiceLogAndErrorHandler())
	{
		accountRoute.POST("login", accountController.Login)
		accountRoute.POST("register", accountController.Register)
	}
	
	heartbeatRoute := router.Group("/heartbeat").Use(serviceLogAndErrorMiddleware.ServiceLogAndErrorHandler())
	{
		heartbeatRoute.GET("/reports", heartBeatController.GetAllReports)
		// Will be using when admin login feature added
		// heartbeatRoute.GET("/clearcache", heartBeatController.ClearCache)
	}
	
	todoRoute:= router.Group("/todo").Use(authorizationMiddleware.AuthorizationHandler()).Use(serviceLogAndErrorMiddleware.ServiceLogAndErrorHandler())
	{
		todoRoute.POST("/add", todoController.AddItem)
		todoRoute.GET("/getall", todoController.GetAllItems)
		todoRoute.POST("/update", todoController.UpdateItem)
		todoRoute.GET("/delete/:todoId", todoController.DeleteItem)
	}

	return router
}
