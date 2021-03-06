package server

import (
	"giligili/api"
	"giligili/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.POST("ping", api.Ping)

		// 用户登录
		v1.POST("user/register", api.UserRegister)

		// 用户登录
		v1.POST("user/login", api.UserLogin)

		// 需要登录保护的
		auth := v1.Group("")
		auth.Use(middleware.AuthRequired())
		{
			// User Routing
			auth.GET("user/me", api.UserMe)
			auth.DELETE("user/logout", api.UserLogout)

			//用户对视频的提交、修改、删除
			auth.POST("/video",api.CreateVideo)
			auth.PUT("/video/:id",api.UpdateVideo)
			auth.DELETE("/video/:id",api.DeleteVideo)
		}

		//查看视频
		v1.GET("/video/:id",api.ShowVideo)
		v1.GET("/videos",api.ListVideos)

		//查看排行榜
		v1.GET("/rank/daily",api.DailyRank)
	}
	return r
}
