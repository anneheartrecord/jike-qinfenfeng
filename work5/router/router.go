package router

import (
	"github.com/gin-gonic/gin"
	"golangstudy/jike/work5/controller"
)

func Router()   { //路由分组
	r:=gin.Default()
	routerGroupUser:=r.Group("/user")
	{
		routerGroupUser.POST("/", controller.Register)
		routerGroupUser.GET("/",controller.Login)
		routerGroupUser.PUT("/",controller.Update)
		routerGroupUser.DELETE("/",controller.Delete)
	}
	routerGroupCourse:=r.Group("/course")
	{
		routerGroupCourse.POST("/",controller.CreateCourse)
		routerGroupCourse.GET("/",controller.ShowCourse)
		routerGroupCourse.PUT("/",controller.UpdateCourse)
		routerGroupCourse.DELETE("/",controller.DeleteCourse)
	}
	routerGroupUserAndCourse:=r.Group("/userCourse")
	{
		routerGroupUserAndCourse.POST("/",controller.CreateRecordInfo)
		routerGroupUserAndCourse.GET("/",controller.ShowStudentInfo)
		routerGroupUserAndCourse.DELETE("/",controller.DelteRecordInfo)
	}
	r.Run(":8080")
}