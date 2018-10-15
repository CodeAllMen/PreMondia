package main

import (
	"fmt"
	"net/url"

	"github.com/MobileCPX/PreMondia/initial"
	_ "github.com/MobileCPX/PreMondia/initial"
	"github.com/MobileCPX/PreMondia/request"
	_ "github.com/MobileCPX/PreMondia/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func main() {

	err, body := request.HTTPRequest("http://payment.mondiamediamena.com/billing-gw/subservice/sendsms?customerId=" + "177090195" + "&message=" + url.QueryEscape("test send sms") + "&lang=pl&operatorId=8")
	fmt.Println(err, string(body), "!!!!!!!!!!!!!")
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
