package route

import (
	"achobeta-svc/internal/achobeta-svc-common/lib/tlog"
	"fmt"
	"github.com/gin-gonic/gin"
)

type PathHandler func(h *gin.RouterGroup)
type Middleware func() gin.HandlerFunc
type Level int32
type Router struct {
	Routes map[Level]*Route
}

type Route struct {
	Group       string
	Path        []PathHandler // 注册的path
	Middlewares []Middleware  // 注册的中间件
}

// 路由级别, 规定数字越大级别越高
// 在某种情况下, 高级别的路由会适用于低级别的路由
const (
	LevelAnonymous Level = 0 // 匿名级别路由
	LevelNormal    Level = 1 // 普通, 需要经过权限校验
	LevelAdmin     Level = 2 // admin
	LevelRoot      Level = 3 // 最高级别
)

var (
	router = &Router{
		Routes: make(map[Level]*Route),
	}
)

func newRoute(level Level) *Route {
	return &Route{
		Group:       fmt.Sprintf("/api/v%d", level),
		Path:        make([]PathHandler, 0),
		Middlewares: make([]Middleware, 0),
	}
}

func (rm *Router) checkRoute(level Level) {
	if _, ok := rm.Routes[level]; !ok {
		rm.Routes[level] = newRoute(level)
	}
}

func GetRouter() *Router {
	return router
}

func Injection(h *gin.Engine) {
	router.register(h)
}

func (rm *Router) register(h *gin.Engine) {
	routeCount, middlewareCount := 0, 0
	for _, route := range rm.Routes {
		v := h.Group(route.Group)
		// 中间件注册
		for _, middleware := range route.Middlewares {
			middlewareCount++
			h.Use(middleware())
		}
		// 路由注册
		for _, r := range route.Path {
			routeCount++
			r(v)
		}
	}
	tlog.Infof("Registering routes, total routes: %d, total middlewares: %d", routeCount, middlewareCount)
}

func (rm *Router) RegisterRouter(level Level, router PathHandler) {
	rm.checkRoute(level)
	rm.Routes[level].Path = append(rm.Routes[level].Path, router)
}

// RegisterMiddleware
// @description 注册中间件
// @param level 路由级别
// @param middleware 中间件
func (rm *Router) RegisterMiddleware(level Level, middleware ...Middleware) {
	rm.checkRoute(level)

	for _, m := range middleware {
		rm.Routes[level].Middlewares = append(rm.Routes[level].Middlewares, m)
	}
}
