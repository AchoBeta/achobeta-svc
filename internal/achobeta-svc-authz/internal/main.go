package main

import (
	"achobeta-svc/internal/achobeta-svc-authz/config"
	"achobeta-svc/internal/achobeta-svc-authz/internal/service"
	"achobeta-svc/internal/achobeta-svc-common/lib/tlog"
	"achobeta-svc/internal/achobeta-svc-common/pkg/utils"
	permissionv1 "achobeta-svc/internal/achobeta-svc-proto/gen/go/authz/permission/v1"
	"fmt"
	"net"

	"flag"

	"google.golang.org/grpc"
)

func main() {
	configPath := flag.String("config", "", "specify config path [config.yaml]")
	logFilePath := flag.String("logs", "./logs/", "log file path")
	flag.Parse()
	// 初始化配置, 需要保证config和log先初始化
	config.InitConfig(*configPath)
	config.InitLog(*logFilePath)
	// 初始化服务
	utils.NewSnowflake()
	/** server 启动要放在最后*/
	// router.RunServer()
	runRPCServer()
}

func runRPCServer() {
	c := config.Get()
	//tcp协议监听指定端口号
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", c.Port))
	if err != nil {
		tlog.Errorf("failed to listen: %v", err)
		panic(err)
	}
	//实例化gRPC服务
	s := grpc.NewServer()
	svcs := service.InitServices()
	permissionv1.RegisterAuthzServiceServer(s, svcs.PermissionService)
	//服务注册
	tlog.Infof("Listen on %s:%d", c.Host, c.Port)
	//启动服务
	if err := s.Serve(lis); err != nil {
		tlog.Errorf("failed to serve: %v", err)
		panic(err)
	}
}
