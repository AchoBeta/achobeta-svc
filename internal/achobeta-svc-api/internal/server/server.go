package server

import (
	"achobeta-svc/internal/achobeta-svc-api/config"
	_ "achobeta-svc/internal/achobeta-svc-api/internal/server/middleware"
	"achobeta-svc/internal/achobeta-svc-api/internal/server/route"
	"achobeta-svc/internal/achobeta-svc-api/internal/server/service"
	"achobeta-svc/internal/achobeta-svc-common/lib/tlog"
	"fmt"

	"github.com/gin-gonic/gin"
)

// RunServer 启动服务
// api层的服务启动 用的gin框架
// 所以在api层的服务启动中，需要注册路由
// 前端与api层的交互走的是http协议, api与service层的交互走的是grpc协议
func RunServer() {
	c := config.Get()
	_ = service.InitServices()
	g := gin.New()
	route.Injection(g)

	// run 在最后
	runServer := fmt.Sprintf("%s:%d", c.Host, c.Port)
	tlog.Infof("gateway server listen on %s", runServer)
	err := g.Run(runServer)
	if err != nil {
		tlog.Errorf("Listen error: %v", err)
		panic(err)
	}
}
