package main

import (
	"fmt"

	"github.com/MobileCPX/PreMondia/initial"
	_ "github.com/MobileCPX/PreMondia/initial"
	"github.com/MobileCPX/PreMondia/request"
	_ "github.com/MobileCPX/PreMondia/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/plugins/cors"
)

func main() {

	// Send 账号
	var requestData request.MondiaRequestData
	requestData.Message = "Witamy w RedLightVideos. Adres URL to http://www.redlightvideos.com/mm/pl. Twój numer konta to " + "51859481"
	requestData.RequestType = "SendSMS"
	requestData.CustomerID = "177090195"
	_, body := request.MondiaHTTPRequest(requestData)
	if string(body) == "OK" {
		logs.Info("订阅成功后发送账号成功")
	} else {
		logs.Info("订阅成功后发送账号失败")
	}

	// err, body := request.HTTPRequest("http://payment.mondiamediamena.com/billing-gw/subservice/sendsms?customerId=" + "177090195" + "&message=" + url.QueryEscape("test send sms") + "&lang=pl&operatorId=8")
	// fmt.Println(err, string(body), "!!!!!!!!!!!!!")
	sd := initial.GetMondiaConf()
	fmt.Println(sd)
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
		ExposeHeaders:   []string{"Content-Length", "Access-Control-Allow-Origin"},
	}))
	beego.Run()
}
