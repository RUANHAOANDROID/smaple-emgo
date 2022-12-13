package udp

import (
	"emcs-relay-go/api"
	"emcs-relay-go/api/entity"
	"emcs-relay-go/db"
	"emcs-relay-go/utils"
	"fmt"
	"net"
)

func Run(address string) {
	// 创建 服务器 UDP 地址结构。指定 IP + port
	udpAddr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		utils.Log.Error(err)
		return
	}
	// 监听 客户端连接
	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		utils.Log.Error(err)
		return
	}
	utils.Log.Info("UDP Listen", udpAddr)
	go func() {
		defer conn.Close()
		for {
			handelUDP(conn)
		}
	}()
}

func handelUDP(conn *net.UDPConn) {
	buf := make([]byte, 1024)
	len, clientAddress, err := conn.ReadFromUDP(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	msg := string(buf[:len])
	//logger.Log.Info(msg)
	fmt.Println(msg)
	bytes := []byte(api.CheckTicket())
	//var str =[]byte(" $F12345678F$")
	go func() {
		log := db.EventLog{Tag: "刷票", Content: msg, Time: utils.NowTimeStr()}
		db.AddEvent(&log)
		api.SendMsg(entity.Pack(entity.TYPE_LOG, log))
	}()
	go checkTicket(conn)
	conn.WriteToUDP(bytes, clientAddress) // 简单回写数据给客户端
}
func checkTicket(conn *net.UDPConn) {

}
