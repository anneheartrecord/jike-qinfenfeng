## gocache

工具包 里面有一个client.go文件和server.go文件

### server.go

创建服务端

主要功能是把lru缓存的限制发给client，比如最多存储5条信息等等

Server接口只有一个Run()方法，打出提示成功运行

### client.go

创建客户端

1.先从服务端把limit读出来

2.实现了具有Set Get Delete方法的client接口



## servertest.go

运行gocache的server端

## clienttest.go

运行gocache的client端

进行gocache的set get delete方法的test 

cache和数据库的增伤改查

# 一些思考

因为缓存和数据库是两个单独的数据源，很显然它们之间是没有办法保证一致性，或者说绝对的一致性的。

对于不同的场景和不同的需求，大致有这么几种方法：

1.先更新库再删缓存

2.先删缓存再更新库

3.延迟双删

//就查到的资料而言，以上方法没有绝对的好坏，主要还是根据场景选择
