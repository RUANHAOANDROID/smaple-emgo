package api

import (
	"emcs-relay-go/flutter"
	"net/http"
)

func RegisterRouters() {
	Gin.Use(Cors())
	//Gin.Use(static.Serve("/", static.LocalFile("./flutter/", true))) //使用 contrib 避免与get path "/"a 冲突
	//Gin.Use(static.Serve("/", static.LocalFile("./flutter/", true))) //使用 contrib 避免与get path "/"a 冲突
	Gin.StaticFS("/gateflow/", http.FS(flutter.Static))
	//tg := Gin.Group("/")
	admin := Gin.Group("/admin")
	devices := Gin.Group("/devices")
	updater := Gin.Group("/updater")
	config := Gin.Group("/config")
	flow := Gin.Group("/ws")
	HandlerAdmin(admin)
	HandlerDeviceManager(devices)
	HandlerVersionManager(updater)
	HandlerConfigManager(config)
	HandlerHoldWS(flow)
}
