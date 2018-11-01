package routers

import (
	"github.com/MobileCPX/PreMondia/controllers/checksubnum"
	"github.com/MobileCPX/PreMondia/controllers/notification"
	"github.com/MobileCPX/PreMondia/controllers/orangesub"
	"github.com/MobileCPX/PreMondia/controllers/searchAPI"
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
	// 通过订阅ID 退订服务
	beego.Router("/unsub/subid", &unsub.SubIDUnsubRequest{})

	// 查询数据接口
	beego.Router("/aff_data", &searchAPI.AffController{}) // 查询网盟转化数据
	beego.Router("/aff_mt", &searchAPI.SearceAffMtController{})
	beego.Router("/quality", &searchAPI.SubscribeQualityController{})                 //渠道质量检查
	beego.Router("/sub/mo_data", &searchAPI.AnytimeProfitAndLossController{})         // 查询任意时间订阅任意时间数据
	beego.Router("/sub/everyday/data", &searchAPI.EverydaySubscribeDataControllers{}) // 每日数据统计查询

	// CheckSubNum 检查订阅数量
	beego.Router("/check/sub/num", &checksubnum.CheckSubNum{})
}
