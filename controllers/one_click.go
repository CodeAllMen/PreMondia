package controllers

import (
	"strconv"

	"github.com/MobileCPX/PreMondia/models"

	"github.com/astaxie/beego"
)

type MondiaReturnIdController struct {
	beego.Controller
}

func (this *MondiaReturnIdController) Get() {
	service := this.GetString("type")
	aff_name := this.GetString("affName")
	click_id := this.GetString("clickId")
	pubId := this.GetString("pubId")
	proId := this.GetString("proId")
	aff_model := new(models.MdId)
	aff_model.ServiceType = service
	aff_model.AffName = aff_name
	aff_model.ClickId = click_id
	aff_model.PubId = pubId
	aff_model.ProId = proId
	err, id := models.InsertIbId(aff_model)
	if err == nil {
		this.Ctx.WriteString(strconv.FormatInt(id, 10))
	} else {
		this.Ctx.WriteString(err.Error())
	}
}
