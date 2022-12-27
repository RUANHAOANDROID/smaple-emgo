package api

import (
	"emcs-relay-go/utils"
	"regexp"
)

//2022-12-12 17:50:37 收到数据：$F384605470533333459544638P21/@ok\@$E
//2022-12-12 17:50:54 发送数据：$F384605470533333459544638111A01000/&\&$E

func isPassed(msg string) (bool, error) {
	pattern2 := "^\\$F[0-9]{24}P21\\/@ok\\\\@\\$E"
	match, err := regexp.MatchString(pattern2, msg)
	utils.Log.Info("Match: ", match, " Error: ", err)
	return match, err
}

func isCheck(msg string) (bool, error) {
	pattern2 := "^\\$F[0-9]{24}[0-9]{3}\\/@.*\\\\@\\$E"
	match, err := regexp.MatchString(pattern2, msg)
	utils.Log.Info("Match: ", match, " Error: ", err)
	return match, err
}
func packOpenGateMsg() {

}
