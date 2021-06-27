package router

import (
	"app/GoSample/controllers"
	"app/GoSample/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.ServiceLogMiddleware())

	accountController := new(controllers.AccountController)
	heartBeatController := new(controllers.HeartBeatController)

	grp1 := router.Group("/account")
	{
		grp1.POST("login", accountController.Login)
		grp1.POST("register", accountController.Register)
	}

	grp2 := router.Group("/heartbeat")
	{
		grp2.GET("/reports", heartBeatController.GetAllReports)
		grp2.GET("/clearCache", heartBeatController.ClearCache)
	}

	return router
}
