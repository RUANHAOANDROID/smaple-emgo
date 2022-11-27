package db

import "database/sql"

type SysLogTime struct {
	Id       uint `gorm:"primaryKey"`
	LastTime sql.NullTime
}
