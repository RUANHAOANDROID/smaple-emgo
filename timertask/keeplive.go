package timertask

import (
	"emcs-relay-go/api"
	"emcs-relay-go/api/entity"
	"emcs-relay-go/db"
	"emcs-relay-go/utils"
	"fmt"
	"time"
)

func RunKeepLive() {
	timer := time.NewTimer(5 * time.Second)
	defer timer.Stop()
	for {
		timer.Reset(5 * time.Second) // 这里复用了 timer
		select {
		case <-timer.C:
			fmt.Println("keeplive .......")
			//sendHardware()
			// event log
			sendEvents()
			//device total
			//total := []db.Device{}
			//db.DevicesList(&total)
			//api.SendMsg(entity.Pack(entity.TYPE_TOTAL, total))
			sendDeviceTotal()
			sendTotal()
		}
	}
}

func sendEvents() {
	event := []db.EventLog{}
	db.GetEvents(&event)
	event[0].Time = utils.Fmt2HMS(time.Now())
	api.SendMsg(entity.Pack(entity.TYPE_EVENT, event))
}

func sendDeviceTotal() {
	devices := []db.Device{}
	db.TotalPassed(&devices)
	api.SendMsg(entity.Pack(entity.TYPE_TOTAL, devices))
}

func sendTotal() {
	count := int64(0)
	db.TotalAllCount(&count)
	api.SendMsg(entity.Pack(entity.TYPE_TOTAL, count))
}

func sendHardware() {
	cpu := entity.CPU()
	memory := entity.Memory()
	disk := entity.Disk()
	logs := entity.Logs()
	hd := []entity.Hardware{cpu, memory, disk, logs}
	api.SendMsg(entity.Pack(entity.TYPE_HARDWARES, hd))
}
