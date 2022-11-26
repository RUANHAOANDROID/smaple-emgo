package main

import (
	"emcs-relay-go/db"
)

const (
	enableSign = true
)

func main() {
	InitLog()
	db.Create()
	//配置GIN
	r := StartRouter()
	//启动UDP，注意先开线程避免无限阻塞
	StartUDPServer("192.168.1.121:60000", db.Instance())
	//启动GIN HTTP服务
	r.SetTrustedProxies([]string{"127.0.0.1", "localhost"})
	r.Run(":8088")
}
