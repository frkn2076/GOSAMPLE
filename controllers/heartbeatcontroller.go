package controllers

import (
	"fmt"
	"strconv"

	"app/GoSample/config/cache"

	"github.com/gin-gonic/gin"
)

type HeartBeatController struct{}

func (u *HeartBeatController) GetAllReports(c *gin.Context) {
	report := fmt.Sprintf("Cache average access time: %s milliseconds", strconv.FormatInt(cache.GetAvaregeAccessTime(), 10))
	c.JSON(200, report)
}

func (u *HeartBeatController) ClearCache(c *gin.Context) {
	cache.Reset()
	c.AbortWithStatus(200)
}
