package icbc

import (
	"bytes"
	"emcs-relay-go/utils"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var checkUrl string
var verifyUrl string
var contentType = "application/json; charset=utf-8"

func init() {
	checkUrl = "http://119.29.194.87:9988"
	verifyUrl = "http://119.29.194.87:9988"
}

// CheckTicket 模拟发起HTTP请求 protocolNo 类型
func CheckTicket(ticket string, protocolNo string) (string, bool) {
	localData := time.Now().Local()
	data := CheckData{
		ClientType:   protocolNo,
		CientTransNo: protocolNo + localData.Format("20060102150405") + random3(),
		UpData:       ticket,
		UpDataLength: string(len(ticket)),
	}
	requestEntity := CheckRequest{
		Data:         data,
		CorpId:       "2000001924",
		CorpId2:      "2000001924",
		StrTESn:      authenticator("yccode", "ss"),
		Version:      "1",
		PrintControl: "0",
		TimeStamp:    string(localData.Unix()),
		ProtocolNo:   protocolNo,
		SystemType:   "1",
	}

	requestBody, err := json.Marshal(requestEntity)
	if err != nil {
		log.Fatal(err)
	}
	clt := http.Client{}
	resp, err := clt.Post(checkUrl, contentType, bytes.NewBuffer(requestBody))
	if err != nil {
		utils.Log.Error(err)
		utils.Log.Error(resp)
	}
	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)
	retMsg := res["retMsg"]
	retCode := res["retCode"]
	utils.Log.Info(retCode, retMsg)
	return fmt.Sprintln(retMsg), retCode == "0"
}
func authenticator(ycCode string, deviceId string) string {
	return ycCode + deviceId
}

/**
 * 随机三位数数
 */
func random3() string {
	random := rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(100)
	return fmt.Sprintf("%03v", random)
}
