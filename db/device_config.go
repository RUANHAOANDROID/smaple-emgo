package db

import "database/sql"

type DeviceConfig struct {
	ID              string
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

func SaveConfig() {
	var sql = "insert into device_config(id,manufacturer_id1,manufacturer_id2,buffer,delayed_time,invalid_time,check_url,write_off_url,num_up_url,num_up_time, heartbeat_url, heartbeat_time, true_voice1, false_voice1, true_voice2, false_voice2, true_voice3, false_voice3, true_voice4, false_voice4, config_url ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?,?, ?, ?, ?, ?,?, ?, ?, ?, ?, ?, ? )"
	db.Prepare(sql)
	db.Exec(sql)
}
