package models

import (
	"github.com/MobileCPX/PreMondia/util"
	"github.com/astaxie/beego/orm"
)

func InsertIbId(gmid *MdId) (error, int64) {
	o := orm.NewOrm()
	o.Using("default")
	_, nowTime := util.GetDatetime()
	gmid.Sendtime = nowTime
	id_int, err := o.Insert(gmid)
	return err, id_int
}
