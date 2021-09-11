package utils

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/url"
	"strconv"
	"strings"

	"github.com/go-resty/resty/v2"
)

func HeartBeat(loginStruct respStruct, num *int) bool {
	client := resty.New()
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	headers := map[string]string{
		"User-Agent":      "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.99 Safari/537.36",
		"Accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8",
		"Accept-Language": "zh-CN,zh;q=0.9",
		"Origin":          Config.AuthURL,
		"Referer":         fmt.Sprintf("%vpage/online_heartBeat.jsp?pl=%v&custompath=&uamInitCustom=0&uamInitLogo=H3C", Config.LoginURL, loginStruct.PL),
		"Cookie":          fmt.Sprintf("hello1=%v,hello2=false", Config.Username),
	}
	formData := map[string]string{
		"userip":      "",
		"basip":       "",
		"userDevPort": url.QueryEscape(loginStruct.UserDevPort),
		"userStatus":  strconv.Itoa(loginStruct.UserStatus),
		"serialNo":    strconv.Itoa(loginStruct.SerialNo),
		"language":    "Chinese",
		"e_d":         "",
		"t":           "hb",
	}
	resp, _ := client.R().
		SetHeaders(headers).
		SetFormData(formData).
		Post(Config.LoginURL + "page/doHeartBeat.jsp")
	if resp.StatusCode() == 200 && strings.Contains(resp.String(), "v_failedTimes=0") {
		*num++
		if *num >= 5 {
			log.Println("已发送5次心跳包")
			*num = 0
		}
		return true
	} else {
		log.Println("心跳包发送失败")
		return false
	}
}
