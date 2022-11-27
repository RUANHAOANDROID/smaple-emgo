package api

import (
	"emcs-relay-go/db"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandlerAdmin(r *gin.RouterGroup) {
	r.Use(SignHandler())
	r.POST("/login", login())
	r.POST("/updatePassword", updatePassword())
	r.POST("/getConfig", getConfig())
	r.POST("/saveConfig", saveConfig())
	r.POST("/addDevice", addDevice())
	r.POST("/deleteDevice", deleteDevice())
	r.POST("/deviceList", deviceList())
	r.POST("/getCurrentConfig", getCurrentConfig())
	r.POST("/openGateTest", openGateTest())
	r.POST("/numUploadTest", numUploadTest())
	r.POST("/checkTicketTest", checkTicketTest())
	r.POST("/getRecentOperateLog", getRecentOperateLog())
	r.POST("/heartbeat", heartbeat())
	r.POST("/getNumToday", getNumToday())
}
func updatePassword() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "保存成功")
	}
}
func addDevice() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "保存成功")
	}
}
func getCurrentConfig() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "保存成功")
	}
}
func numUploadTest() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "保存成功")
	}
}
func getRecentOperateLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "保存成功")
	}
}
func heartbeat() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "保存成功")
	}
}
func getNumToday() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "保存成功")
	}
}

func checkTicketTest() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "保存成功")
	}
}
func getConfig() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "保存成功")
	}
}
func saveConfig() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "保存成功")
	}
}
func deleteDevice() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "保存成功")
	}
}
func deviceList() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "保存成功")
	}
}
func openGateTest() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "保存成功")
	}
}

func login() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestUser := LoginRequest{}
		err := c.ShouldBind(&requestUser)

		if err != nil {
			c.JSON(http.StatusOK, ResponseError("参数错误"))
			return
		}
		dbUser := db.SysUser{}
		db.DB.Where(&db.SysUser{
			UserName: requestUser.UserName,
			UserPwd:  requestUser.Password,
		},
		).Take(&dbUser)
		if requestUser.UserName != dbUser.UserName {
			c.JSON(http.StatusOK, ResponseError("用户不存在"))
			return
		}
		response := ResponseSuccess(any("ok"))
		fmt.Println(response)
		c.JSON(http.StatusOK, response)
	}
}
