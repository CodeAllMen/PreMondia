package routers

import (
	"github.com/MobileCPX/PreMondia/controllers"
	"github.com/MobileCPX/PreMondia/controllers/checksubnum"
	"github.com/MobileCPX/PreMondia/controllers/mondia"
	"github.com/MobileCPX/PreMondia/controllers/searchAPI"
	"github.com/astaxie/beego"
)

func init() {

	// 查询数据接口
	beego.Router("/aff_data", &searchAPI.AffController{}) // 查询网盟转化数据
	beego.Router("/aff_mt", &searchAPI.SearceAffMtController{})
	beego.Router("/quality", &searchAPI.SubscribeQualityController{})                 // 渠道质量检查
	beego.Router("/sub/mo_data", &searchAPI.AnytimeProfitAndLossController{})         // 查询任意时间订阅任意时间数据
	beego.Router("/sub/everyday/data", &searchAPI.EverydaySubscribeDataControllers{}) // 每日数据统计查询

	beego.Router("/", &controllers.MainController{})
	// CheckSubNum 检查订阅数量
	beego.Router("/check/sub/num", &checksubnum.CheckSubNum{})

	// 订阅续订退订通知
	beego.Router("/mondia/notification", &mondia.NotificationController{})

	// 记录每次点击
	beego.Router("/returnid", &mondia.SubFlowController{}, "Get:AffTrack")
	// 跳转到AOC页面 POST
	beego.Router("/customer/req/:trackID", &mondia.SubFlowController{}, "Get:GetCustomerRedirect")
	// GetCustomer 通知
	beego.Router("/get/customer/:trackID", &mondia.SubFlowController{}, "Get:CustomerResultAndStartSub")
	// 订阅之后的回调
	beego.Router("/get/sub_result/:trackID", &mondia.SubFlowController{}, "Get:SubResult")

	// 退订请求
	beego.Router("/unsub/cookie/sub_id/:serviceID", &mondia.UnsubController{}, "Get:UnsubByCookie")
	beego.Router("/unsub/get/customer/:serviceID", &mondia.UnsubController{}, "Get:UnsubGetCustomerIDResult")

	beego.Router("/unsub/:serviceID", &mondia.UnsubController{}, "Get:UnsubPage")
	beego.Router("/unsubPin/:serviceID", &mondia.UnsubController{}, "Post:UnsubSendPin") // 退订请求发送pin
	beego.Router("/getCust", &mondia.UnsubController{}, "Post:UnsubRequest")             // 获取pin之后判断CustomerID 然后退订
	// 通过订阅ID 退订服务
}
