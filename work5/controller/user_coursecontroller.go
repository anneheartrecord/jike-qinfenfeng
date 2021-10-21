package controller

import (
	"github.com/gin-gonic/gin"
	"golangstudy/jike/work5/view"
	"net/http"
)

func CreateRecordInfo(c *gin.Context)  { //提交选课
	Username,Coursename,Number,res:=view.CreateRecordInfo(c)
	if !res {
		c.JSON(http.StatusNotFound,gin.H{
			"code": 404,
			"msg": "选课失败",
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"code": 200,
		"msg": "选课成功",
		"username": Username,
		"coursename": Coursename,
		"number": Number,
	})
}
func ShowStudentInfo(c *gin.Context)  {  //查看学生选课信息
	Number,u,res:=view.ShowStudentInfo(c)
	if !res {
		c.JSON(http.StatusNotFound,gin.H{
			"code": 404,
			"msg": "查询学生选课信息失败",
		})
		return
	}
		c.JSON(http.StatusOK,gin.H{
			"code": 200,
			"msg": "查询信息成功",
			"number": Number,
			"course": u,
		})
}
func DelteRecordInfo(c *gin.Context)  { //删除学生选课信息
	res:=view.DelteRecordInfo(c)
	if !res {
		c.JSON(http.StatusNotFound,gin.H{
			"code": 404,
			"msg": "删除学生选课信息失败",
		})
	}
	c.JSON(http.StatusOK,gin.H{
		"code": 200,
		"msg": "删除学生选课信息成功",
	})
}
