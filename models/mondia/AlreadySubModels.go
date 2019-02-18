package mondia

import (
	"github.com/MobileCPX/PreMondia/util"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

// AlreadySub // 网盟重复推送已经订阅过的用户统计
type AlreadySub struct {
	ID         int64  `orm:"pk;auto;column(id)"` //自增ID
	CustomerID string `orm:"column(customer_id)"`
	ServiceID  string `orm:"column(service_id)"`
	AffName    string `orm:"column(aff_name);size(30)"`  // 网盟名称
	PubID      string `orm:"column(pub_id);size(100)"`   // 子渠道
	ProID      string `orm:"column(pro_id);size(30)"`    // 服务id（可有可无）
	ClickID    string `orm:"column(click_id);size(100)"` // 点击
	TrackID    int64  `orm:"column(track_id)"`           //自增ID
	Sendtime   string `orm:"column(sendtime);size(30)"`  // 扣费时间
}

func InsertAlreadSubData(track *AffTrack) {
	alreadySub := new(AlreadySub)
	o := orm.NewOrm()
	alreadySub.ServiceID = track.ServiceID
	alreadySub.TrackID = track.TrackID
	alreadySub.CustomerID = track.CustomerID
	alreadySub.AffName = track.AffName
	alreadySub.PubID = track.PubID
	alreadySub.Sendtime, _ = util.GetNowTimeFormat()
	_, err := o.Insert(alreadySub)
	if err != nil {
		logs.Error("InsertAlreadSubData  插入重复订阅数据失败，trackID: ", track.TrackID, "ERROR: ", err.Error())
	}
}
