package unsub

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/MobileCPX/PreMondia/models"
	"github.com/MobileCPX/PreMondia/models/unsub"
	"github.com/MobileCPX/PreMondia/request"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// UnsubRequestControllers 退订请求
type UnsubRequestControllers struct {
	beego.Controller
}

// SendPINControllers 退订发送验证码
type SendPINControllers struct {
	beego.Controller
}

type UnsubGetCustomer struct {
	beego.Controller
}

// UnsubPage 退订页面
type UnsubPage struct {
	beego.Controller
}

//Get 退订首页
func (c *UnsubPage) Get() {
	// c.TplName = "unsub.tpl"
	c.TplName = "success.tpl"
}

// Post 退订发送PIN
func (c *SendPINControllers) Post() {
	var requestData request.MondiaRequestData
	msisdn := c.GetString("msisdn")
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	pin := fmt.Sprintf("%03v", rnd.Intn(1000))
	logs.Info("PIN: ", pin)
	// message := url.QueryEscape("[RedLightVideos] Your unsubscribe PIN code is " + pin)
	requestData.RequestType = "SendSMS"
	requestData.Message = "[RedLightVideos] Kod PIN, który anulowałeś swoją subskrypcję, to " + pin
	requestData.Msisdn = msisdn
	// getPinURL := "http://payment.mondiamediamena.com/billing-gw/subservice/sendsms?operatorId=8&lang=pl&msisdn=" + msisdn + "&message=" + message
	status, body := request.MondiaHTTPRequest(requestData)
	fmt.Println(string(body))
	if status == "error" {
		c.Data["error"] = "0"
		c.TplName = "unsub.tpl"
		return
	}
	unsubPin := new(models.UnsubPin)
	unsubPin.Msisdn = msisdn
	unsubPin.Pin = pin
	unsubPin.PinStatus = string(body)
	id, err := unsub.InsertPinData(unsubPin)
	if string(body) == "OK" {
		if err == nil {
			c.Data["id"] = strconv.FormatInt(id, 10)
			c.Data["phone"] = msisdn
			c.TplName = "pin.tpl"
		} else {
			c.Data["error"] = "0"
			c.TplName = "unsub.tpl"
		}
	} else {
		c.Data["error"] = "0"
		c.TplName = "unsub.tpl"
	}
}

// Get 退订请求  0 表示退订成功 1 表示用户不存在  2 表示退订失败
func (c *UnsubRequestControllers) Get() {
	customerID := c.GetString("cusID")
	returnStr := ""
	if customerID != "" {
		mo, isHas := unsub.GetUnsubMoData(customerID)
		if !isHas { // 此用户不存在
			returnStr = "1"
		} else {
			unsubURL := "http://payment.mondiamediamena.com/billing-gw/subservice/unsubscribe?subid=" + mo.SubscriptionID + "&operatorId=8"
			status, body := MondiaHTTPRequest(unsubURL)
			fmt.Println(string(body))
			if status != "error" {
				unsubNotification := new(models.MondiaCharge)
				xmlErr := xml.Unmarshal(body, unsubNotification)
				if xmlErr != nil {
					returnStr = "2"
					c.TplName = "fail.tpl"
					return
				}
				err := unsub.InsertUnsubData(unsubNotification)
				if err == nil && unsubNotification.ResponseCode == "1001" {
					unsub.UpdateUnsubMoTable(mo.SubscriptionID)
					returnStr = "0"
					c.TplName = "success.tpl"
					return
				} else {
					returnStr = "2"
					c.TplName = "fail.tpl"
					return
				}
			} else {

			}
		}
	}
	c.Ctx.WriteString(returnStr)
}

// Post 退订获取CustomerId Post请求
func (c *UnsubGetCustomer) Post() {
	var requestData request.MondiaRequestData
	pin := c.GetString("pin")
	id := c.GetString("id")
	isHasPin, msisdn, _ := unsub.CheckPIN(pin, id)
	if isHasPin {
		requestData.Msisdn = msisdn
		requestData.RequestType = "GetCustomer"
		// getCustomerURL := "http://payment.mondiamediamena.com/billing-gw/service/getcustomer?msisdn=" + msisdn + "&operatorId=8"
		status, body := request.MondiaHTTPRequest(requestData)
		fmt.Println(string(body))
		if status == "error" {
			c.TplName = "fail.tpl"
			return
		}
		fmt.Println(string(body))
		customerData := new(models.UnsubGetCustomer)
		customerErr := xml.Unmarshal(body, customerData)
		if customerErr == nil {
			if customerData.ResponseCode == "1001" {
				subID := unsub.CustomerToGetSubID(customerData.CustomerId, msisdn)
				if subID != "" {
					requestData.SubscriptionID = subID
					requestData.RequestType = "Unsub"

					status, body := request.MondiaHTTPRequest(requestData)
					fmt.Println(string(body))
					if status != "error" {
						unsubNotification := new(models.MondiaCharge)
						xmlErr := xml.Unmarshal(body, unsubNotification)
						if xmlErr != nil {
							c.TplName = "fail.tpl"
							return
						}
						err := unsub.InsertUnsubData(unsubNotification)
						if err == nil && unsubNotification.ResponseCode == "1001" { // 取消订阅成功
							unsub.UpdateUnsubMoTable(subID)
							c.TplName = "success.tpl"
							return
						} else if err == nil && unsubNotification.ResponseCode == "3029" { // 之前已经取消
							c.TplName = "success.tpl"
							return
						}
					} else {
						//  "用户不存在"
						c.TplName = "fail.tpl"
						return
					}
				} else {
					//  "用户不存在"
					c.TplName = "fail.tpl"
					return
				}
			} else {
				c.TplName = "fail.tpl"
				return
			}
		} else {
			c.TplName = "fail.tpl"
			return
		}
	} else {
		c.Data["error"] = "201"
		c.Data["id"] = id
		c.Data["phone"] = msisdn
		c.TplName = "pin.tpl"
	}
}

//MondiaHTTPRequest 向mondia 发起http请求
func MondiaHTTPRequest(URL string) (status string, body []byte) {
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
