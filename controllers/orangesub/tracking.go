package orangesub

import (
	"strconv"

	"github.com/MobileCPX/PreMondia/models"
	"github.com/MobileCPX/PreMondia/models/sub"

	"github.com/astaxie/beego"
)

// LPTrackClickControllers LP页面存储点击
type LPTrackClickControllers struct {
	beego.Controller
}

// Get 请求 LP页面存储点击
func (c *LPTrackClickControllers) Get() {
	affTrack := new(models.AffTrack)
	affTrack.ServiceType = c.GetString("type")
	affTrack.AffName = c.GetString("affName")
	affTrack.ClickID = c.GetString("clickId")
	affTrack.PubID = c.GetString("pubId")
	affTrack.ProID = c.GetString("proId")

	id, err := sub.InsertTrack(affTrack)
	if err == nil {
		c.Ctx.WriteString(strconv.FormatInt(id, 10))
	} else {
		c.Ctx.WriteString("false")
	}
}
