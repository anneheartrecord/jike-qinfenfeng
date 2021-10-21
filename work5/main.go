package main

import (
	"golangstudy/jike/work5/dao"
	"golangstudy/jike/work5/router"
)

func main()  {
	dao.MySQLInit()
	router.Router()
}