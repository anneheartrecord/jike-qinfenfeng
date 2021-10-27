package main

import (
"bufio"
"bytes"
"encoding/binary"
"fmt"
"io"
"math"
"os"
"strconv"
"time"
)
func oddParity(result byte) []byte  {  //奇校验 把每个字符都变成2进制字符串
	//首先我们要把一个整数变成二进制
	s:=strconv.FormatInt(int64(result),2) //把result转换为2进制数
	var count=0 //计数器:记录二进制数中1的个数
	var i int
	for i=0;i<len(s);i++ {  //遍历得到二进制数中1的个数
		if s[i]=='1' {
			count++
		}
	}
	ret:=[] byte(s)
	//奇校验的定义:如果一串二进制数中1的个数是偶数 奇校验位就是1 1的个数位技术 奇校验位就是0
	if count%2==0 {
		ret=append(ret,'1')
	}  else {
		ret=append(ret,'0')   //校验位
	}
	return ret
}
func check(msg string) (bool) { //进行奇校验的检查
	var count=0
	var ret=0.0
	for i:=0;i<len(msg);i++{
		ret+=float64(msg[i]-48)*math.Pow(2.0,float64(len(msg)-i-2))  //先变成二进制字符串
		if msg[i]=='1' {
			count+=1
		}
	}
	if count%2==0 {  //判断二进制数中1的个数 如果是偶数则失败
		return false
	}
	fmt.Printf("%c",int(ret))
	return true
}
func producer(out chan<- string)  {  //生产者
	var str1 string
	for i:=0;i<5;i++{  //一共可以输入五条任务
		fmt.Scan(&str1)
		out<-str1
		result:=[] byte(str1)
		for i:=0;i<len(result);i++ {
			a:=oddParity(result[i])  //进行奇校验
			msg,_:=Encode(a)
			f,err:=os.OpenFile("./mq1.mq",os.O_CREATE|os.O_APPEND|os.O_WRONLY,0666)  //打开文件
			defer f.Close()
			if err!=nil {
				fmt.Printf("failed to openfile ,err:%v\n",err)
				return
			}
			_,err=f.WriteString(string(msg))  //写东西
			if err!=nil {
				fmt.Printf("failed to writestring,err:%v\n",err)
				return
			}
		}
	}
	close(out)
}
//编码函数
func Encode(msg [] byte)  ([]byte,error) {//因为我们规定了result 是int 64大小的，也就是八个字节
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
		fmt.Println( )
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

	return string(msg[8:]), err
}
func consumer (in<-chan string)  {
	for s:=range in {
		fmt.Println(s)
	}
	for {
		f, err := os.Open("./mq1.mq")
		defer f.Close()
		if err == io.EOF {
			fmt.Println("文件读完了")
			return
		}
		if err != nil {
			fmt.Printf("failed to open the file,err:%v\n", err)
			return
		}
		reader := bufio.NewReader(f)
		for {
			msg, err := Decode(reader)  //先解码
			if !check(msg) {  //再奇校验
				fmt.Println("failed to check")
			}
			if err != nil {
				fmt.Println("failed to read the file")
				return
			}

		}
		time.Sleep(1*time.Second)  //间隔一秒钟扫一下
	}
}

func main()  {
	ch:=make(chan string)
	go producer(ch)
	consumer(ch)
	err:=os.Remove("./mq1.mq")
	if err!=nil {
		fmt.Printf("failed to remove file,err:%v",err)
	}
}


