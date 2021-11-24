package main

import "golangstudy/jike/work6/gocache"

func main()  {
     s:=gocache.NewServer("127.0.0.1:8080",5)
     s.Run()
}
