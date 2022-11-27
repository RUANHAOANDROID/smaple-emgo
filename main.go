package main

import (
	"emcs-relay-go/api"
	"emcs-relay-go/udp"
)

func main() {
	//配置GIN
	//r := api.StartRouter()
	//启动UDP，注意先开线程避免无限阻塞
	udp.StartUDPServer("192.168.1.121:60000")
	//启动GIN HTTP服务
	trustedProxies := []string{"127.0.0.1", "localhost"}
	api.BindAddress(trustedProxies)
	api.Run(":8088")
	//api.Gin.Run(":8088")
}
