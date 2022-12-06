package db

import (
	"emcs-relay-go/utils"
)

type Device struct {
	ID uint `json:"id"gorm:"primaryKey"`
	// 设备主板id
	DeviceId string `json:"deviceId"`
	// 设备编号
	DeviceNo string `json:"deviceNo"`
	// 设备ip
	DeviceIp string `json:"deviceIp"`
	// 设备主板id
	SerialNumber string `json:"serialNumber"`
	// 设备版本
	DeviceVersion string `json:"deviceVersion"`
	// 设备状态
	DeviceStatus string `json:"deviceStatus"`
	// 状态 1-正常 2-删除
	Status string `json:"status"`
	// 新增时间
	AddTime string `json:"addTime"`
	// 修改时间
	UpdateTime string `json:"updateTime"`
}

func AddDevice(device Device) error {
	err := DB.Create(&device).Error

	utils.Log.Info(device)
	return err
}
func DeleteDevice(id uint) error {
	//err := UnscopedDb().Delete(&Device{}, id)
	err := UnscopedDb().Where("id = ?", id).Delete(&Device{}).Error
	return err
}
func UpdateDevice(device Device) error {
	return DB.Model(&device).Updates(device).Error
}
