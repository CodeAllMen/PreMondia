package models

import (
	"fmt"
	//"sort"

	"github.com/MobileCPX/PreMondia/util"

	"strconv"

	"github.com/astaxie/beego/orm"
)

type ClickNumInfo struct {
	Datetime    string
	Aff_name    string
	Pub_id      string
	ServiceType string
	ClickType   string
	Click_num   int
}

func InsertClickData() {
	o := orm.NewOrm()
	var click_info []ClickNumInfo
	var max_date_click ClickData

	max_sql := "select * from click_data order by datetime desc limit 1"
	o.Raw(max_sql).QueryRow(&max_date_click)
	hoursTime := util.GetFormatHoursTime()
	sql := fmt.Sprintf("select left(sendtime,13) as Datetime,aff_name, service_type,pub_id,count(id) as "+
		"Click_num,click_type from md_id where left(sendtime,13)>'%s' and left(sendtime,13)<'%s' group by "+
		"aff_name, pub_id, service_type,left(sendtime,13),click_type order by Datetime", max_date_click.Datetime, hoursTime)

	o.Raw(sql).QueryRows(&click_info)
	for _, v := range click_info {
		if v.Click_num > 1 {
			var click_data ClickData
			click_data.ClickNum = v.Click_num
			click_data.Aff_name = v.Aff_name
			click_data.Datetime = v.Datetime
			click_data.Pub_id = v.Pub_id
			click_data.ClickType = v.ClickType
			click_data.ServiceType = v.ServiceType
			o.Insert(&click_data)
		}
	}
}

type TotalDayData struct {
	Date        string
	AffName     string
	SubType     string
	Operator    string
	Num         int
	PostbackNum int
	ClickType   string
}

type AffDayData struct {
	Date        string
	AffName     string
	SubType     string
	SubNum      int
	PostbackNum int
	UnsubNum    int
	MtNum       int
	FailedMt    int
	ClickType   string
}

type OperatorDayData struct {
	Date      string
	Operator  string
	MtNum     int
	ClickType string
}

func InsertEveryDaySubData() {
	o := orm.NewOrm()
	var data EveryDaySubDatas
	var affData []AffDayData
	var operatorData []OperatorDayData
	var allData []EveryDaySubDatas
	_, newDate := util.GetFormatTime()
	o.QueryTable("every_day_sub_datas").OrderBy("-date").One(&data)
	new_table_sql := fmt.Sprintf("select distinct b.aff_name,b.operator,a.sub_type,b.click_type,b.postback_status,"+
		"left(a.sendtime,10) as Date from tn_notification as a left join tn_mo as b on a.subscription_id = "+
		"b.subscription_id where ((a.sub_type in ('subRenew','subStart','subCancel') and charging_status='success') or a.sub_type='subMtFailed') and "+
		"left(sendtime,10)>'%s' and left(sendtime,10)<'%s' and a.sendtime>=b.subtime and b.service_type='game_w' ", data.Date, newDate)

	aff_sql := fmt.Sprintf("select date,aff_name,sub_type,click_type,count(case when sub_type='subStart' then 1 else null "+
		"end) as Sub_num,count(case when sub_type='subMtFailed' then 1 else null end) as Failed_mt, count(case when sub_type='subStart' and postback_status=1 then 1 else null end) as "+
		"postback_num, count(case when sub_type='subCancel' then 1 else null end ) as unsub_num,count(case when "+
		"sub_type='subRenew' then 1 else null end) as Mt_num from (%s) as t group by t.aff_name,t.date,click_type,t.sub_type"+
		" order by t.date,t.aff_name", new_table_sql)

	operator_sql := fmt.Sprintf("select date,operator,count(operator) as Mt_num "+
		"from (%s) as t where t.sub_type ='subRenew' group by t.operator,t.date order by t.date,t.operator", new_table_sql)

	o.Raw(aff_sql).QueryRows(&affData)
	o.Raw(operator_sql).QueryRows(&operatorData)
	active := data.Active
	totalRevenue := data.GrandTotalRevenue
	grandTotalSubNum := data.GrandTotalSub
	grandTotalSpend := data.GrandTotalSpend
	totalSuccessMt := data.GrandTotalSuccessMt
	totalFailedMtNum := data.GrandTotalFailedMtNum
	mtCharges := data.GrandTotalMtCharges
	date := ""
	oneDayData := EveryDaySubDatas{}
	for i, v := range operatorData {
		if v.Date != date {
			if date != "" {
				oneDayData.GrandTotalRevenue = totalRevenue
				allData = append(allData, oneDayData)
			}
			date = v.Date
			oneDayData = EveryDaySubDatas{}
			oneDayData.Date = v.Date
		}
		switch v.Operator {
		case "EEUK":
			oneDayData.Ee = v.MtNum
			totalRevenue += float32(v.MtNum) * 2.834 * 1.3516
		case "O2UK":
			oneDayData.O2 = v.MtNum
			totalRevenue += float32(v.MtNum) * 2.834 * 1.3516
		case "ORANGEUK":
			oneDayData.Orange = v.MtNum
			totalRevenue += float32(v.MtNum) * 2.834 * 1.3516
		case "THREEUK":
			oneDayData.Three = v.MtNum
			totalRevenue += float32(v.MtNum) * 2.834 * 1.3516
		case "TMOBILEUK":
			oneDayData.Tmobile = v.MtNum
			totalRevenue += float32(v.MtNum) * 2.834 * 1.3516
		case "VIRGINUK":
			oneDayData.Virgin = v.MtNum
			totalRevenue += float32(v.MtNum) * 2.834 * 1.3516
		case "VODAFONEUK":
			oneDayData.Vodafone = v.MtNum
			totalRevenue += float32(v.MtNum) * 2.834 * 1.3516
		}
		if len(operatorData) == i+1 {
			oneDayData.GrandTotalRevenue = totalRevenue
			allData = append(allData, oneDayData)
		}
	}

	date = ""

	var sub_1, unsub_1, postback_1, sub_2, unsub_2, postback_2, mt_1, mt_2, postbackSpend_1, postbackSpend_2 string
	//var failedMtNum,successMtNum,daySubNum, grandTotalSubNum, grandTotalSpend,price,daySpend,totalSuccessMt int
	var failedMtNum, successMtNum, daySubNum, price, daySpend int

	for j, aff := range affData {
		if aff.Date != date {
			if date != "" {
				for i, one := range allData {
					if one.Date == date {
						allData[i].SubData_1click = sub_1
						allData[i].UnsubData_1click = unsub_1
						allData[i].PostbackData_1click = postback_1
						allData[i].MtData_1click = mt_1
						allData[i].SubData_2click = sub_2
						allData[i].UnsubData_2click = unsub_2
						allData[i].PostbackData_2click = postback_2
						allData[i].PostbackSpend_1click = postbackSpend_1
						allData[i].PostbackSpend_2click = postbackSpend_2
						allData[i].MtData_2click = mt_2
						allData[i].ServiceType = "game_w"
						allData[i].Active = active
						allData[i].FailedMt = failedMtNum
						allData[i].SuccessMt = successMtNum
						allData[i].GrandTotalSub = grandTotalSubNum
						allData[i].GrandTotalSpend = grandTotalSpend
						allData[i].GrandTotalProfitAndLoss = allData[i].GrandTotalRevenue - float32(allData[i].GrandTotalSpend)
						allData[i].GrandTotalMtCharges = mtCharges
						allData[i].SubNum = daySubNum
						allData[i].DaySpend = daySpend
						allData[i].GrandTotalSuccessMt = totalSuccessMt
						allData[i].GrandTotalFailedMtNum = totalFailedMtNum
						allData[i].MtRate = fmt.Sprintf("%.2f", float32(successMtNum-daySubNum)/float32(successMtNum-daySubNum+failedMtNum)*100) + "%" // 扣费成功率
						allData[i].GrandTotalMtRate = fmt.Sprintf("%.2f", float32(totalSuccessMt)/float32(totalSuccessMt+totalFailedMtNum)*100) + "%"  // 扣费成功率
						o.Insert(&allData[i])
					}
				}
			}
			date = aff.Date
			sub_1 = ""
			unsub_1 = ""
			postback_1 = ""
			postbackSpend_1 = ""
			sub_2 = ""
			unsub_2 = ""
			postback_2 = ""
			postbackSpend_2 = ""
			mt_2 = ""
			mt_1 = ""
			failedMtNum = 0
			successMtNum = 0
			daySubNum = 0
			daySpend = 0
		}

		switch aff.SubType {
		case "subStart":
			grandTotalSubNum += aff.SubNum //累计订阅
			active += aff.SubNum
			daySubNum += aff.SubNum
			if aff.ClickType == "1_click" {
				sub_1 += aff.AffName + "-" + strconv.Itoa(aff.SubNum) + "|"
				if aff.AffName != "" && aff.AffName != "test_affName" {
					price = 8
					postbackSpend_1 += aff.AffName + "-" + strconv.Itoa(aff.PostbackNum*8) + "|"
				}
				postback_1 += aff.AffName + "-" + strconv.Itoa(aff.PostbackNum) + "|"
			} else if aff.ClickType == "2_click" && aff.AffName != "null" && aff.AffName != "test_affName" {
				sub_2 += aff.AffName + "-" + strconv.Itoa(aff.SubNum) + "|"

				if aff.AffName != "" && aff.AffName != "test_affName" {
					price = 10
					if aff.AffName == "Gotzha" {
						if aff.Date < "2017-10-28" && aff.Date >= "2017-10-20" {
							price = 12
						} else if aff.Date < "2017-10-20" {
							price = 10
						} else {
							price = 11
						}
					}
					postbackSpend_2 += aff.AffName + "-" + strconv.Itoa(aff.PostbackNum*price) + "|"
				}
				postback_2 += aff.AffName + "-" + strconv.Itoa(aff.PostbackNum) + "|"
			}
			grandTotalSpend = grandTotalSpend + aff.PostbackNum*price
			daySpend = daySpend + aff.PostbackNum*price
		case "subCancel":
			active = active - aff.UnsubNum
			if aff.ClickType == "1_click" {
				unsub_1 += aff.AffName + "-" + strconv.Itoa(aff.UnsubNum) + "|"
			} else {
				unsub_2 += aff.AffName + "-" + strconv.Itoa(aff.UnsubNum) + "|"
			}
		case "subRenew":
			totalSuccessMt += aff.MtNum
			successMtNum += aff.MtNum
			mtCharges = mtCharges + float32(aff.MtNum)*4.5*1.3516
			if aff.ClickType == "1_click" {
				mt_1 += aff.AffName + "-" + strconv.Itoa(aff.MtNum) + "|"
			} else {
				mt_2 += aff.AffName + "-" + strconv.Itoa(aff.MtNum) + "|"
			}
		case "subMtFailed":
			failedMtNum += aff.FailedMt
			totalFailedMtNum += aff.FailedMt
		}
		if len(affData) == j+1 {
			for i, one := range allData {
				if one.Date == date {
					allData[i].SubData_1click = sub_1
					allData[i].UnsubData_1click = unsub_1
					allData[i].PostbackData_1click = postback_1
					allData[i].MtData_1click = mt_1
					allData[i].SubData_2click = sub_2
					allData[i].UnsubData_2click = unsub_2
					allData[i].PostbackData_2click = postback_2
					allData[i].PostbackSpend_1click = postbackSpend_1
					allData[i].PostbackSpend_2click = postbackSpend_2
					allData[i].MtData_2click = mt_2
					allData[i].ServiceType = "game_w"
					allData[i].Active = active
					allData[i].FailedMt = failedMtNum
					allData[i].SuccessMt = successMtNum
					allData[i].GrandTotalSub = grandTotalSubNum
					allData[i].GrandTotalSpend = grandTotalSpend
					allData[i].GrandTotalProfitAndLoss = allData[i].GrandTotalRevenue - float32(allData[i].GrandTotalSpend)
					allData[i].SubNum = daySubNum
					allData[i].DaySpend = daySpend
					allData[i].GrandTotalMtCharges = mtCharges
					allData[i].GrandTotalSuccessMt = totalSuccessMt
					allData[i].GrandTotalFailedMtNum = totalFailedMtNum
					allData[i].MtRate = fmt.Sprintf("%.2f", float32(successMtNum-daySubNum)/float32(successMtNum-daySubNum+failedMtNum)*100) + "%" // 扣费成功率
					allData[i].GrandTotalMtRate = fmt.Sprintf("%.2f", float32(totalSuccessMt)/float32(totalSuccessMt+totalFailedMtNum)*100) + "%"  // 扣费成功率
					o.Insert(&allData[i])
				}
			}
		}
	}
}
