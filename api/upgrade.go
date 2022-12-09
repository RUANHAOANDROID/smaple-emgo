package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandlerVersionManager(r *gin.RouterGroup) {
	r.POST("/getCurrVersion", getCurrVersion())
	r.POST("/getLatestVersion", getLatestVersion())
	r.POST("/download", download())
	r.POST("/getPercent", getPercent())
	r.POST("/quitAndReboot", quitAndReboot())
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
