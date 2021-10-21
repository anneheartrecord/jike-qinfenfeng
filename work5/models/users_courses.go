package models

import (
	"fmt"
	"golangstudy/jike/work5/dao"
)
const UserMaxCredit = 10
type UsersAndCourses struct {
	Username string `form:"username" json:"username" binding:"required"`
	Coursename string `form:"coursename" json:"coursename" binding:"required"`
	Number string `form:"number" json:"number" binding:"required"`
}

func CreateRecordInfo(Username,Coursename,Number string) bool { //提交选课
   stmt,err:=dao.DB.Query("select coursename from users_courses where number=?",Number) //首先查学生是不是已经选了这个课
   if err!=nil {
   	  fmt.Printf("failed to query tabel users_courses,err:%v\n",err)
   	  return false
   }
   defer stmt.Close()
   var u UsersAndCourses
   for stmt.Next(){
	   err:=stmt.Scan(&u.Coursename)
	   if err!=nil {
	   	fmt.Printf("failed to scan,err:%v\n",err)
	   	return false
	   }
	   if u.Coursename==Coursename {    //遍历学生选的所有课 如果已经选了返回false
		   return false
	   }
   }
	//查一下这门科目的信息
   stmt1,err:=dao.DB.Query("select * from courses where coursename=?",Coursename)
	if err!=nil {
		fmt.Printf("failed to query credit from courses,err:%v\n",err)
		return false
	}
	defer stmt1.Close()
    var c CourseInfo
    if  stmt1.Next(){
   	  stmt1.Scan(&c.CourseName,&c.Credit,&c.NowNumber,&c.MaxNumber)   //写入结构体
   	  fmt.Println(c)
   }
	var u1 UserInfo
   //查一下这名学生的信息
	stmt2,err:=dao.DB.Query("select * from users where number=?",Number)
	if err!=nil {
		fmt.Printf("failed to query credit from users,err:%v\n",err)
		return false
	}
	defer stmt2.Close()
	if  stmt2.Next(){
		stmt2.Scan(&u1.Username,&u1.Number,&u1.Credit,&u1.Password,&u1.MaxCredit)  //写入结构体
		fmt.Println(u1)
	}
	//如果这门科目的学分加上学生现在的学分超过最大学分则返回false
	if u1.Credit+c.Credit>u1.MaxCredit {
		fmt.Println(u1.Credit,c.Credit,u1.MaxCredit)
		return false
	}
	//如果这门科目选课人数已经达到上限
	if c.NowNumber==c.MaxNumber{
		fmt.Println(2)
		return false
	}
	//以上问题都不存在 我们就开始进行更新操作
	//把学生的学分加上科目的学分
	_,err=dao.DB.Exec("update users set credit=credit+? where number=?",c.Credit,Number)
	if err!=nil {
		fmt.Printf("failed to add the credit,err:%v\n",err)
		return false
	}
	//把该课程的选课人数加1
   _,err=dao.DB.Exec("update courses set nownumber=nownumber+1 where coursename=? ",Coursename)
   if err!=nil {
   	  fmt.Printf("failed to add the nownumber,err:%v\n",err)
   	  return false
   }
	//往users_courses表单里面插入数据
	_,err=dao.DB.Exec("insert into users_courses (username,coursename,number) values(?,?,?)",Username,Coursename,Number)
	if err!=nil {
		fmt.Printf("failed to insert into user_courses,err:%v\n",err)
		return false
	}

   return true
}
func ShowStudentInfo(Username string)  ( []string,bool) { //查询选课
	stmt,err:=dao.DB.Query("select coursename from users_courses where username=? ",Username)
	if err!=nil {
		fmt.Printf("failed to query from users_courses,err:%v\n",err)
		return nil ,false
	}
	defer stmt.Close()
	c:=make([]string,10)

	for stmt.Next() {
        i:=0
		stmt.Scan(c[i])
        i++
	}
	return c,true
}
func DelteRecordInfo(Username string) (bool) {  //删除选课
	stmt,err:=dao.DB.Query("select * from users_courses where username=?",Username) //查询是否有这个课程
	defer stmt.Close()
	if !stmt.Next() {
		return false
	}
	_,err=dao.DB.Exec("delete from users_courses where username=?",Username)
	if err!=nil {
		fmt.Printf("failed to exec,err:%v\n",err)
		return false
	}
	return true
}