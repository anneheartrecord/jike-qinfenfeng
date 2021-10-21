package main

import (
	"bufio"
	"fmt"
	"golangstudy/jike/work3/protocol"
	"io"
	"net"
)


func GetInfo(conn net.Conn)  {
	defer conn.Close()
	reader:=bufio.NewReader(conn) //读出conn里面的消息
	for{
        msg,err:=protocol.Decode(reader) //解码函数
         if err!=nil {
		fmt.Printf("decode failed,err:%v\n",err)
		return
	}
        if err==io.EOF{
        	return
		}
		fmt.Println("client info:",msg)
	}
}
func main()  {
	listener,err:=net.Listen("tcp","127.0.0.1:8080") //listen参数：协议和端口，返回值listener 结构体和error
	if err!=nil {
		fmt.Printf("failed to listen,err:%v\n",err)
		return
	}
	defer listener.Close() //记得关闭
	for{
		conn,err:=listener.Accept()  //accpet返回一个coon连接和err
		if err!=nil {
			fmt.Printf("accept failed,err:%v\n",err)
			continue
		}
		go GetInfo(conn)   //开启goroutine 持续读取client info
	}


}
