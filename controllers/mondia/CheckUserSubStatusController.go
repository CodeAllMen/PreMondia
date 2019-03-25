package mondia

import "github.com/MobileCPX/PreMondia/models/mondia"

type CheckUserSubStatusController struct {
	BaseController
}

func (c *CheckUserSubStatusController) Get() {
	user := c.GetString("user")
	if user != "" {
		mo := new(mondia.Mo)
		_ = mo.GetMoByCustomerID(user)
		if mo.ID != 0 {
			c.Ctx.WriteString("YES")
			c.StopRun()
		}
	}
	c.Ctx.WriteString("NO")
}
