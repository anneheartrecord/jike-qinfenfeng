package view

import (
	"github.com/gin-gonic/gin"
	"golangstudy/jike/work5/models"
)

func Register(c *gin.Context) (string ,string ,string,int ,bool)  { //学生注册
	var u models.UserInfo
	c.ShouldBind(&u)
	res:=models.Register(u.Username,u.Password,u.Number,u.MaxCredit) //res是models层数据库操作的返回值
	return u.Username,u.Password,u.Number,u.MaxCredit,res
}
func Login(c * gin.Context) (string ,string ,string ,bool) { //学生登录
	username:=c.PostForm("username")
	password:=c.PostForm("password")
	number:=c.PostForm("number")
	res:=models.Login(number,password)
	return username,password,number,res
}
func Update(c *gin.Context) (string ,string ,string ,bool) { //学生改密码
	username:=c.PostForm("username")
	password:=c.PostForm("password")
	number:=c.PostForm("number")
	res:=models.Update(number,password)
	return username,password,number,res
}
func Delete(c *gin.Context) (bool)  { //学生删除
	number:=c.PostForm("number")
	res:=models.Delete(number)
	return res
}
