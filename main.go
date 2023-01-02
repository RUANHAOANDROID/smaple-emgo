package main

import (
	"emcs-relay-go/api"
	"emcs-relay-go/configs"
	"emcs-relay-go/desktop"
	"emcs-relay-go/static"
	"emcs-relay-go/timertask"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	api.RunUDP(configs.UDPListenAddr)

	gin.SetMode(gin.ReleaseMode)
	api.Gin.StaticFS("/flutter", http.FS(static.FlutterWeb))
	//api.Gin.StaticFS("/index", http.Dir("./web/"))
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
