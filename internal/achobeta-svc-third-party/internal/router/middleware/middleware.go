package middleware

import (
	"achobeta-svc/internal/achobeta-svc-common/pkg/tlog"
	"achobeta-svc/internal/achobeta-svc-common/pkg/web"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func init() {
	web.RouteHandler.RegisterMiddleware(web.LEVEL_GLOBAL, AddTraceId, false)
}

func AddTraceId() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 假设 Trace ID 存在于 HTTP Header "X-Trace-ID" 中
		traceID := c.GetHeader("traceId")
		if traceID == "" {
			traceID = uuid.New().String()
		}
		ctx := tlog.NewContext(c, zap.String("traceId", traceID))
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
