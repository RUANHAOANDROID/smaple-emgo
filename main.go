package main

import (
	_ "embed"
	"emcs-relay-go/api"
	"emcs-relay-go/configs"
	"emcs-relay-go/desktop"
	"emcs-relay-go/static"
	"emcs-relay-go/udp"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	api.Gin.StaticFS("/index", http.FS(static.Static))
	udp.Run(configs.UDPListenAddr)
	gin.SetMode(gin.ReleaseMode)
	if configs.EnableDesktop {
		go func() {
			api.Run()
		}()
		desktop.Run()
	} else {
		api.Run()
	}
	//api.Run()
}
