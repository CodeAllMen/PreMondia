package mondia

import (
	"encoding/xml"
	"fmt"
	"github.com/MobileCPX/PreMondia/models"
	"github.com/MobileCPX/PreMondia/models/mondia"
	"github.com/MobileCPX/PreMondia/models/unsub"
	"github.com/astaxie/beego/logs"
	"math/rand"
	"strconv"
	"time"
)

type UnsubController struct {
	BaseController
}

func (c *UnsubController) UnsubPage() {
	c.Data["service_id"] = c.GetString("service_id")
	c.TplName = "unsub.tpl"
}

func (c *UnsubController) UnsubSendPin() {
	var requestData mondia.MondiaRequestData
	msisdn := c.GetString("msisdn")
	serviceID := c.GetString("service_id")
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	pin := fmt.Sprintf("%03v", rnd.Intn(1000))
	logs.Info("PIN: ", pin)
	// message := url.QueryEscape("[RedLightVideos] Your unsubscribe PIN code is " + pin)
	requestData.ServiceID = serviceID
	requestData.RequestType = "SendSMS"
	requestData.Message = mondia.GetPINUnsubMessage(serviceID, pin)
	requestData.Msisdn = msisdn

	status, body := requestData.Request()
	fmt.Println(string(body))
	if status == "error" {
		c.Data["error"] = "0"
		c.TplName = "unsub.tpl"
		return
	}
	unsubPin := new(mondia.UnsubPin)
	unsubPin.Msisdn = msisdn
	unsubPin.Pin = pin
	unsubPin.PinStatus = string(body)
	id, err := unsubPin.Insert()
	if string(body) == "OK" {
		if err == nil {
			c.Data["id"] = strconv.FormatInt(id, 10)
			c.Data["service_id"] = serviceID
			c.Data["phone"] = msisdn
			c.TplName = "pin.tpl"
		} else {
			c.Data["error"] = "0"
			c.Data["service_id"] = serviceID
			c.TplName = "unsub.tpl"
		}
	} else {
		c.Data["error"] = "0"
		c.TplName = "unsub.tpl"
	}
}

func (c *UnsubController) UnsubRequest() {
	requestData := new(mondia.MondiaRequestData)
	pin := c.GetString("pin")
	id := c.GetString("id")
	serviceID := c.GetString("service_id")
	unsubPIN := new(mondia.UnsubPin)
	err := unsubPIN.CheckPIN(id)
	if err == nil {
		if unsubPIN.Pin == pin {

		}
	}
	if err == nil && unsubPIN.Pin == pin {
		requestData.Msisdn = unsubPIN.Msisdn
		requestData.RequestType = "GetCustomer"
		status, body := requestData.Request()
		fmt.Println(string(body))
		if status == "error" {
			c.UnsubFailed(serviceID)
			return
		}
		fmt.Println(string(body))
		customerData := new(mondia.UnsubGetCustomer)
		customerErr := xml.Unmarshal(body, customerData)
		if customerErr == nil {
			if customerData.ResponseCode == "1001" {
				mo := new(mondia.Mo)
				mo.UnsubGetMoByCustomerID(customerData.CustomerId, serviceID)

				if mo.SubscriptionID != "" {
					mo.Msisdn = requestData.Msisdn
					_ = mo.UpdateMO()
					requestData.SubscriptionID = mo.SubscriptionID
					requestData.RequestType = "Unsub"

					status, body := requestData.Request()
					fmt.Println(string(body))
					if status != "error" {
						unsubNotification := new(models.MondiaCharge)
						xmlErr := xml.Unmarshal(body, unsubNotification)
						if xmlErr != nil {
							//c.TplName = "fail.tpl"
							c.UnsubFailed(serviceID)
							return
						}
						err := unsub.InsertUnsubData(unsubNotification)
						if err == nil && unsubNotification.ResponseCode == "1001" { // 取消订阅成功
							_, _ = mo.UnsubUpdateMo(mo.SubscriptionID)
							//c.Data["url"] = mondia.GetContentURL(mo.ServiceID)
							c.Data["url"] = mondia.GetContentURL(mo.ProductCode)
							c.TplName = "success.tpl"
							return
						} else if err == nil && unsubNotification.ResponseCode == "3029" { // 之前已经取消
							c.Data["url"] = mondia.GetContentURL(serviceID)
							c.TplName = "success.tpl"
							return
						}
					} else {
						//  "用户不存在"
						//c.TplName = "fail.tpl"
						c.UnsubFailed(serviceID)

						return
					}
				} else {
					//  "用户不存在"
					//c.TplName = "fail.tpl"
					c.UnsubFailed(serviceID)
					return
				}
			} else {
				//c.TplName = "fail.tpl"
				c.UnsubFailed(serviceID)
				return
			}
		} else {
			//c.TplName = "fail.tpl"
			c.UnsubFailed(serviceID)
			return
		}
	} else {
		c.Data["error"] = "201"
		c.Data["id"] = id
		c.Data["phone"] = requestData.Msisdn
		c.TplName = "pin.tpl"
	}
}

type request struct {
	UserIdToken    userIDToken `xml:"userIdToken"`
	SubscriptionId string      `xml:"subscriptionId"`
	// Service        string      `xml:"service"`
}
type userIDToken struct {
	Username string `xml:"username"`
	Password string `xml:"password"`
}

////  UnsubWap3G WAP 3G环境下退订
//func (c *UnsubController) UnsubWap3G() {
//	subID := c.GetString("user_id")
//	serviceName := c.GetString("service_name", "PinkCity4K")
//
//	unsubReqData := new(request)
//	unsubResult := unsubReqData.unsubRequest(subID, serviceName)
//
//	if unsubResult == "0" {
//		fmt.Println("退订成功")
//		c.TplName = "unsub_success.html"
//	} else {
//		c.Data["display"] = "none"
//		c.TplName = "unsub_failed.html"
//	}
//}

//// 电话号码退订
//func (c *UnsubController) MsisdnUnsub() {
//	msisdn := c.GetString("msisdn")
//	serviceName := c.GetString("service_name", "PinkCity4K")
//	mo := new(mondia.Mo)
//	// 通过电话号码和ServiceName查询Mo信息
//	mo.GetMoByMsisdnAndServiceName(msisdn, serviceName)
//	if mo.SubscriptionID != "" {
//		unsubReqData := new(request)
//		unsubResult := unsubReqData.unsubRequest(mo.SubscriptionID, serviceName)
//		if unsubResult == "0" {
//			fmt.Println("退订成功")
//			c.TplName = "unsub_success.html"
//		} else {
//			c.Data["error_text"] = "* O número de telefone está errado, por favor, reinsira"
//			c.TplName = "unsub_failed.html"
//		}
//	} else {
//		c.Data["error_text"] = `Este número de telefone não está inscrito no serviço, verifique seu número de telefone: ` + msisdn
//		c.TplName = "unsub_success.html"
//	}
//}

//func (unsubReqData *request) unsubRequest(subID, serviceName string) string {
//	var token userIDToken
//	token.Username = "leaderapi"
//	token.Password = "LDgo1@1@"
//	unsubReqData.SubscriptionId = subID
//	unsubReqData.UserIdToken = token
//	data, _ := xml.MarshalIndent(unsubReqData, "", "  ") // 组装xml 数据  返回 []byte类型
//	fmt.Println(string(data))
//	xmlRequestData := `<?xml version="1.0" encoding="UTF-8"?>
//	<SOAP-ENV:Envelope xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ns1="http://www.go4mobility.com/ws/skysms/wapBilling/v2_0/service">
//		<SOAP-ENV:Body>
//			<ns1:cancelSubscription>` + string(data) + `</ns1:cancelSubscription>
//			</SOAP-ENV:Body>
//		</SOAP-ENV:Envelope>`
//
//	v2XMLRequestURL := mondia.ServiceData[serviceName].WapBillingV2URL
//	respData := util.HttpPostRequest([]byte(xmlRequestData), v2XMLRequestURL, "text/xml;charset=utf-8")
//
//	unsubResult := strings.Split(string(respData), "<result>")[1]
//	result := strings.Split(unsubResult, "</result>")[0]
//	return result
//}

//func unsubRequest(subID, serviceName string) string {
//unsubRequest := new(request)
//var token userIDToken
//token.Username = "leaderapi"
//token.Password = "LDgo1@1@" //西班牙测试
//unsubRequest.SubscriptionId = subID
//// unsubRequest.Service = "TEST_MNT"
//unsubRequest.UserIdToken = token
//data, _ := xml.MarshalIndent(&unsubRequest, "", "  ") // 组装xml 数据  返回 []byte类型
//fmt.Println(string(data))
//xmlRequestData := `<?xml version="1.0" encoding="UTF-8"?>
//<SOAP-ENV:Envelope xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ns1="http://www.go4mobility.com/ws/skysms/wapBilling/v2_0/service">
//	<SOAP-ENV:Body>
//		<ns1:cancelSubscription>` + string(data) + `</ns1:cancelSubscription>
//		</SOAP-ENV:Body>
//	</SOAP-ENV:Envelope>`
//
//v2XMLRequestURL := mondia.ServiceData[serviceName].WapBillingV2URL
//respData := util.HttpPostRequest([]byte(xmlRequestData), v2XMLRequestURL, "text/xml;charset=utf-8")
//
//unsubResult := strings.Split(string(respData), "<result>")[1]
//result := strings.Split(unsubResult, "</result>")[0]
//return result
//}
