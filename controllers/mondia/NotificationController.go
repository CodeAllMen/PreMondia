package mondia

import (
	"encoding/xml"
	"github.com/MobileCPX/PreMondia/models/mondia"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"strconv"
)

type NotificationController struct {
	BaseController
}

func (c *NotificationController) Post() {
	body := c.Ctx.Request.Body
	data, _ := ioutil.ReadAll(body)
	// 打印通知信息
	logs.Info("notification", string(data))
	modiaNotification := new(mondia.Notification)
	err := xml.Unmarshal(data, modiaNotification)
	mo := new(mondia.Mo)
	notificationType := ""
	if err == nil { // 更新mo表（新增订阅，退订，续订）
		if modiaNotification.Action == "SUBSCRIBE" {
			track := new(mondia.AffTrack)
			if modiaNotification.CampaignID == "" {
				subResult := new(mondia.MdSubscribe)
				subResult.GetSubResultBySubID(modiaNotification.SubscriptionID)
				if subResult.ID != 0 && err == nil {
					trackIntID, _ := strconv.Atoi(subResult.TrackID)
					_ = track.GetAffTrackByTrackID(int64(trackIntID))
				}
			} else {
				logs.Info("modiaNotification.CampaignID  不为空：", modiaNotification.CampaignID)
				trackIntID, _ := strconv.Atoi(modiaNotification.CampaignID)
				_ = track.GetAffTrackByTrackID(int64(trackIntID))
			}

			//track := new(mondia.AffTrack)
			////_ = track.GetAffTrackByCustomerID(modiaNotification.CustomerID)
			//subResult := new(mondia.MdSubscribe)
			//subResult.GetSubResultBySubID(modiaNotification.SubscriptionID)

			//if subResult.TrackID != "" {
			//	trackIDInt, err := strconv.Atoi(subResult.TrackID)
			//	if err == nil {
			//		_ = track.GetAffTrackByTrackID(int64(trackIDInt))
			//	}
			//}

			notificationType = c.NewInsertMo(modiaNotification, track)
		} else if modiaNotification.Action == "UNSUBSCRIBE" {
			notificationType, _ = mo.UnsubUpdateMo(modiaNotification.SubscriptionID)
		} else if modiaNotification.Action == "RENEW" || modiaNotification.Action == "RETRY" {
			if modiaNotification.SubscriptionStatus == "ACTIVE" {
				notificationType, _ = mo.SuccessMTUpdateMO(modiaNotification.SubscriptionID)
			}
		} else if modiaNotification.Action == "STATUS_CHANGE" {
			if modiaNotification.SubscriptionStatus == "SUSPENDED" {
				notificationType, _ = mo.FailedMTUpdateMo(modiaNotification.SubscriptionID)
			}
		}
	}

	modiaNotification.NotificationType = notificationType
	err = modiaNotification.Insert()

	if err == nil {
		c.Ctx.WriteString("ok")
	} else {
		c.Ctx.WriteString("ERROR")
	}
}
