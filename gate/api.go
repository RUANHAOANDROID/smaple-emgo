package gate

import (
	"emcs-relay-go/api"
	"net"
)

func OpenGate(number string, ip string) {
	udpAddr := net.UDPAddr{
		IP: net.ParseIP(ip), Port: 6000,
	}
	bytes := []byte("$F" + number + "111A01000/&\\&$E")
	api.Conn.WriteToUDP(bytes, &udpAddr)
}
