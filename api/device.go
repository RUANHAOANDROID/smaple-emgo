package api

import (
	"emcs-relay-go/api/entity"
	"emcs-relay-go/db"
	"emcs-relay-go/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// HandlerDeviceManager path /devices
func HandlerDeviceManager(r *gin.RouterGroup) {
	//r.Use(LogHandler())
	r.POST("/delete", delete())
	r.POST("/list", list())
	r.POST("/add", add())
	r.POST("/update", update())
	r.POST("/openGateTest", openGateTest())
	r.POST("/getEvents", getEvents())
}

func getEvents() gin.HandlerFunc {
	return func(c *gin.Context) {
		//var res map[string]string
		//c.ShouldBindJSON(&res)
		//date := res["date"]
		//pageNo, _ := strconv.Atoi(res["pageNo"])
		//pageSize, _ := strconv.Atoi(res["pageSize"])
		//deviceName := res["deviceName"]
		//print(date, pageNo, pageSize, deviceName)
		count := int64(0)
		var events []db.EventLog
		//err := db.GetEventsPage(&count, &events, date, deviceName, pageNo, pageSize)
		request := entity.GetEventsPage{}
		err := c.ShouldBindJSON(&request)
		if err != nil {
			print(err.Error())
		}
		err = db.GetEventsPage(&count, &events, request.Date, request.DeviceName, request.Type, request.PageNo, request.PageSize)
		response := entity.Page{Count: count, Data: events}
		c.JSON(http.StatusOK, response)
	}
}
func openGateTest() gin.HandlerFunc {
	return func(c *gin.Context) {
		device := db.Device{}
		c.ShouldBindJSON(&device)
		OpenGate(device.Number, device.Ip)
		c.String(http.StatusOK, "ok")
	}
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
		db.DevicesList(&devices)
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
		err = db.AddDevice(&device)
		if err != nil {
			utils.Log.Error(err)
		}
		response := ResponseSuccess(any("ok"))
		fmt.Println(response)
		c.JSON(http.StatusOK, response)
	}
}
func update() gin.HandlerFunc {
	return func(c *gin.Context) {
		device := db.Device{}
		err := c.BindJSON(&device)
		if err != nil {
			utils.Log.Error(err.Error())
		}
		err = db.UpdateDevice(device)
		if err != nil {
			utils.Log.Error(err)
		}
		response := ResponseSuccess(any("ok"))
		fmt.Println(response)
		c.JSON(http.StatusOK, response)
	}
}
