package models

import (
	"fmt"
	"strings"

	"github.com/MobileCPX/PreMondia/util"

	"github.com/astaxie/beego/orm"
)

func InsertCustomer(customer *MdCustomer) {
	o := orm.NewOrm()
	customer.Sendtime, _ = util.GetFormatTime()
	o.Insert(customer)
}

func InsertSubscribe(sub *MdSubscribe) {
	o := orm.NewOrm()
	sub.Sendtime, _ = util.GetFormatTime()
	o.Insert(sub)
}

func GetCustomer(status, operator_country, serviceTypeId, id, productCd, subPackage, imgPath, prodPrice, service_type string) string {
	operator_dict := map[string]string{"Vodafone Egypt": "1", "Orange Egypt": "2", "Orange Jordan": "3", "Etisalat Egypt": "4",
		"Etisilat UAE": "5", "Orange Tunisia": "6", "Vodacom South Africa": "7", "Orange Poland": "8", "VIVA Bahrain": "9"}
	var url string

	langCode := ""
	switch service_type {
	case "video_d", "video_w":
		langCode = "&langCode=pl"
	}

	switch status {
	case "SUCCESS":
		operator_code := operator_dict[operator_country]
		operator := strings.Split(operator_country, " ")[0]
		if operator == "Orange" {
			url = "http://sso.orange.com/mondiamedia_subscription/?method=subscribe&merchantId=93&redirect=" +
				"http%3a%2f%2fcpx3.allcpx.com%2fsubs%2fres%2f" + id + "&productCd=" + productCd + "&subPackage=" + subPackage + "" +
				"&operatorId=" + operator_code + "&imgPath=" + imgPath + langCode
		} else {
			url = "http://login.mondiamediamena.com/billinggw-lcm/billing?method=subscribe&merchantId=93&redirect=" +
				"http%3a%2f%2fcpx3.allcpx.com%2fsubs%2fres%2f" + id + "&productCd=" + productCd + "&subPackage=" + subPackage + "" +
				"&operatorId=" + operator_code + "&imgPath=" + imgPath + langCode
		}
	case "ERROR":
		if len(strings.Split(serviceTypeId, "|")) == 2 {
			url = "http://sso.orange.com/mondiamedia_subscription/?method=getcustomer&merchantId=93&redirect=" +
				"http://cpx3.allcpx.com/subs/getcust/" + serviceTypeId + "|again" + langCode
		} else {
			url = "http://za.mobpre.com/static/operator/index.html?type=" + service_type + "&id=" + id
			if service_type == "video_d" || service_type == "video_w" {
				url = "http://sso.orange.com/mondiamedia_subscription/?method=subscribe&merchantId=93&redirect=" +
					"http%3a%2f%2fcpx3.allcpx.com%2fsubs%2fres%2f" + id + "&imgPath=" + imgPath + "&productCd=" +
					productCd + "&subPackage=" + subPackage + "&operatorId=8" + langCode
			}
		}
	}
	fmt.Println(url)
	return url
}
