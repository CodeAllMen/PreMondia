package unsub

import (
	"github.com/astaxie/beego"
)

type SubIDUnsubRequest struct {
	beego.Controller
}

func (c *SubIDUnsubRequest) Get() {
	subID := c.GetString("subid")
	unsubURL := "http://payment.mondiamediamena.com/billing-gw/subservice/unsubscribe?subid=" + subID + "&operatorId=8"
	status, body := MondiaHTTPRequest(unsubURL)
	c.Ctx.WriteString(status + "####" + string(body))
}
