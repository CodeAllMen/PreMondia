package mondia

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"
)

// Config 内容站配置
type Config struct {
	Service map[string]ServiceInfo
}

type ServiceInfo struct {
	ServiceID   string `yaml:"service_id" orm:"pk;column(service_id)"`
	ServiceName string `yaml:"service_name"`
	ProductCode string `yaml:"product_code"`
	RequestURL  string `yaml:"request_url" orm:"column(request_url)"`
	MrchantID   string `yaml:"mrchant_id" orm:"column(mrchant_id)"`
	ProdPrice   string `yaml:"prod_price"`
	ImgPath     string `yaml:"img_path"`
	OperatorID  string `yaml:"operator_id"`
	SubPackage  string `yaml:"sub_package"`
	LpURL       string `yaml:"LP_url" orm:"column(lp_url)"`
	DeleteURL   string `yaml:"delete_url" orm:"column(delete_url)"`
	ContentURL  string `yaml:"content_url" orm:"column(content_url)"`
	UnsubURL    string `yaml:"unsub_url" orm:"column(unsub_url)"`
	RegisterURL string `yaml:"register_url" orm:"column(register_url)"`
}

var ServiceData = make(map[string]ServiceInfo)

func InitServiceConfig() {
	filename, _ := filepath.Abs("resource/config/conf.yaml")
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	config := new(Config)
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		panic(err)
	}
	ServiceData = config.Service

	o := orm.NewOrm()
	for _, v := range ServiceData {
		_, _ = o.Insert(&v)
	}

}

func (server *ServiceInfo) TableName() string {
	return "server_info"
}

func (server *ServiceInfo) GetServiceInfo() (*ServiceInfo, error) {
	mapServer := ServiceData[server.ServiceID]
	if mapServer.ServiceID != "" {
		return &mapServer, nil
	} else {
		o := orm.NewOrm()
		err := o.Read(server)
		if err != nil {
			logs.Error("查询服务失败 serverID: ", server.ServiceID)
		}
		return server, err
	}

}

//HttpRequest 注册
func ServiceRegisterRequest(subscriptionID, customerID, serviceID, types string) (contentURL string) {
	requestURL := ""
	service := ServiceData[serviceID]
	if types == "register" {
		registerURL := service.RegisterURL
		registerURL = strings.Replace(registerURL, "{subID}", subscriptionID, -1)
		registerURL = strings.Replace(registerURL, "{msisdn}", customerID, -1)
		requestURL = registerURL
	}
	fmt.Println(requestURL)
	resp, err := http.Get(requestURL)
	if err == nil {
		logs.Info(fmt.Sprintf("HttpRequest Success %s service %s subID: %s  CustomerId: %s ",
			types, serviceID, subscriptionID, customerID))
		resp.Body.Close()
	} else {
		logs.Error(fmt.Sprintf("HttpRequest Failed %s service %s subID: %s  CustomerId: %s "+
			"  error: %s ", types, serviceID, subscriptionID, customerID, err.Error()))
	}
	// 获取内容站URL sub 自动登录
	contentURL = strings.Replace(service.ContentURL, "{subID}", subscriptionID, -1)

	return
}

func GetPaymentURL(serviceID string, trackID string) (paymentURL string, isExist bool) {
	service, isExist := ServiceData[serviceID]
	if !isExist {
		logs.Error("GetPaymentURL product_code 不存在 ", serviceID)
		return
	}
	paymentURL = "http://login.mondiamediamena.com/billinggw-lcm/billing?method=subscribe&merchantId=247&redirect=http%3" +
		"A%2F%2Fmm-eg.leadernethk.com/get/sub_result/" + trackID + "&productCd=" + service.ProductCode + "&subPackage=" +
		service.SubPackage + "&operatorId=1&&imgPath=" + service.ImgPath

	//paymentURL = "http://sso.orange.com/mondiamedia_subscription/?method=subscribe&merchantId=93&redirect=" +
	//	"http%3a%2f%2fcpx3.allcpx.com:8085%2fsubs%2fres%2f" + trackID + "&imgPath=" + service.ImgPath + "&productCd=" +
	//	service.ProductCode + "&subPackage=" + service.SubPackage + "&operatorId=8&langCode=pl"
	fmt.Println("paymentURL: ", paymentURL)
	return
}

func GetLpURL(serviceID string) (LpURL string, isExist bool) {
	service, isExist := ServiceData[serviceID]
	if !isExist {
		logs.Error("GetLpURL product_code 不存在 ", serviceID)
		return
	}
	LpURL = service.LpURL
	return
}

func GetContentURL(serviceID string) (contentURL string) {
	service, isExist := ServiceData[serviceID]
	if !isExist {
		logs.Error("GetLpURL product_code 不存在 ", serviceID)
		return
	}
	contentURL = service.ContentURL
	return
}

func GetPINUnsubMessage(serviceID, PIN string) (message string) {
	service, isExist := ServiceData[serviceID]
	if !isExist {
		message = "Kod PIN, który anulowałeś swoją subskrypcję, to " + PIN
		logs.Error("GetPINUnsubMessage product_code 不存在 ", serviceID)
		return
	}

	message = fmt.Sprintf("[%s] Kod PIN, który anulowałeś swoją subskrypcję, to "+PIN, service.ServiceID)
	return

}
