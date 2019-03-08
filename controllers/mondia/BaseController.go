package mondia

import (
	"github.com/MobileCPX/PreMondia/enums"
	"github.com/MobileCPX/PreMondia/models"
	"github.com/MobileCPX/PreMondia/models/mondia"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"strconv"
	"strings"
	"time"
)

type BaseController struct {
	beego.Controller
}

// CheckError 检查是否有错 msg 定义日志信息
func (c *BaseController) CheckError(err error, errorCode enums.ErrorCode, msg ...string) {
	if err != nil {
		// 打印日志信息
		if len(msg) != 0 {
			logs.Error(msg, " ERROR: ", err.Error())
		}
		switch errorCode {
		case enums.RedirectGoogle:
			c.redirect("https://wwww.google.com")
		case enums.Error502:
			c.Ctx.ResponseWriter.WriteHeader(502)
			c.StopRun()
		default:
			c.redirect("https://wwww.google.com")
		}
	}
}

func (c *BaseController) NewInsertMo(notification *mondia.Notification, affTrack *mondia.AffTrack) (notificationType string) {
	mo := new(mondia.Mo)
	isExist := mo.CheckSubIDIsExist(notification.SubscriptionID)
	serviceConfig := c.getServiceConfig(notification.ProductCode)
	// 判断用户是否已经存在
	if !isExist {
		mo = new(mondia.Mo)
		// 初始化MO数据
		mo.InitNewSubMO(notification, affTrack)
		// 查询次网盟今天的订阅数及postback回传数，根据概率判断次数书是否回传
		subNum, postbackNum := mo.GetAffNameTodaySubInfo()
		// 根据概率判断次数书是否回传
		isSuccess, code, payout := mondia.StartPostback(mo, subNum, postbackNum)
		mo.PostbackCode = code
		contentURL := mondia.GetContentURL(mo.ServiceID)
		//订阅成功发送短信
		mondia.SubSceessSendSMS(serviceConfig, contentURL, mo.CustomerID, mo.SubscriptionID)

		if mo.AffName != "" {
			// 有转化后发邮件
			//util.BeegoEmail("", "葡萄牙 NOS 有新的转化", "网盟名称： "+mo.AffName,
			//	[]string{"tengjiaqing@mobilecpx.com", "wangangui@mobilecpx.com"})
		}
		// 判断是否回传成功
		if isSuccess {
			mo.Payout = payout
			mo.PostbackStatus = 1
		}
		err := mo.InsertNewMo()
		if err == nil {
			notificationType = "SUB"
		}
	}
	return
}

// setCookie
func (c *BaseController) setCookie(trackID string) string {
	// 获取cookie
	userId, ok := c.GetSecureCookie("user_cookie", "8A66b76dbd3759445fe924d28a5F6856")
	if !ok {
		userId = "PinkCity__" + trackID + "_" + "1"
	} else {
		userIdList := strings.Split(userId, "_")
		if len(userIdList) != 3 {
			userId = "PinkCity__" + trackID + "_" + "1"
		} else {
			vistTimes, err := strconv.Atoi(userIdList[2])
			if err != nil {
				c.Ctx.ResponseWriter.ResponseWriter.WriteHeader(404)
				c.StopRun()
			} else {
				userId = userIdList[0] + "_" + userIdList[1] + "_" + strconv.Itoa(vistTimes+1)
			}
		}
	}
	// 设置cookie
	c.SetSecureCookie("user_cookie", "8A66b76dbd3759445fe924d28a5F6856", userId, 61622400*time.Second)
	return userId
}

func (c *BaseController) redirect(url string) {
	if url == "" {
		url = "http://google.com"
	}
	c.Redirect(url, 302)
	c.StopRun()
}

func (c *BaseController) jsonResult(code enums.JsonResultCode, msg string, obj interface{}) {
	r := &models.JsonResult{code, msg, obj}
	c.Data["json"] = r
	c.ServeJSON()
	c.StopRun()
}

func (c *BaseController) UnsubSuccess() {

}

func (c *BaseController) UnsubFailed(serviceID string) {
	c.Data["service_id"] = serviceID
	c.Data["contentURL"] = mondia.GetContentURL(serviceID)
	c.TplName = "fail.tpl"
}

func (c *BaseController) getServiceConfig(serviceID string) mondia.ServiceInfo {
	serviceConfig, isExist := c.serviceCofig(serviceID)
	if !isExist {
		logs.Error("服务ID不存在，请检查服务信息，servideID: ", serviceID, "ERROR")
		c.redirect("http://www.google.com")
	}
	return serviceConfig
}

func (c *BaseController) serviceCofig(serviceID string) (mondia.ServiceInfo, bool) {
	serviceConfig, isExist := mondia.ServiceData[serviceID]
	return serviceConfig, isExist
}

// 将string 类型的trackID 转为 INT 类型
func (c *BaseController) trackIDStrToInt(trackID string) int {
	trackIDInt, err := c.strToInt(trackID)
	if err != nil {
		logs.Error("trackID string to int 错误，ERROR: ", err.Error(), " trackID: ", trackID)
		c.redirect("http://google.com")
	}
	return trackIDInt
}

func (c *BaseController) strToInt(str string) (int, error) {
	return strconv.Atoi(str)
}
