package main

import (
	"achobeta-svc/internal/achobeta-svc-authz/config"
	"achobeta-svc/internal/achobeta-svc-authz/internal/service"
	server "achobeta-svc/internal/achobeta-svc-common/lib"
	"achobeta-svc/internal/achobeta-svc-common/lib/tlog"
	"achobeta-svc/internal/achobeta-svc-common/pkg/utils"
	permissionv1 "achobeta-svc/internal/achobeta-svc-proto/gen/go/authz/permission/v1"
	"log"
	"os"
	"os/signal"
	"syscall"

	"flag"

	ants "github.com/panjf2000/ants/v2"
)

func main() {
	configPath := flag.String("config", "", "specify config path [config.yaml]")
	logFilePath := flag.String("logs", "./logs/", "log file path")
	flag.Parse()
	// 初始化配置, 需要保证config和log先初始化
	config.InitConfig(*configPath)
	tlog.InitLog(*logFilePath)
	config.InitDatabase()
	// 初始化服务
	utils.NewSnowflake()
	/** server 启动要放在最后*/
	// router.RunServer()
	runRPCServer()
}

func runRPCServer() {
	c := config.Get()
	//实例化gRPC服务
	s := server.NewServer(server.NewConfig(server.WithPort(c.Port)))
	svcs := service.InitServices()
	permissionv1.RegisterAuthzServiceServer(s, svcs.PermissionService)
	//服务注册
	tlog.Infof("Listen on %s:%d", c.Host, c.Port)

	s.Start()
	tlog.Infof("server started")

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	<-ch

	s.Stop()
	ants.Release()
	log.Println("server stopped")
}
