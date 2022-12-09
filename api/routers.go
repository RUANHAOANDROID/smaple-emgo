package api

func RegisterRouters() {
	//tg := Gin.Group("/")
	admin := Gin.Group("/admin")
	devices := Gin.Group("/devices")
	updater := Gin.Group("/updater")
	config := Gin.Group("/config")
	hard := Gin.Group("/hardware")
	HandlerAdmin(admin)
	HandlerDeviceManager(devices)
	HandlerVersionManager(updater)
	HandlerConfigManager(config)
	HandlerConfigManager(hard)
}
