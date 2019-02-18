package sub

import (
	"fmt"

	"github.com/MobileCPX/PreMondia/models"
	"github.com/MobileCPX/PreMondia/util"
	"github.com/astaxie/beego/orm"
)

// InsertSubscribe 插入订阅通知
func InsertSubscribe(sub *models.MdSubscribe) {
	o := orm.NewOrm()
	fmt.Println("1234566425346")
	sub.Sendtime, _ = util.GetDatetime()
	o.Insert(sub)
}

func InsertSubResult(result models.SubResult) {
	o := orm.NewOrm()
	nowTime, _ := util.GetDatetime()
	result.Sendtime = nowTime
	o.Insert(&result)
}
