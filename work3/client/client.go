package main

import (
	"encoding/json"
	"fmt"
	"golangstudy/jike/work3/protocol"
	"net"
)

type GetUserReq struct {
	UserID int64 `json:"userId"`
}

func main()  {
	user:=GetUserReq{
		UserID: 10086,
	}
	b,err:=json.Marshal(user) //编码为字节数组，返回值为[]byte
	if err!=nil {
		fmt.Printf("failed to Marshal user,err:%v\n",err)
		return
	}
	conn,err:=net.Dial("tcp","127.0.0.1:8080") //Dial函数，参数是协议和端口，返回值是连接和err
    if err!=nil {
    	fmt.Printf("failed to dial,err:%v\n",err)
    	return
	}
    defer conn.Close() //记得关闭是个好习惯
    for i:=0;i<10;i++ {
		msg,err:=protocol.Encode(b) //把json字节数组传入编码函数
		if err!=nil {
			fmt.Printf("faile to encode,err:%v\n",err)
			return
		}
		conn.Write(msg)
	}    //发送十次信息测试

}
