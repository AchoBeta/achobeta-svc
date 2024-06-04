package main

import (
	"achobeta-svc/internal/achobeta-svc-common/pkg/utils"

	"achobeta-svc/internal/achobeta-svc-website/config"
	"achobeta-svc/internal/achobeta-svc-website/config/casbin"
	"achobeta-svc/internal/achobeta-svc-website/internal/router"

	"flag"
)

func main() {
	configPath := flag.String("config", "", "specify config path [config.yaml]")
	logFilePath := flag.String("logs", "./logs/", "log file path")
	casbinPath := flag.String("casbin", "", "specify casbin model path [casbin_model.conf]")
	flag.Parse()
	// 初始化配置
	config.InitConfig(*configPath)
	config.InitLog(*logFilePath)
	casbin.Init(*casbinPath)
	config.InitDatabase()
	config.InitRedis()
	// 初始化服务
	utils.NewSnowflake()
	// service.LoadService()
	/** server 启动要放在最后*/
	router.RunServer()
}
