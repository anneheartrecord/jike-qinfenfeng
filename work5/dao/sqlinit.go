package dao

import (
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
)
 //dao层 实现数据库的初始化
var DB *sql.DB
func MySQLInit() *sql.DB {
	dsn:="root:cxs20030416@tcp(127.0.0.1:3306)/qinfenfeng"  //连接数据库
	db,err:=sql.Open("mysql",dsn)
	if err!=nil {
		fmt.Printf("failed to connect database,err:%v\n",err)
		return nil
	}

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)   //设置最大连接数和闲置连接数
	err =db.Ping()   //Open只能验证参数格式是否正确,ping能检查是否连接上数据库
	if err!=nil {
		fmt.Printf("failed to connect database,err:%v\n",err)
		return nil
	}
	DB=db
	return DB
}
