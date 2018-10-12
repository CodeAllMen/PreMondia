package routers

import (
	"github.com/MobileCPX/PreMondia/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/notification", &controllers.MondiaNotificationController{})

	beego.Router("/subs/getcust/?:id", &controllers.MondiaGetCustomerController{}) //接收用户id及储存click_id及发起订阅请求
	beego.Router("/subs/res/?:id", &controllers.MondiaSubscribeController{})       //储存订阅通知

	beego.Router("/unsubPin", &controllers.UnsubUserSendMsisdnGetPinController{}) //退订请求发送pin
	beego.Router("/getCust", &controllers.UnsubGetCustomer{})                     // 获取pin之后判断CustomerID 然后退订
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/returnid", &controllers.MondiaReturnIdController{})

	beego.Router("/click", &controllers.AffClickDataController{})         // 每一个小时存一次点击数据
	beego.Router("/update/subdata", &controllers.EveryDayInsertSubData{}) // 每天存一次订阅信息

	beego.Router("/world_play/quality", &controllers.SubscribeQualityController{}) //渠道质量检查
	beego.Router("/aff_data", &controllers.AffController{})                        //获取渠道订阅信息
	beego.Router("/mo_data", &controllers.SearceMoController{})                    //
	beego.Router("/aff_mt", &controllers.SearceAffMtController{})                  //获取渠道推广下发MT的质量，下发成功率
	beego.Router("/get_pubid", &controllers.GetPubIdController{})                  //获取子渠道列表
	beego.Router("/sub/mo_data", &controllers.SubscribeTotalDataController{})      // 查询任意时间订阅任意时间数据
	beego.Router("/sub/everyday/data", &controllers.EverydaySubscribeStatistics{}) // 每日数据统计查询

	beego.Router("/unsub", &controllers.UnsubPage{}) // 退订页面

	beego.Router("/mondia/login/check", &controllers.CheckLoginMsisdnController{})
}
