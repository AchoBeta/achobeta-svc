package logic

import (
	"achobeta-svc/internal/achobeta-svc-authz/config"
	"achobeta-svc/internal/achobeta-svc-authz/internal/logic/service"
	"achobeta-svc/internal/achobeta-svc-common/pkg/tlog"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

func RunRPCServer() {
	c := config.Get()
	//tcp协议监听指定端口号
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", c.Port))
	if err != nil {
		tlog.Errorf("failed to listen: %v", err)
		panic(err)
	}
	//实例化gRPC服务
	s := grpc.NewServer()
	//服务注册
	service.New(s)
	tlog.Infof("Listen on %s:%d", c.Host, c.Port)
	//启动服务
	if err := s.Serve(lis); err != nil {
		tlog.Errorf("failed to serve: %v", err)
		panic(err)
	}
}
