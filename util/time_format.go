package util

import "time"

// GetFormatTime 获取当前格式化时间
// nowDatetime 当前详细时间  nowDate当前日期
func GetFormatTime() (nowDatetime, nowDate string) {
	time.LoadLocation("UTC")
	//h, _ := time.ParseDuration("1h")
	nowDatetime = time.Now().UTC().Format("2006-01-02 15:04:05")
	nowDate = time.Now().UTC().Format("2006-01-02")
	return
}

// GetFormatHoursTime 获取当前格式化时间  格式为 2006-01-02 15
func GetFormatHoursTime() string {
	time.LoadLocation("UTC")
	newFormat := time.Now().UTC().Format("2006-01-02 15")
	return newFormat
}

// GetLastMonth 获取上个月月份
func GetLastMonth(num int) string {
	year, month, _ := time.Now().Date()
	thisMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	mouth := thisMonth.AddDate(0, num, 0).Format("2006-01")
	return mouth
}

// GetDateList 获取时间列表
func GetDateList(startDate, endDate string) (dateList []string) {
	time.LoadLocation("UTC")
	d, _ := time.ParseDuration("24h")
	start, _ := time.Parse("2006-01-02", startDate)
	for i := 1; ; i++ {
		if start.Format("2006-01-02") <= endDate {
			dateList = append(dateList, start.Format("2006-01-02"))
			start = start.Add(d)
		} else {
			break
		}
	}
	return
}
