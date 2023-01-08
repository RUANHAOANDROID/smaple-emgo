package timertask

import (
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
			//testData()
			//sendHardware()
			//sendEvents()
			//sendDeviceTotal()
			//sendTotal()
			//senTotal2()
		}
	}
}
