package request

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// MondiaRequestData Mondia网络请求数据
type MondiaRequestData struct {
	RequestType    string
	SubscriptionID string
	CustomerID     string
	Msisdn         string
	Message        string
}

// MondiaHTTPRequest 向mondia 发起http请求
func MondiaHTTPRequest(requestData MondiaRequestData) (status string, body []byte) {
	URL := GetMondiaHTTPRequst(requestData)
	client := &http.Client{}
	fmt.Println(URL)
	reqest, err := http.NewRequest("GET", URL, nil) //建立一个请求
	if err != nil {
		status = "error"
		return
	}
	//Add 头协议
	reqest.Header.Add("Username", "opcpx")
	reqest.Header.Add("Password", "cpx22334")
	response, err := client.Do(reqest) //提交
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
func GetMondiaHTTPRequst(requestData MondiaRequestData) string {
	requestURL := ""
	switch requestData.RequestType {
	case "Unsub":
		requestURL = "http://payment.mondiamediamena.com/billing-gw/subservice/unsubscribe?subid=" + requestData.SubscriptionID + "&operatorId=8"
	case "SendSMS":
		if requestData.CustomerID != "" {
			requestURL = "http://payment.mondiamediamena.com/billing-gw/subservice/sendsms?customerId=" + requestData.CustomerID + "&message=" + url.QueryEscape(requestData.Message) + "&lang=pl&operatorId=8"
		} else {
			requestURL = "http://payment.mondiamediamena.com/billing-gw/subservice/sendsms?msisdn=" + requestData.Msisdn + "&message=" + url.QueryEscape(requestData.Message) + "&lang=pl&operatorId=8"
		}
	case "GetCustomer":
		requestURL = "http://payment.mondiamediamena.com/billing-gw/service/getcustomer?msisdn=" + requestData.Msisdn + "&operatorId=8"
	}
	return requestURL
}
