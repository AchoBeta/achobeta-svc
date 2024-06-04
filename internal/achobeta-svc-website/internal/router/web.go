package router

import (
	"achobeta-svc/internal/achobeta-svc-common/pkg/tlog"
	"achobeta-svc/internal/achobeta-svc-common/pkg/web"
	"achobeta-svc/internal/achobeta-svc-website/config"

	"fmt"

	_ "achobeta-svc/internal/achobeta-svc-website/internal/api"
	_ "achobeta-svc/internal/achobeta-svc-website/internal/router/middleware"

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
