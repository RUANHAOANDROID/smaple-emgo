package db

import "database/sql"

type Device struct {
	ID            uint `gorm:"primaryKey"`
	DeviceNo      sql.NullString
	DeviceIp      sql.NullString
	SerialNumber  sql.NullString
	DeviceVersion sql.NullString
	DeviceStatus  sql.NullString
	Status        sql.NullString
	AddTime       sql.NullString
	UpdateTime    sql.NullString
}
