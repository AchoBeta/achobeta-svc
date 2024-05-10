package middleware

import (
	"achobeta-svc/internal/achobeta-svc-common/pkg/tlog"
	"achobeta-svc/internal/achobeta-svc-third-party/inernal/router/manager"
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func init() {
	manager.RouteHandler.RegisterMiddleware(manager.LEVEL_GLOBAL, AddTraceId, false)
}

func AddTraceId() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		// 假设 Trace ID 存在于 HTTP Header "X-Trace-ID" 中
		traceID := ctx.Request.Header.Get("X-Request-ID")
		if traceID == "" {
			traceID = uuid.New().String()
		}
		c = tlog.NewContext(c, zap.String("traceId", traceID))
		ctx.Next(c)
	}
}
