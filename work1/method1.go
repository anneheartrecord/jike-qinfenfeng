package main

import (
	"fmt"
)

func main()  {   //第一种方法
	ch1,ch2,ch3,ch4:=make(chan string),make(chan string),make(chan  string),make(chan string)
	for i:=0;i<10;i++{  //goroutine里面往channel里面塞东西 在main里面读出来
      go func() {
      	ch1<-"张三"
	  }()
      go func() {
      	ch2<-"李四"
	  }()
      go func() {
      	ch3<-"王五"
	  }()
      go func() {
      	ch4<-"赵六"
	  }()
	}
   for i:=0;i<10;i++{
	   fmt.Println(<-ch1)
	   fmt.Println(<-ch2)
	   fmt.Println(<-ch3)
	   fmt.Println(<-ch4)
   }
}
