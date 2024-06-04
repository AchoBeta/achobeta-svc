package router

import (
	"achobeta-svc/internal/achobeta-svc-common/pkg/tlog"
	"achobeta-svc/internal/achobeta-svc-common/pkg/web"
	"achobeta-svc/internal/achobeta-svc-website/config"
	_ "achobeta-svc/internal/achobeta-svc-website/internal/api"
	_ "achobeta-svc/internal/achobeta-svc-website/internal/router/middleware"
	"fmt"
	"log"
	"net"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func RunServer() {
	c := config.Get()
	g := gin.New()
	// tlog.CtxInfof(context.Background(), "Listen on %s:%d", c.Host, c.Port)
	web.RouteHandler.Register(g)
	// run 在最后
	err := g.Run(fmt.Sprintf("%s:%d", c.Host, c.Port))
	if err != nil {
		tlog.Errorf("Listen error: %v", err)
		panic(err)
	}
}

func RunRPCServer() {
	c := config.Get()
	//tcp协议监听指定端口号
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", c.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	tlog.Infof("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
