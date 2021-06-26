package controllers

import (
	"fmt"
	"strconv"

	"app/GoSample/config/cache"
	"app/GoSample/db/nosql"
	_ "app/GoSample/infra/customeError"

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

func (u *HeartBeatController) GetLogs(c *gin.Context) {
	result := nosql.GetLogRecord()
	c.JSON(200, result)
}
