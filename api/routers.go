package api

import (
	"github.com/gin-contrib/static"
)

func RegisterRouters() {
	Gin.Use(Cors())
	Gin.Use(static.Serve("/", static.LocalFile("./flutter", true))) //使用 contrib 避免与get path "/"a 冲突
	//tg := Gin.Group("/")
	admin := Gin.Group("/admin")
	devices := Gin.Group("/devices")
	updater := Gin.Group("/updater")
	config := Gin.Group("/config")
	hard := Gin.Group("/hardware")
	flow := Gin.Group("/ws")
	HandlerAdmin(admin)
	HandlerDeviceManager(devices)
	HandlerVersionManager(updater)
	HandlerConfigManager(config)
	HandlerConfigManager(hard)
	HandlerHoldWS(flow)
}
