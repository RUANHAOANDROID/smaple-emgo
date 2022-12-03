package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandlerVersionManager(r *gin.RouterGroup) {
	r.Use(SignHandler())
	r.POST("/getCurrVersion", login())
	r.POST("/getLatestVersion", saveConfig())
	r.POST("/download", saveConfig())
	r.POST("/getPercent", saveConfig())
	r.POST("/quitAndReboot", saveConfig())
}
func getCurrVersion() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "保存成功")
	}
}
func getLatestVersion() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "保存成功")
	}
}
func download() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "保存成功")
	}
}
func getPercent() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "保存成功")
	}
}
func quitAndReboot() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "保存成功")
	}
}
