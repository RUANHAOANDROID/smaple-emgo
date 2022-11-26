package main

var success = 1
var fail = 0

type ApiResponse[T interface{}] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

func ResponseError(msg string) ApiResponse[string] {
	return ApiResponse[string]{Code: fail, Message: msg, Data: ""}
}
func ResponseSuccess(data any) ApiResponse[any] {
	return ApiResponse[any]{Code: success, Message: "ok", Data: data}
}
