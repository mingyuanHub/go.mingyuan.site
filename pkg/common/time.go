package common

import "time"

func GetNowSecond() int64 {
	return time.Now().Unix()
}

func GetNowMillisecond() int64 {
	return time.Now().UnixNano() / 1e6
}

func GetZeroTime() int64 {
	nowTime := GetNowSecond()
	return nowTime - nowTime%86400
}