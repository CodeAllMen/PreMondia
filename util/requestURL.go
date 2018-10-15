package util

import (
	"fmt"
	"net/http"

	"github.com/MobileCPX/PreMondia/initial"
	"github.com/astaxie/beego/logs"
)

// GetSubRequestURL 获取3G或者wifi状态下的订阅URL
func GetSubRequestURL(status string) {
	modiaConf := initial.GetMondiaConf()
	redirect := ""
	URL := ""
	if status == "SUCCESS" {
		URL = fmt.Sprintf("http://login.mondiamediamena.com/billinggw-lcm/billing?"+
			"method=subscribe&merchantId=%s&redirect=%s&productCd=%s&subPackage=%s"+
			"&operatorId=1&campaignId=CMP001", modiaConf.MrchantID, redirect,
			modiaConf.ProductCode, modiaConf.SubPackage)
	} else {
		// http: //<host>/billinggw-lcm/billing?method=subscribe&merchantId=<Merchant ID>&redirect=<REDIRECT URL>&productCd=<PRODUCT CODE>&subPackage=<SUBSCRIPTION PACKAGE>& imgPath=<IMAGE URL>&operatorId=<OPERATOR ID>&campaignId=<CAMPAIGN ID>
		// http: //<host>/billinggw-lcm/billing?method=subscribe&merchantId=<Merchant ID>&redirect=<REDIRECT URL>&productCd=<PRODUCT CODE>&subPackage=<SUBSCRIPTION PACKAGE>& imgPath=<IMAGE URL>&operatorId=<OPERATOR ID>&campaignId=<CAMPAIGN ID>
		URL = fmt.Sprintf("http://login.mondiamediamena.com/billinggw-lcm/billing?"+
			"method=subscribe&merchantId=%s&redirect=%s&productCd=%s&subPackage=%s"+
			"&operatorId=1&campaignId=CMP001", modiaConf.MrchantID, redirect,
			modiaConf.ProductCode, modiaConf.SubPackage)
	}

	fmt.Println(URL)
}

//HttpRequest 注册
func HttpRequest(subID, types, period, CustomerId, unsubtime string) {
	urls := ""
	if types == "register" {
		urls = fmt.Sprintf("http://www.redlightvideos.com/addsubs?uiid=%s&sign=mondia", subID)
	}
	if types == "delete" {
		urls = fmt.Sprintf("http://www.redlightvideos.com/delete/user?phone=%s&time=%s", subID, unsubtime)
	}
	resp, err := http.Get(urls)
	if err == nil {
		logs.Info(fmt.Sprintf("HttpRequest Success %s service %s subID: %s  CustomerId: %s ", types, period, subID, CustomerId))
		resp.Body.Close()
	} else {
		logs.Error(fmt.Sprintf("HttpRequest Failed %s service %s subID: %s  CustomerId: %s     error: %s ", types, period, subID, CustomerId, err.Error()))
	}
}
