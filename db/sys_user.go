package db

type SysUser struct {
	ID       uint `gorm:"primaryKey"`
	UserName string
	UserPwd  string
}
