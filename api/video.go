package api

import (
	"giligili/service"
	"github.com/gin-gonic/gin"
)

//模板
/*func (c *gin.Context){
	service := service.Serveice{}
	if err := c.ShouldBind(&service); err==nil{
		res := service.Create()
		c.JSON(200,res)
	}else{
		c.JSON(200,ErrorResponse(err))
	}
}*/

func CreateVideo(c *gin.Context) {
	service := service.CreateVideoService{}
	if err := c.ShouldBind(&service); err==nil{
		res := service.Create()
		c.JSON(200,res)
	}else{
		c.JSON(200,ErrorResponse(err))
	}
}

func ShowVideo(c *gin.Context) {}
func ListVideos(c *gin.Context) {}
func UpdateVideo(c *gin.Context) {}
func DeleteVideo(c *gin.Context) {}
