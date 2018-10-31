package postbackutil

import (
	"fmt"

	"github.com/MobileCPX/PreMondia/models"
	"github.com/astaxie/beego/orm"
)

func PostbackRequest() {
	o := orm.NewOrm()
	var mo []models.Mo
	o.QueryTable("mo").All(&mo)
	fmt.Println(mo)
	for _, oneMo := range mo {
		// trackID := getMdSubscribeTableTrackID(oneMo.CustomerID)
		// fmt.Println(trackID, "11111111111")
		// if trackID == "" {
		// 	return
		// }
		// intTrackID, _ := strconv.Atoi(trackID)
		trackData := getMdSubscribeTableTrackID(oneMo.CustomerID) // 获取用户点击信息
		oneMo.ServiceType = trackData.ServiceType
		oneMo.ClickID = trackData.ClickID
		oneMo.ProID = trackData.ProID
		oneMo.PubID = trackData.PubID
		oneMo.AffName = trackData.AffName
		o.Update(&oneMo)

	}

}

func getMdSubscribeTableTrackID(customerID string) models.AffTrack {
	o := orm.NewOrm()
	var track models.AffTrack
	o.QueryTable("aff_track").Filter("\"customer_id\"", customerID).OrderBy("-id").One(&track)
	return track
}
