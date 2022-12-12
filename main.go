package main

import (
	_ "embed"
	"emcs-relay-go/api"
	"emcs-relay-go/configs"
	"emcs-relay-go/desktop"
	"emcs-relay-go/static"
	"emcs-relay-go/timertask"
	"emcs-relay-go/udp"
	"emcs-relay-go/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	udp.Run(configs.UDPListenAddr)

	gin.SetMode(gin.ReleaseMode)
	api.Gin.StaticFS("/index", http.FS(static.Static))
	go timertask.RunKeepLive()
	if configs.EnableDesktop {
		go api.Run()
		desktop.Run()
	} else {
		udp.OpenGate()
		utils.GetCpuPercent()
		api.Run()
	}
}
