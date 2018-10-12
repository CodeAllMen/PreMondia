package models

import (
	"github.com/astaxie/beego/orm"
)

// Notification 订阅通知信息
type Notification struct {
	Id                 int64  `orm:"pk;auto"`
	TransactionId      string `xml:"TransactionId"`
	TransactionTime    string `xml:"TransactionTimestamp"`
	CustomerId         string `xml:"CustomerID"`
	SubscriptionId     string `xml:"SubscriptionID"`
	ServiceId          string `xml:"ServiceID"`
	ProductCode        string `xml:"ProductCode"`
	PackageCode        string `xml:"PackageCode"`
	Operator           string `xml:"OperatorName"`
	SubscriptionStatus string `xml:"SubscriptionStatus"`
	Action             string `xml:"Action"`
	Price              string `xml:"Price"`
	Channel            string `xml:"Channel"`
	NextAction         string `xml:"NextAction"`
	NextActionDate     string `xml:"NextActionDate"`
	Sendtime           string
	ServiceType        string
	AffName            string
	PubId              string
	ProId              string
	ClickId            string
	ClickType          string
}

// UnsubXml 退订通知信息
type UnsubXml struct {
	Id                 int64  `orm:"pk;auto"`
	TransactionId      string `xml:"TransactionId"`
	TransactionTime    string `xml:"TransactionTimestamp"`
	CustomerID         string `xml:"CustomerID"`
	SubscriptionId     string `xml:"SubscriptionID"`
	ServiceID          string `xml:"ServiceID"`
	ProductCode        string `xml:"ProductCode"`
	PackageCode        string `xml:"PackageCode"`
	Operator           string `xml:"OperatorName"`
	SubscriptionStatus string `xml:"SubscriptionStatus"`
	Action             string `xml:"Action"`
	Price              string `xml:"Price"`
	Channel            string `xml:"Channel"`
	NextAction         string `xml:"NextAction"`
	NextActionDate     string `xml:"NextActionDate"`
	Sendtime           string
}

// MonMo mondia MO信息表
type MonMo struct {
	Id             int64 `orm:"pk;auto"`
	CustomerId     string
	SubscriptionId string
	ServiceId      string
	ProductCode    string
	PackageCode    string
	Operator       string
	Msisdn         string
	Price          string
	Channel        string
	Subtime        string
	Unsubtime      string
	SubNum         int
	SubStatus      int
	FailNum        int
	ServiceType    string
	AffName        string
	PubId          string
	ProId          string
	ClickId        string
	ClickType      string
	PostbackStatus int
	PostbackCode   string `orm:"size(255)"`
	PostbackTime   string `orm:"size(255)"`
}

// MdId 每次点击信息表
type MdId struct {
	Id          int64 `orm:"pk;auto"`
	ServiceType string
	AffName     string
	PubId       string
	ProId       string
	ClickId     string
	ClickType   string
	Sendtime    string
}

type MdCustomer struct {
	Id          int64 `orm:"pk;auto"`
	NewId       string
	Status      string
	Operator    string
	CustomerId  string
	ErrorDesc   string
	ErrorCode   string
	Sendtime    string
	ServiceType string
	AffName     string
	PubId       string
	ProId       string
	ClickId     string
	ClickType   string
}

type MdSubscribe struct {
	Id             int64 `orm:"pk;auto"`
	Sendtime       string
	ClientId       string
	Status         string
	CustomerId     string
	SubId          string
	NextAction     string
	SubStatus      string
	NextActionDate string
	ErrorCode      string
	ErrorDesc      string
	ServiceType    string
	ViewName       string
	AffName        string
	PubId          string
	ClickId        string
}

type PostbackUrl struct {
	Id           int64  `orm:"pk;auto"`
	Aff_name     string `orm:"size(255)"`
	Post_back    string `orm:"size(255)"`
	PostbackRate int    `orm:"default(70)"`
}

type EveryDaySubDatas struct {
	Id                   int64 `orm:"pk;auto"`
	Date                 string
	SubData_1click       string `orm:"size(500)"`
	SubData_2click       string `orm:"size(500)"`
	UnsubData_1click     string `orm:"size(500)"`
	UnsubData_2click     string `orm:"size(500)"`
	PostbackData_1click  string `orm:"size(500)"`
	PostbackData_2click  string `orm:"size(500)"`
	PostbackSpend_1click string `orm:"size(500)"`
	PostbackSpend_2click string `orm:"size(500)"`
	MtData_1click        string `orm:"size(500)"`
	MtData_2click        string `orm:"size(500)"`

	Active    int
	SubNum    int
	FailedMt  int
	SuccessMt int
	MtRate    string
	Expend    string
	//O2  float32
	//Ee   float32
	//Orange   float32
	//Three  float32
	//Vodafone  float32
	//Tmobile  float32
	//Virgin  float32

	O2       int
	Ee       int
	Orange   int
	Three    int
	Vodafone int
	Tmobile  int
	Virgin   int

	GrandTotalSuccessMt     int
	DaySpend                int //每日花费
	GrandTotalSub           int //累计订阅
	GrandTotalFailedMtNum   int //累计扣费失败次数
	GrandTotalProfitAndLoss float32
	GrandTotalRevenue       float32
	GrandTotalMtCharges     float32 // 累计扣费
	GrandTotalMtRate        string
	GrandTotalSpend         int
	ServiceType             string
	ClickType               string
}

type ClickData struct {
	Id          int64  `orm:"pk;auto"`
	Date        string `orm:"size(255)"`
	Datetime    string `orm:"size(255)"`
	Aff_name    string `orm:"size(255)"`
	Pub_id      string `orm:"size(255)"`
	ServiceType string `orm:"size(255)"`
	ClickType   string `orm:"size(255)"`
	ClickNum    int    `orm:"size(255)"`
}

type RegisterJson struct {
	Name string `json:"name"`
	Sign string `json:"sign"`
}

// UnsubPin 用户获取的PIN
type UnsubPin struct {
	Id        int64 `orm:"pk;auto"`
	Msisdn    string
	Pin       string
	PinStatus string
}

type UnsubGetCustomer struct {
	Id            int64  `orm:"pk;auto"`
	TransactionId string `xml:"TransactionId"`
	ResponseCode  string `xml:"ResponseCode"`
	Description   string `xml:"Description"`
	CustomerId    string `xml:"CustomerId"`
}

// MondiaCharge Mondia xml格式通知
type MondiaCharge struct {
	Id                 int64  `orm:"pk;auto"`
	TransactionId      string `xml:"TransactionId"`
	ResponseCode       string `xml:"ResponseCode"`
	Description        string `xml:"Description"`
	SubscriptionStatus string `xml:"SubscriptionStatus"`
	SubscriptionID     string `xml:"SubscriptionID"`
	CustomerID         string `xml:"CustomerID"`
	NextAction         string `xml:"NextAction"`
	NextActionDate     string `xml:"NextActionDate"`
	BillingStatus      string `xml:"BillingStatus"`
}

func init() {
	orm.RegisterModel(new(PostbackUrl), new(Notification), new(MonMo), new(MdId), new(MdSubscribe),
		new(MdCustomer), new(ClickData), new(EveryDaySubDatas), new(UnsubPin), new(UnsubGetCustomer), new(MondiaCharge))
}
