package controller

import (
	"github.com/gin-gonic/gin"
	"golangstudy/jike/work5/view"
	"net/http"
)

func CreateCourse(c *gin.Context)  {  //创建课程
   CourseName,Credit,MaxNumber,res:=view.CreateCourse(c)
   if CourseName=="" {
   	c.JSON(http.StatusOK,gin.H{
   		"code": 200,
   		"msg": "课程名不能为空",
	})
   	return
   }
   if Credit<=0 {
		  c.JSON(http.StatusOK,gin.H{
			"code":200 ,
			"msg": "学分不能小于1",
		  })
   	  return
   }
   if MaxNumber<=0 {
   	  c.JSON(http.StatusOK,gin.H{
   	  	"code":200,
   	  	"msg": "最大选课人数不能小于1",
	  })
   	  return
   }
   if !res {
   	 c.JSON(http.StatusNotFound,gin.H{
   	 	"code":404,
   	 	"msg": "增加课程失败",
	 })
   	 return
   }
   c.JSON(http.StatusOK,gin.H{
   	"code":200,
   	"msg": "增加课程成功",
   	"coursename": CourseName,
   	"credit": Credit,
   	"maxnumber":MaxNumber,
   })
}
func ShowCourse(c *gin.Context)  { //查询课程
	CourseName,Credit,MaxNumber,res:=view.ShowCourse(c)
	if !res {
		c.JSON(http.StatusNotFound,gin.H{
			"code": 404,
			"msg": "查询失败",
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"msg": "查询课程成功",
		"coursename": CourseName,
		"credit": Credit,
		"maxnumber":MaxNumber,
	})
}
func UpdateCourse(c *gin.Context)  { //修改课程
   CourseName,Credit,MaxNumber,res:=view.UpdateCourse(c)
   if !res{
   	 c.JSON(http.StatusNotFound,gin.H{
   	 	"code": 404,
   	 	"msg": "修改课程信息失败",
	 })
   	 return
   }
   c.JSON(http.StatusOK,gin.H{
   	"code": 200,
   	"msg": "修改课程成功",
   	"coursename": CourseName,
   	"newcredit": Credit,
   	"newmaxnumber": MaxNumber,
   })
}
func DeleteCourse(c *gin.Context)  { //删除课程
   CourseName,res:=view.DeleteCourse(c)
    if !res {
    	c.JSON(http.StatusNotFound,gin.H{
    		"code": 404,
    		"msg": "删除课程信息失败",
		})
    	return
	}
	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"msg": "删除名为:"+CourseName+"的课程成功",
	})
}