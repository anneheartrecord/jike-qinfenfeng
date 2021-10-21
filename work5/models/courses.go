package models

import (
	"fmt"
	"golangstudy/jike/work5/dao"

)

type CourseInfo struct {
	CourseName string `form:"coursename" json:"coursename" binding:"required"`
	Credit int `form:"credit" json:"credit" binding:"required"`
	NowNumber int `form:"nownumber" json:"nownumber" `
	MaxNumber int `form:"maxnumber" json:"maxnumber" binding:"required"`
}
func CreateCourse(CourseName string,Credit,MaxNumber int ) bool { //创建课程
	if CourseName==""||Credit<=0||MaxNumber<=0{
		return false
	}
    _,err:=dao.DB.Exec("insert into courses (coursename,credit,maxnumber) values (?,?,?)",CourseName,Credit,MaxNumber)
   if err!=nil {
   	fmt.Printf("failed to insert into the courses,err:%v\n",err)
   	return false
   } else {
   	  return true
   }


}
func ShowCourse(CourseName string) (string,int,int,bool)  { //查询课程
   stmt,err:=dao.DB.Query("select * from courses where coursename=?",CourseName)
   if err!=nil {
   	  fmt.Printf("failed to query the courses,err:%v\n",err)
   	  return "",-1,-1,false
   }
   defer stmt.Close()
   var course CourseInfo
   if !stmt.Next(){
   	 err:= stmt.Scan(&course.CourseName,&course.Credit,&course.MaxNumber)
   	 if err!=nil {
   	 	fmt.Printf("failed to scan course,err:%v\n",err)
   	 	return "",-1,-1,false
	 }
   }
   return course.CourseName, course.Credit,course.MaxNumber,true
}
func UpdateCourse(CourseName string,Credit,MaxNumber int) bool { //更新课程学分
	stmt,err:=dao.DB.Query("select * from courses where coursename=?",CourseName)
	if err!=nil {
		fmt.Printf("failed to query the courses,err:%v\n", err)
		return false
	}
	defer stmt.Close()
	if !stmt.Next(){
		return false
	}
    _,err=dao.DB.Exec("update courses set credit =? where coursename=?",Credit,CourseName)
    if err!=nil {
    	fmt.Printf("failed to update,err:%v\n",err)
    	return false
	}
	_,err=dao.DB.Exec("update courses set maxnumber=? where coursename=?",MaxNumber,CourseName)
	if err!=nil {
		fmt.Printf("failed to update,err:%v\n",err)
		return false
	}
	return true
}
func DeleteCourse(CourseName string) bool   { //删除课程
	stmt,err:=dao.DB.Query("select * from courses where coursename=?",CourseName)
	if err!=nil {
		fmt.Printf("failed to query the courses,err:%v\n", err)
		return false
	}
	defer stmt.Close()
	if !stmt.Next(){
		return false
	}
    _,err=dao.DB.Exec("delete from courses where coursename=?",CourseName)
    if err!=nil {
    	fmt.Printf("failed to delete course,err:%v\n",err)
    	return false
	}
	return true
}