package api

import (
	"achobeta-svc/internal/achobeta-svc-common/pkg/constant"
	"achobeta-svc/internal/achobeta-svc-third-party/inernal/router/manager"
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/route"
)

func init() {
	manager.RouteHandler.RegisterRouter(manager.LEVEL_GLOBAL, func(r *route.RouterGroup) {
		r.GET("/txcloud", ping)
	})
}

func ping(ctx context.Context, c *app.RequestContext) {
	c.JSON(constant.SUCCESS, "pong")
}
