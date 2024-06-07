package auth_api

import (
	"achobeta-svc/internal/achobeta-svc-api/internal/logic/auth"
	"achobeta-svc/internal/achobeta-svc-api/internal/server/manager"

	"github.com/gin-gonic/gin"
)

type AuthzApi struct {
	authLogic *auth.AuthzLogic
}

func NewAuthApi(al *auth.AuthzLogic) *AuthzApi {
	RegisterRouter()
	return &AuthzApi{
		authLogic: al,
	}
}

func RegisterRouter() {
	manager.RouteHandler.RegisterRouter(manager.LEVEL_GLOBAL, func(h *gin.RouterGroup) {
		h.GET("/ping", Ping)
	})
}

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
