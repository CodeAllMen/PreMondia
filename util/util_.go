package util

import (
	"fmt"
	"net/http"
	"reflect"
	"sort"
	"time"

	"github.com/astaxie/beego/logs"
)

func UserRequest(CustomerId, username, request_type string) {
	var url string
	if request_type == "register_game" {
		url = fmt.Sprintf("http://www.gogamehub.com/addusername?username=%s", username)
	} else {
		url = fmt.Sprintf("http://www.prepornvideo.com/addsubs?phone=%s&sign=wp", username)
	}
	resp, err := http.Get(url)
	if err != nil {
		logs.Error(fmt.Sprintf("CustomerId: %s   username: %s %s error  url:%s", CustomerId, username, request_type, url))
	} else {
		resp.Body.Close()
	}

}

func GetFormatTime() (string, string) {
	time.LoadLocation("UTC")
	//h, _ := time.ParseDuration("1h")
	newFormat := time.Now().UTC().Format("2006-01-02 15:04:05")
	newDate := time.Now().UTC().Format("2006-01-02")
	return newFormat, newDate
}

func Duplicate(a []string) (ret []string) { // 列表去重
	sort.Strings(a)
	va := reflect.ValueOf(a)
	for i := 0; i < va.Len(); i++ {
		if i > 0 && reflect.DeepEqual(va.Index(i-1).Interface(), va.Index(i).Interface()) {
			continue
		}
		ret = append(ret, va.Index(i).String())
	}
	return ret
}

func GetLastMonth(num int) string {
	year, month, _ := time.Now().Date()
	thisMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	mouth := thisMonth.AddDate(0, num, 0).Format("2006-01")
	return mouth
}

func GetFormatHoursTime() string {
	time.LoadLocation("UTC")
	newFormat := time.Now().UTC().Format("2006-01-02 15")
	return newFormat
}

//HttpRequest 注册
func HttpRequest(subID, types, period, CustomerId, unsubtime string) {
	urls := ""
	if types == "register" {
		urls = fmt.Sprintf("http://www.redlightvideos.com/addsubs?uiid=%s&sign=mondia", subID)
	}
	if types == "delete" {
		urls = fmt.Sprintf("www.redlightvideos.com/delete/user?phone=%s&time=%s", subID, unsubtime)
	}
	resp, err := http.Get(urls)
	if err == nil {
		logs.Info(fmt.Sprintf("HttpRequest Success %s service %s subID: %s  CustomerId: %s ", types, period, subID, CustomerId))
		resp.Body.Close()
	} else {
		logs.Error(fmt.Sprintf("HttpRequest Failed %s service %s subID: %s  CustomerId: %s     error: %s ", types, period, subID, CustomerId, err.Error()))
	}
}
