package db

type EventLog struct {
	Id         uint   `gorm:"primaryKey"`
	Tag        string `json:"tag"`
	Content    string `json:"content"`
	DeviceName string `json:"deviceName"`
	Time       string `json:"time"`
}

func AddEvent(log *EventLog) error {
	return DB.Save(log).Error
}
func GetEvents(event *[]EventLog) error {
	return DB.Raw("SELECT * FROM event_logs ORDER BY id LIMIT 10").Scan(event).Error
}
