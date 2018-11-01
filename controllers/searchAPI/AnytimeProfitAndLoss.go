package searchAPI

import (
	"github.com/MobileCPX/PreMondia/models/searchAPI"

	"github.com/astaxie/beego"

	"fmt"
)

type AnytimeProfitAndLossController struct {
	beego.Controller
}

func (this *AnytimeProfitAndLossController) Get() {
	startSubDate := this.GetString("start_sub")
	endSubDate := this.GetString("end_sub")
	startDate := this.GetString("start_date")
	endDate := this.GetString("end_date")
	serviceType := this.GetString("service_type")
	operator := this.GetString("operator")
	aff_name := this.GetString("aff_name")
	pubId := this.GetString("pub_id")
	fmt.Println(startSubDate, endSubDate, startDate, endDate)
	status, tableData, chartData := searchAPI.GetSubscribeQualityModels1(startSubDate, endSubDate, startDate, endDate, aff_name, serviceType, pubId, operator)
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
