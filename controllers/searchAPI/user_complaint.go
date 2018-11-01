package searchAPI

import (
	"encoding/json"
	"io/ioutil"

	"github.com/MobileCPX/PreMondia/models"
	"github.com/MobileCPX/PreMondia/models/searchAPI"

	"github.com/astaxie/beego"
)

type GetNewUserComplaint struct {
	beego.Controller
}

type AddNewUserComplaint struct {
	beego.Controller
}

type SerachUserComplaintController struct {
	beego.Controller
}

func (this *GetNewUserComplaint) Get() {
	msisdn := this.GetString("msisdn")
	status, data := searchAPI.GetComplaintsData(msisdn)
	if status != true || msisdn == "" {
		var oneData models.ComplaintData
		data = []models.ComplaintData{}
		oneData.Msisdn = "没有此电话号码信息"
		data = append(data, oneData)
	}
	this.Data["json"] =
		map[string]interface{}{
			"code":    1,
			"data":    data,
			"message": "failed",
		}
	this.ServeJSON()
}

func (this *AddNewUserComplaint) Post() {
	data := this.Ctx.Request.Body
	body, _ := ioutil.ReadAll(data)
	var complaintData searchAPI.GetComplaintData
	json.Unmarshal(body, &complaintData)
	if complaintData.Email == "" || complaintData.UserName == "" {
		this.Data["json"] =
			map[string]interface{}{
				"code":    0,
				"message": "邮箱或者用户名为空",
			}
	} else {
		status := searchAPI.AddUserComplaint(complaintData, "")
		if status == true {
			this.Data["json"] =
				map[string]interface{}{
					"code":    1,
					"message": "插入数据成功",
				}
		} else {
			this.Data["json"] =
				map[string]interface{}{
					"code":    0,
					"message": "没有此电话号码信息: " + complaintData.Msisdn,
				}
		}
	}
	this.ServeJSON()
}

func (this *SerachUserComplaintController) Get() {
	startDate := this.GetString("start")
	endDate := this.GetString("end")
	serviceType := this.GetString("service_type")
	operator := this.GetString("operator")
	aff_name := this.GetString("aff_name")
	pubId := this.GetString("pub_id")
	clickType := this.GetString("clickType")
	msisdn := this.GetString("msisdn")
	level := this.GetString("level")
	result, totalData := searchAPI.SearchComplaintsData(startDate, endDate, aff_name, operator, pubId, serviceType, clickType, msisdn, level)
	this.Data["json"] =
		map[string]interface{}{
			"code":      1,
			"data":      result,
			"totalData": totalData,
		}
	this.ServeJSON()
}
