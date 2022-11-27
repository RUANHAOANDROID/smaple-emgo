package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	RegisterRouters()
}
func RegisterRouters() {
	tg := Gin.Group("/")
	ag := Gin.Group("/admin")
	vmg := Gin.Group("/updater")
	handlerHtml(tg)
	HandlerAdmin(ag)
	HandlerVersionManager(vmg)
}
func handlerHtml(r *gin.RouterGroup) {
	r.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", nil)
	})
}
