package util

import "time"

// GetDatetime 获取时间
func GetDatetime() (string, string) {
	time.LoadLocation("UTC")
	h, _ := time.ParseDuration("1h")
	nowDate := time.Now().UTC().Add(2 * h).Format("2006-01-02")
	nowTime := time.Now().UTC().Add(2 * h).Format("2006-01-02 15:04:05")
	return nowTime, nowDate
}
