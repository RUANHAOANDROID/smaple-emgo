package db

import "database/sql"

type DeviceNumUpload struct {
	ID        uint `gorm:"primaryKey"`
	DeviceNo  sql.NullString
	BizDate   sql.NullString
	PeopleNum sql.NullString
	IsUpload  sql.NullString
}
