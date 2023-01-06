package entity

type DeviceDelete struct {
	Id uint `json:"id"`
}
type GetEventsPage struct {
	Date       string `json:"date"`
	PageNo     int    `json:"pageNo"`
	PageSize   int    `json:"pageSize"`
	DeviceName string `json:"deviceName"`
}
