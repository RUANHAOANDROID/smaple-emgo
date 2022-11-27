package db

import "database/sql"

type SysLog struct {
	Id      uint `gorm:"primaryKey"`
	Content sql.NullString
	AddTime sql.NullTime
}
