package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var success = 1
var fail = 0

type WebResponse[T interface{}] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

func ResponseError(msg string) WebResponse[string] {
	return WebResponse[string]{Code: fail, Message: msg, Data: ""}
}
func ResponseSuccess(data any) WebResponse[any] {
	return WebResponse[any]{Code: success, Message: "SUCCESS", Data: data}
}
func ResError(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, ResponseError(msg))
	return
}
