package api

import (
	"giligili/service"
	"github.com/gin-gonic/gin"
)

func DailyRank(c *gin.Context) {
	service := service.DailyRankService{}
	c.JSON(200,service.Get())
}
