package api

import (
	"emcs-relay-go/db"
	"emcs-relay-go/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandlerDeviceManager(r *gin.RouterGroup) {
	r.Use(LogHandler())
	r.POST("/delete", delete())
	r.POST("/list", list())
	r.POST("/add", add())
}
func delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		var id uint
		err := c.Bind(&id)
		if err != nil {
			utils.Log.Error(err)
		}
		//id, err := strconv.Atoi(string(json["id"]))
		err = db.DeleteDevice(id)
		if err != nil {
			utils.Log.Error(err)
		}
		response := ResponseSuccess(any("ok"))
		fmt.Println(response)
		c.JSON(http.StatusOK, response)
	}
}
func list() gin.HandlerFunc {
	return func(c *gin.Context) {
		var devices []db.Device
		db.DB.Find(&devices)
		responseSuccess := ResponseSuccess(devices)
		utils.Log.Debug(responseSuccess)
		c.JSON(http.StatusOK, responseSuccess)
	}
}
func add() gin.HandlerFunc {
	return func(c *gin.Context) {
		device := db.Device{}
		err := c.BindJSON(&device)
		if err != nil {
			utils.Log.Error(err.Error())
		}
		//deviceId := json["deviceId"].(string)
		//deviceNo := json["deviceNo"].(string)
		//deviceIp := json["deviceIp"].(string)
		//serialNumber := json["serialNumber"].(string)
		//deviceVersion := json["deviceVersion"].(string)
		//id, err := strconv.Atoi(string(json["id"]))
		//device := db.Device{
		//	DeviceId:      deviceId,
		//	DeviceNo:      deviceNo,
		//	DeviceIp:      deviceIp,
		//	SerialNumber:  serialNumber,
		//	DeviceVersion: deviceVersion,
		//}
		err = db.AddDevice(device)
		if err != nil {
			utils.Log.Error(err)
		}
		response := ResponseSuccess(any("ok"))
		fmt.Println(response)
		c.JSON(http.StatusOK, response)
	}
}
