package api

import (
	"emcs-relay-go/api/entity"
	"emcs-relay-go/api/icbc"
	"emcs-relay-go/db"
	"emcs-relay-go/utils"
	"fmt"
	"net"
)

var Conn *net.UDPConn

func RunUDP(address string) {
	// 创建 服务器 UDP 地址结构。指定 IP + port
	udpAddr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		utils.Log.Error(err)
		return
	}
	// 监听 客户端连接
	conn, err := net.ListenUDP("udp", udpAddr)
	Conn = conn
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
	utils.Log.Info(msg)
	isCheck, err := isCheck(msg)
	if err != nil {
		utils.Log.Info(err.Error())
	}
	if isCheck {
		checkTicket(nil, nil, msg)
	}

	isPassed, err := isPassed(msg)
	if err != nil {
		utils.Log.Info(err.Error())
	}
	if isPassed {
		utils.Log.Info(msg)
	}
	bytes := []byte(icbc.CheckTicket())
	//var str =[]byte(" $F12345678F$")
	go func() {
		log := db.EventLog{Tag: "刷票", Content: msg, Time: utils.NowTimeStr()}
		db.AddEvent(&log)
		SendMsg(entity.Pack(entity.TYPE_LOG, log))
	}()
	checkTicket(conn, clientAddress, msg)
	conn.WriteToUDP(bytes, clientAddress) // 简单回写数据给客户端
}
func OpenGate(number string, ip string) {
	udpAddr := net.UDPAddr{
		IP: net.ParseIP(ip), Port: 60066,
	}
	cmd := "$F" + number + "111A01000/&\\&$E"
	bytes := []byte(cmd)
	utils.Log.Info("编号：", number)
	utils.Log.Info("ip：", ip)
	utils.Log.Info("报文：", cmd)
	Conn.WriteToUDP(bytes, &udpAddr)
}
