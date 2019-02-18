package sub

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/logs"

	"github.com/MobileCPX/PreMondia/models"
	"github.com/MobileCPX/PreMondia/util"
	"github.com/astaxie/beego/orm"
)

// CheckUserSubStatus 检查用户的订阅状态  传输数据：customerId
func CheckUserSubStatus(customerID string) (isSub bool, subID string) {
	o := orm.NewOrm()
	var mo models.Mo
	o.QueryTable("mo").Filter("customer_id", customerID).OrderBy("-id").One(&mo)
	if mo.ID != 0 {
		isSub = true
		subID = mo.SubscriptionID

	}
	return
}

// CheckTodaySubNum 检查今日订阅数量
func CheckTodaySubNum(serviceID string, limitNum int) (limitSub bool) {
	o := orm.NewOrm()
	var todaySub struct {
		SubNum      int // 今日订阅数量
		PostbackNum int // 回传数量
	}
	nowTime, todayDate := util.GetDatetime()
	SQL := fmt.Sprintf("SELECT COUNT(1) as sub_num,COUNT(CASE WHEN postback_status = 1 and postback_code ='200'"+
		" THEN 1 ELSE null END) as postback_num FROM mo WHERE service_id='%s' and left(sub_time,10) = '%s' ", serviceID, todayDate)
	o.Raw(SQL).QueryRow(&todaySub)
	if todaySub.SubNum >= limitNum {
		limitSub = true
	} else {
		limitSub = CheckDaySubNum(limitNum)
	}
	logs.Info(nowTime, "： 今日订阅数 ", todaySub.SubNum, " 限制订阅数量： ", limitNum, " postback num", todaySub.PostbackNum)
	return
}

func CheckDaySubNum(limitSubNum int) (isLimit bool) {
	o := orm.NewOrm()
	nowTime, todayDate := util.GetDatetime()
	subNum, _ := o.QueryTable("sub_result").Filter("sendtime__gt", todayDate).Count()
	if int(subNum) >= limitSubNum {
		isLimit = true
	}
	logs.Info(nowTime, "： 今日订阅数 ", subNum, " 限制订阅数量： ", limitSubNum)
	return
}

func LimitTenMinutesSubNum(limitSubNum int) (isLimit bool) {
	o := orm.NewOrm()
	nowTime, _ := util.GetDatetime()
	nowMinutes := nowTime[0:15]
	subNum, _ := o.QueryTable("mo").Filter("sub_time__gt", nowMinutes).Count()
	if int(subNum) > limitSubNum {
		logs.Info(nowMinutes, "   十分钟的订阅已经满3个，十分钟之内最多3个订阅")
		return true
	}
	return false
}

// InsertHaveSubData 插入已经订阅的用户数据
func InsertHaveSubData(trackID, customerID string) {
	o := orm.NewOrm()
	var track models.AffTrack
	o.QueryTable("aff_track").Filter("track_id", trackID).One(&track)
	if track.TrackID != 0 {
		var alreadySub models.AlreadySub
		alreadySub.AffName = track.AffName
		alreadySub.ClickID = track.ClickID
		alreadySub.PubID = track.PubID
		alreadySub.CustomerID = customerID
		alreadySub.CanvasID = track.CanvasID
		alreadySub.TrackID = track.TrackID
		alreadySub.Sendtime, _ = util.GetDatetime()
		o.Insert(&alreadySub)
	}
}

func getPostbackURL(affName string) (models.Postback, error) {
	var postback models.Postback
	o := orm.NewOrm()
	o.Using("default")
	affNameLower := strings.ToLower(affName)
	err := o.QueryTable("postback").Filter("aff_name", affNameLower).One(&postback)
	if err != nil {
		logs.Error("Postback url error: aff_name :" + affName + "   " + err.Error())
	}
	return postback, err
}

func postbackRequest(mo models.Mo, PostbackURL string) string { // postback请求
	var urls, code string
	code = "400"
	if PostbackURL != "" {
		urls = strings.Replace(PostbackURL, "##clickid##", mo.ClickID, -1)
		urls = strings.Replace(urls, "##pro_id##", mo.ProID, -1)
		urls = strings.Replace(urls, "##pub_id##", mo.PubID, -1)
		urls = strings.Replace(urls, "##operator##", mo.Operator, -1)
	}
	code = request(PostbackURL)
	if code != "200" {
		logs.Error("postback Error , CustomerId : " + mo.CustomerID + " aff_name : " + mo.AffName + " error " + code)
		for i := 0; i < 3; i++ {
			code = request(PostbackURL)
			if code == "200" {
				break
			}
			time.Sleep(5 * 1e9)
		}
	}
	return code
}

func request(urls string) (code string) {
	resp, err := http.Get(urls)
	if err == nil {
		defer resp.Body.Close()
		code = strconv.Itoa(resp.StatusCode)
	} else {
		code = err.Error()
	}
	return
}
