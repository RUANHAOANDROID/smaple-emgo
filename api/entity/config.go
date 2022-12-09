package entity

type success struct {
	Normal    string `json:"normal"`
	Other     string `json:"other"`
	YearCard1 string `json:"yearCard1"`
	YearCard2 string `json:"yearCard2"`
}
type fail struct {
	Other   string `json:"other"`
	Invalid string `json:"invalid"`
	Used    string `json:"used"`
	Error   string `json:"error"`
}
type FlowConfig struct {
	HeartbeatTime     int `json:"heartbeatTime"`
	Success           success
	ConfigNo          string `json:"configNo"`
	WriteOffDelayTime int    `json:"writeOffDelayTime"`
	InvalidTime       int    `json:"invalidTime"`
	EquipmentNo       string `json:"equipmentNo"`
	Yccode            string `json:"yccode"`
	ManufacturerId2   string `json:"manufacturerId2"`
	ManufacturerId1   string `json:"manufacturerId1"`
	Fail              fail
	WriteOffUrl       string `json:"writeOffUrl"`
	NumUpUrl          string `json:"numUpUrl"`
	HeartbeatUrl      string `json:"heartbeatUrl"`
	FalseVoice3       string `json:"falseVoice3"`
	NumUpTime         int    `json:"numUpTime"`
	FalseVoice4       string `json:"falseVoice4"`
	WelcomeMsg        string `json:"welcomeMsg"`
	CheckUrl          string `json:"checkUrl"`
}
