package mondia

import (
	"github.com/MobileCPX/PreMondia/util"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

// Mo mo表数据
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
	PostbackCode   string `orm:"column(postback_code);size(3000)"`  // 回传是否成功
	Payout         float32
	SuccessMT      int    `orm:"column(success_mt)"`            // 扣费成功次数
	FailedMT       int    `orm:"column(failed_mt)"`             // 失败扣费次数
	SubTime        string `orm:"column(sub_time);size(20)"`     // 订阅 时间
	UnsubTime      string `orm:"column(unsub_time);size(20)"`   // 退订 时间
	ServiceType    string `orm:"column(service_type);size(30)"` // 服务类型
	AffName        string `orm:"column(aff_name);size(30)"`     // 网盟名称
	PubID          string `orm:"column(pub_id);size(100)"`      // 子渠道
	ProID          string `orm:"column(pro_id);size(30)"`       // 服务id（可有可无）
	ClickID        string `orm:"column(click_id);size(300)"`    // 点击
	IP             string `orm:"column(ip);size(20)"`           // 用户IP地址
	UserAgent      string `orm:"column(user_agent)"`            // 用户user_agent
	Refer          string `orm:"column(refer)"`                 // 网页来源
	AffTrackID     string `orm:"column(aff_track_id)"`          // AffTrack 表自增ID
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
	WeekMtNum          int                                         // 每周已经扣费次数  扣费失败每周最多只能扣费两次
	ModifyDate         string `orm:"column(modify_date);size(20)"` // 最后一次扣费成功时间
	FinalMtType        string                                      //最后一次扣费的类型
	CanvasID           string `orm:"column(canvas_id)"`            // 帆布ID
}

func (mo *Mo) TableName() string {
	return "mo"
}

func (mo *Mo) MoQuery() orm.QuerySeter {
	o := orm.NewOrm()
	return o.QueryTable(mo.TableName())
}

// 插入新订阅数据
func (mo *Mo) InitNewSubMO(response *Notification, affTrack *AffTrack) *Mo {
	// AffTrack init
	mo.AffName = affTrack.AffName
	mo.ClickID = affTrack.ClickID
	mo.ProID = affTrack.ProID
	mo.PubID = affTrack.PubID
	mo.IP = affTrack.IP
	mo.UserAgent = affTrack.UserAgent

	// Notification init
	mo.ServiceID = response.ProductCode
	mo.CustomerID = response.CustomerID
	mo.SubscriptionID = response.SubscriptionID
	mo.Operator = response.Operator
	mo.ProductCode = response.ProductCode
	return mo
}

func (mo *Mo) GetMoByCustomerID(customerID string) error {
	o := orm.NewOrm()
	err := o.QueryTable(MoTBName()).Filter("customer_id", customerID).One(mo)
	if err != nil{
		//logs.Error("GetMoByCustomerID ERROR: ",)
	}
	return err
}
func (mo *Mo) UnsubGetMoByCustomerID(customerID, serviceID string) {
	o := orm.NewOrm()
	SQL := o.QueryTable(MoTBName()).Filter("customer_id", customerID)
	if serviceID != "" {
		_ = SQL.Filter("product_code", serviceID).One(mo)
		if mo.ID == 0 {
			subNum, _ := SQL.Count()
			if subNum > 1 {
				_ = SQL.Filter("sub_status", 1).OrderBy("-id").One(mo)
				if mo.ID == 0 {
					_ = SQL.OrderBy("-id").One(mo)
				}
			} else if subNum == 1 {
				_ = SQL.One(mo)
			}
		}
	} else {
		subNum, _ := SQL.Count()
		if subNum > 1 {
			_ = SQL.Filter("sub_status", 1).OrderBy("-id").One(mo)
			if mo.ID == 0 {
				_ = SQL.OrderBy("-id").One(mo)
			}
		} else if subNum == 1 {
			_ = SQL.One(mo)
		}
	}

	//err := o.QueryTable(MoTBName()).Filter("customer_id", customerID).One(mo)
	return
}

// CheckSubIDIsExist 通过SubId 查询用户是否已经订阅过
func (mo *Mo) CheckSubIDIsExist(SubID string) bool {
	o := orm.NewOrm()
	isExist, err := o.QueryTable(MoTBName()).Filter("subscription_id", SubID).Count()
	if err != nil {
		logs.Error("CheckSubIDIsExist 查询数据失败，ERROR: ", err.Error())
	}
	if isExist != 0 {
		logs.Info("CheckSubIDIsExist 次订阅用户已经存在，subscription_id: ", SubID)
		return true
	}
	return false
}

// InsertNewMo 插入新订阅数据
func (mo *Mo) InsertNewMo() error {
	o := orm.NewOrm()
	nowTime, _ := util.GetNowTimeFormat()
	mo.SubTime = nowTime
	_, err := o.Insert(mo)
	if err != nil {
		logs.Error("新插入订阅数据失败 ERROR: ", err.Error())
	}
	return err
}

func (mo *Mo) UpdateMO() error {
	o := orm.NewOrm()
	_, err := o.Update(mo)
	if err != nil {
		logs.Error("更新订阅数据失败 ERROR: ", err.Error())
	}
	return err
}

// 通过电话号码和ServiceName查询Mo信息
func (mo *Mo) GetMoByMsisdnAndServiceName(msisdn, serviceName string) *Mo {
	o := orm.NewOrm()
	_ = o.QueryTable(MoTBName()).Filter("msisdn", msisdn).Filter("service_name", serviceName).
		OrderBy("-id").One(mo)
	return mo
}

// 成功扣费更新MO表
func (mo *Mo) SuccessMTUpdateMO(subscriptionID string) (notificationType string, err error) {
	o := orm.NewOrm()
	_, nowDate := util.GetNowTimeFormat()
	err = o.QueryTable(MoTBName()).Filter("subscription_id", subscriptionID).One(mo)
	if err != nil {
		logs.Error("SuccessMTUpdateMO 收到扣费通知后更新MO表失败，ERROR: ", err.Error())
		return
	}
	if mo.ID != 0 && mo.ModifyDate != nowDate {
		mo.ModifyDate = nowDate
		mo.SuccessMT++
		_ = mo.UpdateMO()
		notificationType = "MT_SUCCESS"
	}
	return
}

// 退订更新MO表
func (mo *Mo) UnsubUpdateMo(subscriptionID string) (notificationType string, err error) {
	o := orm.NewOrm()
	nowTime, nowDate := util.GetNowTimeFormat()
	err = o.QueryTable(MoTBName()).Filter("subscription_id", subscriptionID).One(mo)
	if err != nil {
		logs.Error("SuccessMTUpdateMO 收到扣费通知后更新MO表失败，ERROR: ", err.Error())
		return
	}
	if mo.ID != 0 {
		mo.ModifyDate = nowDate
		mo.UnsubTime = nowTime
		mo.SubStatus = 0
		_ = mo.UpdateMO()
		notificationType = "UNSUB"
	}
	return

}

//FailedMTUpdateMo 扣费失败更新MO表
func (mo *Mo) FailedMTUpdateMo(subscriptionID string) (notificationType string, err error) {
	o := orm.NewOrm()
	_, nowDate := util.GetNowTimeFormat()
	err = o.QueryTable(MoTBName()).Filter("subscription_id", subscriptionID).One(mo)
	if err != nil {
		logs.Error("SuccessMTUpdateMO 收到扣费通知后更新MO表失败，ERROR: ", err.Error())
		return
	}
	if mo.ID != 0 && mo.ModifyDate != nowDate {
		mo.ModifyDate = nowDate
		mo.FailedMT++
		_ = mo.UpdateMO()
		notificationType = "MT_FAILED"
	}
	return

}

// GetMoBySubscriptionID 根据SubID 查询订阅信息
func GetMoBySubscriptionID(subscriptionID string) (*Mo, error) {
	mo := &Mo{SubscriptionID: subscriptionID}
	o := orm.NewOrm()
	err := o.Read(mo)
	if err != nil {
		logs.Error("根据subscription_id 查询订阅信息失败 Subscript ID ", subscriptionID, err.Error())
	}
	return mo, err
}

// IsSubByCanvasID 通过CanvasID检查用户是否订阅
func (mo *Mo) IsSubByCanvasID() bool {
	o := orm.NewOrm()
	err := o.Read(mo)
	if err != nil {
		logs.Error("通过CanvasID 查询mo信息失败，ERROR: ", err.Error())
	}
	if mo.ID != 0 {
		return true
	} else {
		return false
	}
}

// CheckTodaySubNumMoreLimit 检查今日订阅数量是否超过订阅限制
func (mo *Mo) CheckTodaySubNumMoreLimit(serviceID string) (isCanSub bool) {
	o := orm.NewOrm()
	limitSubNum, err := beego.AppConfig.Int("limitSubNum")
	_, nowDate := util.GetFormatTime()
	subNum, err := o.QueryTable(MoTBName()).Filter("subtime__gt", nowDate).Filter("service_id", serviceID).Count()
	if err != nil {
		logs.Error("CheckTodaySubNumMoreLimit 获取今日的订阅数量失败 ERROR: ", err.Error())
	}
	// 检查今日订阅数是否超过了订阅限制
	if int(subNum) >= limitSubNum {
		logs.Error("CheckTodaySubNumMoreLimit 超过了订阅限制 limitSubNum: ", limitSubNum, " today sub num: ", subNum)
	} else {
		isCanSub = true
	}
	return
}

func (mo *Mo) LimitTenMinutesSubNum(serviceID string, limitSubNum int) (isLimit bool) {
	o := orm.NewOrm()
	nowTime, _ := util.GetNowTimeFormat()
	nowMinutes := nowTime[0:15]
	subNum, _ := o.QueryTable("mo").Filter("sub_time__gt", nowMinutes).
		Filter("service_id", serviceID).Count()
	if int(subNum) > limitSubNum {
		logs.Info(nowMinutes, "   十分钟的订阅已经满3个，十分钟之内最多3个订阅")
		return true
	}
	return false
}

func (mo *Mo) GetAffNameTodaySubInfo() (subNum, postbackNum int64) {
	o := orm.NewOrm()
	_, nowDate := util.GetFormatTime()
	subNum, _ = o.QueryTable(MoTBName()).Filter("aff_name", mo.AffName).Filter("subtime__gt", nowDate).Count()
	postbackNum, _ = o.QueryTable(MoTBName()).Filter("aff_name", mo.AffName).Filter("postback_status", 1).
		Filter("subtime__gt", nowDate).Count()
	logs.Info(mo.AffName, nowDate, "sub_num: ", subNum, " postback_num: ", postbackNum)
	return
}

// 获取今日的订阅数量
func GetTodayMoNum(serviceID string) (int64, error) {
	o := orm.NewOrm()
	_, nowDate := util.GetFormatTime()
	subNum, err := o.QueryTable(MoTBName()).Filter("subtime__gt", nowDate).Filter("service_id", serviceID).Count()
	if err != nil {
		logs.Error("GetTodaySubNum serviceID ", serviceID, "获取今日的订阅数量失败 ERROR: ", err.Error())
	}
	logs.Info("GetTodaySubNum  serviceID ", serviceID, " 今日的订阅数量: ", subNum)
	return subNum, err
}
