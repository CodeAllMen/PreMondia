package notification

import (
	"encoding/xml"
	"io/ioutil"

	"github.com/astaxie/beego/logs"

	"github.com/MobileCPX/PreMondia/models"
	"github.com/MobileCPX/PreMondia/models/notification"

	"github.com/astaxie/beego"
)

// MondiaNotificationController 接收订阅退订续订通知post请求
type MondiaNotificationController struct {
	beego.Controller
}

//Post 接收订阅退订续订通知post请求
func (c *MondiaNotificationController) Post() {
	body := c.Ctx.Request.Body
	data, _ := ioutil.ReadAll(body)
	logs.Info("notification", string(data))
	modiaNotification := models.Notification{}
	err := xml.Unmarshal(data, &modiaNotification)
	mo := new(models.Mo)
	notificationType := ""
	if err == nil { // 更新mo表（新增订阅，退订，续订）
		mo.Price = modiaNotification.Price
		mo.Operator = modiaNotification.Operator
		mo.SubscriptionID = modiaNotification.SubscriptionID
		mo.ServiceID = modiaNotification.ServiceID
		mo.CustomerID = modiaNotification.CustomerID
		mo.Channel = modiaNotification.Channel
		mo.PackageCode = modiaNotification.PackageCode
		mo.ProductCode = modiaNotification.ProductCode
		notificationType = notification.UpdateOrInsertMo(modiaNotification.Action, modiaNotification.SubscriptionStatus, modiaNotification.Price, mo)
	}
	modiaNotification.NotificationType = notificationType
	notification.InsertCharge(modiaNotification)
	c.Ctx.WriteString("ok")

	// 打印通知信息
	logs.Info("notification: ", string(data))
}
