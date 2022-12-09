package timertask

import (
	"fmt"
	"time"
)

func RunKeepLive() {
	timer := time.NewTimer(10 * time.Second)
	for {
		timer.Reset(10 * time.Second) // 这里复用了 timer
		select {
		case <-timer.C:
			fmt.Println("keeplive .......")
		}
	}
}
