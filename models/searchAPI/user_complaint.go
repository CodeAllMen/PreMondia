package searchAPI

import (
	"fmt"

	"github.com/MobileCPX/PreMondia/models"
	"github.com/MobileCPX/PreMondia/util"
	"github.com/astaxie/beego/orm"
	//"database/sql"
)

type GetComplaintData struct {
	Msisdn         string
	Email          string
	UserName       string
	DealWithTime   string
	EquipmentModel string
	GuiltyAffName  string
	GuiltyPubid    string
	Level          string
	Description    string
}

func GetComplaintsData(msisdn string) (status bool, moData []models.ComplaintData) {
	o := orm.NewOrm()
	var complaint_data models.ComplaintData
	sql := fmt.Sprintf("select b.click_id,b.postback_status,b.aff_name,b.pub_id, b.operator,b.msisdn,a.subscription_id as sub_id,"+
		"b.sub_num as mt_num, round(b.sub_num::numeric*4.5,1) as Amount, b.subtime,b.unsubtime from"+
		" nth_charge as a left join nth_mo as b on a.subscription_id=b.subscription_id where a.command = 'deliverSessionState' AND a.status_number = '2' "+
		"and b.msisdn='%s';", msisdn)
	o.Raw(sql).QueryRows(&moData)
	o.QueryTable("complaint_data").Filter("msisdn", msisdn).One(&complaint_data)
	if complaint_data.SubId != "" {
		for i, _ := range moData {
			moData[i].UserName = complaint_data.UserName
			moData[i].Email = complaint_data.Email
			moData[i].GuiltyPubid = complaint_data.GuiltyPubid
			moData[i].GuiltyAffName = complaint_data.GuiltyAffName
			moData[i].EquipmentModel = complaint_data.EquipmentModel
			moData[i].Description = complaint_data.Description
			moData[i].Level = complaint_data.Level
			moData[i].DealWithTime = complaint_data.DealWithTime
		}
	}
	if len(moData) != 0 {
		status = true
	}
	return
}

func AddUserComplaint(addDate GetComplaintData, complainTime string) (status bool) {
	o := orm.NewOrm()
	_, date := util.GetFormatTime()
	if complainTime != "" {
		date = complainTime
	}
	var moData []models.ComplaintData
	sql := fmt.Sprintf("select b.click_id,b.service_type as click_type,b.postback_status,b.aff_name,b.pub_id, b.operator,b.msisdn,a.subscription_id as sub_id,"+
		"b.sub_num as mt_num, round(b.sub_num::numeric*4.5,1) as Amount,b.service_type, b.subtime,b.unsubtime from"+
		" nth_charge as a left join nth_mo as b on a.subscription_id=b.subscription_id where a.command = 'deliverSessionState' AND a.status_number = '2' "+
		"and b.msisdn='%s';", addDate.Msisdn)
	o.Raw(sql).QueryRows(&moData)
	if len(moData) != 0 {
		for _, v := range moData {
			var oneComplainDate models.ComplaintData
			o.QueryTable("complaint_data").Filter("sub_id", v.SubId).One(&oneComplainDate)
			if oneComplainDate.SubId != "" {
				oneComplainDate.Email = addDate.Email
				oneComplainDate.UserName = addDate.UserName
				oneComplainDate.GuiltyPubid = addDate.GuiltyPubid
				oneComplainDate.GuiltyAffName = addDate.GuiltyAffName
				oneComplainDate.EquipmentModel = addDate.EquipmentModel
				oneComplainDate.Description = addDate.Description
				oneComplainDate.Level = addDate.Level
				oneComplainDate.DealWithTime = addDate.DealWithTime
				o.Update(&oneComplainDate)
			} else {
				v.Date = date
				v.Email = addDate.Email
				v.UserName = addDate.UserName
				v.GuiltyPubid = addDate.GuiltyPubid
				v.GuiltyAffName = addDate.GuiltyAffName
				v.EquipmentModel = addDate.EquipmentModel
				v.Description = addDate.Description
				v.Level = addDate.Level
				v.DealWithTime = addDate.DealWithTime
				o.Insert(&v)
			}
		}
		status = true
	}
	return
}

type ComplaintResult struct {
	Date          string
	ComplaintList []models.ComplaintData
}

type TotalAffComplaint struct {
	AffName      string
	ComplaintNum int
	PostbackNum  int
	MtNum        int
	Amount       float32
	Level_1      int
	Level_2      int
	Level_3      int
}

func SearchComplaintsData(StartTime, endTime, affName, operator, pubId, serviceType, clickType,
	msisdn, level string) ([]ComplaintResult, []TotalAffComplaint) {
	o := orm.NewOrm()
	var (
		complaintData     []models.ComplaintData
		resultData        []ComplaintResult
		oneData           ComplaintResult
		date              string
		DateComplaintlist []models.ComplaintData
	)

	sql := o.QueryTable("complaint_data")
	filterSql := ""
	if msisdn == "" {
		sql = sql.Filter("date__gte", StartTime).Filter("date__lte", endTime)
		if operator != "All" {
			filterSql += fmt.Sprintf(" and operator='%s'", operator)
			sql = sql.Filter("operator", operator)
		}
		if serviceType != "All" {
			filterSql += fmt.Sprintf(" and service_type='%s'", serviceType)
			sql = sql.Filter("service_type", serviceType)
		}
		if affName != "All" {
			filterSql += fmt.Sprintf(" and aff_name='%s'", affName)
			sql = sql.Filter("aff_name", affName)
			if pubId != "All" {
				filterSql += fmt.Sprintf(" and pub_id='%s'", pubId)
				sql = sql.Filter("pub_id", pubId)
			}
		}
		if level != "All" {
			filterSql += fmt.Sprintf(" and level='%s'", level)
			sql = sql.Filter("level", level)
		}
		if clickType != "All" {
			filterSql += fmt.Sprintf(" and click_type='%s'", clickType)
			sql = sql.Filter("click_type", clickType)
		}
	} else {
		sql = sql.Filter("msisdn__contains", msisdn)
	}
	sql.OrderBy("-date").All(&complaintData)

	for i, v := range complaintData {
		if v.Unsubtime == "" { // 如果查询到用户还没有退订，再重新查一下mo表，更新数据
			var mo models.NthMo
			o.QueryTable("nth_mo").Filter("subscription_id", v.SubId).One(&mo)
			if mo.Unsubtime != "" {
				v.Unsubtime = mo.Unsubtime
				v.MtNum = mo.SubNum
				v.Amount = float32(mo.SubNum) * 4.5
				o.Update(&v)
			}
		}
		if date == "" {
			date = v.Date
		}
		if date != v.Date {
			oneData.Date = date
			oneData.ComplaintList = DateComplaintlist
			resultData = append(resultData, oneData)
			DateComplaintlist = []models.ComplaintData{}
			date = v.Date
			DateComplaintlist = append(DateComplaintlist, v)
		} else {
			DateComplaintlist = append(DateComplaintlist, v)
		}
		if len(complaintData) == i+1 { //查询最后一条数据
			oneData.Date = v.Date
			oneData.ComplaintList = DateComplaintlist
			resultData = append(resultData, oneData)
		}
	}

	var totalAffComplain []TotalAffComplaint
	sqlTotal := fmt.Sprintf("select aff_name, count(*) as Complaint_num, count(case when postback_status='1' then 1 else null end) as Postback_num,"+
		" sum(mt_num) as Mt_num, sum(mt_num) * 4.5 as Amount, count(case when level='1' then 1 else null end) as Level_1,"+
		"count(case when level='2' then 1 else null end) as level_2,count(case when level='3' then 1 else null end) "+
		"as level_3 from complaint_data where date>='%s' and date<='%s' %s group by aff_name", StartTime, endTime, filterSql)
	o.Raw(sqlTotal).QueryRows(&totalAffComplain)
	var allComplainNum TotalAffComplaint
	for _, v := range totalAffComplain {
		allComplainNum.Amount += v.Amount
		allComplainNum.PostbackNum += v.PostbackNum
		allComplainNum.MtNum += v.MtNum
		allComplainNum.ComplaintNum += v.ComplaintNum
		allComplainNum.Level_1 += v.Level_1
		allComplainNum.Level_2 += v.Level_2
		allComplainNum.Level_3 += v.Level_3
	}
	allComplainNum.AffName = "Total"
	totalAffComplain = append(totalAffComplain, allComplainNum)
	copy(totalAffComplain[1:], totalAffComplain[0:len(totalAffComplain)-1])
	totalAffComplain[0] = allComplainNum
	return resultData, totalAffComplain
}
