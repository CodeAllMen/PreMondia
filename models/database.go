package models

import "github.com/astaxie/beego/orm"

// AffTrack 用户每次点击的完整信息
type AffTrack struct {
	// AffTrack必须要有的参数
	TrackID     int64  `orm:"pk;auto;column(track_id)"`      //自增ID
	Sendtime    string `orm:"column(sendtime);size(30)"`     // 点击时间
	AffName     string `orm:"column(aff_name);size(30)"`     // 网盟名称
	PubID       string `orm:"column(pub_id);size(100)"`      // 子渠道
	ProID       string `orm:"column(pro_id);size(30)"`       // 服务id（可有可无）
	ClickID     string `orm:"column(click_id);size(100)"`    // 点击
	ServiceType string `orm:"column(service_type);size(30)"` // 服务类型
	IP          string `orm:"column(ip);size(20)"`           // 用户IP地址
	UserAgent   string `orm:"column(user_agent)"`            // 用户user_agent
	Refer       string `orm:"column(refer)"`                 // 网页来源
	CanvasID    string `orm:"column(canvas_id)"`             // 帆布ID

	// 根据 需求添加的参数
	// SubscriptionID string `orm:"column(subscription_id)"`     // 订阅id
	Operator   string `orm:"column(operator)"`   // 运营商
	ErrorCode  string `orm:"column(error_code)"` // 订阅失败代码
	ErrorDesc  string `orm:"column(error_desc)"`
	Status     string
	CustomerID string `orm:"column(customer_id)"`
	Channel    string `orm:"size(20)"` // 订阅渠道

}

// Notification 订阅通知信息
type Notification struct {
	// Notification 必须要有的参数
	ID               int64  `orm:"pk;auto;column(id)"`                 //自增ID
	NotificationType string `orm:"column(notification_type);size(15)"` // 通知类型
	Sendtime         string `orm:"column(sendtime);size(30)"`          // 点击时间

	SubscriptionID  string `orm:"column(subscription_id)" xml:"SubscriptionID"` // 订阅id
	TransactionID   string `orm:"column(transaction_id)" xml:"TransactionId"`
	TransactionTime string `xml:"TransactionTimestamp"`
	CustomerID      string `orm:"column(customer_id)" xml:"CustomerID"`

	ServiceID          string `orm:"column(service_id)" xml:"ServiceID"`
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

// Mo 订阅信息mo 表
type Mo struct {
	// MO 必须要有的参数
	ID             int64  `orm:"pk;auto;column(id)"`              //自增ID
	SubscriptionID string `orm:"column(subscription_id)"`         // 订阅id
	Operator       string `orm:"column(operator);size(15)"`       // 运营商
	Msisdn         string `orm:"column(msisdn);size(20)"`         // 电话号码
	Price          string `orm:"column(price);size(10)"`          // 扣费单价
	SubStatus      int    `orm:"column(sub_status);size(1)"`      // 订阅状态 1（订阅）0 （退订）
	PostbackStatus int    `orm:"column(postback_status);size(1)"` // postback 状态
	PostbackTime   string `orm:"column(postback_time);size(20)"`  // postback 时间
	PostbackCode   string `orm:"column(postback_code);size(20)"`  // 回传是否成功
	SuccessMT      int    `orm:"column(success_mt)"`              // 扣费成功次数
	FailedMT       int    `orm:"column(failed_mt)"`               // 失败扣费次数
	SubTime        string `orm:"column(sub_time);size(20)"`       // 订阅 时间
	UnsubTime      string `orm:"column(unsub_time);size(20)"`     // 退订 时间
	ServiceType    string `orm:"column(service_type);size(30)"`   // 服务类型
	AffName        string `orm:"column(aff_name);size(30)"`       // 网盟名称
	PubID          string `orm:"column(pub_id);size(100)"`        // 子渠道
	ProID          string `orm:"column(pro_id);size(30)"`         // 服务id（可有可无）
	ClickID        string `orm:"column(click_id);size(100)"`      // 点击
	IP             string `orm:"column(ip);size(20)"`             // 用户IP地址
	UserAgent      string `orm:"column(user_agent)"`              // 用户user_agent
	Refer          string `orm:"column(refer)"`                   // 网页来源
	AffTrackID     string `orm:"column(aff_track_id)"`            // AffTrack 表自增ID
	// ModifyDate     string `orm:"column(modify_date);size(20)"`    // 最后一次扣费成功时间

	// MO 根据需求添加参数有的参数
	ServiceID        string `orm:"column(service_id);size(50)"`
	MsisdnStatusCode string `orm:"size(5)"`
	ReferenceID      string
	CustomerID       string `orm:"column(customer_id)"`
	Channel          string `orm:"column(channel)"`
	PackageCode      string `orm:"column(package_code)"`
	ProductCode      string `orm:"column(product_code)"`

	MTSuccessDatetimes string
	MTFailedDatetimes  string
	WeekMtNum          int    // 每周已经扣费次数  扣费失败每周最多只能扣费两次
	ModifyDate         string `orm:"column(modify_date);size(20)"` // 最后一次扣费成功时间
	FinalMtType        string //最后一次扣费的类型
	CanvasID           string `orm:"column(canvas_id)"` // 帆布ID
}

//
type MondiaCharge struct {
	ID                 int64  `orm:"pk;auto;column(id)"` //自增ID
	TransactionID      string `orm:"column(transaction_id)" xml:"TransactionId"`
	ResponseCode       string `xml:"ResponseCode"`
	Description        string `xml:"Description"`
	SubscriptionStatus string `xml:"SubscriptionStatus"`
	SubscriptionID     string `xml:"SubscriptionID"`
	CustomerID         string `xml:"CustomerID"`
	NextAction         string `xml:"NextAction"`
	NextActionDate     string `xml:"NextActionDate"`
	BillingStatus      string `xml:"BillingStatus"`
}

// AlreadySub // 网盟重复推送已经订阅过的用户统计
type AlreadySub struct {
	ID         int64  `orm:"pk;auto;column(id)"` //自增ID
	CustomerID string `orm:"column(customer_id)"`
	AffName    string `orm:"column(aff_name);size(30)"`  // 网盟名称
	PubID      string `orm:"column(pub_id);size(100)"`   // 子渠道
	ProID      string `orm:"column(pro_id);size(30)"`    // 服务id（可有可无）
	ClickID    string `orm:"column(click_id);size(100)"` // 点击
	CanvasID   string `orm:"column(canvas_id)"`          // 帆布ID
	TrackID    int64  `orm:"column(track_id)"`           //自增ID
	Sendtime   string `orm:"column(sendtime);size(30)"`  // 扣费时间
}

type MdSubscribe struct {
	ID             int64 `orm:"pk;auto;column(id)"` //自增ID
	Sendtime       string
	TrackID        string `orm:"column(track_id)"`
	Status         string
	CustomerID     string `orm:"column(customer_id)"`
	SubscriptionID string `orm:"column(subscription_id)"`
	NextAction     string
	SubStatus      string
	NextActionDate string
	ErrorCode      string
	ErrorDesc      string
	ServiceType    string
	ViewName       string
}

//Postback 网盟信息
type Postback struct {
	ID           int64  `orm:"pk;auto;column(id)"`                //自增ID
	AffName      string `orm:"column(aff_name);size(30)"`         // 网盟名称
	PostbackURL  string `orm:"column(postback_url);size(180)"`    // postback URL
	PostbackRate int    `orm:"column(postback_rate);default(70)"` // 回传概率
}

// UnsubPin 用户获取的PIN
type UnsubPin struct {
	ID        int64 `orm:"pk;auto;column(id)"` //自增ID
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

//BillingHistory 订阅成功后发起http扣费请求状态数据
type BillingHistory struct {
	ID             int64  `orm:"pk;auto;column(id)"`        //自增ID
	SubscriptionID string `orm:"column(subscription_id)"`   // 订阅id
	AffTrackID     string `orm:"column(aff_track_id)"`      // AffTrack 表自增ID
	BillingStatus  int    `orm:"column(billing_status)"`    // 0表示扣费失败，1表示发起扣费请求成功
	Error          string `orm:"column(error);size(10)"`    // 订阅失败代码
	Sendtime       string `orm:"column(sendtime);size(30)"` // 扣费时间
}

type SubData struct {
	ID          int64 `orm:"pk;auto;column(id)"`
	Date        string
	AffName     string
	PubId       string
	Operator    string
	ServiceType string
	ClickType   string
	SubNum      int
	UnsubNum    int
	PostbackNum int
	SuccessMt   int
	FailedMt    int
}

type EveryDaySubDatas struct {
	ID            int64 `orm:"pk;auto;column(id)"`
	Date          string
	SubData       string `orm:"size(500)"`
	UnsubData     string `orm:"size(500)"`
	PostbackData  string `orm:"size(500)"`
	PostbackSpend string `orm:"size(500)"`
	MtData        string `orm:"size(500)"`

	Active    int
	SubNum    int
	FailedMt  int
	SuccessMt int
	MtRate    string
	Expend    string

	Telfort  int
	KPN      int
	Vodafone int
	Tmobile  int
	Tele2    int

	DayRevenue              float32
	GrandTotalSuccessMt     int
	DaySpend                float32 //每日花费
	GrandTotalSub           int     //累计订阅
	GrandTotalFailedMtNum   int     //累计扣费失败次数
	GrandTotalProfitAndLoss float32
	GrandTotalRevenue       float32
	GrandTotalMtCharges     float32 // 累计扣费
	GrandTotalMtRate        string
	GrandTotalSpend         float32
	ServiceType             string
	ClickType               string
}

type ComplaintData struct {
	SubID          string `orm:"pk;column(sub_id)"`
	Msisdn         string
	Date           string
	Operator       string
	AffName        string
	ServiceType    string
	PubId          string
	ClickId        string
	MtNum          int
	Amount         float32
	PostbackStatus string
	Subtime        string
	Unsubtime      string
	RackingCode    string
	Email          string
	UserName       string
	DealWithTime   string
	EquipmentModel string
	GuiltyAffName  string
	GuiltyPubid    string
	ClickType      string
	Level          string
	Description    string
}

type ClickData struct {
	Id          int64  `orm:"pk;auto"`
	Date        string `orm:"size(255)"`
	Datetime    string `orm:"size(255)"`
	AffName     string `orm:"size(255)"`
	PubId       string `orm:"size(255)"`
	ServiceType string `orm:"size(255)"`
	ClickType   string `orm:"size(255)"`
	ClickNum    int    `orm:"size(255)"`
}

func init() {
	orm.RegisterModel(new(AffTrack), new(Notification), new(Postback), new(Mo),
		new(BillingHistory), new(ClickData), new(EveryDaySubDatas), new(MdSubscribe), new(UnsubPin), new(MondiaCharge), new(AlreadySub))
}
