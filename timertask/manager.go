package timertask

import "emcs-relay-go/utils"

// 1.一个单线程负责多个计时
// 2.执行符合时间间隔的事件
func RunTimer() {
	keepLiveTime := 10
	keepUploadPassed := 10
	hardwareQuery := 10
	utils.Log.Info(keepLiveTime)
	utils.Log.Info(keepUploadPassed)
	utils.Log.Info(hardwareQuery)

}
