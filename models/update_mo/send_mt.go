package update_mo

import (
	"fmt"

	"github.com/MobileCPX/PreMondia/models"
	"github.com/MobileCPX/PreMondia/request"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

func SendSMS() {
	// Send 账号
	o := orm.NewOrm()
	var mos []models.Mo
	o.QueryTable("mo").Filter("id__gt", 150).All(&mos)
	for _, mo := range mos {
		fmt.Println(mo.ID)

		var requestData request.MondiaRequestData
		requestData.Message = "Witamy w RedLightVideos. Adres URL to http://www.redlightvideos.com/mm/pl. Twój numer konta to " + mo.SubscriptionID
		requestData.RequestType = "SendSMS"
		requestData.CustomerID = mo.CustomerID
		_, body := request.MondiaHTTPRequest(requestData)
		if string(body) == "OK" {
			logs.Info("订阅成功后发送账号成功")
		} else {
			logs.Info("订阅成功后发送账号失败")
		}
	}

}
