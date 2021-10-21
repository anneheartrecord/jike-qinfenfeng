package models

import (
	"fmt"
	"golangstudy/jike/work5/dao"
)

type UserInfo struct {  //models层的user实现了对user的CRUD
	Username string `form:"username" json:"username" binding:"required"`
	Password string  `form:"password" json:"password" binding:"required"`
	Number string `form:"number" json:"number" binding:"required"`
	Credit int `form:"credit" json:"credit" `
	MaxCredit int `form:"maxcredit" json:"maxcredit" `
}

func Register(Username,Password,Number string,Maxcredit int) bool {    //插入 因为number是主键，所以不用query number是否存在
	if Username==""||len(Password)<6||len(Number)!=10||Maxcredit<=0 {
		return false
	}
	sqlStr2:="insert into users (username,password,number,maxcredit) value (?,?,?,?)"
	_,err:=dao.DB.Exec(sqlStr2,Username,Password,Number,Maxcredit)
	if err!=nil {
		fmt.Printf("failed to exec ,err:%v\n",err)
		return  false
	}
	return true
}
func Login(Number,Password string,) bool {    //query 数据库中的密码是否与发送请求的密码一致
	if len(Password)<6||len(Number)!=10 {
		return false
	}
	sqlStr := "select password from users where number =? "
    stmt,err:=dao.DB.Query(sqlStr,Number)
    if err!=nil {
    	fmt.Printf("failed to query,err:%v\n",err)
    	return false
	}
	defer stmt.Close()  //要记得关闭连接
	var u UserInfo
    for stmt.Next() {
    	err:=stmt.Scan(&u.Password)
    	if err!=nil {
    		fmt.Printf("failed to scan,err:%v\n",err)
    		return false
		}
	}
	if u.Password==Password{
		return true
	}
	return false
}
func Update(Number,Password string) bool  { //改用户密码
	sqlstr:="Update users set password =? where number = ?"
	_,err:=dao.DB.Exec(sqlstr,Password,Number)
	if err!=nil {
		fmt.Printf("failed to update password ,err:%v\n",err)
		return false
	}
	return true
}
func Delete(Number string) bool {//删除用户
	stmt,err:=dao.DB.Query("select * from users where number=?",Number)
	if err!=nil {
		fmt.Printf("failed to query info,err:%v\n",err)
		return false
	}
	defer stmt.Close()
	if !stmt.Next(){
		return false
	}
	sqlStr:="delete from users where number=?"
	_,err=dao.DB.Exec(sqlStr,Number)
	if err!=nil {
		fmt.Printf("failed to delete user,err:%v\n",err)
		return false
	}
	return true
}