package backData

import (
	"fmt"
	"github.com/MobileCPX/PreBaseLib/splib/admindata"
	"github.com/MobileCPX/PreMondia/models/mondia"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"strings"
)

var ch = make(chan int, 1)

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

var subIDMap = make(map[string]int)

func SendMo() {
	//UpdateMO()
	o := orm.NewOrm()
	mos := new([]mondia.Mo)
	o.QueryTable("mo").OrderBy("id").All(mos)
	fmt.Println(len(*mos))
	for i, mo := range *mos {
		sendNoti := new(admindata.Notification)
		promoterID := 1
		if mo.SubTime < "2019-04-18" {
			promoterID = 1
		} else if strings.ToLower(mo.ProID) == "zyy" {
			promoterID = 2
		}

		sendNoti.PostbackPrice = 2.8

		sendNoti.OfferID = 0
		sendNoti.SubscriptionID = mo.SubscriptionID
		sendNoti.ServiceID = mo.ProductCode
		sendNoti.ClickID = mo.ClickID
		sendNoti.Msisdn = mo.CustomerID
		logs.Info("mo.CustomerID", mo.CustomerID)
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

		logs.Info("!!!!!!!!!!!!!!!!!", i)
		ch <- 1
		//defer func() {
		//	<-ch
		//}()
		//sendNoti.SendData(admindata.PROD)

		go sendData(sendNoti)

	}

}

func sendData(sendNoti *admindata.Notification) {
	defer func() {
		d := <-ch
		fmt.Println(d)
	}()
	if num, isEixt := subIDMap[sendNoti.SubscriptionID]; isEixt {
		subIDMap[sendNoti.SubscriptionID]++
		logs.Info(sendNoti.SubscriptionID, "!!!!!!!", num)
	}

	sendNoti.SendData(admindata.PROD)
}
