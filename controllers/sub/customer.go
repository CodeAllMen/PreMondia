package sub

import (
	"strconv"

	"github.com/MobileCPX/PreMondia/models"
	"github.com/MobileCPX/PreMondia/models/sub"
	"github.com/astaxie/beego"
)

// GetCustomerControllers 获取CustomerID
type GetCustomerControllers struct {
	beego.Controller
}

// Get 请求
func (c *GetCustomerControllers) Get() {
	trackID := c.Ctx.Input.Param("id") // //124|game_d 服务类型及id
	status := c.GetString("status")
	customerID := c.GetString("customerId")
	operatorCountry := c.GetString("operator")
	errorDesc := c.GetString("errorDesc")
	errorCode := c.GetString("errorCode")
	trackIDInt, err := strconv.Atoi(trackID)
	if err != nil {
		c.Redirect("http://google.com", 302)
		return
	}
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
}

func checkSubStatus(trackData *models.AffTrack) {
	var redirectURL ""
	if trackData.Status == "SUCCESS" { // 3G网络环境下订阅
		
	}
}
