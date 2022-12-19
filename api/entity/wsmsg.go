package entity

const TYPE_HARDWARE = 1
const TYPE_HARDWARES = 4
const TYPE_LOG = 2
const TYPE_TOTAL = 6
const TYPE_EVENT = 3

type Msg[T interface{}] struct {
	Type int8 `json:"type"`
	Data T    `json:"data"`
}

func Pack(aType int8, data interface{}) Msg[interface{}] {
	return Msg[interface{}]{Type: aType,
		Data: data}
}
