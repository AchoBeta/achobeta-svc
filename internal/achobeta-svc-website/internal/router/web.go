package router

import (
	"achobeta-svc/internal/achobeta-svc-common/lib/tlog"
	"achobeta-svc/internal/achobeta-svc-common/pkg/web"
	"achobeta-svc/internal/achobeta-svc-website/config"
	_ "achobeta-svc/internal/achobeta-svc-website/internal/api"
	_ "achobeta-svc/internal/achobeta-svc-website/internal/router/middleware"
	"context"
	"fmt"
	"net"

	hello "achobeta-svc/internal/achobeta-svc-proto/gen/go/website/hello/v1"

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

type HelloService struct {
	// UnimplementedHelloServiceServer这个结构体是必须要内嵌进来的
	// 也就是说我们定义的这个结构体对象必须继承UnimplementedHelloServiceServer。
	// 嵌入之后，我们就已经实现了GRPC这个服务的接口，但是实现之后我们什么都没做，没有写自己的业务逻辑，
	// 我们要重写实现的这个接口里的函数，这样才能提供一个真正的rpc的能力。
	hello.UnimplementedHelloServiceServer
}

// Hello 重写实现的接口里的Hello函数
func (p *HelloService) Hello(ctx context.Context, req *hello.HelloRequest) (*hello.HelloResponse, error) {
	resp := &hello.HelloResponse{}
	resp.Value = "hello:" + req.Value
	return resp, nil
}

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
	hello.RegisterHelloServiceServer(s, &HelloService{})
	tlog.Infof("Listen on %s:%d", c.Host, c.Port)
	//启动服务
	if err := s.Serve(lis); err != nil {
		tlog.Errorf("failed to serve: %v", err)
		panic(err)
	}
}
