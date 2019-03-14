package routers

import (
	"github.com/MobileCPX/PreMondia/controllers"
	"github.com/MobileCPX/PreMondia/controllers/mondia"
	"github.com/MobileCPX/PreMondia/controllers/searchAPI"
	"github.com/astaxie/beego"
)

func init() {

	//// 跳转到AOC页面 POST
	//beego.Router("/payment", &orangesub.GetPostRequestControlelr{})
	//// 订阅续订退订通知
	//beego.Router("/notification", &notification.MondiaNotificationController{})
	//
	//// 记录每次点击
	//beego.Router("/returnid", &orangesub.LPTrackClickControllers{})
	//// GetCustomer 通知
	//beego.Router("/subs/getcust/?:trackID", &orangesub.GetCustomerControllers{})
	//// 订阅之后的回调
	//beego.Router("/subs/res/?:trackID", &orangesub.MondiaSubscribeController{})
	//
	////退订请求
	//beego.Router("/unsub", &unsub.UnsubPage{})
	//beego.Router("/unsubPin", &unsub.SendPINControllers{}) //退订请求发送pin
	//beego.Router("/getCust", &unsub.UnsubGetCustomer{})    // 获取pin之后判断CustomerID 然后退订
	//// 通过订阅ID 退订服务
	//beego.Router("/unsub/subid", &unsub.SubIDUnsubRequest{})

	// 查询数据接口
	beego.Router("/aff_data", &searchAPI.AffController{}) // 查询网盟转化数据
	beego.Router("/aff_mt", &searchAPI.SearceAffMtController{})
	beego.Router("/quality", &searchAPI.SubscribeQualityController{})                 //渠道质量检查
	beego.Router("/sub/mo_data", &searchAPI.AnytimeProfitAndLossController{})         // 查询任意时间订阅任意时间数据
	beego.Router("/sub/everyday/data", &searchAPI.EverydaySubscribeDataControllers{}) // 每日数据统计查询

	beego.Router("/", &controllers.MainController{})
	// CheckSubNum 检查订阅数量
	beego.Router("/check/sub/num", &mondia.SubFlowController{}, "Get:CheckTodaySubNum")

	// 跳转到AOC页面 POST
	//beego.Router("/payment", &mondia.SubFlowController{}, "Post:GetCustomerRedirect")
	// 订阅续订退订通知
	beego.Router("/notification", &mondia.NotificationController{})

	// 记录每次点击
	beego.Router("/returnid", &mondia.SubFlowController{}, "Get:AffTrack")
	// GetCustomer 通知
	beego.Router("/subs/getcust/?:trackID", &mondia.SubFlowController{}, "Get:CustomerResultAndStartSub")
	// 订阅之后的回调
	beego.Router("/subs/res/?:trackID", &mondia.SubFlowController{}, "Get:SubResult")

	//退订请求
	beego.Router("/unsub", &mondia.UnsubController{}, "Get:UnsubPage")
	beego.Router("/unsubPin", &mondia.UnsubController{}, "Post:UnsubSendPin") //退订请求发送pin
	beego.Router("/getCust", &mondia.UnsubController{}, "Post:UnsubRequest")  // 获取pin之后判断CustomerID 然后退订
	// 通过订阅ID 退订服务
	//beego.Router("/unsub/subid", &unsub.SubIDUnsubRequest{})
}
