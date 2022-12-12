package api

import (
	"bytes"
	"emcs-relay-go/configs"
	"emcs-relay-go/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"io"
	"net/http"
)

func SignHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !configs.EnableSign {
			return
		}
		url := c.Request.URL.Path
		if url == "/admin/getConfig" {
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
				fmt.Println("request body format error")
				c.JSON(http.StatusOK, ResponseError("request body format error"))
				return
			}
			c.Request.Body = io.NopCloser(bytes.NewBuffer(marshal)) //new body
			c.Next()
		} else {
			c.JSON(http.StatusOK, ResponseError("sign error"))
			return
		}

	}
}
func LogHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		url := c.Request.URL.Path
		if url == "/admin/getConfig" {
			return
		}
		utils.Log.Info(url)
		body, _ := io.ReadAll(c.Request.Body)
		apiRequest := string(body)
		utils.Log.Info(apiRequest)
		fmt.Println(body)
		c.Request.Body = io.NopCloser(bytes.NewBuffer(body)) //new body
		c.Next()
	}
}
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
