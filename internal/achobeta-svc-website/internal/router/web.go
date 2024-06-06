package router

import (
	"achobeta-svc/internal/achobeta-svc-common/lib/tlog"
	"achobeta-svc/internal/achobeta-svc-common/pkg/web"
	"achobeta-svc/internal/achobeta-svc-website/config"
	_ "achobeta-svc/internal/achobeta-svc-website/internal/api"
	_ "achobeta-svc/internal/achobeta-svc-website/internal/router/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
)

func RunServer() {
	c := config.Get()
	g := gin.New()
	web.RouteHandler.Register(g)
	// run 在最后
	err := g.Run(fmt.Sprintf("%s:%d", c.Host, c.Port))
	if err != nil {
		tlog.Errorf("Listen error: %v", err)
		panic(err)
	}
}
