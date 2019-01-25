package orangesub

import (
	"strconv"

	"github.com/MobileCPX/PreMondia/models"
	"github.com/MobileCPX/PreMondia/models/sub"
	"github.com/MobileCPX/PreMondia/util"

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

	affTrack.Refer = c.Ctx.Input.Refer()           // 获取refer信息
	affTrack.IP = util.GetIPAddress(c.Ctx.Request) // 用户的ip地址
	affTrack.UserAgent = c.Ctx.Input.UserAgent()   //用户设备信息

	id, err := sub.InsertTrack(affTrack)
	c.Ctx.WriteString("false")
	return
	if err == nil {
		c.Ctx.WriteString(strconv.FormatInt(id, 10))
	} else {
		c.Ctx.WriteString("false")
	}
}
