package api

import (
	"emcs-relay-go/db"
	"emcs-relay-go/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"time"
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
	r.POST("/updateDevice", updateDevice())
}
func updateDevice() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := db.Device{}
		err := c.ShouldBind(&req)
		if err != nil {
			utils.Log.Error(err)
			ResError(c, "参数错误")
			return
		}
		req.UpdateTime = time.Now().Format(utils.TimeFormat)
		err = db.UpdateDevice(req)
		if err != nil {
			ResError(c, err.Error())
			return
		}
		c.JSON(http.StatusOK, ResponseSuccess("OK"))
	}
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
		utils.Log.Debug(responseSuccess)
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
		updUser := UpdateUser{}
		err := c.ShouldBind(&updUser)
		if err != nil {
			ResError(c, "参数错误")
			return
		}
		db.UpdatePwd(updUser.UserName, updUser.Password)
		if err != nil {
			ResError(c, "更改密码失败")
			return
		}
		c.JSON(http.StatusOK, ResponseSuccess("OK"))
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
		configs := db.DeviceConfig{}
		err := db.GetConfig(&configs)
		if err != nil {
			if err.Error() == "record not found" {
				c.JSON(http.StatusOK, ResponseSuccess(err.Error()))
				return
			}
			ResError(c, "查询失败")
			return
		}
		//response := []db.DeviceConfig{
		//	configs,
		//}
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
		sc1 := SaveConfig{}
		err1 := c.ShouldBindBodyWith(&sc1, binding.JSON)
		if err1 == nil {
			utils.Log.Error(err1.Error())
			utils.Log.Info(sc1)
			err := db.SaveConfig(db.DeviceConfig{
				Id:          sc1.Id,
				Buffer:      sc1.Buffer,
				CheckUrl:    sc1.CheckUrl,
				ConfigUrl:   sc1.ConfigUrl,
				DeFalseText: sc1.DeFalseText,
				//DeFalseVoice    :sc1.DeFalseVoice.Error,
				DeTrueText:      sc1.DeTrueText,
				DeTrueVoice:     sc1.TrueVoice1,
				DelayedTime:     sc1.DelayedTime,
				FalseVoice1:     sc1.FalseVoice1,
				FalseVoice2:     sc1.FalseVoice2,
				FalseVoice3:     sc1.FalseVoice3,
				FalseVoice4:     sc1.FalseVoice4,
				HeartbeatTime:   sc1.HeartbeatTime,
				HeartbeatUrl:    sc1.HeartbeatUrl,
				InvalidTime:     sc1.InvalidTime,
				ManufacturerId1: sc1.ManufacturerId1,
				ManufacturerId2: sc1.ManufacturerId2,
				NumUpTime:       sc1.NumUpTime,
				NumUpUrl:        sc1.NumUpUrl,
				TrueVoice1:      sc1.TrueVoice1,
				TrueVoice2:      sc1.TrueVoice2,
				TrueVoice3:      sc1.TrueVoice3,
				TrueVoice4:      sc1.TrueVoice4,
				WriteOffUrl:     sc1.WriteOffUrl,
			})
			if err != nil {
				ResError(c, err.Error())
				return
			}
			c.JSON(http.StatusOK, ResponseSuccess("ok"))
		} else {
			sc2 := db.DeviceConfig{}
			err := c.ShouldBindBodyWith(&sc2, binding.JSON)
			if err != nil {
				ResError(c, err.Error())
				return
			}
			err = db.SaveConfig(sc2)
			if err != nil {
				ResError(c, err.Error())
				return
			}
			c.JSON(http.StatusOK, ResponseSuccess("0k"))
		}
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
