package main

import (
	_ "database/sql"
	"math/rand"

	"github.com/MobileCPX/PreMondia/controllers"
	_ "github.com/MobileCPX/PreMondia/initial"
	_ "github.com/MobileCPX/PreMondia/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	_ "github.com/astaxie/beego/session/postgres"
	_ "github.com/lib/pq"

	"fmt"

	_ "github.com/astaxie/beego/cache/redis"
	//"time"
	"time"
)

func main() {

	// redis, err := cache.NewCache("redis", `{"conn":"127.0.0.1:6379", "key":"1112221"}`)
	// fmt.Println(redis)
	// s := redis.Get("ip")
	//if s == ""{
	//	redis.Put("ip",1,1000*time.Second)
	//}else {
	//	if s > 3{
	//		url := "google"
	//	}else{
	//		s =  s + 1
	//		redis.Put("ip",s,1000*time.Second)
	//	}
	//}
	// redis.Put("ip", "4", 1000*time.Second)
	// s = redis.Get("124124")
	// fmt.Println(redis, s)
	// newId, ok := s.(string)
	// fmt.Println(newId, ok)
	// if err != nil {
	// 	log.Println(err)
	// }
	// http://payment.mondiamediamena.com/billing-gw/subservice/sendsms?msisdn=48506541080&message=send_test&lang=pl
	// http://payment.mondiamediamena.com/billing-gw/service/getcustomer?msisdn=48506541080&operatorId=8
	controllers.MondiaHTTPRequest("http://payment.mondiamediamena.com/billing-gw/subservice/getcustomer?msisdn=48506541080&operatorId=8")
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprintf("%03v", rnd.Int31n(1000))
	fmt.Println(vcode)

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
		ExposeHeaders:   []string{"Content-Length", "Access-Control-Allow-Origin"},
	}))
	beego.SetStaticPath("/game", "/root/go/src/game")
	beego.Run()
}
