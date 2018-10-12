package models

import (
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/MobileCPX/PreMondia/util"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

func UpdateOrInsertMo(actionType, subStatus, price string, mo *MonMo) *MonMo {
	o := orm.NewOrm()
	var one_mo MonMo
	new_time, _ := util.GetFormatTime()
	if actionType == "SUBSCRIBE" {
		CustomerData := GetCustomerAffData(mo.CustomerId) // 获取用户点击信息
		mo.ServiceType = CustomerData.ServiceType
		mo.ClickId = CustomerData.ClickId
		mo.ProId = CustomerData.ProId
		mo.ClickType = CustomerData.ClickType
		mo.PubId = CustomerData.PubId
		mo.AffName = CustomerData.AffName
		if subStatus == "ACTIVE" {
			_, postback := Get_postback_url(mo.AffName)
			rate := postback.PostbackRate
			IfIsPostback := PostbackRate(mo, rate) //扣量比例  70表示扣百分之三十的量  YES 表示确定不扣量  NO表示扣量
			if IfIsPostback == "YES" {
				mo.PostbackStatus = 1
				mo.PostbackCode = PostbackRequest(mo, postback.Post_back)
				mo.PostbackTime = new_time
			}
			mo.Subtime = new_time
			mo.SubNum = 1
			mo.SubStatus = 1
			o.Insert(mo)
		} else if subStatus == "SUSPENDED" {
			mo.Subtime = new_time
			mo.SubNum = 0
			mo.SubStatus = 1
			o.Insert(mo)
		}
		return mo
	} else {
		o.QueryTable("mon_mo").Filter("subscription_id", mo.SubscriptionId).Filter("sub_status", 1).One(&one_mo)
		if actionType == "UNSUBSCRIBE" {
			if one_mo.Id != 0 {
				one_mo.SubStatus = 0
				one_mo.Unsubtime = new_time
				if one_mo.Msisdn != "" {
					util.HttpRequest(mo.Msisdn, "register", "video", mo.CustomerId, "")
				} else {
					util.HttpRequest(mo.SubscriptionId, "register", "video", mo.CustomerId, "")
				}
			}
		} else if actionType == "RENEW" || actionType == "RETRY" {
			if one_mo.Id != 0 {
				if subStatus == "ACTIVE" {
					one_mo.Price = price
					one_mo.SubNum += 1
				} else if subStatus == "SUSPENDED" {
					one_mo.FailNum += 1
				}
			}
		} else if actionType == "STATUS_CHANGE" {
			if one_mo.Id != 0 {
				//one_mo.SubscriptionStatus = mo.SubscriptionStatus
			}
		}
		o.Update(&one_mo)
		return &one_mo
	}
}

func GetAffDataId(id int) MdId {
	o := orm.NewOrm()
	var md_id MdId
	o.QueryTable("md_id").Filter("id", id).One(&md_id)
	return md_id
}

func InsertCharge(charge Notification) error {
	o := orm.NewOrm()
	charge.Sendtime, _ = util.GetFormatTime()

	_, err := o.Insert(&charge)
	return err
}

func IdGetServiceType(subId string) MdId {
	o := orm.NewOrm()
	var sub_data MdId
	idInt64, _ := strconv.ParseInt(subId, 10, 64)
	o.QueryTable("md_id").Filter("id", idInt64).One(&sub_data)
	return sub_data
}

func GetCustomerAffData(customerId string) MdCustomer { // 根据用户id查询出用户网盟信息
	o := orm.NewOrm()
	var customer MdCustomer
	o.QueryTable("md_customer").Filter("customer_id", customerId).OrderBy("-id").One(&customer)
	return customer
}

func Get_postback_url(aff_name string) (error, PostbackUrl) {
	var postback PostbackUrl
	o := orm.NewOrm()
	o.Using("default")
	err := o.QueryTable("postback_url").Filter("aff_name", aff_name).One(&postback)
	if err != nil {
		logs.Error("Postback url error: aff_name :" + aff_name + "   " + err.Error())
	}
	return err, postback
}

func PostbackRate(mo *MonMo, rate int) string {
	var status string
	rand.Seed(time.Now().UnixNano())
	result := rand.Intn(100)
	if result < rate {
		status = "YES"
	} else {
		status = "NO"
	}

	logs.Info("postback status rate: %s   randNum: %s    status: %s  subId:%s   msisdn: %s"+
		"", strconv.Itoa(rate), strconv.Itoa(result), status, mo.SubscriptionId, mo.CustomerId)
	return status
}

func PostbackRequest(mo_data *MonMo, postback_url string) string { // postback请求
	var urls, code string
	code = "400"
	if postback_url != "" {
		urls = strings.Replace(postback_url, "##clickid##", mo_data.ClickId, -1)
		urls = strings.Replace(urls, "##pro_id##", mo_data.ProId, -1)
		urls = strings.Replace(urls, "##pub_id##", mo_data.PubId, -1)
		urls = strings.Replace(urls, "##operator##", mo_data.Operator, -1)
	}
	post_result, err := http.Get(urls)
	if err == nil {
		code = strconv.Itoa(post_result.StatusCode)
		post_result.Body.Close()
	} else {
		code = err.Error()
		logs.Error("postback Error , CustomerId : " + mo_data.CustomerId + " aff_name : " + mo_data.AffName + " error " + err.Error())
	}
	return code
}
