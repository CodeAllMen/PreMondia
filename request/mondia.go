package request

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// GetMondiaHTTPRequst  像mondia后台发起网络请求
func GetMondiaHTTPRequst(requestType, SubID, customerID, msisdn, message string) {
	requestURL := ""
	switch requestType {
	case "Unsub":
		requestURL = "http://payment.mondiamediamena.com/billing-gw/subservice/unsubscribe?subid=" + SubID + "&operatorId=8"
	case "SendSMS":
		if customerID != "" {
			requestURL = "http://payment.mondiamediamena.com/billing-gw/subservice/sendsms?customerId=" + customerID + "&message=" + message + "&lang=pl&operatorId=8"
		} else {
			requestURL = "http://payment.mondiamediamena.com/billing-gw/subservice/sendsms?msisdn=" + msisdn + "&message=" + message + "&lang=pl&operatorId=8"
		}
	case "GetCustomer":
		requestURL = "http://payment.mondiamediamena.com/billing-gw/service/getcustomer?msisdn=" + msisdn + "&operatorId=8"
	}
	fmt.Println(requestURL)
}

//HTTPRequest 向mondia 发起http请求
func HTTPRequest(URL string) (status string, body []byte) {
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
