package mondia

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// MondiaRequestData Mondia网络请求数据
type MondiaRequestData struct {
	RequestType    string
	ServiceID      string
	SubscriptionID string
	CustomerID     string
	Msisdn         string
	Message        string
}

func SubSceessSendSMS(serviceConfig ServiceInfo, contentURL, customerID, subID string) {
	// Send 账号
	requestData := new(MondiaRequestData)
	requestData.Message = strings.Replace(serviceConfig.ContentURL, "{subID}", subID, -1)
	requestData.RequestType = "SendSMS"
	requestData.CustomerID = customerID

	_, body := requestData.Request(serviceConfig)
	if string(body) == "OK" {
		logs.Info("订阅成功后发送账号成功")
	} else {
		logs.Info("订阅成功后发送账号失败")
	}
}

// Request 向mondia 发起http请求
func (mondiaRequest *MondiaRequestData) Request(serviceConfig ServiceInfo) (status string, body []byte) {
	URL := mondiaRequest.GetMondiaHTTPRequst(serviceConfig)
	client := &http.Client{}
	fmt.Println(URL)
	req, err := http.NewRequest("GET", URL, nil) //建立一个请求
	if err != nil {
		status = "error"
		return
	}
	//Add 头协议
	req.Header.Add("Username", serviceConfig.Username)
	req.Header.Add("Password", serviceConfig.Password)
	response, err := client.Do(req) //提交
	if err != nil {
		status = "error"
		return
	}
	defer response.Body.Close()
	body, err = ioutil.ReadAll(response.Body)
	fmt.Println(string(body), "#########")
	if err != nil {
		status = "error"
		return
	}
	return
}

// GetMondiaHTTPRequst  像mondia后台发起网络请求
func (mondiaRequest *MondiaRequestData) GetMondiaHTTPRequst(serviceConfig ServiceInfo) string {
	requestURL := ""
	switch mondiaRequest.RequestType {
	case "Unsub":
		requestURL = "http://payment.mondiamediamena.com/billing-gw/subservice/unsubscribe?subid=" + mondiaRequest.SubscriptionID + "&operatorId=" + serviceConfig.OperatorID
	case "SendSMS":
		if mondiaRequest.CustomerID != "" {
			requestURL = "http://payment.mondiamediamena.com/billing-gw/subservice/sendsms?customerId=" + mondiaRequest.CustomerID + "&message=" + url.QueryEscape(mondiaRequest.Message) + "&lang=" + serviceConfig.Language + "&operatorId=" + serviceConfig.OperatorID
		} else {
			requestURL = "http://payment.mondiamediamena.com/billing-gw/subservice/sendsms?msisdn=" + mondiaRequest.Msisdn + "&message=" + url.QueryEscape(mondiaRequest.Message) + "&lang=" + serviceConfig.Language + "&operatorId=" + serviceConfig.OperatorID
		}
	case "GetCustomer":
		requestURL = "http://payment.mondiamediamena.com/billing-gw/service/getcustomer?msisdn=" + mondiaRequest.Msisdn + "&operatorId=" + serviceConfig.OperatorID
	}
	return requestURL
}
