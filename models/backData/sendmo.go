package backData

import (
	"github.com/MobileCPX/PreBaseLib/sp/admindata"
	"github.com/MobileCPX/PreMondia/models/mondia"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"strings"
)

var ServiceCamp = map[string]int{"REDLIGHTVIDEOS": 21, "GOGAMEHUB": 19, "IFUNNY": 20, "PREPRON4K": 22, "PLEASURECLICK": 23}

func UpdateMO() {
	o := orm.NewOrm()
	mos := new([]mondia.Mo)
	o.QueryTable("mo").Filter("aff_name", "hyperclick").All(mos)
	for _, mo := range *mos {
		mo.ProID = "zyy"
		o.Update(&mo)
	}

	mos = new([]mondia.Mo)
	o.QueryTable("mo").Filter("aff_name", "skrmobi").All(mos)
	for _, mo := range *mos {
		mo.ProID = "jg"
		o.Update(&mo)
	}

	mos = new([]mondia.Mo)
	o.QueryTable("mo").Filter("sub_time__gt", "2019-04-18").Filter("pro_id", "").All(mos)
	for _, mo := range *mos {
		if mo.ID%2 == 0 {
			mo.ProID = "zyy"
		} else {
			mo.ProID = "jq"
		}
		o.Update(&mo)
	}
}

func SendMo() {
	UpdateMO()
	o := orm.NewOrm()
	mos := new([]mondia.Mo)
	o.QueryTable("mo").OrderBy("id").All(mos)
	for _, mo := range *mos {
		sendNoti := new(admindata.Notification)
		promoterID := 1
		if mo.SubTime < "2019-04-18" {
			promoterID = 1
		} else if strings.ToLower(mo.ProID) == "zyy" {
			promoterID = 2
		}

		sendNoti.OfferID = 0
		sendNoti.SubscriptionID = mo.SubscriptionID
		sendNoti.ServiceID = mo.ProductCode
		sendNoti.ClickID = mo.ClickID
		sendNoti.Msisdn = mo.CustomerID
		logs.Info("mo.CustomerID",mo.CustomerID)
		sendNoti.CampID = ServiceCamp[mo.ProductCode]
		sendNoti.PubID = mo.PubID
		sendNoti.PostbackStatus = mo.PostbackStatus
		sendNoti.PostbackMessage = mo.PostbackCode
		sendNoti.TransactionID = ""
		sendNoti.AffName = mo.AffName
		if sendNoti.AffName == "" {
			sendNoti.AffName = "未知"
		}
		sendNoti.Operator = mo.Operator
		sendNoti.PromoterID = promoterID
		sendNoti.Sendtime = mo.SubTime
		sendNoti.NotificationType = "SUB"
		sendNoti.SendData()

	}

}
