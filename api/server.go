package api

import (
	"emcs-relay-go/configs"
	"emcs-relay-go/utils"
	"github.com/gin-gonic/gin"
	_ "github.com/sirupsen/logrus"
)

var Gin *gin.Engine

func init() {
	if Gin == nil {
		utils.Log.Info("start web server gin init ")
		Gin = ginConfig()
	}
}
func ginConfig() *gin.Engine {
	r := gin.Default()
	Gin = r
	//路由重定向
	r.GET("/", func(c *gin.Context) {
		//跳转到/luyou2对应的路由处理函数
		c.Request.URL.Path = "/index" //把请求的URL修改
		r.HandleContext(c)            //继续后续处理
	})
	r.Static("/index", "./static") //制定static 目录
	//r.LoadHTMLGlob("/static/index.html")
	RegisterRouters()
	return r
}

func Run() {
	trustedProxies := []string{configs.HttpLoopAddr, configs.Localhost}
	Gin.SetTrustedProxies(trustedProxies)
	utils.Log.Info("http" + "://" + configs.HttpLoopAddr + configs.HttpListenPort + "/index")
	utils.Log.Info("http" + "://" + configs.Localhost + configs.HttpListenPort + "/index")
	Gin.Run(configs.HttpListenPort)
}
