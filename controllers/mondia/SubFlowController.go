package mondia

import (
	"fmt"
	"github.com/MobileCPX/PreMondia/enums"
	"github.com/MobileCPX/PreMondia/models/mondia"
	"github.com/MobileCPX/PreMondia/util"
	"github.com/astaxie/beego/logs"
	"net/url"
	"strconv"
)

// LPTrackControllers 存储点击
type SubFlowController struct {
	BaseController
}

// AffTrack 存储点击
func (c *SubFlowController) AffTrack() {
	logs.Info(c.Ctx.Input.URL())
	track := new(mondia.AffTrack)
	track.ServiceID = c.GetString("type")
	track.AffName = c.GetString("affName")
	track.ClickID = c.GetString("clickId")
	track.PubID = c.GetString("pubId")
	track.ProID = c.GetString("proId")
	track.IP = util.GetIPAddress(c.Ctx.Request)
	track.Refer = c.Ctx.Input.Refer()
	track.UserAgent = c.Ctx.Input.UserAgent() //用户设备信息

	// 插入点击数据
	trackID, err := track.Insert()
	// 获取今日订阅数量，判断是否超过订阅限制
	todaySubNum, err1 := mondia.GetTodayMoNum(track.ServiceID)

	if err != nil || err1 != nil || int(todaySubNum) >= enums.DayLimitSub || track.AffName == "pocketmedia" {
		c.Ctx.WriteString("false")
	} else {
		c.Ctx.WriteString(strconv.FormatInt(trackID, 10))
	}
}

func (c *SubFlowController) GetCustomerRedirect() {
	trackID := c.Ctx.Input.Param(":trackID") // id
	track := new(mondia.AffTrack)
	trackIDInt, err := strconv.Atoi(trackID)
	if err != nil {
		logs.Info("trackID string to int failed ,redirect google page")
		c.redirect("http://www.google.com")
	}

	_ = track.GetAffTrackByTrackID(int64(trackIDInt))
	serviceConfig := c.getServiceConfig(track.ServiceID)

	customerRedirectURL := serviceConfig.MondiaRequestURL + "?method=getcustomer&merchantId=" +
		serviceConfig.MrchantID + "&redirect=" + url.QueryEscape(serviceConfig.GetCustomerCallbackURL+trackID) + "&operatorId=" + serviceConfig.OperatorID
	fmt.Println(customerRedirectURL)
	c.redirect(customerRedirectURL)
}

func (c *SubFlowController) CustomerResultAndStartSub() {
	track := new(mondia.AffTrack)
	trackID := c.Ctx.Input.Param(":trackID") // id
	track.Status = c.GetString("status")
	track.CustomerID = c.GetString("customerId")
	track.Operator = c.GetString("operator")
	track.ErrorDesc = c.GetString("errorDesc")
	track.ErrorCode = c.GetString("errorCode")
	//data, _ := json.Marshal(track)
	//fmt.Println(string(data))
	trackIDInt, err := strconv.Atoi(trackID)
	if err != nil {
		logs.Error("trackID string to int failed ,redirect google page ERROR")
		c.redirect("http://www.google.com")
	}

	err = track.GetAffTrackByTrackID(int64(trackIDInt))
	if err != nil {
		logs.Info("没有找到点击信息，跳转到google页面")
		c.redirect("http://www.google.com")
	}

	serviceConfig := c.getServiceConfig(track.ServiceID)

	mo := new(mondia.Mo)
	// 检查用户之前是否已经订阅过服务
	c.checkUserSubStatus(track, mo)

	err = track.Update()
	if err != nil {
		logs.Info("update track error，跳转到google页面")
		c.redirect("http://www.google.com")
	}

	isCanSub := mo.CheckTodaySubNumMoreLimit(track.ServiceID)
	// 能订阅
	if !isCanSub {
		logs.Info("超过每日订阅限制，跳转到google页面")
		c.redirect("http://www.google.com")
	}
	// 获取跳转到支付页面的URL
	paymentURL := mondia.GetPaymentURL(serviceConfig, trackID)

	//限制每分钟只能产生3个订阅
	isLimit := mo.LimitTenMinutesSubNum(track.ServiceID, 2)

	if isLimit {
		logs.Info("十分钟之内超过3个订阅，跳转到google页面")
		c.redirect("http://www.google.com")
	}

	c.redirect(paymentURL)

}

// SubResult 订阅结果通知
func (c *SubFlowController) SubResult() {
	logs.Info("")
	subResult := new(mondia.MdSubscribe)
	subResult.TrackID = c.Ctx.Input.Param(":trackID")
	subResult.SubStatus = c.GetString("subStatus")
	subResult.CustomerID = c.GetString("customerId")
	subResult.Status = c.GetString("status")
	subResult.ErrorCode = c.GetString("errorCode")
	subResult.NextAction = c.GetString("nextAction")
	subResult.NextActionDate = c.GetString("nextActionDate")
	subResult.SubscriptionID = c.GetString("subId")
	subResult.ViewName = c.GetString("viewName")
	subResult.ErrorDesc = c.GetString("errorDesc")
	// 存入通知
	subResult.Insert()
	// trackID string to int
	trackIDInt := c.trackIDStrToInt(subResult.TrackID)

	track := new(mondia.AffTrack)

	err := track.GetAffTrackByTrackID(int64(trackIDInt))
	if err != nil {
		logs.Info("SubResult 通过trackId查询点击失败，跳转到google页面")
		c.redirect("http://www.google.com")
	}

	// 3001 已经注册过     "SUCCESS" 表示订阅成功
	if (subResult.Status == enums.SubStatusSuccess || subResult.ErrorCode == "3001") && subResult.SubStatus == enums.SubStatusActive {
		contentURL := mondia.ServiceRegisterRequest(subResult.SubscriptionID, subResult.CustomerID,
			track.ServiceID, "register")
		c.redirect(contentURL)
	} else if subResult.SubStatus == enums.SubStatusSuspended { // 扣费失败，但是已经成功订阅
		contentURL := mondia.ServiceRegisterRequest(subResult.SubscriptionID, subResult.CustomerID,
			track.ServiceID, "register")
		c.redirect(contentURL)
	} else if subResult.SubStatus == enums.SubStatusUnsubscribed {
		c.redirect("https://www.google.com")
	} else {
		LpURL, isExist := mondia.GetLpURL(track.ServiceID)

		if isExist {
			c.Data["ErrorCode"] = subResult.ErrorCode
			c.Data["ErrorMessage"] = subResult.ErrorDesc
			c.Data["Lp"] = LpURL
			c.TplName = "views/error.html"
		} else {
			c.redirect("https://www.google.com")
		}
	}

}

func (c *SubFlowController) checkUserSubStatus(track *mondia.AffTrack, mo *mondia.Mo) {
	if track.CustomerID != "" {
		err := mo.GetMoByCustomerID(track.CustomerID)
		// 检查用户是否已经订阅  如果mo ID 不为0表示已经订阅过
		if err != nil && mo.ID != 0 {
			contentURL := mondia.ServiceRegisterRequest(mo.SubscriptionID, mo.CustomerID, track.ServiceID, "register")
			// 记录网盟重复送量的数据
			mondia.InsertAlreadSubData(track)
			// 跳转到内容站
			c.redirect(contentURL)
		}
	}
}
