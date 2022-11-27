package db

import "database/sql"

type SysLogTime struct {
	ID       uint `gorm:"primaryKey"`
	LastTime sql.NullTime
}
