package main

import (
	"achobeta-svc/internal/achobeta-svc-api/config"
	"achobeta-svc/internal/achobeta-svc-api/internal/server"
	"achobeta-svc/internal/achobeta-svc-common/lib/tlog"
	"flag"
)

func main() {
	configPath := flag.String("config", "", "specify config path [config.yaml]")
	logFilePath := flag.String("logs", "./logs/", "log file path")
	flag.Parse()
	// 初始化配置, 需要保证config和log先初始化
	config.InitConfig(*configPath)
	tlog.InitLog(*logFilePath)
	/** server 启动要放在最后*/
	server.RunServer()
}
