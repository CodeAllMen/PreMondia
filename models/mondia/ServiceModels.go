package mondia

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"
)

// Config 内容站配置
type Config struct {
	Service map[string]ServiceInfo
}

type ServiceInfo struct {
	ServiceID              string `yaml:"service_id"`
	ServiceName            string `yaml:"service_name"`
	ProductCode            string `yaml:"product_code"`
	Password               string `yaml:"password"`
	Username               string `yaml:"username"`
	RequestURL             string `yaml:"request_url"`
	MrchantID              string `yaml:"mrchant_id" `
	ProdPrice              string `yaml:"prod_price"`
	ImgPath                string `yaml:"img_path"`
	OperatorID             string `yaml:"operator_id"`
	SubPackage             string `yaml:"sub_package"`
	LpURL                  string `yaml:"LP_url"`
	DeleteURL              string `yaml:"delete_url"`
	ContentURL             string `yaml:"content_url"`
	UnsubURL               string `yaml:"unsub_url" `
	RegisterURL            string `yaml:"register_url"`
	MondiaRequestURL       string `yaml:"mondia_request_url"`        // mondia 订阅url
	GetCustomerCallbackURL string `yaml:"get_customer_callback_url"` // 获取CustomerID 之后的回到URL
	SubResultRedirect      string `yaml:"sub_result_redirect"`
	UnsubPINMessage        string `yaml:"unsub_pin_message"`
	SuccessSubMessage      string `yaml:"success_sub_message"`
	Language               string `yaml:"language"`
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

func GetPaymentURL(serviceConfig ServiceInfo, trackID string) (paymentURL string) {

	paymentURL = fmt.Sprintf("%s?method=subscribe&merchantId=%s&redirect=%s&productCd=%s&"+
		"subPackage=%s&operatorId=%s&imgPath=%s", serviceConfig.MondiaRequestURL, serviceConfig.MrchantID,
		url.QueryEscape(serviceConfig.SubResultRedirect+trackID), serviceConfig.ProductCode, serviceConfig.SubPackage,
		serviceConfig.OperatorID, url.QueryEscape(serviceConfig.ImgPath))
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
	fmt.Println(serviceID)
	service, isExist := ServiceData[serviceID]
	fmt.Println(service)
	if !isExist {
		logs.Error("GetPINUnsubMessage product_code 不存在 ", serviceID)
		return
	}
	message = strings.Replace(service.UnsubPINMessage, "{PIN}", PIN, -1)
	fmt.Println(message)
	return

}
