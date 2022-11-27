package db

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"strings"
)

var DB *gorm.DB

// Set UTC as the default for created and updated timestamps.
func init() {
	if DB == nil {
		db, err := gorm.Open(sqlite.Open("emcs.db"), &gorm.Config{
			NowFunc: UTC,
		})
		if err != nil {
			fmt.Println("open db error")
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
	user := SysUser{UserName: sql.NullString{
		String: "admin", Valid: true,
	}, UserPwd: sql.NullString{String: "666666", Valid: true}}
	if !DB.Migrator().HasTable(&user) {
		DB.AutoMigrate(&user)
		DB.Create(&user)
	}
}
func createDevice() {
	dst := &Device{}
	if !DB.Migrator().HasTable(dst) {
		DB.AutoMigrate(dst)
	}
}
func createDeviceConfig() {
	dst := &DeviceConfig{}
	if !DB.Migrator().HasTable(dst) {
		DB.AutoMigrate(&DeviceConfig{})
	}
}
func createNumUpLod() {
	dst := &DeviceNumUpload{}
	if !DB.Migrator().HasTable(dst) {
		DB.AutoMigrate(dst)
	}
}
func createSysLog() {
	dst := &SysLog{}
	if !DB.Migrator().HasTable(dst) {
		DB.AutoMigrate(dst)
	}
}
func createSysLogTime() {
	dst := &SysLogTime{}
	if !DB.Migrator().HasTable(dst) {
		DB.AutoMigrate(dst)
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
