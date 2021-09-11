package main

import (
	"flag"
	"fmt"
	"jnu_network/utils"
	"time"
)

func main() {
	configPath := flag.String("config", "./config.yml", "配置文件地址, 默认为\"./config.yml\"")
	logFlag := flag.Bool("log", false, "日志输出到文件,默认为false即输出到终端, 打开则输入到文件")
	flag.Parse()
	utils.LogInit(*logFlag)
	fmt.Println("使用配置:", *configPath)
	utils.Config.GetConf(*configPath)
	loginFlag, loginStruct := utils.Login()
	count := 0
	hbcount := 0
	for loginFlag && count < 5 {
		hbFlag := utils.HeartBeat(loginStruct, &hbcount)
		if !hbFlag {
			count++
		}
		time.Sleep(time.Minute * time.Duration(utils.Config.HBTime))
	}
}
