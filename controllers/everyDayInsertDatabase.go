package controllers

import (
	"github.com/MobileCPX/PreMondia/models"
)

// EveryDayInsertSubData 每天存一下盈亏，累计花费等数据
func EveryDayInsertSubData() {
	models.InsertEveryDaySubData()
}

// DailyInsertChartSubData 每天存一下订阅数据
func DailyInsertChartSubData() {
	models.InserSubData()
	//this.Ctx.WriteString("ok")
}
