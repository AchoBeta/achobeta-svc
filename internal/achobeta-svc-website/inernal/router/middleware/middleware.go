package middleware

import (
	"achobeta-svc/internal/achobeta-svc-common/pkg/constant"
	"achobeta-svc/internal/achobeta-svc-common/pkg/tlog"
	"achobeta-svc/internal/achobeta-svc-common/pkg/web"
	"achobeta-svc/internal/achobeta-svc-website/config"
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func init() {
	web.RouteHandler.RegisterMiddleware(web.LEVEL_GLOBAL, AddTraceId, true)
	web.RouteHandler.RegisterMiddleware(web.LEVEL_GLOBAL, ErrorHandler, true)
	web.RouteHandler.RegisterMiddleware(web.LEVEL_V1, CheckToken, true)
}

func AddTraceId() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 假设 Trace ID 存在于 HTTP Header "X-Trace-ID" 中
		traceId := c.GetHeader("traceId")
		if traceId == "" {
			traceId = uuid.New().String()
		}
		ctx := tlog.NewContext(c, zap.String("traceId", traceId))
		c.Request = c.Request.WithContext(ctx)
		c.Keys = map[string]any{
			"traceId": traceId,
		}
		c.Next()
	}
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		for _, e := range c.Errors {
			err := e.Err
			if err != nil {
				r := web.NewResponse(c)
				tlog.Infof("error: %v", err)
				r.ErrorTrace(constant.COMMON_FAIL, err.Error(), c.Request.Context().Value("traceId").(string))
				return
			}
		}
	}
}

func CheckToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader(string(constant.RequestHeaderKeyToken))
		if token == "" {
			r := web.NewResponse(c)
			r.ErrorCode(constant.TOKEN_IS_NULL)
			c.Abort()
		}
		userId, err := config.GetRedis().Get(token).Int()
		if err != nil {
			r := web.NewResponse(c)
			r.ErrorCode(constant.TOKEN_IS_INVALID)
			c.Abort()
		}
		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), constant.RequestHeaderKeyUserId, uint(userId)))
		c.Next()
	}
}
