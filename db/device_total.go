package db

import (
	"emcs-relay-go/utils"
	_ "emcs-relay-go/utils"
)

type DeviceTotal struct {
	//自增ID
	ID uint `json:"_"gorm:"primaryKey"`
	//日期到天
	Date string `json:"date"`
	// 设备标签
	DTag string `json:"deviceTag"`
	// 编号
	DId string `json:"deviceId"`
	// 设备过闸次数
	DCount int64 `json:"_" gorm:"default:0"`
}
type DeviceTotalVO struct {
	DeviceTotal
	Sum        int64   `json:"device_sum"`
	Proportion float32 `json:"proportion"`
}

func TotalUpAdd(number string) error {
	err := Db().Exec("UPDATE devices SET count=(count+1) , last_time=? WHERE number=?", utils.NowTimeStr(), number)
	return err.Error
}

func TotalDeviceCountByDay(devices *[]DeviceTotalVO, day string) error {
	err := Db().Raw("SELECT d.d_id, d.d_tag, SUM(d_count) as sum, d.date FROM  device_totals d WHERE d.date =? GROUP BY d.d_tag ORDER BY sum DESC", day).Scan(devices)
	return err.Error
}

func TotalDeviceCount(devices *[]DeviceTotalVO) error {
	err := Db().Raw("SELECT d.d_id,d.d_tag, SUM(d_count) as sum, d.date FROM  device_totals d GROUP BY d.d_tag").Scan(devices)
	return err.Error
}
func TotalSumByDay(count *int64, day string) error {
	err := Db().Raw("SELECT SUM(d_count) FROM  device_totals WHERE date=?", day).Scan(count)
	return err.Error
}
