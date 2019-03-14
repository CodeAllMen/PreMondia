package main

import (
	"fmt"
	"github.com/MobileCPX/PreMondia/controllers/searchAPI"
	_ "github.com/MobileCPX/PreMondia/initial"
	"github.com/MobileCPX/PreMondia/models/mondia"
	_ "github.com/MobileCPX/PreMondia/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/robfig/cron"
	"strconv"
	"strings"
	"time"
)

func init() {
	mondia.InitServiceConfig()
}

/*
现在是测试环境，上线后需要替换MO表中的service_ID
 */

func main() {

	logs.SetLogger(logs.AdapterFile, `{"filename":"/mondia/logs/mondia.log","level":6,"maxlines":100000000,"daily":true,"maxdays":10000}`)
	logs.Async(1e3)
	logs.EnableFuncCallDepth(true)
	//sub.CheckDaySubNum(30)
	// postbackutil.PostbackRequest()
	postbackURL := "http://##auto_id##"
	postbackURL = strings.Replace(postbackURL, "##auto_id##", strconv.Itoa(int(time.Now().Unix())), -1)
	fmt.Println(postbackURL)

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
		ExposeHeaders:   []string{"Content-Length", "Access-Control-Allow-Origin"},
	}))
	task()
	searchAPI.AffClickData()
	beego.Run()
}

// 定时任务
func task() {
	cr := cron.New()
	// cr.AddFunc("0 5 7 */1 * ?", dcb.EveryDayBillingRequest)
	cr.AddFunc("0 0 */1 * * ?", searchAPI.AffClickData)

	// cr.AddFunc("0 2 */1 * * ?", dcb.StartBillingRequest) // 每一个小时统一扣一次费用
	// cr.AddFunc("0 1 0 */1 * ?", models.InsertEveryDaySubData)
	// cr.AddFunc("0 1 */1 * * ?", util.TimedTaskDeleteIPlist)
	// cr.AddFunc("0 5 0 */1 * ?", controllers.DailyInsertChartSubData)
	cr.Start()
}
