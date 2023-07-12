package utils

import "time"

func GetCurrentTime() (int64, string) {
	now := time.Now()
	timestamp := time.Now().Unix()
	nowDate := now.Format("2006-01-02")
	return timestamp, nowDate
}
