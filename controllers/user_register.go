package controllers

import (
	"github.com/MobileCPX/PreMondia/models"
	"github.com/MobileCPX/PreMondia/util"

	"github.com/astaxie/beego"
	//"encoding/json"
	//"fmt"
)

type RegisterController struct {
	beego.Controller
}

type RegisterJson struct {
	Name string
	Sign string
}

func (this *RegisterController) Get() {
	//var user_data RegisterJson
	//json.Unmarshal(this.Ctx.Input.RequestBody, &user_data)
	//fmt.Println()
	//fmt.Println(user_data.Sign)
	user_name := this.GetString("name")
	clickId := this.GetString("sign")
	if clickId != "" {
		status, sub_data := models.CheckUserClickId(clickId)
		if status {
			if sub_data.ServiceType == "game_w" || sub_data.ServiceType == "game_d" {
				util.UserRequest(sub_data.CustomerId, user_name, "register_game") // 测试，先所以账号都注册

				util.UserRequest(sub_data.CustomerId, user_name, "register_video")
			} else {
				util.UserRequest(sub_data.CustomerId, user_name, "register_video")

				util.UserRequest(sub_data.CustomerId, user_name, "register_game")
			}
			this.Ctx.WriteString("yes")
		} else {
			this.Ctx.WriteString("no")
		}
	} else {
		this.Ctx.WriteString("no")
	}
}
