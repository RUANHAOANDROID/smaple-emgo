package udp

import (
	"emcs-relay-go/api"
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
			HandelUDP(conn)
		}
	}()
}

func HandelUDP(conn *net.UDPConn) {
	buf := make([]byte, 1024)
	len, clientAddress, err := conn.ReadFromUDP(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	msg := string(buf[:len])
	//logger.Log.Info(msg)
	utils.Log.Info(msg)
	isCheck, err := isCheck(msg)
	if err != nil {
		utils.Log.Info(err.Error())
	}
	if isCheck {
		checkTicket(msg)
	}

	isPass, err := isPassed(msg)
	if err != nil {
		utils.Log.Info(err.Error())
	}
	if isPass {
		utils.Log.Info(msg)
	}
	bytes := []byte(api.CheckTicket())
	//var str =[]byte(" $F12345678F$")
	conn.WriteToUDP(bytes, clientAddress) // 简单回写数据给客户端
}
