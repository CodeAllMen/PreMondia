package main

import (
	_ "github.com/MobileCPX/PreMondia/initial"
	_ "github.com/MobileCPX/PreMondia/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/plugins/cors"
)

func main() {

	logs.SetLogger(logs.AdapterFile, `{"filename":"/mondia/logs/mondia.log","level":6,"maxlines":100000000,"daily":true,"maxdays":10000}`)
	logs.Async(1e3)
	logs.EnableFuncCallDepth(true)
	// postbackutil.PostbackRequest()
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
		ExposeHeaders:   []string{"Content-Length", "Access-Control-Allow-Origin"},
	}))
	beego.Run()
}
