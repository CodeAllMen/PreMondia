package mondia

import (
	"github.com/MobileCPX/PreMondia/util"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

// AffTrack 网盟点击追踪
type AffTrack struct {
	TrackID     int64  `orm:"pk;auto;column(track_id)" json:"track_id"`    //自增ID
	Sendtime    string `orm:"column(sendtime);size(30)" json:"sendtime"`   // 点击时间
	AffName     string `orm:"column(aff_name);size(30)" json:"aff_name"`   // 网盟名称
	PubID       string `orm:"column(pub_id);size(100)" json:"pub_id"`    // 子渠道
	ProID       string `orm:"column(pro_id);size(30)" json:"pro_id"`     // 服务id（可有可无）
	ClickID     string `orm:"column(click_id);size(100)" json:"click_id"`  // 点击
	ServiceID   string `orm:"column(service_id);size(30)" json:"service_id"` // 服务类型
	ServiceName string `orm:"column(service_name)" json:"service_name"`
	IP          string `orm:"column(ip);size(255)" json:"ip"` // 用户IP地址
	UserAgent   string `orm:"column(user_agent)" json:"user_agent"`  // 用户user_agent
	Refer       string `orm:"column(refer)" json:"refer"`       // 网页来源

	ErrorCode  string
	ErrorDesc  string
	Operator   string
	CustomerID string `orm:"column(customer_id)"`
	Status     string
}

func (track *AffTrack) TableName() string {
	return "aff_track"
}

func (track *AffTrack) TrackQuery() orm.QuerySeter {
	o := orm.NewOrm()
	return o.QueryTable(track.TableName())
}

func (track *AffTrack) Insert() (int64, error) {
	o := orm.NewOrm()
	track.Sendtime, _ = util.GetNowTimeFormat()
	trackID, err := o.Insert(track)
	if err != nil {
		logs.Error("新插入点击错误 ", err.Error())
	}
	return trackID, err
}

func (track *AffTrack) Update() error {
	o := orm.NewOrm()
	_, err := o.Update(track)
	if err != nil {
		logs.Error("AffTrack Update 更新点击数据失败，ERROR ", err.Error())
	}
	return err
}

func (track *AffTrack) GetAffTrackByTrackID(trackID int64) error {
	o := orm.NewOrm()
	err := o.QueryTable(track.TableName()).Filter("track_id", trackID).One(track)
	if err != nil {
		logs.Error("通过trackID 查询点击信息失败，未找到此trackID： ", trackID)
	}
	return err
}

// 通过 customerID 查询点击信息
func (track *AffTrack) GetAffTrackByCustomerID(customerID string) error {
	o := orm.NewOrm()
	err := o.QueryTable(AffTrackTBName()).Filter("customer_id", customerID).OrderBy("-track_id").One(track)

	if err != nil {
		logs.Error("通过customerID查询点击信息失败，未找到此customerID： ", customerID, "ERROR: ", err.Error())
	}
	return err
}
