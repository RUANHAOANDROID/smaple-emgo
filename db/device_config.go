package db

import (
	"database/sql"
)

type DeviceConfig struct {
	Id              uint `gorm:"primaryKey"`
	ManufacturerId1 sql.NullString
	ManufacturerId2 sql.NullString
	Buffer          sql.NullString
	DelayedTime     sql.NullString
	InvalidTime     sql.NullString
	CheckUrl        sql.NullString
	WriteOffUrl     sql.NullString
	NumUpUrl        sql.NullString
	NumUpTime       sql.NullString
	HeartbeatUrl    sql.NullString
	HeartbeatTime   sql.NullString
	TrueVoice1      sql.NullString
	FalseVoice1     sql.NullString
	TrueVoice2      sql.NullString
	FalseVoice2     sql.NullString
	TrueVoice3      sql.NullString
	FalseVoice3     sql.NullString
	TrueVoice4      sql.NullString
	FalseVoice4     sql.NullString
	ConfigUrl       sql.NullString
}

func SaveConfig(dcf DeviceConfig) {
	//var sql = "insert into device_config(" +
	//	"manufacturer_id1," +
	//	"manufacturer_id2," +
	//	"buffer," +
	//	"delayed_time," +
	//	"invalid_time," +
	//	"check_url," +
	//	"write_off_url," +
	//	"num_up_url," +
	//	"num_up_time," +
	//	"heartbeat_url," +
	//	"heartbeat_time," +
	//	"true_voice1," +
	//	"false_voice1," +
	//	"true_voice2," +
	//	"false_voice2," +
	//	"true_voice3," +
	//	"false_voice3," +
	//	"true_voice4," +
	//	"false_voice4," +
	//	"config_url" +
	//	") VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?,?, ?, ?, ?, ?,?, ?, ?, ?, ?, ?, ? )"
	//stmt, err := db.Prepare(sql)
	//if err != nil {
	//	logger.E(err)
	//}
	//stmt.Exec(sql,
	//	dcf.ManufacturerId1,
	//	dcf.ManufacturerId2,
	//	dcf.Buffer,
	//	dcf.DelayedTime,
	//	dcf.InvalidTime,
	//	dcf.CheckUrl,
	//	dcf.ConfigUrl,
	//	dcf.WriteOffUrl,
	//	dcf.NumUpUrl,
	//	dcf.NumUpTime,
	//	dcf.HeartbeatUrl,
	//	dcf.HeartbeatTime,
	//	dcf.TrueVoice1,
	//	dcf.FalseVoice1,
	//	dcf.TrueVoice2,
	//	dcf.FalseVoice2,
	//	dcf.TrueVoice3,
	//	dcf.FalseVoice3,
	//	dcf.TrueVoice4,
	//	dcf.FalseVoice4,
	//	dcf.ConfigUrl,
	//)
}
