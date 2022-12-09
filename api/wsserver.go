package api

import (
	"emcs-relay-go/timertask"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
	return true
}} // use default options

func HandlerHoldWS(r *gin.RouterGroup) {
	r.GET("/flow", func(c *gin.Context) {
		wsConn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Print("upgrade:", err)
			return
		}
		go timertask.RunKeepLive(wsConn)
		defer wsConn.Close()
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
