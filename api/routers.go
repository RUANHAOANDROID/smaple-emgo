package api

func RegisterRouters() {
	Gin.Use(Cors())
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
