package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golangstudy/jike/work5/view"
	"net/http"
)

func Register(c *gin.Context)  { //学生注册
	username,password,number,maxcredit,res:=view.Register(c)
    if len(number)!=10 {
    	c.JSON(http.StatusOK,gin.H{
    		"code":200,
    		"msg": "学号不为10位",
		})
	}
	if len(password)<6 {
		c.JSON(http.StatusOK,gin.H{
			"code": 200,
			"msg":  "密码不能少于6位",
		})
		return
	}
	if username==""{
		c.JSON(http.StatusOK,gin.H{
			"code": 200,
			"msg": "用户名不能为空",
		})
		return
	}
	if maxcredit<=0 {
		c.JSON(http.StatusOK,gin.H{
			"code": 200,
			"msg": "最大学分不能小于1",
		})
	}
	//检验Models层的数据库操作是否成功
	if res {
		c.JSON(http.StatusOK,gin.H{
			"code": 200,
			"msg": "注册成功",
			"username": username,
			"password": password,
			"number": number,
			"maxcredit": maxcredit,
		})
	} else {
		c.JSON(http.StatusNotFound,gin.H{
			"code" : 404,
			"msg": "注册失败",
		})
	}

}
func Login(c *gin.Context)  { //学生登陆
	 username,password,number,res:=view.Login(c)
	 //参数验证
	if len(number)!=10 {
		c.JSON(http.StatusOK,gin.H{
			"code":200,
			"msg": "学号不为10位",
		})
	}
	if len(password)<6 {
		c.JSON(http.StatusOK,gin.H{
			"code": 200,
			"msg":  "密码不能少于6位",
		})
		return
	}
	if res {
		c.SetCookie("username",username,3600,"/","localhost",false,true)
		c.SetCookie("password",password,3600,"/","localhost",false,true)
		c.JSON(http.StatusOK,gin.H{
			"code" : 200,
			"msg":  "登陆成功",
			"username": username,
			"password": password,
			"telephone": number,
		})
	}  else {
		c.JSON(http.StatusNotFound,gin.H{
			"code": 404,
			"msg": "登陆失败",
		})
	}
}
func Update(c *gin.Context)  { //学生改密码
	//参数验证
   OldPassword,err:=c.Cookie("password")
   if err!=nil {
   	 fmt.Printf("failed to get cookie(password),err:%v\n",err)
   	 return
   }
   username,password,number,res:=view.Update(c)
   if OldPassword==password {
   	   c.JSON(http.StatusOK,gin.H{
   	   	   "code": 200,
   	   	   "msg": "新密码不能与原密码一致",
	   })
   	   return
   }
   if res {
   	   c.JSON(http.StatusOK,gin.H{
   	   	"code":200,
   	   	"msg": "修改密码成功",
   	   	"username": username,
   	   	"newpassword": password,
   	   	"number": number,
	   })
   } else {
   	c.JSON(http.StatusNotFound,gin.H{
   		"code" : 404,
   		"msg": "修改密码失败",
	})
   }
}
func Delete(c *gin.Context)  { //删除学生
	username,err:=c.Cookie("username")
	if err!=nil {
		fmt.Printf("failed to get cookie(username),err:%v\n")
		return
	}
	res:=view.Delete(c)
	if !res{
		c.JSON(http.StatusNotFound,gin.H{
			"code" : 404,
			"msg": "删除用户失败",
		})
	} else {
		c.JSON(http.StatusOK,gin.H{
			"code" :200,
			"msg": "删除名为:"+username+"的用户成功",
		})
	}
	c.Redirect(307,"/user/login")  //删除之后重定向到登录界面
}