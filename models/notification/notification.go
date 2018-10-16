package notification

import (
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/MobileCPX/PreMondia/models"
	"github.com/MobileCPX/PreMondia/request"
	"github.com/MobileCPX/PreMondia/util"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

// InsertCharge 插入所有通知
func InsertCharge(charge models.Notification) error {
	o := orm.NewOrm()
	charge.Sendtime, _ = util.GetDatetime()

	_, err := o.Insert(&charge)
	if err != nil {
		logs.Error("插入交易通知表错误 notification，subid：", charge.SubscriptionID)
	}
	return err
}

// UpdateOrInsertMo 更新或者插入MO
func UpdateOrInsertMo(actionType, subStatus, price string, mo *models.Mo) {
	o := orm.NewOrm()
	// var mo models.Mo
	nowTime, _ := util.GetDatetime()
	if actionType == "SUBSCRIBE" {
		trackID := getMdSubscribeTableTrackID(mo.CustomerID)
		trackID = "1"
		if trackID == "" {
			return
		}
		intTrackID, _ := strconv.Atoi(trackID)

		trackData := GetCustomerAffData(intTrackID) // 获取用户点击信息
		mo.ServiceType = trackData.ServiceType
		mo.ClickID = trackData.ClickID
		mo.ProID = trackData.ProID
		mo.PubID = trackData.PubID
		mo.AffName = trackData.AffName
		if subStatus == "ACTIVE" {
			// Send 账号
			var requestData request.MondiaRequestData
			requestData.Message = "Witamy w RedLightVideos. Adres URL to http://www.redlightvideos.com/mm/pl. Twój numer konta to " + mo.CustomerID
			requestData.RequestType = "SendSMS"
			requestData.CustomerID = mo.CustomerID
			_, body := request.MondiaHTTPRequest(requestData)
			if string(body) == "OK" {
				logs.Info("订阅成功后发送账号成功")
			} else {
				logs.Info("订阅成功后发送账号失败")
			}

			postback, _ := getPostbackURL(mo.AffName)
			rate := postback.PostbackRate
			IfIsPostback := postbackRate(mo, rate) //扣量比例  70表示扣百分之三十的量  YES 表示确定不扣量  NO表示扣量
			if IfIsPostback == "YES" {
				mo.PostbackStatus = 1
				mo.PostbackCode = postbackRequest(mo, postback.PostbackURL)
				mo.PostbackTime = nowTime
			}
			mo.SubTime = nowTime
			mo.SubStatus = 1
			o.Insert(mo)
		} else if subStatus == "SUSPENDED" {
			mo.SubTime = nowTime
			mo.SubStatus = 1
			o.Insert(mo)
		}
		return
	} else {
		o.QueryTable("mo").Filter("subscription_id", mo.SubscriptionID).Filter("sub_status", 1).One(&mo)
		if actionType == "UNSUBSCRIBE" {
			if mo.ID != 0 {
				mo.SubStatus = 0
				mo.UnsubTime = nowTime
				// if mo.Msisdn != "" {
				// 	util.HttpRequest(mo.Msisdn, "register", "video", mo.CustomerId, "")
				// } else {
				// 	util.HttpRequest(mo.SubscriptionId, "register", "video", mo.CustomerId, "")
				// }
			}
		} else if actionType == "RENEW" || actionType == "RETRY" {
			if mo.ID != 0 {
				if subStatus == "ACTIVE" {
					mo.Price = price
					mo.SuccessMT++
				} else if subStatus == "SUSPENDED" {
					mo.FailedMT++
				}
			}
		} else if actionType == "STATUS_CHANGE" {
			if mo.ID != 0 {
				//mo.SubscriptionStatus = mo.SubscriptionStatus
			}
		}
		o.Update(&mo)
		return
	}
}

func getMdSubscribeTableTrackID(customerID string) (trackID string) {
	o := orm.NewOrm()
	var mdSubscribe models.MdSubscribe
	o.QueryTable("md_subscribe").Filter("customer_id", customerID).OrderBy("-id").One(&mdSubscribe)
	return mdSubscribe.TrackID
}

// GetCustomerAffData  根据customerID  查询点击信息
func GetCustomerAffData(trackID int) models.AffTrack { // 根据用户id查询出用户网盟信息
	o := orm.NewOrm()
	var trackData models.AffTrack
	o.QueryTable("aff_track").Filter("track_id", trackID).OrderBy("-track_id").One(&trackData)
	return trackData
}

func getPostbackURL(affName string) (models.Postback, error) {
	var postback models.Postback
	o := orm.NewOrm()
	o.Using("default")
	affNameLower := strings.ToLower(affName)
	err := o.QueryTable("postback").Filter("aff_name", affNameLower).One(&postback)
	if err != nil {
		logs.Error("Postback url error: aff_name :" + affName + "   " + err.Error())
	}
	return postback, err
}

func postbackRate(mo *models.Mo, rate int) string {
	var status string
	rand.Seed(time.Now().UnixNano())
	result := rand.Intn(100)
	if result < rate {
		status = "YES"
	} else {
		status = "NO"
	}

	logs.Info("postback status rate: %s   randNum: %s    status: %s  subId:%s   msisdn: %s"+
		"", strconv.Itoa(rate), strconv.Itoa(result), status, mo.SubscriptionID, mo.CustomerID)
	return status
}

func postbackRequest(mo *models.Mo, PostbackURL string) string { // postback请求
	var urls, code string
	code = "400"
	if PostbackURL != "" {
		urls = strings.Replace(PostbackURL, "##clickid##", mo.ClickID, -1)
		urls = strings.Replace(urls, "##pro_id##", mo.ProID, -1)
		urls = strings.Replace(urls, "##pub_id##", mo.PubID, -1)
		urls = strings.Replace(urls, "##operator##", mo.Operator, -1)
	}
	resp, err := http.Get(urls)
	if err == nil {
		defer resp.Body.Close()
		code = strconv.Itoa(resp.StatusCode)
	} else {
		code = err.Error()
		logs.Error("postback Error , CustomerId : " + mo.CustomerID + " aff_name : " + mo.AffName + " error " + err.Error())
	}
	if code != "200" {
		// 发送邮件通知
	}
	return code
}
