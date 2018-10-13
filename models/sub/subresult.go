package sub

import (
	"github.com/MobileCPX/PreMondia/models"
	"github.com/MobileCPX/PreMondia/util"
	"github.com/astaxie/beego/orm"
)

// InsertSubscribe 插入订阅通知
func InsertSubscribe(sub *models.MdSubscribe) {
	o := orm.NewOrm()
	sub.Sendtime, _ = util.GetDatetime()
	o.Insert(sub)
}
