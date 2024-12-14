package main

import (
	"achobeta-svc/backend/api/config"
	"achobeta-svc/backend/api/internal/server"
	"achobeta-svc/backend/common/lib/tlog"
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
