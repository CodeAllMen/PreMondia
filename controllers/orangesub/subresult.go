package orangesub

import (
	"github.com/MobileCPX/PreMondia/models"
	"github.com/MobileCPX/PreMondia/models/sub"
	"github.com/MobileCPX/PreMondia/util"
	"github.com/astaxie/beego"
)

type MondiaSubscribeController struct {
	beego.Controller
}

func (c *MondiaSubscribeController) Get() {
	trackID := c.Ctx.Input.Param(":trackID")
	status := c.GetString("status")
	customerID := c.GetString("customerId")
	subcriptionID := c.GetString("subId")
	nextAction := c.GetString("nextAction")
	subStatus := c.GetString("subStatus")
	nextActionDate := c.GetString("nextActionDate")
	errorCode := c.GetString("errorCode")
	errorDesc := c.GetString("errorDesc")
	viewName := c.GetString("viewName")
	subNotification := new(models.MdSubscribe)
	subNotification.TrackID = trackID
	subNotification.SubStatus = subStatus
	subNotification.CustomerID = customerID
	subNotification.Status = status
	subNotification.ErrorCode = errorCode
	subNotification.NextAction = nextAction
	subNotification.NextActionDate = nextActionDate
	subNotification.SubscriptionID = subcriptionID
	subNotification.ViewName = viewName
	subNotification.ErrorDesc = errorDesc
	sub.InsertSubscribe(subNotification)
	// 3001 已经注册过     "SUCCESS" 表示订阅成功
	if (status == "SUCCESS" || errorCode == "3001") && subStatus == "ACTIVE" {
		util.HttpRequest(subNotification.SubscriptionID, "register", "video", subNotification.SubscriptionID, "")
		c.Redirect("http://www.redlightvideos.com/mm/pl?sub="+subNotification.SubscriptionID, 302)
	} else if subStatus == "SUSPENDED" {
		util.HttpRequest(subNotification.SubscriptionID, "register", "video", subNotification.SubscriptionID, "")
		c.Redirect("http://www.redlightvideos.com/mm/pl?sub="+subNotification.SubscriptionID, 302)
	} else if subStatus == "UNSUBSCRIBED" {
		c.Redirect("https://www.google.com", 302)
	} else {
		c.Redirect("http://www.redlightvideos.com/lp/mm/pl/index.html?affName=Slef", 302)
	}

	if status == "SUCCESS" || status == "SUSPENDED" {
		var subResult models.SubResult
		subResult.Status = subStatus
		subResult.Msisdn = customerID
		// subResult.OperatorID = c.GetString("operator_id")
		subResult.SubsID = subcriptionID
		subResult.TrackID = trackID
		sub.InsertSubResult(subResult)
	}
}
