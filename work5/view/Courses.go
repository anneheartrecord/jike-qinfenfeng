package view

import (
	"github.com/gin-gonic/gin"
	"golangstudy/jike/work5/models"
)

func CreateCourse(c *gin.Context) (string,int,int,bool) {
   var course models.CourseInfo
   c.ShouldBind(&course)
   res:=models.CreateCourse(course.CourseName,course.Credit,course.MaxNumber)
   return course.CourseName,course.Credit,course.MaxNumber,res
}
func ShowCourse(c *gin.Context) (string,int,int,bool) {
    CourseName:=c.PostForm("coursename")
    var course models.CourseInfo
    var res bool
    course.CourseName,course.Credit,course.MaxNumber,res=models.ShowCourse(CourseName)
    return course.CourseName,course.Credit,course.MaxNumber,res
}
func UpdateCourse(c *gin.Context) (string,int,int,bool)  {
   var course models.CourseInfo
   c.ShouldBind(&course)
   res:=models.UpdateCourse(course.CourseName,course.Credit,course.MaxNumber)
   return course.CourseName,course.Credit,course.MaxNumber,res
}
func DeleteCourse(c *gin.Context) (string, bool) {
    coursename:=c.PostForm("coursename")
    res:=models.DeleteCourse(coursename)
    return coursename,res
}