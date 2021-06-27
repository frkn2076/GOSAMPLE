package router

import (
	"app/GoSample/controllers"
	"app/GoSample/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	
	accountController := new(controllers.AccountController)
	heartBeatController := new(controllers.HeartBeatController)
	todoController := new(controllers.TodoController)
	
	accountRoute := router.Group("/account").Use(middleware.ServiceLogAndErrorHandler())
	{
		accountRoute.POST("login", accountController.Login)
		accountRoute.POST("register", accountController.Register)
	}
	
	heartbeatRoute := router.Group("/heartbeat").Use(middleware.ServiceLogAndErrorHandler())
	{
		heartbeatRoute.GET("/reports", heartBeatController.GetAllReports)
		heartbeatRoute.GET("/clearcache", heartBeatController.ClearCache)
	}
	
	todoRoute:= router.Group("/todo").Use(middleware.AuthorizationHandler()).Use(middleware.ServiceLogAndErrorHandler())
	{
		todoRoute.POST("/add", todoController.AddItem)
		todoRoute.GET("/getall", todoController.GetAllItems)
	}
		
	return router
}
