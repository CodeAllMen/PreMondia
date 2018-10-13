package notification

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"

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
	modiaNotification := models.Notification{}
	fmt.Println(string(data))
	err := xml.Unmarshal(data, &modiaNotification)
	fmt.Println(&modiaNotification)
	mo := new(models.Mo)
	if err == nil { // 更新mo表（新增订阅，退订，续订）
		mo.Price = modiaNotification.Price
		mo.Operator = modiaNotification.Operator
		mo.SubscriptionID = modiaNotification.SubscriptionID
		mo.ServiceID = modiaNotification.ServiceID
		mo.CustomerID = modiaNotification.CustomerID
		mo.Channel = modiaNotification.Channel
		mo.PackageCode = modiaNotification.PackageCode
		mo.ProductCode = modiaNotification.ProductCode
		notification.UpdateOrInsertMo(modiaNotification.Action, modiaNotification.SubscriptionStatus, modiaNotification.Price, mo)
	}
	notification.InsertCharge(modiaNotification)
	c.Ctx.WriteString("ok")
}
