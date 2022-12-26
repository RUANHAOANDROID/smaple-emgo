package entity

const TYPE_DEVICES = 1   //设备
const TYPE_LOG = 2       //日志
const TYPE_EVENT = 3     //事件
const TYPE_HARDWARES = 4 // 硬件信息
const TYPE_TOTAL = 6     // 统计

type Msg[T interface{}] struct {
	Type int8 `json:"type"`
	Data T    `json:"data"`
}

func Pack(aType int8, data interface{}) Msg[interface{}] {
	return Msg[interface{}]{Type: aType,
		Data: data}
}
