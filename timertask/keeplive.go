package timertask

import (
	"emcs-relay-go/api"
	"emcs-relay-go/api/entity"
	"fmt"
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
			cpu := entity.CPU()
			api.SendMsg(entity.Pack(entity.TYPE_HARDWARE, cpu))

			memory := entity.Memory()
			api.SendMsg(entity.Pack(entity.TYPE_HARDWARE, memory))

			disk := entity.Disk()
			api.SendMsg(entity.Pack(entity.TYPE_HARDWARE, disk))

			logs := entity.Logs()
			api.SendMsg(entity.Pack(entity.TYPE_LOG, logs))

		}
	}
}
