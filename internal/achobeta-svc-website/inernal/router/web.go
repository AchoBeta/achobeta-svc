package router

import (
	"achobeta-svc/internal/achobeta-svc-common/pkg/tlog"
	"achobeta-svc/internal/achobeta-svc-third-party/config"
	"achobeta-svc/internal/achobeta-svc-third-party/inernal/router/manager"
	"fmt"

	_ "achobeta-svc/internal/achobeta-svc-third-party/inernal/api"
	_ "achobeta-svc/internal/achobeta-svc-third-party/inernal/router/middleware"

	"github.com/gin-gonic/gin"
)

func RunServer() {
	c := config.Get()
	g := gin.New()
	// tlog.CtxInfof(context.Background(), "Listen on %s:%d", c.Host, c.Port)
	manager.RouteHandler.Register(g)
	// run 在最后
	err := g.Run(fmt.Sprintf("%s:%d", c.Host, c.Port))
	if err != nil {
		tlog.Errorf("Listen error: %v", err)
		panic(err)
	}
}
