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

func RunDeviceStatic() {
	timer := time.NewTimer(10 * time.Second)
	defer timer.Stop()
	for {
		timer.Reset(10 * time.Second) // 这里复用了 timer
		select {
		case <-timer.C:
			fmt.Println("timer task:hardware static info .......")
			sendHardware()
			//testData()
			//sendEvents()
			//sendDeviceTotal()
			//sendTotal()
			//senTotal2()
		}
	}
}

func testData() {
	numbers := [3]string{"384605470533333459544638", "1", "8"}
	randIndex := rand.Int()
	print(randIndex)
	db.TotalUpAdd(numbers[randIndex%3])

}

func sendEvents() {
	var events []db.EventLog
	err := db.GetEvents(&events)
	count, err := db.GetEventsByDay(&events, "2023-01-06", "武汉测试闸", 1, 10)
	if err != nil {
		utils.Log.Error(err)
	}
	print(count)
	if events != nil {
		events[0].Time = utils.Fmt2HMS(time.Now())
		api.SendMsg(entity.Pack(entity.TYPE_EVENT, events))
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
		api.SendMsg(entity.Pack(entity.TYPE_TOTAL, totalVo))
		return
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
