package api

import (
	"emcs-relay-go/configs"
	"emcs-relay-go/logger"
	"github.com/gin-gonic/gin"
	_ "github.com/sirupsen/logrus"
)

var Gin *gin.Engine

func init() {
	if Gin == nil {
		logger.Log.Info("start web server gin init ")
		Gin = ginConfig()
	}
}
func ginConfig() *gin.Engine {
	r := gin.Default()
	Gin = r
	r.Static("/static/js", "./static/static/js")
	r.Static("/static/css", "./static/static/css")
	r.Static("/static/img", "./static/static/img")
	r.Static("/static/fonts", "./static/static/fonts")
	//r.LoadHTMLGlob("/static/index.html")
	return r
}

func Run() {
	trustedProxies := []string{configs.HttpLoopAddr, configs.Localhost}
	Gin.SetTrustedProxies(trustedProxies)
	logger.Log.Info("http" + "://" + configs.HttpLoopAddr + configs.HttpListenPort + "/index")
	logger.Log.Info("http" + "://" + configs.Localhost + configs.HttpListenPort + "/index")
	Gin.Run(configs.HttpListenPort)
}
