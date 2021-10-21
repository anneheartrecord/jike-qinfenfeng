package view

import (
	"github.com/gin-gonic/gin"
	"golangstudy/jike/work5/models"
)

func  CreateRecordInfo(c *gin.Context) (string,string,string,bool)  {  //提交选课
	 Username:=c.PostForm("username")
	 Coursename:=c.PostForm("coursename")
	 Number:=c.PostForm("number")
	 res:=models.CreateRecordInfo(Username,Coursename,Number)
	 return  Username,Coursename,Number,res
}
func ShowStudentInfo(c *gin.Context) (string,[] string,bool)  {//查询选课
	Username:=c.PostForm("username")
	u,res:=models.ShowStudentInfo(Username)
	return Username,u,res
}
func DelteRecordInfo(c *gin.Context) bool   {  //删除选课
	Username:=c.PostForm("username")
	res:=models.DelteRecordInfo(Username)
	return res
}