package searchAPI

import (
	"github.com/MobileCPX/PreMondia/models/searchAPI"

	"github.com/astaxie/beego"
)

type SubscribeQualityController struct {
	beego.Controller
}

func (this *SubscribeQualityController) Get() {
	date := this.GetString("sub_date")
	operator := this.GetString("operator")
	aff_name := this.GetString("aff_name")
	end_date := this.GetString("end_date")
	service_type := this.GetString("serverType")
	clickType := this.GetString("clickType")
	pubId := this.GetString("pub_id")
	status, tableData, chartData := searchAPI.GetSubscribeQualityModels(date, end_date, operator, aff_name, service_type, clickType, pubId)
	if status == false {
		var failedData searchAPI.SubResult
		failedData.Date = "未查询到数据"
		tableData = append(tableData, failedData)
	}

	this.Data["json"] =
		map[string]interface{}{
			"code":      1,
			"chartData": chartData,
			"tableData": tableData,
			"message":   "failed",
		}
	this.ServeJSON()
}
