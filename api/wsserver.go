package api

import (
	"emcs-relay-go/api/entity"
	"emcs-relay-go/utils"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
	return true
}} // use default options
var wsc *websocket.Conn

func SendMsg(msg entity.Msg[interface{}]) {
	if wsc == nil {
		utils.Log.Error("ws server is nil")
		return
	}
	wsc.WriteJSON(msg)
}
func HandlerHoldWS(r *gin.RouterGroup) {
	r.GET("/", func(c *gin.Context) {
		wsConn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Print("upgrade:", err)
			return
		}
		wsc = wsConn
		defer wsConn.Close()
		wsConn.SetCloseHandler(func(code int, text string) error {
			print("SetCloseHandler")
			wsConn.Close()
			return err
		})

		for {
			mt, message, err := wsConn.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				break
			}

			log.Printf("recv: %s", message)
			err = wsConn.WriteMessage(mt, message)
			if err != nil {
				log.Println("write:", err)
				break
			}
		}
	})
}
