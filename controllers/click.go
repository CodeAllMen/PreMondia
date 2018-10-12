package controllers

import (
	"github.com/MobileCPX/PreMondia/models"

	"github.com/astaxie/beego"
)

type AffClickDataController struct {
	beego.Controller
}

type EveryDayInsertSubData struct {
	beego.Controller
}

func (this *AffClickDataController) Get() {
	models.InsertClickData()
	this.Ctx.WriteString("ok")
}

func (this *EveryDayInsertSubData) Get() {
	models.InsertEveryDaySubData()
}
