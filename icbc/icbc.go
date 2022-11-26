package icbc

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// CheckTicket 模拟发起HTTP请求
func CheckTicket() string {
	clt := http.Client{}
	resp, err := clt.Get("https://dev.hao88.cloud/log/get")
	if responseError(err, resp) {
		return "check fail"
	}
	content, err := ioutil.ReadAll(resp.Body)
	respBody := string(content)
	fmt.Println(respBody, err)
	return "check ticket success"
}

func responseError(err error, resp *http.Response) bool {
	if err != nil {
		fmt.Println(resp, err)
		return true
	}
	return false
}
