package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
)

var md5Key = "#!1EMCXu5eIx190"

type ApiRequest[T interface{}] struct {
	Data      T      `json:"data"`
	Sign      string `json:"sign"`
	Timestamp int64  `json:"timestamp"`
}
type LoginRequest struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

func (r ApiRequest[any]) CheckSign() bool {
	body, _ := json.Marshal(r.Data)
	bodyString := string(body)
	md5Str := "data=" + bodyString + "&key=" + md5Key + "&timestamp=" + fmt.Sprint(r.Timestamp)
	hash := md5.Sum([]byte(md5Str))
	return fmt.Sprintf("%x", hash) == r.Sign
}
