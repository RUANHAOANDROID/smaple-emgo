package db

import "database/sql"

type DeviceNumUpload struct {
	ID uint `gorm:"primaryKey"`
	// 设备编号
	DeviceNo sql.NullString
	// 生成时间
	BizDate sql.NullString
	// 人数
	PeopleNum sql.NullString
	// 是否上传 0-否 1-是
	IsUpload sql.NullString
}
