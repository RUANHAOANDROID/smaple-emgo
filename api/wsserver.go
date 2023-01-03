package api

import (
	"emcs-relay-go/api/entity"
	"emcs-relay-go/db"
	"emcs-relay-go/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
	return true
}} // use default options
var wsc *websocket.Conn

func SendMsg(msg entity.Msg[interface{}]) {
	if wsc == nil {
		println("ws server is nil")
		return
	}
	wsc.WriteJSON(msg)
}

func HandlerHoldWS(r *gin.RouterGroup) {
	r.GET("/flow", func(c *gin.Context) {
		wsConn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Print("upgrade:", err)
			return
		}
		wsc = wsConn
		defer wsc.Close()
		wsc.SetCloseHandler(func(code int, text string) error {
			println("conn closed")
			wsConn.CloseHandler()
			wsConn.Close()
			err = errors.New(text)
			return err
		})
		sendTotal()
		sendHardware()
		sendEventLogs()
		for {
			if err != nil {
				break
			}
			handlerMessage(wsConn)
		}
	})
}

func handlerMessage(wsConn *websocket.Conn) {
	mt, message, err := wsConn.ReadMessage()
	if err != nil {
		log.Println("read:", err)
		return
	}
	log.Printf("recv: %s", message)

	err = wsConn.WriteMessage(mt, message)
	if err != nil {
		log.Println("write:", err)
		return
	}
}
func sendHardware() {
	cpu := entity.CPU()
	memory := entity.Memory()
	disk := entity.Disk()
	logs := entity.Logs()
	hd := []entity.Hardware{cpu, memory, disk, logs}
	SendMsg(entity.Pack(entity.TYPE_HARDWARES, hd))
}

type TotalVO struct {
	Sum         int64              `json:"sum"`
	DeviceTotal []db.DeviceTotalVO `json:"deviceTotals"`
}

func sendEventLogs() {
	var event []db.EventLog
	err := db.GetEvents(&event)
	if err != nil {
		utils.Log.Error(err)
	}
	if event != nil {
		//event[0].Time = utils.Fmt2HMS(time.Now())
		SendMsg(entity.Pack(entity.TYPE_EVENT, event))
	}
}

// 发送统计初始值
func sendTotal() {
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
		SendMsg(entity.Pack(entity.TYPE_TOTAL, totalVo))
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
	SendMsg(entity.Pack(entity.TYPE_TOTAL, totalVo))
}
