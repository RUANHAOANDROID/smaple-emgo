package main

import (
	_ "embed"
	"emcs-relay-go/api"
	"emcs-relay-go/configs"
	"emcs-relay-go/desktop"
	"emcs-relay-go/static"
	"emcs-relay-go/udp"
	"emcs-relay-go/utils"
	"emcs-relay-go/ws"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	udp.Run(configs.UDPListenAddr)

	gin.SetMode(gin.ReleaseMode)
	api.Gin.StaticFS("/index", http.FS(static.Static))
	//go timertask.RunKeepLive()
	go ws.Run()
	if configs.EnableDesktop {
		go api.Run()
		desktop.Run()
	} else {
		utils.GetCpuPercent()
		api.Run()
	}
}
