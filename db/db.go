package db

import (
	"emcs-relay-go/utils"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"strings"
)

var DB *gorm.DB

func init() {
	if DB == nil {
		db, err := gorm.Open(sqlite.Open("emcs.db"), &gorm.Config{
			NowFunc: utils.Local,
			Logger:  logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			utils.PrintErr(err)
		}
		DB = db
		createUser()
		createSysLog()
		createSysLogTime()
		createDevice()
		createDeviceConfig()
		createNumUpLod()
	}
}
func createUser() {
	user := SysUser{UserName: "admin", UserPwd: "666666"}
	if !DB.Migrator().HasTable(&user) {
		DB.AutoMigrate(&user)
		DB.Create(&user)
	}
}
func createDevice() {
	dst := &Device{}
	if !DB.Migrator().HasTable(dst) {
		err := DB.AutoMigrate(dst)
		utils.PrintErr(err)
	}
}
func createDeviceConfig() {
	dst := &DeviceConfig{}
	if !DB.Migrator().HasTable(dst) {
		err := DB.AutoMigrate(&DeviceConfig{})
		utils.PrintErr(err)
	}
}
func createNumUpLod() {
	dst := &DeviceNumUpload{}
	if !DB.Migrator().HasTable(dst) {
		err := DB.AutoMigrate(dst)
		utils.PrintErr(err)
	}
}
func createSysLog() {
	dst := &SysLog{}
	if !DB.Migrator().HasTable(dst) {
		err := DB.AutoMigrate(dst)
		utils.PrintErr(err)
	}
}
func createSysLogTime() {
	dst := &SysLogTime{}
	if !DB.Migrator().HasTable(dst) {
		err := DB.AutoMigrate(dst)
		utils.PrintErr(err)
	}
}
func createTab() {
	file, _ := os.ReadFile("./sql/table.sql")
	sqlString := strings.Split(strings.ReplaceAll(string(file), "\n", ""), ";")
	for index, sql := range sqlString {
		fmt.Println(index)
		if sql != "" {
			DB.Exec(sql)
			fmt.Println(sql)
		}
	}
}
func installData() {
	file, _ := os.ReadFile("./sql/data.sql")
	sqlString := strings.Split(strings.ReplaceAll(string(file), "\n", ""), ";")
	for index, sql := range sqlString {
		fmt.Println(index)
		if sql != "" {
			DB.Exec(sql)
			fmt.Println(sql)
		}
	}
}

// Db returns the default *gorm.DB connection.
func Db() *gorm.DB {
	return DB
}

// UnscopedDb returns an unscoped *gorm.DB connection
// that returns all records including deleted records.
func UnscopedDb() *gorm.DB {
	return Db().Unscoped()
}
