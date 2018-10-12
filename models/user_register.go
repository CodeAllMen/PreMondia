package models

import "github.com/astaxie/beego/orm"

func CheckUserClickId(clickId string) (bool, MdSubscribe) {
	o := orm.NewOrm()
	var subscribe_data MdSubscribe
	o.QueryTable("md_subscribe").Filter("click_id", clickId).One(&subscribe_data)
	if subscribe_data.Id != 0 {
		return true, subscribe_data
	} else {
		return false, subscribe_data
	}
}
