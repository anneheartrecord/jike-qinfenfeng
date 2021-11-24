package main

import (
	"database/sql"
	"fmt"
	"golangstudy/jike/work6/gocache"
	_"github.com/go-sql-driver/mysql"
)
type BlogArtical struct {
	Name string
	Id int
}
var DB *sql.DB
func InitDB() *sql.DB {
	db,err:=sql.Open("mysql","root:cxs20030416@tcp(127.0.0.1:3306)/work6")
	if err!=nil {
		fmt.Println("failed to connect database,err:",err)
		return nil
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	DB=db
	return DB
}
func CreateArtical(c gocache.Client,articalname string)  error {  //增  先增数据库 再增缓存
	sqlStr:="insert into artical(name) values (?)"
	_,err:=DB.Exec(sqlStr,articalname)
	if err!=nil {
		fmt.Println("failed to db exec,err:",err)
		return err
	}
	c.Set(articalname,articalname)
	return nil
}
func QueryArtical(c gocache.Client,articalname string)  {  //查 先查缓存 再查数据库
	res,err:=c.Get(articalname)
	if err!=nil {
		fmt.Println("failed to get,err:",err)
	}
	if res!=""{   //缓存中查到了
		fmt.Println(res)
		return
	} else {  //缓存中没有  在MySQL里面接着查
		sqlStr:="select id from artical where name=?"
		rows,err:=DB.Query(sqlStr,0)
		if err!=nil {
              fmt.Println("failed to query,err:",err)
              return
		}
		defer rows.Close()
		for rows.Next() {
			var a BlogArtical
			err:=rows.Scan(&a.Id)
			if err!=nil {
				fmt.Println("failed to scan,err:",err)
			}
			fmt.Println("Blog Artical id:",a.Id)
			//更新缓存
			c.Set(articalname,articalname)
		}
	}
}

func ChangeArtical(c gocache.Client,id int,articalname string)  { //更新
   sqlStr:="update artical set name=? where id=?"

      //更新MySQL
	_,err:=DB.Exec(sqlStr,articalname,id)
   if err!=nil {
   	  fmt.Println("failed to exec ,err:",err)
   	  return
   }
   //删除缓存
   c.Delete(articalname)
}
func DeleteArtical(c gocache.Client,articalname string)  {//删除
    //先删除MySQL
	sqlStr:="delete from artical where name=?"
	_,err:=DB.Exec(sqlStr,articalname)
	if err!=nil {
		fmt.Println("failed to exec,err:",err)
		return
	}
	//删除缓存
	c.Delete(articalname)
}
func main()  {
	InitDB()
	c:=gocache.Newclient("127.0.0.1:8080")
/*	c.Set("1","11")
	c.Set("2","22")
	s1,_:=c.Get("1")
	s2,_:=c.Get("2")
	fmt.Println(s1)
	fmt.Println(s2)
	c.Delete("1")
	s3,_:=c.Get("1")
    fmt.Println(s3)  */ //  Set Get Delete 的test
    articalname:="pingfandeshijie"

    CreateArtical(c,articalname)
    QueryArtical(c,articalname)
    ChangeArtical(c,2,"abc")
    //DeleteArtical(c,articalname)
}
