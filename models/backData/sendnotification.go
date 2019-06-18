package backData

import (
	"github.com/MobileCPX/PreBaseLib/splib/admindata"
	"github.com/MobileCPX/PreMondia/models/mondia"
	"github.com/astaxie/beego/orm"
)

func SendNotification() {
	o := orm.NewOrm()
	notifis := new([]mondia.Notification)
	notificationTypeList := []string{"FAILED_MT", "MT_FAILED", "MT_SUCCESS", "SUCCESS_MT", "UNSUB"}
	o.QueryTable("notification").Filter("notification_type__in", notificationTypeList).OrderBy("id").All(notifis)
	for _, one := range *notifis {
		sendNoti := new(admindata.Notification)
		notificationType := one.NotificationType
		if notificationType == "MT_SUCCESS" {
			notificationType = "SUCCESS_MT"
		} else if notificationType == "MT_FAILED" {
			notificationType = "FAILED_MT"
		}

		sendNoti.SubscriptionID = one.SubscriptionID
		sendNoti.ServiceID = one.ProductCode
		sendNoti.CampID = ServiceCamp[one.ProductCode]

		sendNoti.TransactionID = one.TransactionID

		sendNoti.Sendtime = one.Sendtime

		sendNoti.NotificationType = notificationType

		if sendNoti.Sendtime > "2018-12-15" {
			ch <- 1
			defer func() {
				<-ch
			}()
			//sendNoti.SendData(admindata.PROD)


			go sendData(sendNoti)
		}

	}

}
