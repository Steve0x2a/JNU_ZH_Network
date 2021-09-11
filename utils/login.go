package utils

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/go-resty/resty/v2"
)

type respStruct struct {
	PL          string `json:"portalLink"`
	SerialNo    int    `json:"serialNo"`
	UserStatus  int    `json:"userStatus"`
	UserDevPort string `json:"userDevPort"`
}

func Login() (bool, respStruct) {
	client := resty.New()
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	loginFlag := false
	var resps respStruct
	formData := map[string]string{
		"userName":            Config.Username,
		"userPwd":             base64.StdEncoding.EncodeToString([]byte(Config.Password)),
		"userDynamicPwd":      "",
		"userDynamicPwdd":     "",
		"serviceTypeHIDE":     "",
		"serviceType":         "",
		"userurl":             "",
		"userip":              "",
		"basip":               "",
		"language":            "Chinese",
		"usermac":             "null",
		"wlannasid":           "",
		"wlanssid":            "",
		"entrance":            "null",
		"loginVerifyCode":     "",
		"userDynamicPwddd":    "",
		"customPageId":        "105",
		"pwdMode":             "0",
		"portalProxyIP":       Config.AuthIP,
		"portalProxyPort":     "50200",
		"dcPwdNeedEncrypt":    "1",
		"assignIpType":        "0",
		"appRootUrl":          Config.LoginURL,
		"manualUrl":           "",
		"manualUrlEncryptKey": "",
	}
	headers := map[string]string{
		"User-Agent":      "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.99 Safari/537.36",
		"Accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8",
		"Accept-Language": "zh-CN,zh;q=0.9",
	}
	res, err := client.R().
		EnableTrace().
		SetHeaders(headers).
		SetFormData(formData).
		Post(Config.LoginURL + "pws?t=li")
	if err != nil {
		fmt.Println(nil)
	}
	respJson := decodeResp(res.String())
	if strings.Contains(respJson, "errorNumber") && strings.Contains(respJson, "heartBeatTimeoutMaxTime") {
		log.Println("成功登录")
		loginFlag = true
	} else if strings.Contains(respJson, "E63032:密码错误") {
		// 密码错误 	{"portServIncludeFailedCode":"63032","portServIncludeFailedReason":"E63032:密码错误，您还可以重试8次。","e_c":"portServIncludeFailedCode","e_d":"portServIncludeFailedReason","errorNumber":"7"}
		fmt.Println("密码错误")
		log.Println("用户密码错误")
		loginFlag = false
	} else if strings.Contains(respJson, "E63018:用户不存在或者用户没有申请该服务") {
		// 用户名不存在 {"portServIncludeFailedCode":"63018","portServIncludeFailedReason":"E63018:用户不存在或者用户没有申请该服务。","e_c":"portServIncludeFailedCode","e_d":"portServIncludeFailedReason","errorNumber":"7"}
		log.Println("用户不存在或者用户没有申请该服务")
		fmt.Println("用户不存在或者用户没有申请该服务")
		loginFlag = false
	} else if strings.Contains(respJson, "设备拒绝请求") {
		// 已经登录		{"portServErrorCode":"1","portServErrorCodeDesc":"设备拒绝请求","e_c":"portServErrorCode","e_d":"portServErrorCodeDesc","errorNumber":"7"}
		log.Println("用户已经登录, 尝试强制下线后重新登录")
		loginFlag, resps = Login()
	} else {
		fmt.Println("登录失败, 错误信息如下: " + respJson)
		log.Println("登录失败, 错误信息如下: " + respJson)
		loginFlag = false
	}
	json.Unmarshal([]byte(respJson), &resps)
	return loginFlag, resps
}

func decodeResp(msg string) string {
	temp, _ := base64.StdEncoding.DecodeString(msg)
	if string(temp) != "" {
		respJson, _ := url.QueryUnescape(string(temp))
		if respJson != "" {
			return respJson
		}
	}
	temp, _ = base64.StdEncoding.DecodeString(msg + "=")
	if string(temp) != "" {
		respJson, _ := url.QueryUnescape(string(temp))
		if respJson != "" {
			return respJson
		}
	}
	temp, _ = base64.StdEncoding.DecodeString(msg + "==")
	if string(temp) != "" {
		respJson, _ := url.QueryUnescape(string(temp))
		if respJson != "" {
			return respJson
		}
	}
	return "null"
}
