package timertask

import (
	"emcs-relay-go/api"
	"emcs-relay-go/api/entity"
	"emcs-relay-go/db"
	"emcs-relay-go/utils"
	"fmt"
	"math/rand"
	"time"
)

func RunKeepLive() {
	timer := time.NewTimer(3 * time.Second)
	defer timer.Stop()
	for {
		timer.Reset(3 * time.Second) // 这里复用了 timer
		select {
		case <-timer.C:
			fmt.Println("keeplive .......")
			testData()
			sendHardware()
			sendEvents()
			//sendDeviceTotal()
			//sendTotal()
			senTotal2()
		}
	}
}

func testData() {
	numbers := [3]string{"384605470533333459544638", "1", "2"}
	randIndex := rand.Int()
	print(randIndex)
	db.TotalUpAdd(numbers[randIndex%3])

}

func sendEvents() {
	var event []db.EventLog
	err := db.GetEvents(&event)
	if err != nil {
		utils.Log.Error(err)
	}
	if event != nil {
		event[0].Time = utils.Fmt2HMS(time.Now())
		api.SendMsg(entity.Pack(entity.TYPE_EVENT, event))
	}
}

func sendDeviceTotal() {
	var devices []db.Device
	err := db.TotalPassed(&devices)
	if err != nil {
		utils.Log.Error(err)
	}
	api.SendMsg(entity.Pack(entity.TYPE_DEVICES, devices))
}

func sendTotal() {
	count := int64(0)
	err := db.TotalAllCount(&count)
	if err != nil {
		utils.Log.Error(err)
	}
	api.SendMsg(entity.Pack(entity.TYPE_TOTAL, count))
}

type TotalVO struct {
	Sum         int64              `json:"sum"`
	DeviceTotal []db.DeviceTotalVO `json:"deviceTotals"`
}

func senTotal2() {
	var data []db.DeviceTotalVO
	ymd := utils.Fmt2Day(time.Now().Local())
	err := db.TotalDeviceCountByDay(&data, ymd)
	if err != nil {
		utils.Log.Error(err)
		return
	}
	totalVo := TotalVO{Sum: 0}
	err = db.TotalSumByDay(&totalVo.Sum, ymd)
	if totalVo.Sum == 0 { // Toady, no one passed
		var devices []db.Device
		db.DevicesList(&devices)
	}
	for i := 0; i < len(data); i++ {
		data[i].Proportion = float32(data[i].Sum) / float32(totalVo.Sum) * 100
	}
	if err != nil {
		utils.Log.Error(err)
		return
	}
	totalVo.DeviceTotal = data
	api.SendMsg(entity.Pack(entity.TYPE_TOTAL, totalVo))
}
func sendHardware() {
	cpu := entity.CPU()
	memory := entity.Memory()
	disk := entity.Disk()
	logs := entity.Logs()
	hd := []entity.Hardware{cpu, memory, disk, logs}
	api.SendMsg(entity.Pack(entity.TYPE_HARDWARES, hd))
}
