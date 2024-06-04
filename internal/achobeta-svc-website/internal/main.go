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
	// 初始化配置, 需要保证config和log先初始化
	config.InitConfig(*configPath)
	config.InitLog(*logFilePath)
	// 其他初始化, 顺序自行调整
	config.InitDatabase()
	config.InitRedis()
	// 初始化casbin, 因为要用到数据库的配置，所以要放在数据库初始化之后
	casbin.Init(*casbinPath)
	// 初始化服务
	utils.NewSnowflake()
	/** server 启动要放在最后*/
	// router.RunServer()
	router.RunRPCServer()
}
