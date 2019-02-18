package mondia

import (
	"github.com/MobileCPX/PreMondia/util"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"strconv"
)

// UnsubPin 用户获取的PIN
type UnsubPin struct {
	ID        int64 `orm:"pk;auto;column(id)"` //自增ID
	Msisdn    string
	Pin       string
	Sendtime  string
	PinStatus string
}
type UnsubGetCustomer struct {
	Id            int64  `orm:"pk;auto"`
	TransactionId string `xml:"TransactionId"`
	ResponseCode  string `xml:"ResponseCode"`
	Description   string `xml:"Description"`
	CustomerId    string `xml:"CustomerId"`
}

func (unsubPin *UnsubPin) Insert() (int64, error) {
	o := orm.NewOrm()
	nowTime, _ := util.GetNowTimeFormat()
	unsubPin.Sendtime = nowTime
	id, err := o.Insert(unsubPin)
	if err != nil {
		logs.Error("UnsubPin Insert 插入PIN数据错误，ERROR: ", err.Error())
	}
	return id, err
}

func (unsubPin *UnsubPin) CheckPIN(id string) error {
	o := orm.NewOrm()
	idInt, err := strconv.Atoi(id)
	if err != nil {
		logs.Error("CheckPIN id string to int ERROR: ", err.Error())
		return err
	}
	unsubPin.ID = int64(idInt)
	err = o.Read(unsubPin)
	return err

}
