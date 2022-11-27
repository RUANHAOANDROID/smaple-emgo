package main

import (
	"emcs-relay-go/api"
	"github.com/gin-gonic/gin"
	_ "github.com/sirupsen/logrus"
	"net/http"
)

func StartRouter() *gin.Engine {
	r := gin.Default()
	//把静态资源static/js目录挂在到相对路径js
	r.Static("/static/js", "./static/static/js")
	r.Static("/static/css", "./static/static/css")
	r.Static("/static/img", "./static/static/img")
	r.Static("/static/fonts", "./static/static/fonts")
	r.Use()
	//挂载templates目录下html资源
	r.LoadHTMLGlob("static/*.html")
	topGroup := r.Group("/")        //页面组
	adminGroup := r.Group("/admin") //api组
	handlerHtml(topGroup)           //处理html组
	api.HandlerAdmin(adminGroup)    //处理API配置组
	return r
}
func handlerHtml(r *gin.RouterGroup) {
	r.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", nil)
	})
}
