package sub

import (
	"github.com/MobileCPX/PreMondia/models"
	"github.com/MobileCPX/PreMondia/util"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

// InsertTrack 插入点击
func InsertTrack(trackData *models.AffTrack) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	_, trackData.Sendtime = util.GetDatetime()

	idInt, err := o.Insert(trackData)
	if err != nil {
		logs.Error("插入点击错误 ", err.Error())
	}
	return idInt, err
}

// GetAffTrackData 根据track自增id查询此次点击信息
func GetAffTrackData(trackID int) (*models.AffTrack, bool) {
	o := orm.NewOrm()
	o.Using("default")
	trackData := new(models.AffTrack)
	var isHas bool
	o.QueryTable("aff_track").Filter("track_id", trackID).One(trackData)
	if trackData.TrackID == 0 {
		logs.Error("未查询到次trackID信息")
		isHas = false
		return trackData, isHas
	}
	return trackData, true
}

// UpdateTrackData 更新aff_track 数据
func UpdateTrackData(trackData *models.AffTrack) error {
	o := orm.NewOrm()
	o.Using("default")
	_, err := o.Update(trackData)
	if err != nil {
		logs.Error("更新数据track数据失败")
	}
	return err
}
