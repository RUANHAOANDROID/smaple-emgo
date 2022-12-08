package api

func RegisterRouters() {
	//tg := Gin.Group("/")
	ag := Gin.Group("/admin")
	dg := Gin.Group("/devices")
	vmg := Gin.Group("/updater")
	cf := Gin.Group("/config")
	HandlerAdmin(ag)
	HandlerDeviceManager(dg)
	HandlerVersionManager(vmg)
	HandlerVersionManager(cf)
}
