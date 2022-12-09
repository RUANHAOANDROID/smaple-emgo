package entity

type EmcsConfig struct {
	TrueVoice1    string `json:"trueVoice1"`
	HeartbeatTime int    `json:"heartbeatTime"`
	TrueVoice3    string `json:"trueVoice3"`
	TrueVoice2    string `json:"trueVoice2"`
	DeFalseVoice  struct {
		Other   string `json:"other"`
		Invalid string `json:"invalid"`
		Used    string `json:"used"`
		Error   string `json:"error"`
	} `json:"deFalseVoice"`
	ConfigNo          string `json:"configNo"`
	WriteOffDelayTime int    `json:"writeOffDelayTime"`
	InvalidTime       int    `json:"invalidTime"`
	EquipmentNo       string `json:"equipmentNo"`
	Yccode            string `json:"yccode"`
	ManufacturerId2   string `json:"manufacturerId2"`
	ManufacturerId1   string `json:"manufacturerId1"`
	DeTrueVoice       struct {
		Normal    string `json:"normal"`
		Other     string `json:"other"`
		YearCard1 string `json:"yearCard1"`
		YearCard2 string `json:"yearCard2"`
	} `json:"deTrueVoice"`
	WriteOffUrl  string `json:"writeOffUrl"`
	NumUpUrl     string `json:"numUpUrl"`
	HeartbeatUrl string `json:"heartbeatUrl"`
	FalseVoice3  string `json:"falseVoice3"`
	NumUpTime    int    `json:"numUpTime"`
	FalseVoice4  string `json:"falseVoice4"`
	WelcomeMsg   string `json:"welcomeMsg"`
	TrueVoice4   string `json:"trueVoice4"`
	CheckUrl     string `json:"checkUrl"`
	FalseVoice1  string `json:"falseVoice1"`
	FalseVoice2  string `json:"falseVoice2"`
}
