package routers

import (
	"github.com/MobileCPX/PreMondia/controllers/notification"
	"github.com/MobileCPX/PreMondia/controllers/orangesub"
	"github.com/MobileCPX/PreMondia/controllers/unsub"

	"github.com/astaxie/beego"
)

func init() {

	// 跳转到AOC页面 POST
	beego.Router("/payment", &orangesub.GetPostRequestControlelr{})
	// 订阅续订退订通知
	beego.Router("/notification", &notification.MondiaNotificationController{})

	// 记录每次点击
	beego.Router("/returnid", &orangesub.LPTrackClickControllers{})
	// GetCustomer 通知
	beego.Router("/subs/getcust/?:trackID", &orangesub.GetCustomerControllers{})
	// 订阅之后的回调
	beego.Router("/subs/res/?:trackID", &orangesub.MondiaSubscribeController{})

	//退订请求
	beego.Router("/unsub", &unsub.UnsubPage{})
	beego.Router("/unsubPin", &unsub.SendPINControllers{}) //退订请求发送pin
	beego.Router("/getCust", &unsub.UnsubGetCustomer{})    // 获取pin之后判断CustomerID 然后退订
}
