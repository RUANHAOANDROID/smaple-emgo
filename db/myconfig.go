package db

type MyConfig struct {
	Id          uint   `json:"id"gorm:"primaryKey"`
	JsonContent string `json:"json_content"`
	Version     int    `json:"version"`
}

func GetMyConfig(config *MyConfig) error {
	return DB.Last(config).Error
}
func SaveMyConfig(dcf *MyConfig) error {
	return DB.Save(&dcf).Error
}
func DeleteMyOtherConfig() error {
	return DB.Exec("DELETE  From my_configs WHERE id<((select max(id) id from my_configs))").Error
}
