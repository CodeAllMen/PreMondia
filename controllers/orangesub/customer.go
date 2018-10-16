package orangesub

import (
	"fmt"
	"strconv"

	"github.com/MobileCPX/PreMondia/initial"
	"github.com/MobileCPX/PreMondia/models/sub"
	"github.com/astaxie/beego"
)

// GetCustomerControllers 获取CustomerID
type GetCustomerControllers struct {
	beego.Controller
}

// GetPostResquest
type GetPostRequestControlelr struct {
	beego.Controller
}

func (c *GetPostRequestControlelr) Post() {
	frmlp := c.GetString("frmlp")
	fmt.Println(frmlp)
	c.Redirect("http://sso.orange.com/mondiamedia_subscription/?method=getcustomer&merchantId=93&langCode=pl&redirect=http://cpx3.allcpx.com/subs/getcust/"+frmlp, 302)
}

// Get 请求
func (c *GetCustomerControllers) Get() {
	trackID := c.Ctx.Input.Param(":trackID") // id
	status := c.GetString("status")
	customerID := c.GetString("customerId")
	operatorCountry := c.GetString("operator")
	errorDesc := c.GetString("errorDesc")
	errorCode := c.GetString("errorCode")

	// isSub, subID := sub.CheckUserSubStatus(customerID)
	// if isSub { // 用户已经订阅
	// 	// 需要将customerID注册一次
	// 	util.HttpRequest(subID, "register", "video", "", "")
	// 	c.Redirect("http://www.redlightvideos.com/mm/pl?sub="+subID, 302)
	// 	return
	// }

	trackIDInt, err := strconv.Atoi(trackID)
	if err != nil {
		c.Redirect("http://google.com", 302)
		return
	}
	// 获取 tracking 数据
	trackData, isHas := sub.GetAffTrackData(trackIDInt)
	if !isHas {
		c.Redirect("http://google.com", 302)
		return
	}
	trackData.ErrorCode = errorCode
	trackData.ErrorDesc = errorDesc
	trackData.Operator = operatorCountry
	trackData.CustomerID = customerID
	trackData.Status = status
	err = sub.UpdateTrackData(trackData)
	if err != nil {
		c.Redirect("http://google.com", 302)
		return
	}
	subURL := redirectSubURL(trackID)
	c.Redirect(subURL, 302)
}

// 获取支付页面
func redirectSubURL(trackID string) string {
	orangeConf := initial.GetMondiaConf()
	url := "http://sso.orange.com/mondiamedia_subscription/?method=subscribe&merchantId=93&redirect=" +
		"http%3a%2f%2fcpx3.allcpx.com%2fsubs%2fres%2f" + trackID + "&imgPath=" + orangeConf.ImgPath + "&productCd=" +
		orangeConf.ProductCode + "&subPackage=" + orangeConf.SubPackage + "&operatorId=8&&langCode=pl"
	fmt.Println(url)
	return url
}
