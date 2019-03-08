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
	c.Data["service_id"] = c.Ctx.Input.Param(":serviceID")
	c.TplName = "unsub.html"
}

func (c *UnsubController) UnsubSendPin() {
	var requestData mondia.MondiaRequestData
	msisdn := c.GetString("msisdn")
	serviceID := c.Ctx.Input.Param(":serviceID")
	c.Data["service_id"] = serviceID
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	pin := fmt.Sprintf("%04v", rnd.Intn(10000))
	logs.Info("PIN: ", pin)

	serviceConfig := c.getServiceConfig(serviceID)

	requestData.ServiceID = serviceID
	requestData.RequestType = "SendSMS"
	requestData.Message = mondia.GetPINUnsubMessage(serviceID, pin)
	requestData.Msisdn = msisdn

	if requestData.Message == "" {
		c.Data["error"] = "0"
		c.TplName = "unsub.html"
		return
	}

	status, body := requestData.Request(serviceConfig)
	fmt.Println(string(body))
	if status == "error" {
		c.Data["error"] = "0"
		c.TplName = "unsub.html"
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
			c.TplName = "unsub.html"
		}
	} else {
		c.Data["error"] = "0"
		c.TplName = "unsub.html"
	}
}

func (c *UnsubController) UnsubRequest() {
	requestData := new(mondia.MondiaRequestData)
	pin := c.GetString("pin")
	id := c.GetString("id")
	serviceID := c.GetString("service_id")
	serviceConfig := c.getServiceConfig(serviceID)
	unsubPIN := new(mondia.UnsubPin)
	err := unsubPIN.CheckPIN(id)
	if err == nil {
		if unsubPIN.Pin == pin {

		}
	}
	if err == nil && unsubPIN.Pin == pin {
		requestData.Msisdn = unsubPIN.Msisdn
		requestData.RequestType = "GetCustomer"
		status, body := requestData.Request(serviceConfig)
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

					status, body := requestData.Request(serviceConfig)
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



func (c *UnsubController)setPINPage(){

}
