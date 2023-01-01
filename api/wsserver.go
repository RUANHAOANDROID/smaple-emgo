package api

import (
	"emcs-relay-go/api/entity"
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
		defer wsConn.Close()
		wsConn.SetCloseHandler(func(code int, text string) error {
			print("SetCloseHandler")
			wsConn.Close()
			return err
		})
		for {
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
