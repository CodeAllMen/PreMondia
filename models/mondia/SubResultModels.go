package mondia

import (
	"github.com/MobileCPX/PreMondia/util"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type MdSubscribe struct {
	ID             int64 `orm:"pk;auto;column(id)"` //自增ID
	Sendtime       string
	TrackID        string `orm:"column(track_id)"`
	Status         string
	CustomerID     string `orm:"column(customer_id)"`
	SubscriptionID string `orm:"column(subscription_id)"`
	NextAction     string
	SubStatus      string
	NextActionDate string
	ErrorCode      string
	ErrorDesc      string
	ServiceType    string
	ViewName       string
}



func (subResult *MdSubscribe) Insert() {
	o := orm.NewOrm()
	nowTime, _ := util.GetNowTimeFormat()
	subResult.Sendtime = nowTime
	_, err := o.Insert(subResult)
	if err != nil {
		logs.Error("MdSubscribe Insert  存入订阅结果通知失败，ERROR: ", err.Error())
	}
}
