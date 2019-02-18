package initial

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/astaxie/beego/logs"
)

var MondiaConf *MondiaOrangeConfData

// MondiaOrangeConfData Mondia 配置struct
type MondiaOrangeConfData struct {
	RequestURL  string `json:"request_url"`
	MrchantID   string `json:"mrchant_id"`
	PaymentURL  string `json:"payment_url"`
	ProductCode string `json:"product_code"`
	ProdPrice   string `json:"prodPrice"`
	ImgPath     string `json:"img_path"`
	OperatorID  string `json:"operator_id"`
	SubPackage  string `json:"subPackage"`
}

// InitMondiaConf 初始化话数据操作
func InitMondiaConf() {
	file, d := os.Open("source/conf/mondia_conf.json")
	fmt.Println(d)
	defer file.Close()

	decoder := json.NewDecoder(file)
	conf := new(MondiaOrangeConfData)
	// fmt.Println(s)
	err := decoder.Decode(conf)
	if err != nil {
		logs.Emergency("mondia 数据初始化错误")
		fmt.Println("Error:", err)
	} else {
		MondiaConf = conf
	}
	fmt.Println("!!!!!!!!!!!!!!!!!!!")
}

// GetMondiaConf 获取Mondia配置文件
func GetMondiaConf() *MondiaOrangeConfData {
	return MondiaConf
}
