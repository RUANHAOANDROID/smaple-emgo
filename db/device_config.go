package db

type DeviceConfig struct {
	Id              string `json:"id" gorm:"primaryKey"`
	Buffer          string `json:"buffer"`
	CheckUrl        string `json:"checkUrl"`
	ConfigUrl       string `json:"configUrl"`
	DeFalseText     string `json:"deFalseText"`
	DeFalseVoice    string `json:"deFalseVoice"`
	DeTrueText      string `json:"deTrueText"`
	DeTrueVoice     string `json:"deTrueVoice"`
	DelayedTime     string `json:"delayedTime"`
	FalseVoice1     string `json:"falseVoice1"`
	FalseVoice2     string `json:"falseVoice2"`
	FalseVoice3     string `json:"falseVoice3"`
	FalseVoice4     string `json:"falseVoice4"`
	HeartbeatTime   string `json:"heartbeatTime"`
	HeartbeatUrl    string `json:"heartbeatUrl"`
	InvalidTime     string `json:"invalidTime"`
	ManufacturerId1 string `json:"manufacturerId1"`
	ManufacturerId2 string `json:"manufacturerId2"`
	NumUpTime       string `json:"numUpTime"`
	NumUpUrl        string `json:"numUpUrl"`
	TrueVoice1      string `json:"trueVoice1"`
	TrueVoice2      string `json:"trueVoice2"`
	TrueVoice3      string `json:"trueVoice3"`
	TrueVoice4      string `json:"trueVoice4"`
	WriteOffUrl     string `json:"writeOffUrl"`
}

func GetConfig(config []DeviceConfig) error {
	return DB.First(&config).Error
}
func SaveConfig(dcf DeviceConfig) error {
	return DB.Create(&dcf).Error
}
