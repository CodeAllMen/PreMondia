package models

import (
	"fmt"
	//"sort"
	//"strings"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"

	"strings"

	"github.com/MobileCPX/PreMondia/util"
)

type SubCount struct {
	Aff_name   string
	Pub_id     string
	Pro_id     string
	Total_sub  int
	Active_sub int
	PostNum    int
	Unsub_num  int
}

type AffMoMtCilck struct {
	AffName  string
	Aff_data []PubData
}

type PubData struct {
	Pubname  string
	Ser_list []ProData
}

type ProData struct {
	Servername    string
	Total_num     int
	Active_num    int
	Unsub_num     int
	Click_num     int
	SuccessMT_Num int
	FailtMT_Num   int
	PostNum       int
	Churn_rate    string
}

type Data struct {
	Name     string
	Aff_data []Pub_name
}

type Pub_name struct {
	Pubname  string
	Ser_list []server
}

type server struct {
	Servername    string
	Total_num     int
	Active_num    int
	SuccessMT_Num int
	FailtMT_Num   int
	Unsub_num     int
	Click_num     int
	PostNum       int
	Churn_rate    string
}

type TotalSub struct {
	Total_sub int
}

type click struct {
	Click_num int
}

type GetNum struct {
	Num int
}

type SubQualityModels struct {
	Date              string
	Num               int
	Notification_type string
}

type MoMtClickData struct {
	AffName     string
	PubId       string
	ProId       string
	SubNum      int
	SuccessMt   int
	MtFailed    int
	UnsubNum    int
	PostbackNum int
	ClickNum    int
}

type ClickNum struct {
	Click int
}

type SubResult struct {
	Date        string
	TotalSubNum int
	ActivateNum int
	TotalMt     int
	MtFailed    int
	UnsubNum    int
	RenewNum    int
	PostbackNum int
	ClickNum    int
}

type TotalSubData struct {
	TotalSub      int
	TotalUnsub    int
	SuccessMt     int
	FailedMt      int
	TotalPostback int
	MtRate        string
}

type AffQualityData struct {
	Date        string
	SubNum      int
	SuccessMt   int
	MtFailed    int
	UnsubNum    int
	PostbackNum int
	ClickNum    int
}

func GetAffdDate(startTime, endTime, operator, service_type, aff_name, clickType string) (error, []AffMoMtCilck) {
	o := orm.NewOrm()
	o.Using("default")
	var (
		data                    []AffMoMtCilck  //查询返回的结果
		affData                 []MoMtClickData //根据网盟及子渠道分组查询SQL 返回的数据结构体
		total                   MoMtClickData   //  汇总total数据结构体  插入到  data 第一位
		clickFilter, filter_sql string          // sql语句筛选条件
		click_num               ClickNum        // 点击数量
	)

	// 查看是否查询所有运营商
	if operator != "All" {
		filter_sql = fmt.Sprintf(" and operator = '%s'", operator)
	}
	//  是否所有服务
	if service_type != "All" {
		filter_sql = filter_sql + fmt.Sprintf(" and service_type = '%s'", service_type)
		clickFilter = fmt.Sprintf(" and service_type = '%s'", service_type)
	}
	//  是否所有网盟
	if aff_name != "All" {
		filter_sql = filter_sql + fmt.Sprintf(" and aff_name = '%s'", aff_name)
		clickFilter = clickFilter + fmt.Sprintf(" and aff_name = '%s'", aff_name)
	}
	if clickType != "All" {
		filter_sql = filter_sql + fmt.Sprintf(" and click_type = '%s'", clickType)
		clickFilter = clickFilter + fmt.Sprintf(" and click_type = '%s'", clickType)
	}
	sql := fmt.Sprintf("select aff_name,pub_id,count(case when (t.action='SUBSCRIBE' or t.action='RENEW') and t.subscription_status='ACTIVE' then 1 else null end) as Success_Mt,"+
		"count(case when t.action='RENEW' and t.subscription_status<>'ACTIVE'  then 1 else null end) as Mt_Failed,"+
		"count(case when t.action='SUBSCRIBE' then 1 else null end) as Sub_Num,"+
		"count(case when t.action='UNSUBSCRIBE' and t.subscription_status='UNSUBSCRIBED' then 1 else null end) as Unsub_Num,"+
		"count(case when t.action='SUBSCRIBE' and t.postback_status = 1 then 1 else null end) as Postback_Num "+
		" from (select distinct b.aff_name,b.pub_id, action,subscription_status,a.subscription_id,b.postback_status,left(sendtime,10) from notification as a "+
		"left join mon_mo as b on a.subscription_id = b.subscription_id where sendtime>'%s' and sendtime<'%s' and subtime>'%s' and subtime<'%s' %s) as t group by t.aff_name,t.pub_id", startTime, endTime, startTime, endTime, filter_sql)

	//  根据网盟及子渠道分组查询 点击数量  clickNum
	clickGroupByAffSql := fmt.Sprintf("select aff_name,pub_id,sum(click_num) as click_num from click_data where datetime>'%s'"+
		" and datetime<'%s' %s group by aff_name,pub_id", startTime, endTime, clickFilter)

	// 查询数据
	var clickNum []MoMtClickData
	o.Raw(clickGroupByAffSql).QueryRows(&clickNum)

	_, err := o.Raw(sql).QueryRows(&affData)
	// 查询total click
	totalClickNumSql := fmt.Sprintf("select sum(click_num) as Click from click_data where datetime>'%s' and datetime<'%s' %s", startTime, endTime, clickFilter)
	o.Raw(totalClickNumSql).QueryRow(&click_num)

	affData = append(affData, clickNum...)
	//求汇总数据
	for _, subCharge := range affData {
		total.SubNum += subCharge.SubNum
		total.UnsubNum += subCharge.UnsubNum
		total.PostbackNum += subCharge.PostbackNum
		total.SuccessMt += subCharge.SuccessMt
		total.MtFailed += subCharge.MtFailed
	}

	//  将汇总数据插入到第一行
	total.AffName = "Total"
	total.PubId = "Total"
	total.ProId = "Total"
	total.ClickNum = click_num.Click
	affData = append(affData, total)
	copy(affData[1:], affData[0:len(affData)-1])
	affData[0] = total

	// 得到data结构体数据
	var affName, PubName string
	for i, subData := range affData {
		var oneData AffMoMtCilck
		var pubData PubData
		var serviceData ProData
		if affData[i].AffName != affName {
			affName = affData[i].AffName
			PubName = affData[i].PubId
			oneData.AffName = affName
			pubData.Pubname = affData[i].PubId
			serviceData.Servername = affData[i].ProId
			pubData.Ser_list = append(pubData.Ser_list, GetserviceDataList(subData))
			oneData.Aff_data = append(oneData.Aff_data, pubData)
			data = append(data, oneData)
		} else {
			if affData[i].PubId != PubName {
				PubName = affData[i].PubId
				oneData.AffName = affName
				pubData.Pubname = affData[i].PubId
				serviceData.Servername = affData[i].ProId
				pubData.Ser_list = append(pubData.Ser_list, GetserviceDataList(subData))
				for i, _ := range data {
					if data[i].AffName == affName {
						data[i].Aff_data = append(data[i].Aff_data, pubData)
						break
					}
				}
			} else {
				PubName = affData[i].PubId
				oneData.AffName = affName
				pubData.Pubname = affData[i].PubId
				serviceData.Servername = affData[i].ProId
				pubData.Ser_list = append(pubData.Ser_list, GetserviceDataList(subData))
				for i, _ := range data {
					if data[i].AffName == affName {
						for j, _ := range data[i].Aff_data {
							if data[i].Aff_data[j].Pubname == PubName {
								data[i].Aff_data[j].Ser_list = append(data[i].Aff_data[j].Ser_list, GetserviceDataList(subData))
								break
							}
						}
						break
					}
				}
			}
		}
	}
	return err, data
}

func GetserviceDataList(subData MoMtClickData) ProData {
	var service ProData
	service.Servername = subData.ProId
	service.PostNum = subData.PostbackNum
	service.Click_num = subData.ClickNum
	service.Total_num = subData.SubNum
	service.FailtMT_Num = subData.MtFailed
	service.SuccessMT_Num = subData.SuccessMt
	service.Unsub_num = subData.UnsubNum
	churn_rate := float32(service.Unsub_num) / float32(service.Total_num) * 100
	service.Churn_rate = fmt.Sprintf("%.2f", churn_rate) + "%" // 退订率百分比
	return service
}

type PubidList struct {
	PubId string
}

func GetPubIdModels(aff_name string) []string {
	o := orm.NewOrm()
	var pub_list []PubidList
	var pubList []string
	o.Raw(fmt.Sprintf("select DISTINCT pub_id from mon_mo where aff_name='%s'", aff_name)).QueryRows(&pub_list)
	for _, k := range pub_list {
		pubList = append(pubList, k.PubId)
	}
	return pubList
}

type SubscribeData struct {
	Sub_type string
	Num      int
}

func GetAffMTData(service_type, start, end, operator, pub_id, aff_name, clickType string) TotalSubData {
	var total TotalSubData
	o := orm.NewOrm()

	filterSql := ""
	if operator != "All" {
		filterSql = fmt.Sprintf(" and b.operator = '%s'", operator)
	}
	if service_type != "All" {
		filterSql += fmt.Sprintf(" and b.service_type = '%s'", service_type)
	}
	if aff_name != "All" {
		filterSql += fmt.Sprintf(" and b.aff_name = '%s'", aff_name)
		if pub_id != "All" {
			filterSql += fmt.Sprintf(" and b.pub_id = '%s'", pub_id)
		}
	}
	if clickType != "All" {
		filterSql += fmt.Sprintf(" and b.click_type = '%s'", clickType)
	}

	sql := fmt.Sprintf("select count(case when (t.action='SUBSCRIBE' or t.action='RENEW') and t.subscription_status='ACTIVE' then 1 else null end) as Success_mt,"+
		"count(case when t.action='RENEW' and t.subscription_status<>'ACTIVE'  then 1 else null end) as Failed_mt,"+
		"count(case when t.action='SUBSCRIBE' then 1 else null end) as Total_sub,"+
		"count(case when t.action='UNSUBSCRIBE' and t.subscription_status='UNSUBSCRIBED' then 1 else null end) as Total_unsub,"+
		"count(case when t.action='SUBSCRIBE' and t.postback_status = 1 then 1 else null end) as Total_postback "+
		" from (select distinct action,subscription_status,a.subscription_id,b.postback_status,left(sendtime,10) from notification as a "+
		"left join mon_mo as b on a.subscription_id = b.subscription_id where sendtime>'%s' and sendtime<'%s' %s) as t", start, end, filterSql)

	err := o.Raw(sql).QueryRow(&total)
	if err != nil {
		logs.Error("search data error")
	}
	churn_rate := float32(total.SuccessMt) / float32(total.SuccessMt+total.FailedMt) * 100
	total.MtRate = fmt.Sprintf("%.2f", churn_rate) + "%"
	return total
}

func GetSubscribeQualityModels(sub_time, end_time, opeator, aff_name, service_type, clickType string) ([]SubResult, SubResult) {
	var (
		total                SubResult
		result, total_result []SubResult
		charge               []AffQualityData
	)
	o := orm.NewOrm()
	add_filter := ""
	if opeator != "All" {
		add_filter = fmt.Sprintf("and operator = '%s'", opeator)
	}
	if aff_name != "All" {
		add_filter = add_filter + fmt.Sprintf("and aff_name = '%s'", aff_name)
	}
	if service_type != "All" {
		add_filter = add_filter + fmt.Sprintf("and service_type = '%s'", service_type)
	}
	if clickType != "All" {
		add_filter = add_filter + fmt.Sprintf("and click_type = '%s'", clickType)
	}
	sql := fmt.Sprintf("select date,count(case when (t.action='SUBSCRIBE' or t.action='RENEW') and t.subscription_status='ACTIVE' then 1 else null end) as Success_Mt,"+
		"count(case when t.action='RENEW' and t.subscription_status<>'ACTIVE'  then 1 else null end) as Mt_Failed,"+
		"count(case when t.action='SUBSCRIBE' then 1 else null end) as Sub_Num,"+
		"count(case when t.action='UNSUBSCRIBE' and t.subscription_status='UNSUBSCRIBED' then 1 else null end) as Unsub_Num,"+
		"count(case when t.action='SUBSCRIBE' and t.postback_status = 1 then 1 else null end) as Postback_Num "+
		" from (select distinct b.aff_name,b.pub_id, action,subscription_status,a.subscription_id,b.postback_status,left(sendtime,10) as date from notification as a "+
		"left join mon_mo as b on a.subscription_id = b.subscription_id where left(sendtime,10)>='%s' and left(sendtime,10)<='%s' and left(subtime,10)='%s'  %s) as t group by date order by date", sub_time, end_time, sub_time, add_filter)

	o.Raw(sql).QueryRows(&charge)
	var active, total_sub int
	for i, v := range charge {
		var one SubResult
		one.Date = v.Date
		one.TotalSubNum = active
		if i == 0 {
			one.TotalSubNum = v.SubNum
			total_sub = v.SubNum
			active = v.SubNum
		}
		active = active - v.UnsubNum
		one.ActivateNum = active
		one.MtFailed = v.MtFailed
		one.RenewNum = v.SuccessMt
		one.PostbackNum = v.PostbackNum
		one.UnsubNum = v.UnsubNum
		one.TotalMt = one.MtFailed + one.RenewNum
		total_result = append(total_result, one)
	}

	total.Date = "Total"
	total.TotalSubNum = total_sub
	total.ActivateNum = active
	total.UnsubNum = total_sub - active
	for _, v := range charge {
		total.MtFailed += v.MtFailed
		total.PostbackNum += v.PostbackNum
		total.RenewNum += v.SuccessMt
		total.TotalMt += v.MtFailed + v.SuccessMt
	}
	result = append(result, total)
	result = append(result, total_result...)
	return result, total
}

type EveryDateClickData struct {
	Date          string
	SubData       []string
	PostbackData  []string
	UnSubData     []string
	PostbackSpend []string
	MtNumData     []string
	Active        int
	SuccessMt     int
	MtRate        string
	Amout         string

	O2         string
	Ee         string
	DayRevenue string
	Orange     string
	Three      string
	Vodafone   string
	Tmobile    string
	Virgin     string

	TotalSuccessMt          int
	GrandTotalMtCharges     float32
	DaySpend                int
	GrandTotalSub           int //累计订阅
	GrandTotalProfitAndLoss string
	GrandTotalRevenue       string
	GrandTotalSpend         int
	GrandTotalMtRate        string
}

func GetEveryDaySubData(date string) ([]EveryDateClickData, []map[string][]string, string, string, int) {
	o := orm.NewOrm()
	var everyDate []EveryDaySubDatas
	var allClickData []EveryDateClickData
	affNameClickType := make(map[string][]string)

	var lastMonthData, twoLastMonthData EveryDaySubDatas
	lastMonth := util.GetLastMonth(-1)
	twoLastMonth := util.GetLastMonth(-2)
	lastSql := fmt.Sprintf("select * from every_day_sub_datas where left(date,7)='%s' order by date desc limit 1", lastMonth)
	twoLastSql := fmt.Sprintf("select * from every_day_sub_datas where left(date,7)='%s' order by date desc limit 1", twoLastMonth)
	o.Raw(lastSql).QueryRow(&lastMonthData)
	o.Raw(twoLastSql).QueryRow(&twoLastMonthData)
	lastMonthProfitAndLoss := fmt.Sprintf("%.3f", lastMonthData.GrandTotalProfitAndLoss-twoLastMonthData.GrandTotalProfitAndLoss)
	lastMonthRevenue := fmt.Sprintf("%.3f", lastMonthData.GrandTotalRevenue-twoLastMonthData.GrandTotalRevenue)
	lastMouthSpend := lastMonthData.GrandTotalSpend - twoLastMonthData.GrandTotalSpend

	sql := fmt.Sprintf("select * from every_day_sub_datas where left(date,7)='%s' order by date", date)
	click_1_affNameList := [][]string{}
	click_2_affNameList := [][]string{}
	o.Raw(sql).QueryRows(&everyDate)
	for _, i := range everyDate {
		click_1_sub := strings.Split(i.SubData_1click, "|")
		click_1_postback := strings.Split(i.PostbackData_1click, "|")
		click_1_unsub := strings.Split(i.UnsubData_1click, "|")
		click_1_mt := strings.Split(i.MtData_1click, "|")
		click_1_affNameList = append(click_1_affNameList, click_1_sub, click_1_postback, click_1_unsub, click_1_mt)
		click_2_sub := strings.Split(i.SubData_2click, "|")
		click_2_postback := strings.Split(i.PostbackData_2click, "|")
		click_2_unsub := strings.Split(i.UnsubData_2click, "|")
		click_2_mt := strings.Split(i.MtData_2click, "|")
		click_2_affNameList = append(click_2_affNameList, click_2_sub, click_2_postback, click_2_unsub, click_2_mt)
	}
	affNameList_1click := []string{}
	affNameList_2click := []string{}
	for _, v := range click_1_affNameList {
		for _, v1 := range v {
			affNameList_1click = append(affNameList_1click, strings.Split(v1, "-")[0])
		}
	}

	for _, v := range click_2_affNameList {
		for _, v1 := range v {
			affNameList_2click = append(affNameList_2click, strings.Split(v1, "-")[0])
		}
	}
	click_1 := util.Duplicate(affNameList_1click)
	click_2 := util.Duplicate(affNameList_2click)
	for _, v := range click_1 {
		clickType := []string{"1_click"}
		if v == "" {
			v = "null"
		}
		affNameClickType[v] = clickType
	}

	for _, affName2 := range click_2 {
		clickType := []string{"2_click"}
		if affName2 == "" {
			affName2 = "null"
		}
		status := ""
		for i, _ := range affNameClickType {
			if i == affName2 {
				affNameClickType[i] = append(affNameClickType[i], "2_click")
				status = "ok"
				break
			}
		}
		if status != "ok" {
			affNameClickType[affName2] = clickType
		}
	}

	var affNameClickList []map[string][]string

	for i, v := range affNameClickType {
		maps := make(map[string][]string)
		maps[i] = v
		affNameClickList = append(affNameClickList, maps)
	}

	for _, oneData := range everyDate {
		var oneClickData EveryDateClickData
		oneClickData.Date = oneData.Date
		oneClickData.MtRate = oneData.MtRate
		oneClickData.Active = oneData.Active
		oneClickData.SuccessMt = oneData.SuccessMt
		oneClickData.GrandTotalSub = oneData.GrandTotalSub

		oneClickData.Amout = fmt.Sprintf("%.3f", oneData.GrandTotalMtCharges)
		oneClickData.O2 = fmt.Sprintf("%.3f", float32(oneData.O2)*3.104*1.3516)
		oneClickData.Orange = fmt.Sprintf("%.3f", float32(oneData.Orange)*2.773*1.3516)
		oneClickData.Tmobile = fmt.Sprintf("%.3f", float32(oneData.Tmobile)*2.834*1.3516)
		oneClickData.Three = fmt.Sprintf("%.3f", float32(oneData.Three)*2.622*1.3516)
		oneClickData.Virgin = fmt.Sprintf("%.3f", float32(oneData.Virgin)*1.943*1.3516)
		oneClickData.Vodafone = fmt.Sprintf("%.3f", float32(oneData.Vodafone)*2.879*1.3516)
		oneClickData.Ee = fmt.Sprintf("%.3f", float32(oneData.Ee)*2.834*1.3516)

		oneClickData.DayRevenue = fmt.Sprintf("%.3f", (float32(oneData.O2)*3.104+float32(oneData.Orange)*2.773+
			float32(oneData.Tmobile)*2.834+float32(oneData.Three)*2.622+float32(oneData.Virgin)*1.943+
			float32(oneData.Vodafone)*2.879+float32(oneData.Ee)*2.834)*1.3516)
		oneClickData.GrandTotalRevenue = fmt.Sprintf("%.3f", oneData.GrandTotalRevenue)
		oneClickData.GrandTotalSpend = oneData.GrandTotalSpend
		oneClickData.GrandTotalProfitAndLoss = fmt.Sprintf("%.3f", oneData.GrandTotalProfitAndLoss)
		oneClickData.DaySpend = oneData.DaySpend
		oneClickData.GrandTotalMtCharges = oneData.GrandTotalMtCharges
		oneClickData.TotalSuccessMt = oneData.GrandTotalSuccessMt
		oneClickData.GrandTotalMtRate = oneData.GrandTotalMtRate

		click_1_sub := addClickTypeMap(strings.Split(oneData.SubData_1click, "|"))
		click_2_sub := addClickTypeMap(strings.Split(oneData.SubData_2click, "|"))
		click_1_post := addClickTypeMap(strings.Split(oneData.PostbackData_1click, "|"))
		click_2_post := addClickTypeMap(strings.Split(oneData.PostbackData_2click, "|"))
		click_1_unsub := addClickTypeMap(strings.Split(oneData.UnsubData_1click, "|"))
		click_2_unsub := addClickTypeMap(strings.Split(oneData.UnsubData_2click, "|"))
		click_1_spend := addClickTypeMap(strings.Split(oneData.PostbackSpend_1click, "|"))
		click_2_spend := addClickTypeMap(strings.Split(oneData.PostbackSpend_2click, "|"))
		click_1_MtNum := addClickTypeMap(strings.Split(oneData.MtData_1click, "|"))
		click_2_MtNum := addClickTypeMap(strings.Split(oneData.MtData_2click, "|"))

		oneClickData.SubData = getTotalAffClickTypeSub(click_1_sub, click_2_sub, affNameClickList)
		oneClickData.PostbackData = getTotalAffClickTypeSub(click_1_post, click_2_post, affNameClickList)
		oneClickData.UnSubData = getTotalAffClickTypeSub(click_1_unsub, click_2_unsub, affNameClickList)
		oneClickData.PostbackSpend = getTotalAffClickTypeSub(click_1_spend, click_2_spend, affNameClickList)
		oneClickData.MtNumData = getTotalAffClickTypeSub(click_1_MtNum, click_2_MtNum, affNameClickList)

		allClickData = append(allClickData, oneClickData)
	}
	return allClickData, affNameClickList, lastMonthProfitAndLoss, lastMonthRevenue, lastMouthSpend
}

func getTotalAffClickTypeSub(click_1, click_2 map[string]string, affClickData []map[string][]string) []string {
	var affClickList []string
	var num string
	for _, affClickData := range affClickData {
		for aff_name, clickTypeList := range affClickData {
			for _, clickType := range clickTypeList {
				if clickType == "1_click" {
					num = click_1[aff_name]
				} else {
					num = click_2[aff_name]
				}
				if num == "" {
					num = "0"
				}
				affClickList = append(affClickList, num)
			}

		}
	}
	return affClickList
}

func addClickTypeMap(lists []string) map[string]string {
	affDateClick := make(map[string]string)
	for _, v := range lists {
		splitResult := strings.Split(v, "-")
		aff_name := splitResult[0]
		if aff_name == "" {
			aff_name = "null"
		}
		if len(splitResult) == 2 {
			affDateClick[aff_name] = splitResult[1]
		}
	}
	return affDateClick
}

type DateSubDetailed struct {
	AffName     string
	Operator    string
	ServiceName string
	SubNum      int
	UnsubNum    int
	SuccessMt   int
	FailedMt    int
	PostbackNum int
}

type SubOperatorQuality struct {
	Operator    string
	ServiceList []Service
}

type Service struct {
	ServiceName string
	Price       float32
	AffSubData  []AffSub
}

type AffSub struct {
	AffName     string
	SubNum      int
	UnsubNum    int
	SuccessMt   int
	FailedMt    int
	PostbackNum int
	Amount      float32
}

// 查询根据任意时间订阅的用户在任意时间的订阅，扣费，续订情况  例如11月份订阅的用户，查询12月的扣费情况
//     逻辑：
// 		先通过连表查询（worldplay_charge表和world_play_mo表）查询出a.charging_status,a.sendtime,a.msisdn,a.service_name,a.action,a.operator,b.postback_status,a.renewal_times,b.aff_name，：作为新的表
// 	再根据group by  aff_name,operator  分组查询 得出 aff_name,operator,subNum,unsubNum,postbackNum,SuccessMt,FailedMt
func SearceMoDetailedData(affName, pubId, operator_name, serviceType, subStartDate, subEndDate, startDate, endDate, clickType string) []SubOperatorQuality {
	o := orm.NewOrm()
	sqlFilter := ""
	if affName != "All" {
		sqlFilter = fmt.Sprintf(" and b.aff_name = '%s'", affName)
		//if pubId != "All"{
		//	sqlFilter +=  fmt.Sprintf(" and b.pub_id='%s'",pubId)
		//}
	}
	if operator_name != "All" {
		sqlFilter += fmt.Sprintf(" and b.operator = '%s'", operator_name)
	}
	if serviceType != "All" {
		sqlFilter += fmt.Sprintf(" and b.service_type = '%s'", serviceType)
	}
	if clickType != "All" {
		sqlFilter += fmt.Sprintf(" and b.click_type = '%s'", clickType)
	}

	//sql := fmt.Sprintf("select date,count(case when (t.action='SUBSCRIBE' or t.action='RENEW') and t.subscription_status='ACTIVE' then 1 else null end) as Success_Mt," +
	//	"count(case when t.action='RENEW' and t.subscription_status<>'ACTIVE'  then 1 else null end) as Mt_Failed," +
	//	"count(case when t.action='SUBSCRIBE' then 1 else null end) as Sub_Num," +
	//	"count(case when t.action='UNSUBSCRIBE' and t.subscription_status='UNSUBSCRIBED' then 1 else null end) as Unsub_Num," +
	//	"count(case when t.action='SUBSCRIBE' and t.postback_status = 1 then 1 else null end) as Postback_Num " +
	//	" from (select distinct b.aff_name,b.pub_id, action,subscription_status,a.subscription_id,b.postback_status,left(sendtime,10) as date from notification as a " +
	//	"left join mon_mo as b on a.subscription_id = b.subscription_id where left(sendtime,10)>='%s' and left(sendtime,10)<='%s' and left(subtime,10)='%s'  %s) as t group by date order by date",sub_time,end_time,sub_time,add_filter)

	subNum_sql := "count(case when t.action='SUBSCRIBE' then 1 else null end)"
	unsubNum_sql := "count(case when t.action='UNSUBSCRIBE' and t.subscription_status='UNSUBSCRIBED' then 1 else null end)"
	postbackNum_sql := "count(case when t.action='SUBSCRIBE' and t.postback_status = 1 then 1 else null end)"
	successMtNum_sql := "count(case when (t.action='SUBSCRIBE' or t.action='RENEW') and t.subscription_status='ACTIVE' then 1 else null end) "
	failedMtNum_sql := "count(case when t.action='RENEW' and t.subscription_status<>'ACTIVE'  then 1 else null end)"

	total := fmt.Sprintf("select t.aff_name,t.operator,t.service_type as service_name, %s as sub_num, %s as unsub_num, %s as "+
		"success_mt, %s as failed_mt, %s as postback_num from (select distinct b.aff_name,b.pub_id,b.operator,b.service_type, action,subscription_status,a.subscription_id,b.postback_status,left(sendtime,10) as date from notification as a "+
		"left join mon_mo as b on a.subscription_id = b.subscription_id where a.sendtime<'%s' and "+
		"a.sendtime>'%s' and b.subtime>'%s' and b.subtime<'%s' %s ) as t  group by t.aff_name,t.operator,t.service_type order by t.operator,t.service_type;",
		subNum_sql, unsubNum_sql, successMtNum_sql, failedMtNum_sql, postbackNum_sql, endDate, startDate, subStartDate, subEndDate, sqlFilter)
	fmt.Println(total)
	var subData []DateSubDetailed
	o.Raw(total).QueryRows(&subData)
	var operator, operator_service, operator_tatal_name string
	operator = "xx"
	var allgame_W, allvideo_W, gameMtNum_W, videoMtNum_W, allgame_D, allvideo_D, gameMtNum_D, videoMtNum_D, service_index, SubOperatorData_index int

	servicePrice := map[string]float32{"game_w": 4.5, "video_w": 4.5, "game_d": 0.25, "video_d": 0.25}
	lengthSubdata := len(subData)
	var totalOperatorData, AllTotalData AffSub
	var OperatorService Service
	var SubOperatorData []SubOperatorQuality

	var AllTotalOperatorData SubOperatorQuality

	for i, v := range subData {
		var OneOperator SubOperatorQuality
		var OneService Service
		var OneAffData AffSub
		OneAffData.AffName = v.AffName
		OneAffData.SubNum = v.SubNum
		OneAffData.UnsubNum = v.UnsubNum
		OneAffData.SuccessMt = v.SuccessMt
		OneAffData.FailedMt = v.FailedMt
		OneAffData.PostbackNum = v.PostbackNum
		OneAffData.Amount = float32(v.SuccessMt) * servicePrice[v.ServiceName]

		OneService.ServiceName = v.ServiceName
		OneService.Price = servicePrice[v.ServiceName]

		if v.Operator != operator {
			OneService.AffSubData = append(OneService.AffSubData, OneAffData)

			OneOperator.Operator = v.Operator
			OneOperator.ServiceList = append(OneOperator.ServiceList, OneService)
			SubOperatorData = append(SubOperatorData, OneOperator)
			operator_service = v.Operator + v.ServiceName
			service_index = 0

			operator = v.Operator
			SubOperatorData_index += 1

		} else {
			if v.Operator+v.ServiceName != operator_service {
				OneService.AffSubData = append(OneService.AffSubData, OneAffData)
				SubOperatorData[SubOperatorData_index-1].ServiceList = append(SubOperatorData[SubOperatorData_index-1].ServiceList, OneService)
				operator_service = v.Operator + v.ServiceName
				service_index = 1
			} else {
				SubOperatorData[SubOperatorData_index-1].ServiceList[service_index].AffSubData = append(SubOperatorData[SubOperatorData_index-1].ServiceList[service_index].AffSubData, OneAffData)
			}
		}

		if (operator_tatal_name != operator && SubOperatorData_index != 1) || lengthSubdata == i+1 {
			if lengthSubdata == i+1 {
				if v.ServiceName == "game_w" {
					gameMtNum_W += v.SuccessMt
				} else if v.ServiceName == "video_w" {
					videoMtNum_W += v.SuccessMt
				} else if v.ServiceName == "game_d" {
					gameMtNum_D += v.SuccessMt
				} else if v.ServiceName == "video_d" {
					videoMtNum_D += v.SuccessMt
				}

				totalOperatorData.SubNum += v.SubNum
				totalOperatorData.FailedMt += v.FailedMt
				totalOperatorData.PostbackNum += v.PostbackNum
				totalOperatorData.UnsubNum += v.UnsubNum

				SubOperatorData_index += 1
			}
			OperatorService.ServiceName = "Total Service"
			totalOperatorData.SuccessMt = gameMtNum_W + videoMtNum_W + gameMtNum_D + videoMtNum_D
			totalOperatorData.AffName = "Total Aff"
			totalOperatorData.Amount = float32(gameMtNum_W)*servicePrice["game_w"] + float32(videoMtNum_W)*servicePrice["video_w"] + float32(gameMtNum_D)*servicePrice["game_d"] + float32(videoMtNum_D)*servicePrice["video_d"]
			OperatorService.AffSubData = append(OperatorService.AffSubData, totalOperatorData)
			SubOperatorData[SubOperatorData_index-2].ServiceList = append(SubOperatorData[SubOperatorData_index-2].ServiceList, OperatorService)

			totalOperatorData = AffSub{}
			OperatorService = Service{}
			operator_tatal_name = operator
			gameMtNum_D = 0
			videoMtNum_D = 0
			gameMtNum_W = 0
			videoMtNum_W = 0
		}

		if v.ServiceName == "game_w" {
			gameMtNum_W += v.SuccessMt
			allgame_W += v.SuccessMt
		} else if v.ServiceName == "video_w" {
			videoMtNum_W += v.SuccessMt
			allvideo_W += v.SuccessMt
		} else if v.ServiceName == "game_d" {
			gameMtNum_D += v.SuccessMt
			allgame_D += v.SuccessMt
		} else if v.ServiceName == "video_d" {
			videoMtNum_D += v.SuccessMt
			allvideo_D += v.SuccessMt
		}

		totalOperatorData.SubNum += v.SubNum
		totalOperatorData.FailedMt += v.FailedMt
		totalOperatorData.PostbackNum += v.PostbackNum
		totalOperatorData.UnsubNum += v.UnsubNum

		AllTotalData.SubNum += v.SubNum
		AllTotalData.FailedMt += v.FailedMt
		AllTotalData.PostbackNum += v.PostbackNum
		AllTotalData.UnsubNum += v.UnsubNum
	}

	AllTotalOperatorData.Operator = "Total"
	OperatorService = Service{}
	OperatorService.ServiceName = "Total"
	AllTotalData.SuccessMt = allvideo_D + allgame_D + allvideo_W + allgame_W
	AllTotalData.AffName = "Total"
	AllTotalData.Amount = float32(allgame_W)*servicePrice["game_w"] + float32(allvideo_W)*servicePrice["video_w"] + float32(allgame_D)*servicePrice["game_d"] + float32(allvideo_D)*servicePrice["video_d"]
	OperatorService.AffSubData = append(OperatorService.AffSubData, AllTotalData)
	AllTotalOperatorData.ServiceList = append(AllTotalOperatorData.ServiceList, OperatorService)
	SubOperatorData = append(SubOperatorData, AllTotalOperatorData)
	copy(SubOperatorData[1:], SubOperatorData[0:len(SubOperatorData)-1])
	SubOperatorData[0] = AllTotalOperatorData
	return SubOperatorData
}

type newlist []SubQualityModels

func (I newlist) Len() int {
	return len(I)
}
func (I newlist) Less(i, j int) bool {
	return I[i].Date < I[j].Date
}
func (I newlist) Swap(i, j int) {
	I[i], I[j] = I[j], I[i]
}
