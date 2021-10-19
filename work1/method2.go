package main

import (
	"fmt"
	"time"
)

func main()  {
	ch1,ch2,ch3,ch4:=make(chan bool),make(chan bool),make(chan bool),make(chan bool)
	for i:=0;i<10;i++ {          //第二种方法 通过chan阻塞在各个goroutine之间 实现顺序打印
		go func() {
			if <-ch1 {
				fmt.Println("张三")
				ch2 <- true

			}
		}()
	}
		for i := 0; i < 10; i++ {
			go func() {
				if <-ch2 {
					fmt.Println("李四")
					ch3 <- true
				}
			}()
		}

		for i := 0; i < 10; i++ {
			go func() {
				if <-ch3 {
					fmt.Println("王五")
					ch4 <- true
				}
			}()
		}
		for i := 0; i < 10; i++ {
			go func() {
				if <-ch4 {
					fmt.Println("赵六")
					ch1 <- true
				}
			}()
		}
        ch1<-true   //一定要让main协程睡一会 不然main跑完就直接退出了
		time.Sleep(1*time.Millisecond)
	}