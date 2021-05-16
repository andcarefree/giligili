package api

import (
	"giligili/model"
	"giligili/service"
	"github.com/gin-gonic/gin"
)

//模板
/*func (c *gin.Context){
	service := service.Service{}
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
		//绑定作者id
		//TODO 感觉这里不太优雅，看看有没有系统一些使用的方法
		User,_ := c.Get("user")
		service.AuthID = User.(*model.User).ID

		res := service.Create()
		c.JSON(200,res)
	}else{
		c.JSON(200,ErrorResponse(err))
	}
}

func ShowVideo(c *gin.Context) {
	service := service.ShowVideoService{}
	if err := c.ShouldBind(&service); err==nil{
		res := service.Show(c.Param("id"))
		c.JSON(200,res)
	}else{
		c.JSON(200,ErrorResponse(err))
	}
}

func ListVideos(c *gin.Context) {
	service := service.ListVideosService{}
	res := service.List()
	c.JSON(200,res)
}
func UpdateVideo(c *gin.Context) {
	service := service.UpdateVideoService{}
	if err := c.ShouldBind(&service); err==nil{
		res := service.Update(c.Param("id"))
		c.JSON(200,res)
	}else{
		c.JSON(200,ErrorResponse(err))
	}
}

func DeleteVideo(c *gin.Context) {
	service := service.DeleteVideoService{}
	res := service.Delete(c.Param("id"))
	c.JSON(200,res)
}
