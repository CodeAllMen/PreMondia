package sub

import (
	"fmt"

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
func CheckTodaySubNum(limitNum int) (limitSub bool) {
	o := orm.NewOrm()
	var todaySub struct {
		SubNum int // 今日订阅数量
	}
	_, todayDate := util.GetDatetime()
	SQL := fmt.Sprintf("SELECT COUNT(1) as sub_num FROM mo WHERE left(sub_time,10) = '%s' ", todayDate)
	o.Raw(SQL).QueryRow(&todaySub)
	fmt.Println(todaySub.SubNum)
	if todaySub.SubNum >= limitNum {
		limitSub = true
	}
	logs.Info(todayDate, "： 今日订阅数 ", todaySub.SubNum, " 限制订阅数量： ", limitNum)
	return
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
