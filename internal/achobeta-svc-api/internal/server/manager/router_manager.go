package manager

import (
	"achobeta-svc/internal/achobeta-svc-common/lib/tlog"
	"github.com/gin-gonic/gin"
)

type PathHandler func(h *gin.RouterGroup)
type Middleware func() gin.HandlerFunc
type RouteLevel int32
type RouteManager struct {
	Routes map[RouteLevel]*Route
}

type Route struct {
	Group       string
	Path        []PathHandler // 注册的path
	Middlewares []Middleware  // 注册的中间件
}

// 路由级别, 规定数字越大级别越高
// 在某种情况下, 高级别的路由会适用于低级别的路由
const (
	LevelAnonymous RouteLevel = 0 // 匿名级别路由
	LevelNormal    RouteLevel = 1 // 普通, 需要经过权限校验
	LevelAdmin     RouteLevel = 2 // admin
	LevelRoot      RouteLevel = 3 // 最高级别
)

var (
	RouteHandler = &RouteManager{
		Routes: make(map[RouteLevel]*Route),
	}
)

func (rm *RouteManager) Register(h *gin.Engine) {
	routeCount, middlewareCount := 0, 0
	for _, route := range rm.Routes {
		v := h.Group(route.Group)
		// 中间件注册
		for _, middleware := range route.Middlewares {
			middlewareCount++
			v.Use(middleware())
		}
		// 路由注册
		for _, router := range route.Path {
			routeCount++
			router(v)
		}
	}
	tlog.Infof("Registering routes, total routes: %d, total middlewares: %d", routeCount, middlewareCount)
}

func (rm *RouteManager) RegisterRouter(level RouteLevel, router PathHandler) {
	rm.Routes[level].Path = append(rm.Routes[level].Path, router)
}

// RegisterMiddleware
// @description 注册中间件
// @param level RouteLevel 路由级别
// @param middleware Middleware 中间件
// @param iteration bool 是否迭代, 即 v0 等级的中间件会同时注册到 v3, v2, v1 等级 (v3 级别最高, 范围最小)
func (rm *RouteManager) RegisterMiddleware(level RouteLevel, middleware Middleware, iteration bool) {
	if iteration {
		for e := level; e <= 3; e++ {
			rm.Routes[e].Middlewares = append(rm.Routes[e].Middlewares, middleware)
		}
		return
	}
	rm.Routes[level].Middlewares = append(rm.Routes[level].Middlewares, middleware)
}
