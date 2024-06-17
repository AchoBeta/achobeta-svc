package middleware

import (
	"achobeta-svc/internal/achobeta-svc-api/internal/repo/authz"
	"achobeta-svc/internal/achobeta-svc-api/internal/server/route"
	"achobeta-svc/internal/achobeta-svc-common/lib/tlog"
	"achobeta-svc/internal/achobeta-svc-common/pkg/constant"
	"achobeta-svc/internal/achobeta-svc-common/pkg/web"
	permissionv1 "achobeta-svc/internal/achobeta-svc-proto/gen/go/authz/permission/v1"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

var authService = authz.New()

func init() {
	route.GetRouter().RegisterMiddleware(route.LevelAnonymous, AddTraceId)
	route.GetRouter().RegisterMiddleware(route.LevelAnonymous, VerifyToken)
	route.GetRouter().RegisterMiddleware(route.LevelAnonymous, ErrorHandler)
}

func AddTraceId() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 假设 Trace ID 存在于 HTTP Header "X-Trace-ID" 中
		traceId := c.GetHeader("x-trace-id")
		if traceId == "" {
			traceId = uuid.New().String()
		}
		ctx := tlog.NewContext(c, zap.String("x-trace-id", traceId))
		c.Request = c.Request.WithContext(ctx)
		c.Keys = map[string]any{
			"traceId": traceId,
		}
		c.Next()
	}
}
func VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader(string(constant.RequestHeaderKeyToken))
		if token == "" {
			_ = c.AbortWithError(constant.TOKEN_IS_NULL.Code, fmt.Errorf(constant.TOKEN_IS_NULL.Msg))
			return
		}

		resp, err := authService.VerifyToken(c.Request.Context(), &permissionv1.VerifyTokenRequest{
			Token: token,
		})
		if err != nil {
			_ = c.AbortWithError(constant.TOKEN_IS_INVALID.Code, fmt.Errorf(constant.TOKEN_IS_INVALID.Msg))
			return
		}
		if !resp.Valid {
			_ = c.AbortWithError(constant.TOKEN_INSUFFICENT_PERMISSIONS.Code, fmt.Errorf(constant.TOKEN_INSUFFICENT_PERMISSIONS.Msg))
			return
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
				tlog.Infof("error: %+v", c.Keys)
				r.ErrorTrace(constant.COMMON_FAIL, err.Error(), c.Keys["traceId"].(string))
				return
			}
		}
	}
}

// func CheckToken() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		token := c.GetHeader(string(constant.RequestHeaderKeyToken))
// 		if token == "" {
// 			c.AbortWithError(constant.TOKEN_IS_NULL.Code, fmt.Errorf(constant.TOKEN_IS_NULL.Msg))
// 		}
// 		accountId, err := config.GetRedis().Get(token).Int()
// 		if err != nil {
// 			c.AbortWithError(constant.TOKEN_IS_INVALID.Code, fmt.Errorf(constant.TOKEN_IS_INVALID.Msg))
// 		}
// 		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), constant.RequestHeaderKeyAccountId, uint(accountId)))
// 		c.Next()
// 	}
// }
