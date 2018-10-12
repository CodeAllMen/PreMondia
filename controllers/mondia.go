package controllers

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"

	"github.com/MobileCPX/PreMondia/models"
	"github.com/MobileCPX/PreMondia/util"

	"net/url"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

// MondiaNotificationController  接收订阅退订续订通知
type MondiaNotificationController struct {
	beego.Controller
}

// MondiaGetCustomerController 获取Customer Id
type MondiaGetCustomerController struct {
	beego.Controller
}

type MondiaSubscribeController struct {
	beego.Controller
}

//Post 接收订阅退订续订通知post请求
func (c *MondiaNotificationController) Post() {
	body := c.Ctx.Request.Body
	data, _ := ioutil.ReadAll(body)
	notification := models.Notification{}
	fmt.Println(string(data))
	err := xml.Unmarshal(data, &notification)
	fmt.Println(&notification)
	mo := new(models.MonMo)
	if err == nil { // 更新mo表（新增订阅，退订，续订）
		mo.Price = notification.Price
		mo.Operator = notification.Operator
		mo.SubscriptionId = notification.SubscriptionId
		mo.ServiceId = notification.ServiceId
		mo.CustomerId = notification.CustomerId
		mo.Channel = notification.Channel
		mo.PackageCode = notification.PackageCode
		mo.ProductCode = notification.ProductCode
		mo = models.UpdateOrInsertMo(notification.Action, notification.SubscriptionStatus, notification.Price, mo)
	}
	notification.ServiceType = mo.ServiceType
	notification.AffName = mo.AffName
	notification.ClickType = mo.ClickType
	notification.ClickId = mo.ClickId
	notification.PubId = mo.PubId
	notification.ProId = mo.ProId
	models.InsertCharge(notification)
	c.Ctx.WriteString("ok")
}

// Get 获取Customer Id   Get请求
func (c *MondiaGetCustomerController) Get() {
	CustomerModels := new(models.MdCustomer)
	serviceTypeClientID := c.Ctx.Input.Param(":id") //124|game_d 服务类型及id
	id := strings.Split(serviceTypeClientID, "|")[0]
	serviceType := strings.Split(serviceTypeClientID, "|")[1]
	status := c.GetString("status")
	customerID := c.GetString("customerId")
	operatorCountry := c.GetString("operator")
	errorDesc := c.GetString("errorDesc")
	errorCode := c.GetString("errorCode")
	idInt, _ := strconv.Atoi(id)

	affData := models.GetAffDataId(idInt) //根据id 查询点击信息（网盟，子渠道，clickId,点击类型）
	CustomerModels.AffName = affData.AffName
	CustomerModels.PubId = affData.PubId
	CustomerModels.PubId = affData.ProId
	CustomerModels.ClickType = affData.ClickType
	CustomerModels.ClickId = affData.ClickId

	CustomerModels.ErrorDesc = errorDesc
	CustomerModels.ErrorCode = errorCode
	CustomerModels.NewId = id
	CustomerModels.Status = status
	CustomerModels.CustomerId = customerID
	CustomerModels.Operator = operatorCountry
	CustomerModels.ServiceType = serviceType
	models.InsertCustomer(CustomerModels)
	productCd, subPackage, imgPath, prodPrice := GetProductCd(serviceType)
	subURL := models.GetCustomer(status, operatorCountry, serviceTypeClientID, id, productCd, subPackage, imgPath, prodPrice, serviceType)
	c.Redirect(subURL, 302)
}

func (c *MondiaSubscribeController) Get() {
	new_id := c.Ctx.Input.Param(":id")
	status := c.GetString("status")
	customerId := c.GetString("customerId")
	subId := c.GetString("subId")
	nextAction := c.GetString("nextAction")
	subStatus := c.GetString("subStatus")
	nextActionDate := c.GetString("nextActionDate")
	errorCode := c.GetString("errorCode")
	errorDesc := c.GetString("errorDesc")
	viewName := c.GetString("viewName")
	sub := new(models.MdSubscribe)
	sub.ClientId = new_id
	sub.SubStatus = subStatus
	sub.CustomerId = customerId
	sub.Status = status
	sub.ErrorCode = errorCode
	sub.NextAction = nextAction
	sub.NextActionDate = nextActionDate
	sub.SubId = subId
	sub.ViewName = viewName
	sub.ErrorDesc = errorDesc
	aff_data := models.IdGetServiceType(new_id)
	sub.ClickId = aff_data.ClickId
	sub.AffName = aff_data.AffName
	sub.PubId = aff_data.PubId
	sub.ServiceType = aff_data.ServiceType
	fmt.Println(aff_data.ServiceType)
	fmt.Println(sub.ServiceType)
	models.InsertSubscribe(sub)
	if (status == "SUCCESS" || errorCode == "3001") && (aff_data.ServiceType == "game_d" || aff_data.ServiceType == "game_w") && subStatus == "ACTIVE" {
		c.Redirect("http://www.gogamehub.com/mm/eg", 302)
	} else if (status == "SUCCESS" || errorCode == "3001") && (aff_data.ServiceType == "video_d" || aff_data.ServiceType == "video_w") && subStatus == "ACTIVE" {
		util.HttpRequest(sub.SubId, "register", "video", sub.CustomerId, "")
		c.Redirect("http://www.redlightvideos.com/mm/pl?sub="+subId, 302)
	} else if errorCode == "2004" {
		c.Redirect("http://za.mobpre.com/static/operator/index.html?type="+aff_data.ServiceType+"&id="+new_id, 302)
	} else if subStatus == "SUSPENDED" || subStatus == "UNSUBSCRIBED" {
		c.Redirect("https://www.google.com", 302)
	} else {
		if aff_data.ServiceType == "game_d" || aff_data.ServiceType == "game_w" {
			c.Redirect("http://www.gogamehub.com/lp/tn/uk/index.html?affName=Slef", 302)
		} else {
			c.Redirect("http://www.redlightvideos.com/lp/mm/pl/index.html?affName=Slef", 302)
		}
	}
}

func GetProductCd(service_type string) (string, string, string, string) {
	productCd := ""
	subPackage := ""
	imgPath := ""
	prodPrice := ""
	if service_type == "game_d" {
		subPackage = "D"
		prodPrice = "2.0"
		productCd = "GOGAMEHUB"
		imgPath = url.QueryEscape("http://in.mobpre.com/lp/game/aa/video/img/bg.png")
	} else if service_type == "game_w" {
		productCd = "GOGAMEHUB"
		subPackage = "W"
		prodPrice = "2.0"
		imgPath = url.QueryEscape("http://in.mobpre.com/lp/game/aa/video/img/bg.png")
	} else if service_type == "video_d" {
		productCd = "REDLIGHTVIDEOS"
		subPackage = "D"
		prodPrice = "2.0"
		imgPath = url.QueryEscape("http://www.redlightvideos.com/lp/tn/uk/img/banner.jpg")
	} else if service_type == "video_w" {
		productCd = "REDLIGHTVIDEOS"
		subPackage = "W"
		prodPrice = "9.9"
		imgPath = url.QueryEscape("http://www.redlightvideos.com/lp/tn/uk/img/banner.jpg")
	}
	return productCd, subPackage, imgPath, prodPrice
}
