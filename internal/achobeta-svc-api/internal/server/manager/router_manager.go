package manager

import (
	"achobeta-svc/internal/achobeta-svc-common/lib/tlog"
	"fmt"

	"github.com/gin-gonic/gin"
)

type PathHandler func(h *gin.RouterGroup)
type Middleware func() gin.HandlerFunc
type RouteLevel int32
type RouteManager struct {
	Routes map[RouteLevel]*Route
}
type Route struct {
	Url         string
	Path        []PathHandler // 注册的path
	Middlewares []Middleware  // 注册的中间件
}

const (
	// 路由级别, 规定数字越大级别越高
	// 在某种情况下, 高级别的路由会适用于低级别的路由
	LEVEL_GLOBAL RouteLevel = 0 // 匿名级别路由
	LEVEL_V1     RouteLevel = 1
	LEVEL_V2     RouteLevel = 2
	LEVEL_V3     RouteLevel = 3
)

var (
	RouteHandler = &RouteManager{
		Routes: make(map[RouteLevel]*Route, 0),
	}
)

func urlLevel(level RouteLevel) string {
	return fmt.Sprintf("v%d", level)
}

func buildUrl(level RouteLevel) string {
	return fmt.Sprintf("/api/%s", urlLevel(level))
}

func NewRoute(level RouteLevel) *Route {
	return &Route{
		Url:         buildUrl(level),
		Path:        make([]PathHandler, 0),
		Middlewares: make([]Middleware, 0),
	}
}

func (rm *RouteManager) checkRoute(level RouteLevel) {
	if _, ok := rm.Routes[level]; !ok {
		rm.Routes[level] = NewRoute(level)
	}
}

func (rm *RouteManager) Register(h *gin.Engine) {
	routeCount, middlewareCount := 0, 0
	for _, route := range rm.Routes {
		v := h.Group(route.Url)
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
	rm.checkRoute(level)
	rm.Routes[level].Path = append(rm.Routes[level].Path, router)
}

// @title RegisterMiddleware
// @description 注册中间件
// @param level RouteLevel 路由级别
// @param middleware Middleware 中间件
// @param iteration bool 是否迭代, 即 v0 等级的中间件会同时注册到 v3, v2, v1 等级 (v3 级别最高, 范围最小)
func (rm *RouteManager) RegisterMiddleware(level RouteLevel, middleware Middleware, iteration bool) {
	rm.checkRoute(level)
	if iteration {
		for e := level; e <= 3; e++ {
			rm.checkRoute(e)
			rm.Routes[e].Middlewares = append(rm.Routes[e].Middlewares, middleware)
		}
		return
	}
	rm.Routes[level].Middlewares = append(rm.Routes[level].Middlewares, middleware)
}
