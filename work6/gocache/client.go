package gocache

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
)


var head *Node
var end *Node   //头结点和尾结点要专门拿出来

func Newclient(addr string) Client {
	conn,err:=net.Dial("tcp",addr)
	if err!=nil {
		fmt.Println("failed to listen,err:",err)
		return nil
	}
	reader:=bufio.NewReader(conn)
	var buf [128] byte
	n,err:=reader.Read(buf[:])
	if err!=nil {
		fmt.Println("failed to read,err:",err)
		return nil
	}
	res:=string(buf[:n])

	limit,err:=strconv.ParseInt(res,0,10)
	if err!=nil {
		fmt.Println("failed to parse int ")
		return nil
	}
    l:=&LRUCache{limit:int(limit)}
	l.HashMap=make(map[string]*Node,int(limit))
    return l
}
func (l *LRUCache) refreshNode(node *Node)  {  //刷新节点
	if (node==end) {
		return
	}
	l.removeNode(node)
	l.addNode(node)
}
type Client interface {
	//Set() 如果没有就新建,有就更新, 失败返回错误
	Set(key string, value string) error
	//Get() 获取缓存的内容, 失败返回错误
	Get(key string) (string,error)
	//Delete() 删除缓存的内容, 失败返回错误
	Delete(key string ) error
}

func (l *LRUCache)  removeNode(node *Node) string { //删除结点
	if (node==end) {
		end=end.pre
	} else if (node==head) {
		head=head.next
	} else {
		node.pre.next=node.next
		node.next.pre=node.pre
	}
	return node.Key
}
func (l *LRUCache) addNode(node *Node){  //set的时候如果没有 则要新建
	if (end != nil){ //当已经到尾结点了  那么就加一个结点
		end.next = node
		node.pre = end
		node.next = nil
	}
	end = node
	if (head == nil) {
		head = node
	}
}

func (l *LRUCache)Set(key , value string) error {  //设置
	if v,ok := l.HashMap[key];!ok{   //查是否存在
		if(len(l.HashMap) >= l.limit){ //判断是否满了
			oldKey := l.removeNode(head)  // 满了就删掉头
			delete(l.HashMap, oldKey)
		}
		node := Node{Key:key, Value:value}
		l.addNode(&node)    //加一个新的结点
		l.HashMap[key] = &node
	}else {     //如果存在的话 更新就好了
		v.Value = value
		l.refreshNode(v)
	}
	return nil
}
func (l *LRUCache) Get(key string) (string,error) { //Get
	if v,ok:=l.HashMap[key];ok{
		l.refreshNode(v)
		return v.Value,nil
	} else {
		return "",nil
	}
}
func (l *LRUCache)  Delete(key string) error  {
	if _,ok:=l.HashMap[key];ok {
		oldKey:=l.removeNode(head)
		delete(l.HashMap,oldKey)
		return nil
	}
	return nil
}
