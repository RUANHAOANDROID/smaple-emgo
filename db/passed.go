package db

// Passed 过闸统计
type Passed struct {
	ID uint `json:"id"gorm:"primaryKey"`
	// 标签
	DId string `json:"did"`
	// 编号
	Count string `json:"count"`
	//最后过闸时间
	LastTime string `json:"last_time"`
}
