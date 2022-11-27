package db

type SysUser struct {
	Id       uint `gorm:"primaryKey"`
	UserName string
	UserPwd  string
}
