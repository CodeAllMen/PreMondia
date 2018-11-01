package searchAPI

import (
	"fmt"

	"github.com/MobileCPX/PreMondia/models"

	"github.com/MobileCPX/PreMondia/util"

	"github.com/astaxie/beego/orm"
)

type ClickNumInfo struct {
	Datetime    string
	AffName     string
	PubId       string
	ServiceType string
	ClickType   string
	ClickNum    int
}

// InsertClickData 每小时存一次点击
func InsertClickData() {
	o := orm.NewOrm()
	var (
		clickInfo    []ClickNumInfo
		maxDateClick models.ClickData
	)

	maxSQL := "select * from click_data order by datetime desc limit 1"
	o.Raw(maxSQL).QueryRow(&maxDateClick)
	hoursTime := util.GetFormatHoursTime()
	sql := fmt.Sprintf("select left(sendtime,13) as Datetime,aff_name, service_type,pub_id,count(id) as "+
		"Click_num,click_type from nth_id where left(sendtime,13)>'%s' and left(sendtime,13)<'%s' group by "+
		"aff_name, pub_id, service_type,left(sendtime,13),click_type order by Datetime", maxDateClick.Datetime, hoursTime)

	o.Raw(sql).QueryRows(&clickInfo)
	for _, v := range clickInfo {
		var clickData models.ClickData
		clickData.ClickNum = v.ClickNum
		clickData.AffName = v.AffName
		clickData.Datetime = v.Datetime
		clickData.PubId = v.PubId
		clickData.ClickType = v.ClickType
		clickData.ServiceType = v.ServiceType
		o.Insert(&clickData)
	}
}
