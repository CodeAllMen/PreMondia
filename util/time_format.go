package util

import "time"


func GetDatetime()(string,string){
	time.LoadLocation("UTC")
	h, _ := time.ParseDuration("1h")
	billDate := time.Now().UTC().Add(2 * h).Format("20060102")
	now_time := time.Now().UTC().Add(2 * h).Format("2006-01-02 15:04:05")
	return billDate,now_time
}
