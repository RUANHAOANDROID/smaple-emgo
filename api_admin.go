package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func HandlerAdmin(r *gin.RouterGroup) {
	r.Use(signHandler())
	r.POST("/login", func(c *gin.Context) {
		request := LoginRequest{}
		err := c.ShouldBind(&request)
		if err != nil {
			c.JSON(http.StatusOK, ResponseError("参数错误"))
		} else {
			response := ResponseSuccess(any("ok"))
			fmt.Println(response)
			c.JSON(http.StatusOK, response)
		}
	})
	r.POST("/config/save", func(c *gin.Context) {
		c.String(http.StatusOK, "保存成功")
	})
}
func signHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !enableSign {
			return
		}
		body, _ := io.ReadAll(c.Request.Body)
		apiRequest := ApiRequest[map[string]interface{}]{}
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
