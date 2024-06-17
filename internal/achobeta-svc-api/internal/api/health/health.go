package health

import (
	"achobeta-svc/internal/achobeta-svc-api/internal/server/manager"
	"github.com/gin-gonic/gin"
)

type Api struct {
}

func NewHealthApi() *Api {
	api := &Api{}
	RegisterRouter(api)
	return api
}
func RegisterRouter(api *Api) {
	manager.RouteHandler.RegisterRouter(manager.LevelAnonymous, func(h *gin.RouterGroup) {
		h.GET("/ping", api.Ping)
	})
	manager.RouteHandler.RegisterRouter(manager.LevelNormal, func(h *gin.RouterGroup) {
		h.GET("/normal/ping", api.Ping2)
	})
	manager.RouteHandler.RegisterRouter(manager.LevelAdmin, func(h *gin.RouterGroup) {
		h.GET("/normal/ping", api.Ping3)
	})
}

func (api *Api) Ping(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong anonymous"})
}

func (api *Api) Ping2(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong normal"})
}

func (api *Api) Ping3(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong admin"})
}
