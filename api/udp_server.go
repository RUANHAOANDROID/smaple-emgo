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

const voice1 = "1" // 1请通过
const voice2 = "2" // 2无效卡
const voice3 = "3" // 3验票通过
const voice4 = "4" // 4 一路平安
const voice5 = "5" // 5 准许通过

var Conn *net.UDPConn

func RunUDP(port string) {
	ipv4, err := utils.GetOutboundIP()
	if err != nil {
		panic("未获取到本机IP地址")
	}
	// 创建 服务器 UDP 地址结构。指定 IP + port
	udpAddr, err := net.ResolveUDPAddr("udp", ipv4.String()+":"+port)
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
		sn := msg[2:26]                                     // 主板id 设备序列号
		number := msg[26:27]                                // 门号
		portType := msg[27:28]                              // 端口进出 1-进 2-出
		verifyType := msg[28:29]                            // 刷卡类型 默认1
		cardNo := msg[31 : bytes.Count([]byte(msg), nil)-4] // 卡号
		utils.Log.Info(sn, number, portType, verifyType, cardNo)
		device := db.Device{}
		err = db.GetDevice(sn, &device).Error
		if err != nil {
			utils.Log.Error("该设备未注册")
			return
		}
		eventLog.DeviceName = device.Tag
		sendSwipeEvent(eventLog, msg)
		result, ok := icbc.CheckTicket(cardNo, "002")
		utils.Log.Info(result, ok)
		if ok {
			OpenGate(sn, clientAddress.IP.String())
			sendCheckEvent(eventLog, "验票成功")
			if !configs.PassedVerifyMode {
				sendVerifyEvent(eventLog)
			}
		} else {
			ErrorTip(sn, clientAddress.IP.String())
			sendCheckEvent(eventLog, "验票失败,"+result)
		}
		return
	}
	isPassed, err := isPassed(msg)
	if err != nil {
		utils.Log.Info(err.Error())
	}
	if isPassed {
		sn := msg[2:26]
		device := db.Device{}
		err = db.GetDevice(sn, &device).Error
		eventLog.DeviceName = device.Tag
		db.TotalUpAdd(sn)
		sendPassedEvent(eventLog, msg)
		sendTotalMsg()
		utils.Log.Info(msg)
		return
	}
	utils.Log.Error("接收到位置消息", msg)
}

// 统计变更
func sendTotalMsg() {
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

// 过闸事件
func sendPassedEvent(eventLog db.EventLog, msg string) {
	eventLog.Tag = "过闸"
	eventLog.Time = utils.NowTimeStr()
	eventLog.Content = msg
	db.AddEvent(&eventLog)
	SendMsg(entity.Pack(entity.TYPE_EVENT, []db.EventLog{eventLog}))
}

// 刷票事件
func sendSwipeEvent(eventLog db.EventLog, msg string) {
	eventLog.Tag = "刷卡"
	eventLog.Time = utils.NowTimeStr()
	eventLog.Content = msg
	db.AddEvent(&eventLog)
	SendMsg(entity.Pack(entity.TYPE_EVENT, []db.EventLog{eventLog}))
}

// 验票事件
func sendCheckEvent(eventLog db.EventLog, tip string) {
	eventLog.Tag = "验票"
	eventLog.Time = utils.NowTimeStr()
	eventLog.Content = tip
	db.AddEvent(&eventLog)
	SendMsg(entity.Pack(entity.TYPE_EVENT, []db.EventLog{eventLog}))
}

// 核销事件
func sendVerifyEvent(eventLog db.EventLog) {
	eventLog.Tag = "核销"
	eventLog.Time = utils.NowTimeStr()
	eventLog.Content = "核销？"
	db.AddEvent(&eventLog)
	SendMsg(entity.Pack(entity.TYPE_EVENT, []db.EventLog{eventLog}))
}

// OpenGate number 设备号 /*
func OpenGate(number string, ip string) {
	utils.Log.Info("开闸", number, ip)
	udpAddr := net.UDPAddr{
		IP: net.ParseIP(ip), Port: 60066,
	}
	cmd := pack(number, resultSuccess, voice1)
	bytes := []byte(cmd)
	Conn.WriteToUDP(bytes, &udpAddr)
}

func ErrorTip(sn string, ip string) {
	utils.Log.Info("错误播报", sn, ip)
	udpAddr := net.UDPAddr{
		IP: net.ParseIP(ip), Port: 60066,
	}
	cmd := pack(sn, resultFail, voice2)
	Conn.WriteToUDP([]byte(cmd), &udpAddr)
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
