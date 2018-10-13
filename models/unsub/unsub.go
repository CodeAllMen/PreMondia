package unsub

import (
	"github.com/MobileCPX/PreMondia/models"
	"github.com/MobileCPX/PreMondia/util"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

// GetUnsubMoData  通过customerID 获取mo退订信息
func GetUnsubMoData(customerID string) (*models.Mo, bool) {
	o := orm.NewOrm()
	mo := new(models.Mo)
	var isHas bool
	o.QueryTable("mo").Filter("customer_id", customerID).OrderBy("-id").One(mo)
	if mo.ID != 0 {
		isHas = true
	}
	return mo, isHas
}

// InsertUnsubData 插入退订交易通知
func InsertUnsubData(unsubCharge *models.MondiaCharge) (err error) {
	o := orm.NewOrm()
	_, err = o.Insert(unsubCharge)
	if err != nil {
		logs.Error("InsertUnsubData 插入退订数据失败， subid ：" + unsubCharge.SubscriptionID)
	}
	return
}

// UpdateUnsubMoTable  退订修改MO数据表
func UpdateUnsubMoTable(subID string) (udpateStatus bool) {
	o := orm.NewOrm()
	var mo models.Mo
	nowDatetime, _ := util.GetDatetime()
	o.QueryTable("mo").Filter("subscription_id", subID).One(&mo)
	if mo.ID != 0 {
		mo.SubStatus = 0
		mo.UnsubTime = nowDatetime
		_, err := o.Update(&mo)
		if err == nil {
			udpateStatus = true
		}
	}
	return udpateStatus
}
