package controllers

import (
	"fmt"
	"strconv"

	"app/GoSample/config/cache"

	"github.com/gin-gonic/gin"
)

type HeartBeatController struct{}

func (u *HeartBeatController) GetAllReports(context *gin.Context) {
	report := fmt.Sprintf("Cache average access time: %s milliseconds", strconv.FormatInt(cache.GetAvaregeAccessTime(), 10))
	context.JSON(200, report)
}

// Will be using when admin login feature added
// func (u *HeartBeatController) ClearCache(context *gin.Context) {
// 	cache.Reset()
// 	context.JSON(200, response.Success)
// }
