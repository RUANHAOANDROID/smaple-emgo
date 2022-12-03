package db

type SysUser struct {
	UserName string `gorm:"primaryKey"`
	UserPwd  string
}

var user = &SysUser{}

func UserExists(name string) error {
	return DB.Where("user_name = ?", name).Take(&SysUser{}).Error
}
func UserVerify(name string, pwd string) error {
	return DB.Where("user_name=? AND user_pwd=?", name, pwd).Take(&SysUser{}).Error
}
func UpdatePwd(name string, newPwd string) error {
	user.UserName = name
	return DB.Model(user).Update("user_pwd", newPwd).Error
}
