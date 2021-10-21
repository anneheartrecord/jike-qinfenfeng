## 项目目录：

**|--**work3

​     **|--**README.md

​     **|--**client

​           **|--**client.go

​           **|--**fakeserver.go

​     **|--**protocol

​          **|--**protocol.go

​     **|--**server

​          **|--**server.go

​          **|--**fakeclient.go

##     说明

#####  client目录

1.client.go是模拟客户端 向server发送10个值为10086的user结构体作为测试，具体流程先用内置json包对数据进行编码，使其变为字节数组，然后通过net.Dial与server建立连接，最后再使用protocol目录的Encode函数对信息进行加密

2.fakeserver.go是模拟服务端，通过net.Listen与fakeclient建立连接，通过bufio.NewReader和protocol的Decode函数对信息进行解码

##### protocol目录：

protocol.go里面封装了Encode和Decode两个函数，因为UserID是int64类型，所以Encode函数选择规定前八个字节为包头，以小端方式把数据写进buf，进行封包处理，而Decode函数读取包头和真实数据，最终返回真实数据，进行拆包处理.

##### server目录：

1.server.go是模拟服务端，读取client发送的信息，具体流程：通过net.Listen与client建立连接，通过bufio.NewReader和protocol的Decode函数对信息进行解码

2.fakeclient.go是模拟客户端 向server发送10个值为10086，"xiaohuang"的user结构体作为测试，具体流程先用内置json包对数据进行编码，使其变为字节数组，然后通过net.Dial与server建立连接，最后再使用protocol目录的Encode函数对信息进行加密

##### 测试截图：

![img](file:///D:\QQ\2523286318\Image\C2C\S8ELOE58D]B28JZ6NQ~5XVV.png)

![QQ图片20211021234732](C:/Users/86188/Desktop/QQ%E5%9B%BE%E7%89%8720211021234732.png)

