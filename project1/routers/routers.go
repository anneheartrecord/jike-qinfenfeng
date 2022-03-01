package routers

import (
	"golangstudy/jike/project1/controller"
	"golangstudy/jike/project1/logger"
	"golangstudy/jike/project1/middleware"
	"golangstudy/jike/project1/setting"

	"github.com/gin-gonic/gin"
)

func Init(cfg *setting.ChatRoomConfig) *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true)) //调用自己实现的两个日志中间件
	//注册路由
	u1 := r.Group("/user")
	{
		u1.POST("/", controller.Register)
		u1.GET("/", controller.Login)
		u1.PUT("/", middleware.JWTAuthMiddleware(), controller.Update)
		u1.DELETE("/", middleware.JWTAuthMiddleware(), controller.Exit)
		u1.PUT("/changepwd", middleware.JWTAuthMiddleware(), controller.Changepwd)
		u1.POST("/forgetpwd", middleware.JWTAuthMiddleware(), controller.Forgetpwd)
		u1.PUT("/verify", middleware.JWTAuthMiddleware(), controller.Verify)
	}
	v1 := r.Group("/room")
	{
		v1.POST("/", middleware.JWTAuthMiddleware(), controller.CreateRoom)
		v1.PUT("/", middleware.JWTAuthMiddleware(), controller.ChangeRoom)
		v1.POST("/enter", middleware.JWTAuthMiddleware(), controller.EnterRoom)
		v1.DELETE("/", middleware.JWTAuthMiddleware(), controller.ExitRoom)
	}
	return r
}
