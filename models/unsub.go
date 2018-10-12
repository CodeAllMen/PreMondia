package models

import (
	"strconv"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

// InsertPinData 存入pin信息
func InsertPinData(pinData *UnsubPin) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(pinData)
	if err != nil {
		logs.Error("InsertPinData error")
		return
	}
	return
}

// CheckPIN 检查用户输入的PIN
func CheckPIN(pin, id string) (msisdn string, err error) {
	o := orm.NewOrm()
	var pinData UnsubPin
	idInt64, err := strconv.ParseInt(id, 10, 64)
	if err == nil {
		o.QueryTable("unsub_pin").Filter("id", idInt64).Filter("pin", pin).One(&pinData)
		if pinData.Id != 0 {
			msisdn = pinData.Msisdn
			return
		}
	}
	return
}

// CustomerToGetSubID 根据customerID 在mo表里面获取订阅id
func CustomerToGetSubID(customerID, msisdn string) (subID string) {
	o := orm.NewOrm()
	var mo MonMo
	o.QueryTable("mon_mo").Filter("customer_id", customerID).Filter("sub_status", 1).One(&mo)
	if mo.Id != 0 {
		subID = mo.SubscriptionId
		mo.Msisdn = msisdn
		o.Update(&mo)
	}
	return
}

// InsertUnsubData
func InsertUnsubData(unsubCharge *MondiaCharge) (err error) {
	o := orm.NewOrm()
	_, err = o.Insert(unsubCharge)
	if err != nil {
		logs.Error("InsertUnsubData 插入退订数据失败， subid ：" + unsubCharge.SubscriptionID)
	}
	return
}
