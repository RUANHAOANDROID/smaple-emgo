package main

import (
	"emcs-relay-go/api"
	"emcs-relay-go/configs"
	"emcs-relay-go/desktop"
	"emcs-relay-go/timertask"
	"github.com/gin-gonic/gin"
)

// start1234
func main() {
	api.RunUDP(configs.UDPListenPort)
	gin.SetMode(gin.DebugMode)
	go timertask.RunDeviceStatic()
	if configs.EnableDesktop {
		go api.Run()
		desktop.Run()
	} else {
		api.Run()
	}
}
