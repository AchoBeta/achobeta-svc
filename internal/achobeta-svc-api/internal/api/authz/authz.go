package auth_api

import (
	"achobeta-svc/internal/achobeta-svc-api/internal/entity"
	"achobeta-svc/internal/achobeta-svc-api/internal/logic/auth"
	"achobeta-svc/internal/achobeta-svc-api/internal/server/manager"
	"achobeta-svc/internal/achobeta-svc-common/lib/tlog"
	"achobeta-svc/internal/achobeta-svc-common/pkg/web"

	"github.com/gin-gonic/gin"
)

type AuthzApi struct {
	authLogic *auth.AuthzLogic
}

func NewAuthApi(al *auth.AuthzLogic) *AuthzApi {
	api := &AuthzApi{
		authLogic: al,
	}
	RegisterRouter(api)
	return api
}

func RegisterRouter(api *AuthzApi) {
	manager.RouteHandler.RegisterRouter(manager.LEVEL_GLOBAL, func(h *gin.RouterGroup) {
		h.GET("/ping", api.Ping)

		h.POST("/create", api.CreateAccount)
	})
}

func (api *AuthzApi) CreateAccount(c *gin.Context) {
	r := web.NewResponse(c)
	cap := &entity.CreateAccountParams{}
	if err := c.ShouldBindJSON(cap); err != nil {
		tlog.CtxErrorf(c.Request.Context(), "bind json error: %v", err)
		c.Error(err)
		return
	}
	id, err := api.authLogic.CreateAccount(c.Request.Context(), cap)
	if err != nil {
		c.Error(err)
		return
	}
	r.Success(id)
}

func (api *AuthzApi) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong pong",
	})
}
