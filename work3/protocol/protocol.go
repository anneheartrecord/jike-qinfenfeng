package protocol

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
)

func Encode(msg [] byte ) ([]byte,error) { //编码函数
	//因为UserID是int64 即八个字节 所以我们规定前八个字节为包头
	 len:=int64(len(msg))
	 buf:=new(bytes.Buffer)  //开辟缓冲区
	 err:=binary.Write(buf,binary.LittleEndian,len) //以小端方式将我们的包头写入buf
	 if err!=nil { //错误处理
	 	fmt.Printf("failed to write,err:%v\n",err)
	 	return nil ,err
	 }
	 err=binary.Write(buf,binary.LittleEndian,msg) //以小端方式把消息写入buf
	 if err!=nil {
		 fmt.Printf("failed to write,err:%v\n",err)
		 return nil ,err
	 }
	 return buf.Bytes(),nil //如果没毛病就返回buf里的字节数组
}
//解码函数
func Decode(reader *bufio.Reader) (string,error) {
	var len int64
	lenByte,_:=reader.Peek(8) //读取前8个字节的数据
	lenBuff:=bytes.NewBuffer(lenByte)
	err:=binary.Read(lenBuff,binary.LittleEndian,&len) //以小端方式读，要取地址
	if err!=nil {
		fmt.Printf("failed to read,err:%v\n",err)
		return "",err
	}
	//判断缓冲中可读的字节数
	if int64(reader.Buffered())<len+8 {
		return "",err
	}
	//读取数据
    msg:=make([]byte,int(8+len))
    _,err=reader.Read(msg)
    if err!=nil {
    	fmt.Printf("failed to read,err:%v\n",err)
    	return "",err
	}
	return string(msg[8:]),nil
}