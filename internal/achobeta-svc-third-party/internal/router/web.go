package router

import (
	"achobeta-svc/internal/achobeta-svc-common/lib/tlog"
	"achobeta-svc/internal/achobeta-svc-common/pkg/web"
	"achobeta-svc/internal/achobeta-svc-third-party/config"
	"fmt"

	_ "achobeta-svc/internal/achobeta-svc-third-party/internal/api"
	_ "achobeta-svc/internal/achobeta-svc-third-party/internal/router/middleware"

	"github.com/gin-gonic/gin"
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
