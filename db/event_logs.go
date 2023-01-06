package db

import "emcs-relay-go/utils"

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
	return DB.Raw("SELECT * FROM event_logs ORDER BY id LIMIT 20").Scan(event).Error
}
func GetEventsByDay(events *[]EventLog, day string, device string, offset int, limit int) (int64, error) {
	eventTab := DB.Model(&EventLog{})
	eventTab.Where("date(time)=?", day)
	if device != "全部" {
		eventTab.Where("device_name=?", device)
	}
	var count int64 = 0
	err := eventTab.Count(&count).Error
	if err != nil {
		utils.Log.Error(err.Error())
		return count, nil
	}
	err = eventTab.Offset(offset).Limit(limit).Find(&events).Error //查询pageindex页的数据
	return count, err
}
func GetEventsPage(count *int64, events *[]EventLog, day string, device string, offset int, limit int) error {
	eventTab := DB.Model(&EventLog{})
	eventTab.Where("date(time)=?", day)
	if device != "全部设备" {
		eventTab.Where("device_name=?", device)
	}
	err := eventTab.Count(count).Error
	if err != nil {
		utils.Log.Error(err.Error())
	}
	//err = eventTab.Where("id BETWEEN ? AND ?", offset, limit).Find(&events).Error
	err = eventTab.Order("id DESC").Offset(offset).Limit(limit).Find(&events).Error //查询pageindex页的数据
	return err
}
