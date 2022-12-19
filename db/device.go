package db

import (
	"emcs-relay-go/utils"
)

type Device struct {
	ID uint `json:"id"gorm:"primaryKey"`
	// 标签
	Tag string `json:"tag"`
	// 编号
	Number string `json:"number"`
	// ip
	Ip string `json:"ip"`
	//序列号
	Sn string `json:"sn"`
	// 版本
	Version string `json:"version"`
	// 设备过闸次数
	Count int64 `json:"count" gorm:"default:0"`
	// 状态 1-正常 2-删除
	Status string `json:"status" gorm:"default:0"`
	// 新增时间
	AddTime string `json:"addTime"`
	// 修改时间
	UpdateTime string `json:"updateTime"`
	//最后过闸时间
	LastTime string `json:"last_time"`
}

func DevicesList(devices *[]Device) error {
	return DB.Find(devices).Error
}
func AddDevice(device *Device) error {
	device.AddTime = utils.NowTimeStr()
	err := DB.Create(&device).Error
	utils.Log.Info(&device)
	return err
}
func DeleteDevice(id uint) error {
	//err := UnscopedDb().Delete(&Device{}, id)
	err := UnscopedDb().Where("id = ?", id).Delete(&Device{}).Error
	return err
}
func UpdateDevice(device Device) error {
	//{"id":null,"tag":"Tagalog","number":"设备编号","ip":"IP","sn":"sn",
	//"version":"version","status":null,"addTime":null,"updateTime":null}
	//err := DB.Model(&device).Where("id=?", device.ID).Updates(map[string]interface{}{
	//	"tag":         device.Tag,
	//	"number":      device.Number,
	//	"ip":          device.Ip,
	//	"sn":          device.Sn,
	//	"version":     device.Version,
	//	"status":      device.Status,
	//	"update_time": utils.Fmt2HMS(time.Now()),
	//},
	//).Error
	//err := DB.Save(&device).Where("id=?", device.ID).Error
	err := DB.Exec("UPDATE devices SET "+
		" tag = ?,"+
		" number = ?,"+
		" ip = ?,"+
		" sn = ?,"+
		" version = ?,"+
		" status = ?,"+
		" update_time = ?"+
		" WHERE  id = ?",
		device.Tag,
		device.Number,
		device.Ip,
		device.Sn,
		device.Version,
		device.Status,
		utils.NowTimeStr(),
		device.ID).Error
	return err
}
func PassedAddUp(number string) error {
	err := Db().Exec("UPDATE devices SET count=(count+1) , last_time=? WHERE number=?", utils.NowTimeStr(), number)
	return err.Error
}

func TotalPassed(devices *[]Device) error {
	err := Db().Raw("SELECT * FROM  devices ORDER BY count DESC").Scan(devices)
	return err.Error
}
func TotalAllCount(count *int64) error {
	err := Db().Raw("SELECT SUM(count) FROM  devices").Scan(count)
	return err.Error
}
