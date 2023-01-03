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
	RegisterRouters()
	return r
}

func Run() {
	trustedProxies := []string{configs.HttpLoopAddr, configs.Localhost}
	Gin.SetTrustedProxies(trustedProxies)
	utils.Log.Info("http" + "://" + configs.HttpLoopAddr + configs.HttpListenPort)
	utils.Log.Info("http" + "://" + configs.Localhost + configs.HttpListenPort)
	Gin.Run(configs.HttpListenPort)
}
