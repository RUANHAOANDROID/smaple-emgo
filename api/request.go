package api

import (
	"crypto/md5"
	"emcs-relay-go/configs"
	"emcs-relay-go/logger"
	"encoding/json"
	"fmt"
	"strconv"
)

var md5Key = "#!1EMCXu5eIx190"

type EmcsRequest[T interface{}] struct {
	Data      T      `json:"data"`
	Sign      string `json:"sign"`
	Timestamp string `json:"timestamp"`
}
type Request[T interface{}] struct {
	Data      T      `json:"data"`
	Sign      string `json:"sign"`
	Timestamp int64  `json:"timestamp"`
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
	logger.Log.Info(md5Str)
	logger.Log.Info(md5)
	return EmcsRequest[any]{
		Data:      aJson,
		Sign:      md5,
		Timestamp: javaL,
	}
}

type LoginRequest struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type WebGetConfig struct {
	ConfigUrl string `json:"configUrl"`
	Id        string `json:"id"`
}
type GetConfigEmcs struct {
	SerialNum string `json:"serialNum"`
}
