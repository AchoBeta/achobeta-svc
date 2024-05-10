package main

import (
	"achobeta-svc/internal/achobeta-svc-third-party/config"

	"achobeta-svc/internal/achobeta-svc-third-party/inernal/router"
	"flag"
)

func main() {
	configPath := flag.String("config", "", "specify config path [config.yaml]")
	logFilePath := flag.String("l", "./logs/", "log file path")
	flag.Parse()
	// 初始化配置
	config.InitConfig(*configPath)
	config.InitLog(*logFilePath)
	/** server 启动要放在最后*/
	router.RunServer()
}
