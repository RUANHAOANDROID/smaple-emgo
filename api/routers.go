package api

func RegisterRouters() {
	//tg := Gin.Group("/")
	ag := Gin.Group("/admin")
	vmg := Gin.Group("/updater")
	HandlerAdmin(ag)
	HandlerVersionManager(vmg)
}
