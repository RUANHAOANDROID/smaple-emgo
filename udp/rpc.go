package udp

import (
	"bytes"
	"emcs-relay-go/utils"
	"net"
)

func checkTicket(conn *net.UDPConn, address *net.UDPAddr, msg string) {
	//        String id = str.substring(2,26);// 主板id 设备序列号
	//        String number = str.substring(26,27);// 门号
	//        String portType = str.substring(27,28);// 端口进出 1-进 2-出
	//        String verifyType = str.substring(28,29);// 刷卡类型 默认1
	//        String cardNo = str.substring(31,str.length()-4);// 卡号
	id := msg[2:26]                                     // 主板id 设备序列号
	number := msg[26:27]                                // 门号
	portType := msg[27:28]                              // 端口进出 1-进 2-出
	verifyType := msg[28:29]                            // 刷卡类型 默认1
	cardNo := msg[31 : bytes.Count([]byte(msg), nil)-4] // 卡号
	utils.Log.Info(id, number, portType, verifyType, cardNo)

	//conn.WriteToUDP("OK", address, msg)
}
