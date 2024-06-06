package main

import (
	"achobeta-svc/internal/achobeta-svc-website/config"
	"achobeta-svc/internal/achobeta-svc-website/internal/router"

	"flag"
)

func main() {
	configPath := flag.String("config", "", "specify config path [config.yaml]")
	logFilePath := flag.String("logs", "./logs/", "log file path")
	flag.Parse()
	// 初始化配置, 需要保证config和log先初始化
	config.InitConfig(*configPath)
	config.InitLog(*logFilePath)
	/** server 启动要放在最后*/
	router.RunServer()
}
