package db

import (
	"database/sql"
	"testing"
)

func TestSaveConfig(t *testing.T) {
	SaveConfig(DeviceConfig{
		ManufacturerId1: sql.NullString{String: "ManufacturerId1", Valid: true},
		ManufacturerId2: sql.NullString{String: "ManufacturerId2", Valid: true},
		Buffer:          sql.NullString{String: "Buffer", Valid: true},
		DelayedTime:     sql.NullString{String: "ManufacturerId1", Valid: true},
		InvalidTime:     sql.NullString{String: "InvalidTime", Valid: true},
		CheckUrl:        sql.NullString{String: "CheckUrl", Valid: true},
		WriteOffUrl:     sql.NullString{String: "WriteOffUrl", Valid: true},
		NumUpUrl:        sql.NullString{String: "NumUpUrl", Valid: true},
		NumUpTime:       sql.NullString{String: "NumUpTime", Valid: true},
		HeartbeatUrl:    sql.NullString{String: "HeartbeatUrl", Valid: true},
		HeartbeatTime:   sql.NullString{String: "HeartbeatUrl", Valid: true},
		TrueVoice1:      sql.NullString{String: "TrueVoice1", Valid: true},
		FalseVoice1:     sql.NullString{String: "FalseVoice1", Valid: true},
		TrueVoice2:      sql.NullString{String: "TrueVoice2", Valid: true},
		FalseVoice2:     sql.NullString{String: "FalseVoice2", Valid: true},
		TrueVoice3:      sql.NullString{String: "TrueVoice3", Valid: true},
		FalseVoice3:     sql.NullString{String: "FalseVoice3", Valid: true},
		TrueVoice4:      sql.NullString{String: "TrueVoice4", Valid: true},
		FalseVoice4:     sql.NullString{String: "FalseVoice4", Valid: true},
		ConfigUrl:       sql.NullString{String: "ConfigUrl", Valid: true},
	})
	//demos, _ := db.Query("select * from demo")
	//assert.Nil(t, demos)
}
