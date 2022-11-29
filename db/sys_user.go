package db

type SysUser struct {
	UserName string `gorm:"primaryKey"`
	UserPwd  string
}

func UserExists(name string) error {
	return DB.Where("user_name = ?", name).Take(&SysUser{}).Error
}
func UserVerify(name string, pwd string) error {
	return DB.Where("user_name=? AND user_pwd=?", name, pwd).Take(&SysUser{}).Error
}
