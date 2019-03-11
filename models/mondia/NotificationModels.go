package mondia

import (
	"github.com/MobileCPX/PreMondia/util"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

// Notification 订阅通知信息
type Notification struct {
	// Notification 必须要有的参数
	ID               int64  `orm:"pk;auto;column(id)"`                           //自增ID
	NotificationType string `orm:"column(notification_type);size(15)"`           // 通知类型
	Sendtime         string `orm:"column(sendtime);size(30)"`                    // 点击时间
	SubscriptionID   string `orm:"column(subscription_id)" xml:"SubscriptionID"` // 订阅id
	TransactionID    string `orm:"column(transaction_id)" xml:"TransactionId"`
	TransactionTime  string `xml:"TransactionTimestamp"`
	CustomerID       string `orm:"column(customer_id)" xml:"CustomerID"`
	ServiceID        string `orm:"column(service_id)" xml:"ServiceID"`

	ProductCode        string `xml:"ProductCode"`
	PackageCode        string `xml:"PackageCode"`
	Operator           string `xml:"OperatorName"`
	SubscriptionStatus string `xml:"SubscriptionStatus"`
	Action             string `xml:"Action"`
	Price              string `xml:"Price"`
	Channel            string `xml:"Channel"`
	NextAction         string `xml:"NextAction"`
	NextActionDate     string `xml:"NextActionDate"`
}

//// GoNotification 订阅，续订、退订通知
//type Notification struct {
//	Id               int64  `orm:"pk;auto" `
//	ActionType       string `form:"type"`           // 通知类型
//	Mcc              string `form:"mcc"`            // 运营商 Mcc
//	Mnc              string `form:"mnc"`            // 运营商 Mnc
//	SubscriptionId   string `form:"subscriptionId"` // 订阅id
//	Msisdn           string `form:"msisdn"`         // 用户电话号码
//	TransactionId    string `form:"transactionId"`  //唯一交易ID
//	MovementDate     string `form:"movementDate"`
//	Value            string `form:"value"`    // 单价
//	Currency         string `form:"currency"` // 货币类型
//	ActivationDate   string `form:"activationDate"`
//	DeactivationDate string `form:"deactivationDate"`
//	StatusDate       string `form:"statusDate"`
//	ServiceName      string `form:"service"`
//	TrackID          string `from:"dcbExternalId" orm:"column(track_id)"`
//	NotificationType string
//	Sendtime         string
//}

func (notification *Notification) Insert() error {
	o := orm.NewOrm()
	nowTime, _ := util.GetNowTimeFormat()
	notification.Sendtime = nowTime
	_, err := o.Insert(notification)
	if err != nil {
		logs.Error("Notification Insert 数据失败，ERROR: ", err.Error())
	}
	return err
}
