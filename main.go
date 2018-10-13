package main

import (
	"fmt"

	"github.com/MobileCPX/PreMondia/initial"
	_ "github.com/MobileCPX/PreMondia/initial"
	_ "github.com/MobileCPX/PreMondia/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func main() {
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
