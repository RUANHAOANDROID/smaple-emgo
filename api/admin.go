package api

import (
	"emcs-relay-go/db"
	"emcs-relay-go/logger"
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
func getNumToday() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, ResponseSuccess(any(100)))
	}
}
func deviceList() gin.HandlerFunc {
	return func(c *gin.Context) {
		var devices []db.Device
		db.DB.Find(&devices)
		responseSuccess := ResponseSuccess(devices)
		logger.Log.Debug(responseSuccess)
		c.JSON(http.StatusOK, responseSuccess)
	}
}
func getRecentOperateLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, ResponseSuccess(any("保存成功")))
	}
}
func updatePassword() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "保存成功")
	}
}
func addDevice() gin.HandlerFunc {
	return func(c *gin.Context) {
		device := db.Device{}
		err := c.ShouldBind(&device)
		if err != nil {
			ResError(c, "参数错误")
			return
		}
		err = db.AddDevice(device)
		if err != nil {
			ResError(c, err.Error())
			return
		}
		c.JSON(http.StatusOK, ResponseSuccess(any("OK")))
	}
}

func getCurrentConfig() gin.HandlerFunc {
	return func(c *gin.Context) {
		var configs []db.DeviceConfig
		err := db.GetConfig(configs)
		if err != nil {
			ResError(c, "查询失败")
		}
		c.JSON(http.StatusOK, ResponseSuccess(configs))
	}
}
func numUploadTest() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "保存成功")
	}
}

func heartbeat() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, ResponseSuccess(any("")))
	}
}

func checkTicketTest() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "保存成功")
	}
}
func getConfig() gin.HandlerFunc {
	return func(c *gin.Context) {
		request := Request[WebGetConfig]{}
		err := c.BindJSON(&request)
		if err != nil {
			ResError(c, "参数错误")
			return
		}
		wgc := request.Data
		result, err := RequestConfig(wgc.ConfigUrl, wgc.Id, request.Timestamp)
		if err != nil {
			ResError(c, "获取配置失败"+err.Error())
			return
		}
		c.JSON(http.StatusOK, result)
	}
}
func saveConfig() gin.HandlerFunc {
	return func(c *gin.Context) {
		dc := db.DeviceConfig{}
		pErr := c.BindJSON(&dc)
		if pErr != nil {
			ResError(c, pErr.Error())
			return
		}
		dErr := db.SaveConfig(dc)
		if dErr != nil {
			ResError(c, dErr.Error())
			return
		}
		c.JSON(http.StatusOK, ResponseSuccess("ok"))
	}
}
func deleteDevice() gin.HandlerFunc {
	return func(c *gin.Context) {
		device := db.Device{}
		err := c.BindJSON(&device)
		if err != nil {
			ResError(c, err.Error())
			return
		}
		dbError := db.DeleteDevice(device.Id)
		if dbError != nil {
			ResError(c, dbError.Error())
			return
		}
		c.JSON(http.StatusOK, ResponseSuccess("OK"))
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
		err = db.UserExists(requestUser.UserName)
		if err != nil {
			c.JSON(http.StatusOK, ResponseError("用户不存在"))
			return
		}
		err = db.UserVerify(requestUser.UserName, requestUser.Password)
		if err != nil {
			c.JSON(http.StatusOK, ResponseError("密码错误"))
			return
		}
		response := ResponseSuccess(any("ok"))
		fmt.Println(response)
		c.JSON(http.StatusOK, response)
	}
}
