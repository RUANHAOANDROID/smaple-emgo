package api

import (
	"bytes"
	"emcs-relay-go/api/entity"
	"emcs-relay-go/api/icbc"
	"emcs-relay-go/configs"
	"emcs-relay-go/db"
	"emcs-relay-go/utils"
	"fmt"
	"net"
	"time"
)

const resultSuccess = "A"
const resultFail = "5"
const voice1 = "1"
const voice2 = "2"
const voice3 = "3"
const voice4 = "4"
const voice5 = "5"

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
	eventLog := db.EventLog{}
	if isCheck {
		sendSwipeEvent(eventLog, msg)
		id := msg[2:26]                                     // 主板id 设备序列号
		number := msg[26:27]                                // 门号
		portType := msg[27:28]                              // 端口进出 1-进 2-出
		verifyType := msg[28:29]                            // 刷卡类型 默认1
		cardNo := msg[31 : bytes.Count([]byte(msg), nil)-4] // 卡号
		utils.Log.Info(id, number, portType, verifyType, cardNo)
		swipeTicket(msg)
		result, ok := icbc.CheckTicket(cardNo, "002")
		utils.Log.Info(result, ok)
		if ok {
			OpenGate(id, clientAddress.IP.String())
			sendCheckEvent(eventLog, "验票成功")
			if !configs.PassedVerifyMode {
				sendVerifyEvent(eventLog)
			}
		} else {
			ErrorTip(id, clientAddress.IP.String())
			sendCheckEvent(eventLog, "验票失败")
		}
		db.AddEvent(&eventLog)
		db.TotalUpAdd(number)
		return
	}

	isPassed, err := isPassed(msg)
	if err != nil {
		utils.Log.Info(err.Error())
	}
	if isPassed {
		sendPassedEvent(eventLog, msg)
		sendTotalEvent()
		utils.Log.Info(msg)
		return
	}
	//var str =[]byte(" $F12345678F$")

}
func sendTotalEvent() {
	type TotalVO struct {
		Sum         int64              `json:"sum"`
		DeviceTotal []db.DeviceTotalVO `json:"deviceTotals"`
	}
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
func sendPassedEvent(eventLog db.EventLog, msg string) {
	eventLog.Tag = "过闸"
	eventLog.Time = utils.NowTimeStr()
	eventLog.Content = msg
	SendMsg(entity.Pack(entity.TYPE_EVENT, []db.EventLog{eventLog}))
}
func sendSwipeEvent(eventLog db.EventLog, msg string) {
	eventLog.Tag = "刷卡"
	eventLog.Time = utils.NowTimeStr()
	eventLog.Content = msg
	SendMsg(entity.Pack(entity.TYPE_EVENT, []db.EventLog{eventLog}))
}
func sendCheckEvent(eventLog db.EventLog, tip string) {
	eventLog.Tag = "验票"
	eventLog.Time = utils.NowTimeStr()
	eventLog.Content = tip
	SendMsg(entity.Pack(entity.TYPE_EVENT, []db.EventLog{eventLog}))
}
func sendVerifyEvent(eventLog db.EventLog) {
	eventLog.Tag = "核销"
	eventLog.Time = utils.NowTimeStr()
	eventLog.Content = "核销？"
	SendMsg(entity.Pack(entity.TYPE_EVENT, []db.EventLog{eventLog}))
}

func swipeTicket(msg string) {
	go func() {
		log := db.EventLog{Tag: "刷票", Content: msg, Time: utils.NowTimeStr()}
		db.AddEvent(&log)
		SendMsg(entity.Pack(entity.TYPE_LOG, log))
	}()
}

// OpenGate number 设备号 /*
func OpenGate(number string, ip string) {
	utils.Log.Info("开闸", number, ip)
	udpAddr := net.UDPAddr{
		IP: net.ParseIP(ip), Port: 60066,
	}
	cmd := "$F" + number + "111" + "A" + "01000/&\\&$E"
	bytes := []byte(cmd)
	utils.Log.Info("编号：", number)
	utils.Log.Info("ip：", ip)
	utils.Log.Info("报文：", cmd)
	Conn.WriteToUDP(bytes, &udpAddr)
}

func ErrorTip(number string, ip string) {
	utils.Log.Info("错误播报", number, ip)
	udpAddr := net.UDPAddr{
		IP: net.ParseIP(ip), Port: 60066,
	}
	pack := pack(number, resultFail, voice2)
	Conn.WriteToUDP([]byte(pack), &udpAddr)
}

/*
打包回复报文
number 门号
result 验证结果 A通过|5不通过
voice 语音播报 1-5
*/
func pack(number string, result string, voice string) string {
	gateNo := "1"     //门号 1-4
	io := "1"         //进出 进1|出2
	ticketType := "1" //票据类型 默认1
	peopleCount := "01"
	passedCount := "00"
	pack := ("$") + ("F") + number + gateNo + io + ticketType + result + peopleCount + passedCount + voice + ("/&") + ("\\&") + ("$E")
	utils.Log.Info(pack)
	return pack
}
