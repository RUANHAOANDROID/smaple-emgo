package main

import (
	"emcs-relay-go/icbc"
	"fmt"
	"net"
)

func StartUDPServer(address string) {
	// 创建 服务器 UDP 地址结构。指定 IP + port
	udpAddr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		fmt.Println("ResolveUDPAddr err:", err)
		return
	}
	// 监听 客户端连接
	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println("net.ListenUDP err:", err)
		return
	}

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
	fmt.Println(msg)
	bytes := []byte(icbc.CheckTicket())
	//var str =[]byte(" $F12345678F$")
	conn.WriteToUDP(bytes, clientAddress) // 简单回写数据给客户端
}
