package unsub

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/MobileCPX/PreMondia/models"
	"github.com/MobileCPX/PreMondia/models/unsub"
	"github.com/astaxie/beego"
)

// UnsubRequestControllers 退订请求
type UnsubRequestControllers struct {
	beego.Controller
}

// MondiaCharge Mondia xml格式通知

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
