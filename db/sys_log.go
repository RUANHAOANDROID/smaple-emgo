package db

import "database/sql"

type SysLog struct {
	ID      uint `gorm:"primaryKey"`
	Content sql.NullString
	AddTime sql.NullTime
}
