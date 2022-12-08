package api

import (
	"crypto/md5"
	"emcs-relay-go/configs"
	"emcs-relay-go/utils"
	"encoding/json"
	"fmt"
	"strconv"
)

var md5Key = "#!1EMCXu5eIx190"

type EmcsRequest[T interface{}] struct {
	Data      T      `json:"data"`
	Sign      string `json:"sign"`
	Timestamp string `json:"timestamp"` //this is a string
}
type Request[T interface{}] struct {
	Data      T      `json:"data"`
	Sign      string `json:"sign"`
	Timestamp int64  `json:"timestamp"` // this is int64 (java long)
}

func (r Request[any]) CheckSign() bool {
	body, _ := json.Marshal(r.Data)
	bodyString := string(body)
	md5Str := "data=" + bodyString + "&key=" + md5Key + "&timestamp=" + fmt.Sprint(r.Timestamp)
	hash := md5.Sum([]byte(md5Str))
	return fmt.Sprintf("%x", hash) == r.Sign
}
func Sign(aJson any, timeLong int64) EmcsRequest[any] {
	//timestamp := db.TimeStamp().String()
	//timeInt64, _ := strconv.ParseInt(timestamp, 10, 64)
	bodyByte, _ := json.Marshal(&aJson)
	javaL := strconv.FormatInt(timeLong, 10)
	md5Str := "data=" + string(bodyByte) + "&key=" + configs.EmcsMD5Key + "&timestamp=" + javaL
	hash := md5.Sum([]byte(md5Str))
	md5 := fmt.Sprintf("%x", hash)
	utils.Log.Info(md5Str)
	utils.Log.Info(md5)
	return EmcsRequest[any]{
		Data:      aJson,
		Sign:      md5,
		Timestamp: javaL,
	}
}

type LoginRequest struct {
	UserName string `json:"userName"`
	PassWord string `json:"passWord"`
}

type WebGetConfig struct {
	ConfigUrl string `json:"configUrl"`
	Id        string `json:"id"`
}
type GetConfigEmcs struct {
	SerialNum string `json:"serialNum"`
}
type UpdateUser struct {
	Password string `json:"password"`
	UserName string `json:"userName"`
}

type UpdateDevice struct {
	DeviceIp      string `json:"deviceIp"`
	DeviceNo      string `json:"deviceNo"`
	DeviceStatus  string `json:"deviceStatus"`
	DeviceVersion string `json:"deviceVersion"`
	Id            string `json:"id"`
	SerialNumber  string `json:"serialNumber"`
}
type SaveConfig struct {
	Buffer       string `json:"buffer"`
	CheckUrl     string `json:"checkUrl"`
	ConfigUrl    string `json:"configUrl"`
	DeFalseText  string `json:"deFalseText"`
	DeFalseVoice struct {
		Error   string `json:"error"`
		Invalid string `json:"invalid"`
		Other   string `json:"other"`
		Used    string `json:"used"`
	} `json:"deFalseVoice"`
	DeTrueText  string `json:"deTrueText"`
	DeTrueVoice struct {
		Normal    string `json:"normal"`
		Other     string `json:"other"`
		YearCard1 string `json:"yearCard1"`
		YearCard2 string `json:"yearCard2"`
	} `json:"deTrueVoice"`
	DelayedTime     string `json:"delayedTime"`
	FalseVoice1     string `json:"falseVoice1"`
	FalseVoice2     string `json:"falseVoice2"`
	FalseVoice3     string `json:"falseVoice3"`
	FalseVoice4     string `json:"falseVoice4"`
	HeartbeatTime   string `json:"heartbeatTime"`
	HeartbeatUrl    string `json:"heartbeatUrl"`
	Id              string `json:"id"`
	InvalidTime     string `json:"invalidTime"`
	ManufacturerId1 string `json:"manufacturerId1"`
	ManufacturerId2 string `json:"manufacturerId2"`
	NumUpTime       string `json:"numUpTime"`
	NumUpUrl        string `json:"numUpUrl"`
	TrueVoice1      string `json:"trueVoice1"`
	TrueVoice2      string `json:"trueVoice2"`
	TrueVoice3      string `json:"trueVoice3"`
	TrueVoice4      string `json:"trueVoice4"`
	WriteOffUrl     string `json:"writeOffUrl"`
}
