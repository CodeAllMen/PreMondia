package sub

import (
	"github.com/MobileCPX/PreMondia/models"
	"github.com/astaxie/beego/orm"
)

// CheckUserSubStatus 检查用户的订阅状态  传输数据：customerId
func CheckUserSubStatus(customerID string) (isSub bool) {
	o := orm.NewOrm()
	var mo models.Mo
	o.QueryTable("mo").Filter("customer_id", customerID).One(&mo)
	if mo.ID != 0 {
		isSub = true
	}
	return
}
