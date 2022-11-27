package api

import (
	"bytes"
	"emcs-relay-go/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"io"
	"net/http"
)

func SignHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !config.EnableSign {
			return
		}
		body, _ := io.ReadAll(c.Request.Body)
		apiRequest := Request[map[string]interface{}]{}
		json.Unmarshal(body, &apiRequest)
		fmt.Println(body)
		isOK := apiRequest.CheckSign()
		if isOK {
			//c.Request.Body = io.NopCloser(bytes.NewBuffer(body)) //原始的body放回去
			marshal, err := json.Marshal(apiRequest.Data)
			if err != nil {
				fmt.Println("解析data时错误")
			}
			c.Request.Body = io.NopCloser(bytes.NewBuffer(marshal)) //new body
			c.Next()
		} else {
			c.JSON(http.StatusOK, ResponseError("sign error"))
		}

	}
}
