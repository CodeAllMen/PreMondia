package mondia

import (
	"encoding/json"
	"fmt"
	"github.com/MobileCPX/PreMondia/enums"
	"github.com/MobileCPX/PreMondia/models/mondia"
	"github.com/MobileCPX/PreMondia/util"
	"github.com/astaxie/beego/logs"
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

	if err != nil || err1 != nil || int(todaySubNum) >= enums.DayLimitSub {
		c.Ctx.WriteString("false")
	} else {
		c.Ctx.WriteString(strconv.FormatInt(trackID, 10))
	}
}

func (c *SubFlowController) GetCustomerRedirect() {
	trackID := c.GetString("frmlp")
	track := new(mondia.AffTrack)
	trackIDInt, err := strconv.Atoi(trackID)
	if err != nil {
		logs.Info("trackID string to int failed ,redirect google page")
		c.redirect("http://www.google.com")
	}

	err = track.GetAffTrackByTrackID(int64(trackIDInt))

	customerRedirectURL := "http://sso.orange.com/mondiamedia_subscription/?method=getcustomer&merchantId=93&langCode=pl" +
		"&redirect=http://cpx3.allcpx.com:8085/subs/getcust/" + trackID + "?product_code=" + track.ServiceID
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
	data,_ := json.Marshal(track)
	fmt.Println(string(data))
	trackIDInt, err := strconv.Atoi(trackID)
	if err != nil {
		logs.Info("trackID string to int failed ,redirect google page")
		c.redirect("http://www.google.com")
	}

	err = track.GetAffTrackByTrackID(int64(trackIDInt))
	if err != nil {
		logs.Info("没有找到点击信息，跳转到google页面")
		c.redirect("http://www.google.com")
	}

	mo := new(mondia.Mo)
	//if track.CustomerID != "" {
	//	err := mo.GetMoByCustomerID(track.CustomerID)
	//	// 检查用户是否已经订阅  如果mo ID 不为0表示已经订阅过
	//	if err != nil && mo.ID != 0 {
	//		contentURL := mondia.ServiceRegisterRequest(mo.SubscriptionID, mo.CustomerID, track.ServiceID, "register")
	//		// 记录网盟重复送量的数据
	//		mondia.InsertAlreadSubData(track)
	//		// 跳转到内容站
	//		c.redirect(contentURL)
	//	}
	//}

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
	paymentURL, isExist := mondia.GetPaymentURL(track.ServiceID, trackID)
	if !isExist {
		logs.Info("product_code 不存在，跳转到google页面")
		c.redirect("http://www.google.com")
	}
	//限制每分钟只能产生3个订阅
	isLimit := mo.LimitTenMinutesSubNum(track.ServiceID, 2)

	if isLimit {
		logs.Info("十分钟之内超过3个订阅，跳转到google页面")
		c.redirect("http://www.google.com")
	}

	printRedirectAocLog(track)
	c.redirect(paymentURL)

}

// SubResult 订阅结果通知
func (c *SubFlowController) SubResult() {
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

	track := new(mondia.AffTrack)
	trackIDInt, err := strconv.Atoi(subResult.TrackID)
	if err != nil {
		logs.Info("trackID string to int failed ,redirect google page")
		c.redirect("http://www.google.com")
	}
	err = track.GetAffTrackByTrackID(int64(trackIDInt))

	// 3001 已经注册过     "SUCCESS" 表示订阅成功
	if (subResult.Status == "SUCCESS" || subResult.ErrorCode == "3001") && subResult.SubStatus == "ACTIVE" {
		contentURL := mondia.ServiceRegisterRequest(subResult.SubscriptionID, subResult.CustomerID,
			track.ServiceID, "register")
		c.redirect(contentURL)
	} else if subResult.SubStatus == "SUSPENDED" {
		contentURL := mondia.ServiceRegisterRequest(subResult.SubscriptionID, subResult.CustomerID,
			track.ServiceID, "register")
		c.redirect(contentURL)
	} else if subResult.SubStatus == "UNSUBSCRIBED" {
		c.redirect("https://www.google.com")
	} else {
		LpURL, isExist := mondia.GetLpURL(track.ServiceID)
		if isExist {
			c.redirect(LpURL)
		} else {
			c.redirect("https://www.google.com")
		}

	}

}

func printRedirectAocLog(track *mondia.AffTrack) {
	_, nowDate := util.GetDatetime()
	logs.Info(nowDate, " 跳转到aoc页面，affName: ", track.AffName, " PubID: ", track.PubID, " click_id: ", track.PubID)
}
