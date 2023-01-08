package entity

type Page struct {
	Count int64       `json:"count"`
	Data  interface{} `json:"data"`
}
