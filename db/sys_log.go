package db

import "database/sql"

type SysLog struct {
	Id      uint `gorm:"primaryKey"`
	Content string
	AddTime sql.NullTime
}

func AddLog(log SysLog) error {
	return DB.Save(&log).Error
}
func GetLog(log *SysLog) error {
	return DB.Raw("SELECT * FROM sys_users WHERE add_time = ?", "admin").Scan(log).Error
}
