package db

import (
	"emcs-relay-go/utils"
	_ "emcs-relay-go/utils"
	"time"
)

type DeviceTotal struct {
	//自增ID
	ID uint `json:"_"gorm:"primaryKey"`
	// 编号
	DId uint `json:"deviceId"`
	//日期到天
	Date string `json:"date"`
	// 设备标签
	DTag string `json:"deviceTag"`
	// 设备过闸次数
	DCount int64 `json:"_" gorm:"default:0"`
}
type DeviceTotalVO struct {
	DeviceTotal
	Sum        int64   `json:"device_sum"`
	Proportion float32 `json:"proportion"`
}

func TotalUpAdd(number string) error {
	day := utils.Fmt2Day(time.Now().Local())
	count := 0
	device := Device{}
	err := GetDevice(number, &device)
	if err.Error != nil {
		utils.Log.Error(err.Error)
		return err.Error
	}
	total := DeviceTotal{}
	err = Db().Table("device_totals").Select("1").Where("d_id=? AND date=?", device.ID, day).Limit(1).Scan(&count)
	if err.Error != nil {
		utils.Log.Info(err.Error)
	}
	if count != 0 {
		err = Db().Exec("UPDATE device_totals SET d_count=(d_count+1) WHERE date=? AND d_id=?", day, device.ID)
	} else {
		total = DeviceTotal{
			Date:   day,
			DCount: 1,
			DTag:   device.Tag,
			DId:    device.ID,
		}
		err = DB.Create(&total)
		if err.Error != nil {
			utils.Log.Error(err)
		}
	}
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
