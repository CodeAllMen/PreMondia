package controllers

import (
	"github.com/MobileCPX/PreMondia/models"

	"github.com/astaxie/beego"
)

type AffController struct {
	beego.Controller
}

type SearceMoController struct {
	beego.Controller
}

type SearceAffMtController struct {
	beego.Controller
}

type GetPubIdController struct {
	beego.Controller
}

type SubscribeQualityController struct {
	beego.Controller
}

type SubscribeTotalDataController struct {
	beego.Controller
}

type EverydaySubscribeStatistics struct { //每日订阅数据统计
	beego.Controller
}

func (this *AffController) Get() {
	start_time := this.GetString("start_time")
	end_time := this.GetString("end_time")
	telco := this.GetString("operator")
	serverType := this.GetString("serverType")
	aff_name := this.GetString("aff_name")
	clickType := this.GetString("clickType")
	err, data := models.GetAffdDate(start_time, end_time, telco, serverType, aff_name, clickType)
	if err == nil {
		this.Data["json"] =
			map[string]interface{}{
				"code":    "1",
				"data":    data,
				"message": "success",
			}
		this.ServeJSON()
	} else {
		this.Data["json"] =
			map[string]interface{}{
				"code":    0,
				"message": "failed",
			}
		this.ServeJSON()
	}

}

func (this *GetPubIdController) Get() {
	aff_name := this.GetString("aff_name")
	result := models.GetPubIdModels(aff_name)
	this.Data["json"] =
		map[string]interface{}{
			"code":    1,
			"data":    result,
			"message": "success",
		}
	this.ServeJSON()
}

func (this *SearceAffMtController) Get() {
	start_time := this.GetString("start_time")
	end_time := this.GetString("end_time")
	aff_pub := this.GetString("pubid")
	service_type := this.GetString("service_type")
	operator := this.GetString("operator")
	aff_name := this.GetString("aff_name")
	clickType := this.GetString("clickType")
	totalSubData := models.GetAffMTData(service_type, start_time, end_time, operator, aff_pub, aff_name, clickType)
	this.Data["json"] =
		map[string]interface{}{
			"code":  1,
			"datas": totalSubData,
		}
	this.ServeJSON()
}

func (this *SubscribeQualityController) Get() {
	date := this.GetString("sub_date")
	operator := this.GetString("operator")
	aff_name := this.GetString("aff_name")
	end_date := this.GetString("end_date")
	service_type := this.GetString("serverType")
	clickType := this.GetString("clickType")
	QualityResult, total := models.GetSubscribeQualityModels(date, end_date, operator, aff_name, service_type, clickType)
	this.Data["json"] =
		map[string]interface{}{
			"code":    1,
			"data":    QualityResult,
			"total":   total,
			"message": "failed",
		}
	this.ServeJSON()
}

func (this *SubscribeTotalDataController) Get() {
	startSubDate := this.GetString("start_sub")
	endSubDate := this.GetString("end_sub")
	startDate := this.GetString("start_date")
	endDate := this.GetString("end_date")
	serviceType := this.GetString("service_type")
	operator := this.GetString("operator")
	aff_name := this.GetString("aff_name")
	pubId := this.GetString("pub_id")
	clickId := this.GetString("clickType")
	data := models.SearceMoDetailedData(aff_name, pubId, operator, serviceType, startSubDate, endSubDate, startDate, endDate, clickId)
	this.Data["json"] =
		map[string]interface{}{
			"code":    1,
			"data":    data,
			"message": "failed",
		}
	this.ServeJSON()
}

func (this *EverydaySubscribeStatistics) Get() {
	date := this.GetString("date")
	//endDate := this.GetString("endDate")
	//serviceType := this.GetString("serviceType")
	//clickType := this.GetString("clickType")
	data, column, lastMonthProfitAndLoss, lastMonthRevenue, lastMouthSpend := models.GetEveryDaySubData(date)
	this.Data["json"] =
		map[string]interface{}{
			"code":                   1,
			"affClickType":           column,
			"lastMonthProfitAndLoss": lastMonthProfitAndLoss,
			"lastMonthRevenue":       lastMonthRevenue,
			"lastMouthSpend":         lastMouthSpend,
			"data":                   data,
			"message":                "failed",
		}
	this.ServeJSON()
}
