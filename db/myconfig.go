package db

import (
	"emcs-relay-go/api/entity"
	"emcs-relay-go/utils"
	"github.com/goccy/go-json"
)

var EmcsUrl *string
var EmcsCode *string
var EmcsConfig *entity.EmcsConfig

type MyConfig struct {
	Id      uint   `json:"id"gorm:"primaryKey"`
	Code    string `json:"code"`
	Url     string `json:"url"`
	Content string `json:"content"`
	Version int    `json:"version"`
}

func init() {
	one := MyConfig{}
	err := GetMyConfig(&one)
	if err != nil {
		utils.Log.Info("获取配置失败")
		return
	}
	emcsConfig := entity.EmcsConfig{}
	err = json.Unmarshal([]byte(one.Content), emcsConfig)
	if err != nil {
		utils.Log.Error(err)
	}
	EmcsConfig = &emcsConfig
	EmcsUrl = &one.Url
	EmcsCode = &one.Code

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
