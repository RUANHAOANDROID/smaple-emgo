package main

import (
	_ "embed"
	"emcs-relay-go/api"
	"emcs-relay-go/configs"
	"emcs-relay-go/desktop"
	"emcs-relay-go/timertask"
	"github.com/gin-gonic/gin"
)

func main() {
	api.RunUDP(configs.UDPListenAddr)

	gin.SetMode(gin.ReleaseMode)
	//api.Gin.StaticFS("/index", http.FS(static.Static))
	//api.Gin.StaticFS("/web", http.FS(build.Static))
	//go timertask.RunKeepLive()
	if configs.EnableDesktop {
		go api.Run()
		desktop.Run()
	} else {
		go timertask.RunKeepLive()
		api.Run()
	}
}
