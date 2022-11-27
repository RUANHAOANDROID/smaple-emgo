package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var success = 1
var fail = 0

type Response[T interface{}] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

func ResponseError(msg string) Response[string] {
	return Response[string]{Code: fail, Message: msg, Data: ""}
}
func ResponseSuccess(data any) Response[any] {
	return Response[any]{Code: success, Message: "SUCCESS", Data: data}
}
func error(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, ResponseError(msg))
	return
}
