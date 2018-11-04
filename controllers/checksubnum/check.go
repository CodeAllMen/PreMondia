package checksubnum

import (
	"github.com/MobileCPX/PreMondia/models/sub"
	"github.com/astaxie/beego"
)

type CheckSubNum struct {
	beego.Controller
}

func (c *CheckSubNum) Get() {
	limitSub := sub.CheckTodaySubNum(48)
	limitSubStr := "NO"
	if limitSub {
		limitSubStr = "YES"
	}
	c.Ctx.WriteString(string(limitSubStr))
}
