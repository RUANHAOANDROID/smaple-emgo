package db

import (
	"database/sql"
	"emcs-relay-go/logger"
)

type Device struct {

	// 设备主板id
	Id string `json:"id" gorm:"primaryKey"`
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
	Status sql.NullString `json:"status"`
	// 新增时间
	AddTime sql.NullString `json:"addTime"`
	// 修改时间
	UpdateTime sql.NullString `json:"updateTime"`
}

func AddDevice(device Device) error {
	err := DB.Create(&device).Error

	logger.Log.Info(device)
	return err
}
