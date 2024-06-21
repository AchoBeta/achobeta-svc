package middleware

import (
	"achobeta-svc/internal/achobeta-svc-api/internal/repo/authz"
	"achobeta-svc/internal/achobeta-svc-api/internal/server/route"
	"achobeta-svc/internal/achobeta-svc-common/lib/tlog"
	"achobeta-svc/internal/achobeta-svc-common/pkg/constant"
	"achobeta-svc/internal/achobeta-svc-common/pkg/web"
	permissionv1 "achobeta-svc/internal/achobeta-svc-proto/gen/go/authz/permission/v1"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

var authService = authz.New()

func init() {
	// anonymous
	route.GetRouter().RegisterMiddleware(route.LevelAnonymous,
		// middlewares
		AddTraceId, ErrorHandler)

	// normal
	route.GetRouter().RegisterMiddleware(route.LevelNormal,
		// middlewares
		AddTraceId, ErrorHandler, VerifyTokenNormal)

	// admin
	route.GetRouter().RegisterMiddleware(route.LevelAdmin,
		// middlewares
		AddTraceId, ErrorHandler, verifyTokenAdmin)

	// root
	route.GetRouter().RegisterMiddleware(route.LevelRoot,
		// middlewares
		AddTraceId, ErrorHandler, verifyTokenRoot)
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

func VerifyTokenNormal() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := verifyToken(c, permissionv1.VerifyTokenRequest_ROLE_NORMAL); err != nil {
			_ = c.Error(status.Error(codes.PermissionDenied, constant.TOKEN_INSUFFICENT_PERMISSIONS.Msg))
			return
		}

		c.Next()
	}
}

func verifyTokenAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := verifyToken(c, permissionv1.VerifyTokenRequest_ROLE_ADMIN); err != nil {
			_ = c.Error(status.Error(codes.PermissionDenied, constant.TOKEN_INSUFFICENT_PERMISSIONS.Msg))
			return
		}

		c.Next()
	}
}

func verifyTokenRoot() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := verifyToken(c, permissionv1.VerifyTokenRequest_ROLE_ROOT); err != nil {
			_ = c.Error(status.Error(codes.PermissionDenied, constant.TOKEN_INSUFFICENT_PERMISSIONS.Msg))
			return
		}

		c.Next()
	}
}

func verifyToken(c *gin.Context, role permissionv1.VerifyTokenRequest_Role) error {
	ctx := c.Request.Context()
	token := c.GetHeader(string(constant.RequestHeaderKeyToken))
	if token == "" {
		return fmt.Errorf(constant.TOKEN_IS_NULL.Msg)
	}

	resp, err := authService.VerifyToken(ctx, &permissionv1.VerifyTokenRequest{
		Role:  role,
		Token: token,
	})

	if err != nil {
		return fmt.Errorf(constant.TOKEN_IS_INVALID.Msg)
	}
	if !resp.Valid {
		return fmt.Errorf(constant.TOKEN_INSUFFICENT_PERMISSIONS.Msg)
	}
	return nil
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
