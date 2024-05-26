package middleware

import (
	"achobeta-svc/internal/achobeta-svc-third-party/inernal/router/manager"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func init() {
	manager.RouteHandler.RegisterMiddleware(manager.LEVEL_GLOBAL, AddTraceId, false)
}

func AddTraceId() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 假设 Trace ID 存在于 HTTP Header "X-Trace-ID" 中
		traceID := c.Request.Header.Get("X-Request-ID")
		if traceID == "" {
			traceID = uuid.New().String()
		}
	}
}
