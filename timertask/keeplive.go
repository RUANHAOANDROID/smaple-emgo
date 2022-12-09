package timertask

import (
	"emcs-relay-go/api/entity"
	"fmt"
	"github.com/gorilla/websocket"
	"time"
)

func RunKeepLive(conn *websocket.Conn) {
	timer := time.NewTimer(3 * time.Second)
	for {
		timer.Reset(3 * time.Second) // 这里复用了 timer
		select {
		case <-timer.C:
			fmt.Println("keeplive .......")
			conn.WriteJSON(entity.CPU())
			conn.WriteJSON(entity.Memory())
			conn.WriteJSON(entity.Disk())
		}
	}
}
