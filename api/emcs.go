package api

import (
	"bytes"
	"emcs-relay-go/configs"
	"emcs-relay-go/utils"
	"github.com/goccy/go-json"
	"io/ioutil"
	"net/http"
)

var contentType = "application/json; charset=utf-8"
var url = configs.PlatformUrl

const (
	gConfig = "/gateMachine/getConfig"
)

type EmcsResponse[T interface{}] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

func RequestConfig(url string, id string, timestamp int64) (WebResponse[any], error) {
	utils.Log.Info("request getConfig", url, id)
	getConf := GetConfigEmcs{
		SerialNum: id,
	}
	rq := Sign(getConf, timestamp)
	utils.Log.Info(rq)
	reqBody, err := json.Marshal(&rq)
	clt := http.Client{}
	resp, err := clt.Post(url+gConfig, contentType, bytes.NewBuffer(reqBody))
	content, err := ioutil.ReadAll(resp.Body)
	jsonString := string(content)
	utils.Log.Info(jsonString)
	emcsR := EmcsResponse[any]{}
	err = json.Unmarshal(content, &emcsR)
	webR := WebResponse[any]{
		Code:    1,
		Data:    emcsR.Data,
		Message: emcsR.Msg,
	}
	return webR, err
}
