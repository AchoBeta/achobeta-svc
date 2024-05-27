package main

import (
	"achobeta-svc/internal/achobeta-svc-common/pkg/utils"

	"achobeta-svc/internal/achobeta-svc-website/config"
	"achobeta-svc/internal/achobeta-svc-website/inernal/router"

	"flag"
)

func main() {
	configPath := flag.String("config", "", "specify config path [config.yaml]")
	logFilePath := flag.String("logs", "./logs/", "log file path")
	flag.Parse()
	// 初始化配置
	config.InitConfig(*configPath)
	config.InitLog(*logFilePath)
	config.InitDatabase()
	// 初始化服务
	utils.NewSnowflake()
	// service.LoadService()
	/** server 启动要放在最后*/
	router.RunServer()
}
