package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func LogInit(logFlag bool) {
	if logFlag {
		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			log.Fatal(err)
		}
		file := fmt.Sprintf("%v/networkLog.txt", dir)
		fmt.Println("日志保存在:", file)
		logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0766)
		if err != nil {
			fmt.Println("配置文件地址有误, 请重新确认")
			panic(err)
		}
		log.SetOutput(logFile) // 将文件设置为log输出的文件
	}

	log.SetPrefix("[JNU_ZHUHAI]")
	log.SetFlags(log.LstdFlags | log.Ldate)
}
