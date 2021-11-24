package gocache

import (
	"fmt"
	"net"
	"strconv"
)

type Node struct {   //结点  因为是哈希的双向链表 所以字段分别为 k  v pre 和next
	Key string
	Value  string
	pre *Node
	next *Node
}
type  LRUCache struct {   //一个lru数据库
	limit int
	HashMap map[string]*Node
}
type Server interface {
	Run()
}

func (l* LRUCache) Run()  {
	fmt.Println("start run")
}

func NewServer(port string ,limit int) Server  {  //NewServer 作用:创建一个服务端 并且返回一个可以储存指定数量的lru数据库
	listen,err:=net.Listen("tcp",port)
	if err!=nil {
		fmt.Println("failed to dial,err:",err)
		return nil
	}
	defer listen.Close()
	conn,err:=listen.Accept()
	if err!=nil {
		fmt.Println("failed to listen,err:",err)
		return nil
	}
	_,err=conn.Write([]byte(strconv.Itoa(limit)))
	if err!=nil {
		fmt.Println("failed to write,err:",err)
		return nil
	}
	lruCache:=&LRUCache{limit:limit}   //限制数量为limit
	lruCache.HashMap=make(map[string]*Node,limit)

	return lruCache
}

