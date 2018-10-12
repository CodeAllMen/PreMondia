package controllers

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/astaxie/beego/logs"

	"github.com/MobileCPX/PreMondia/models"

	"github.com/astaxie/beego"
)

//UnsubPage 退订页面
type UnsubPage struct {
	beego.Controller
}

// UnsubUserSendMsisdnGetPinController 用户发送电话号码获取pin退订
type UnsubUserSendMsisdnGetPinController struct {
	beego.Controller
}

// UnsubGetCustomer 退订获取CustomerId
type UnsubGetCustomer struct {
	beego.Controller
}

// Post 用户发送电话号码获取pin退订Post请求
func (c *UnsubUserSendMsisdnGetPinController) Post() {
	msisdn := c.GetString("msisdn")
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	pin := fmt.Sprintf("%03v", rnd.Intn(1000))
	logs.Info("PIN: ", pin)
	message := url.QueryEscape("[RedLightVideos] Your unsubscribe PIN code is " + pin)
	// message := pin
	getPinURL := "http://payment.mondiamediamena.com/billing-gw/subservice/sendsms?operatorId=8&lang=pl&msisdn=" + msisdn + "&message=" + message
	status, body := MondiaHTTPRequest(getPinURL)
	fmt.Println(string(body))
	if status == "error" {
		c.Data["error"] = "0"
		c.TplName = "unsub.tpl"
		return
	}
	fmt.Println(string(body))
	unsubPin := new(models.UnsubPin)
	unsubPin.Msisdn = msisdn
	unsubPin.Pin = pin
	unsubPin.PinStatus = string(body)
	id, err := models.InsertPinData(unsubPin)
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

// Post 退订获取CustomerId Post请求
func (c *UnsubGetCustomer) Post() {
	pin := c.GetString("pin")
	id := c.GetString("id")
	msisdn, _ := models.CheckPIN(pin, id)
	if msisdn != "" {
		// getCustomerURL := "http://payment.mondiamediamena.com/billing-gw/service/getcustomer?msisdn=" + msisdn + "&operatorId=8"
		getCustomerURL := "http://payment.mondiamediamena.com/billing-gw/service/getcustomer?msisdn=" + msisdn + "&operatorId=8"
		status, body := MondiaHTTPRequest(getCustomerURL)
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
				subID := models.CustomerToGetSubID(customerData.CustomerId, msisdn)
				if subID != "" {
					unsubURL := "http://payment.mondiamediamena.com/billing-gw/subservice/unsubscribe?subid=" + subID + "operatorId=8"
					status, body := MondiaHTTPRequest(unsubURL)
					fmt.Println(string(body))
					if status != "error" {
						unsubNotification := new(models.MondiaCharge)
						xmlErr := xml.Unmarshal(body, unsubNotification)
						if xmlErr != nil {
							c.TplName = "fail.tpl"
							return
						}
						err := models.InsertUnsubData(unsubNotification)
						if err == nil && unsubNotification.ResponseCode == "1001" {
							c.TplName = "success.tpl"
							return
						} else {
							c.TplName = "fail.tpl"
							return
						}
					} else {

					}
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

//Get 退订首页
func (c *UnsubPage) Get() {
	c.TplName = "unsub.tpl"
}

// CheckLoginMsisdnController 根据电话号码请求mondia后台判断此用户是否订阅服务
type CheckLoginMsisdnController struct {
	beego.Controller
}

// Get 根据电话号码请求mondia后台判断此用户是否订阅服务 Get 请求
func (c *CheckLoginMsisdnController) Get() {
	returnStatus := "ERROR"
	msisdn := c.GetString("msisdn")
	getCustomerURL := "http://payment.mondiamediamena.com/billing-gw/service/getcustomer?msisdn=" + msisdn + "&operatorId=8"
	status, body := MondiaHTTPRequest(getCustomerURL)
	if status != "error" {
		customerData := new(models.UnsubGetCustomer)
		customerErr := xml.Unmarshal(body, customerData)
		if customerErr == nil {
			if customerData.ResponseCode == "1001" {
				subID := models.CustomerToGetSubID(customerData.CustomerId, msisdn)
				logs.Info("subID: ", subID)
				if subID != "" {
					returnStatus = "OK"
				}
			}
		}
	}
	c.Ctx.WriteString(returnStatus)
}
