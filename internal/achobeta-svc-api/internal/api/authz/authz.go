package authz

import (
	"achobeta-svc/internal/achobeta-svc-api/internal/entity"
	"achobeta-svc/internal/achobeta-svc-api/internal/logic/authz"
	"achobeta-svc/internal/achobeta-svc-api/internal/server/manager"
	"achobeta-svc/internal/achobeta-svc-common/lib/tlog"
	"achobeta-svc/internal/achobeta-svc-common/pkg/web"

	"github.com/gin-gonic/gin"
)

type Api struct {
	authLogic *authz.Logic
}

func NewAuthApi(al *authz.Logic) *Api {
	api := &Api{
		authLogic: al,
	}
	RegisterRouter(api)
	return api
}

func RegisterRouter(api *Api) {
	manager.RouteHandler.RegisterRouter(manager.LevelAnonymous, func(h *gin.RouterGroup) {
		h.POST("/create", api.CreateAccount)
		h.POST("/login", api.Login)
	})
}

func (api *Api) CreateAccount(c *gin.Context) {
	r := web.NewResponse(c)
	cap := &entity.CreateAccountParams{}
	if err := c.ShouldBindJSON(cap); err != nil {
		tlog.CtxErrorf(c.Request.Context(), "bind json error: %v", c.Error(err))
		return
	}
	id, err := api.authLogic.CreateAccount(c.Request.Context(), cap)
	if err != nil {
		tlog.CtxErrorf(c.Request.Context(), "create account error: %v", c.Error(err))
		return
	}
	r.Success(id)
}

func (api *Api) Login(c *gin.Context) {
	r := web.NewResponse(c)
	lap := &entity.LoginAccountParams{}
	if err := c.ShouldBindJSON(lap); err != nil {
		tlog.CtxErrorf(c.Request.Context(), "bind json error: %v", c.Error(err))
		return
	}
	token, err := api.authLogic.Login(c.Request.Context(), lap)
	if err != nil {
		tlog.CtxErrorf(c.Request.Context(), "login error: %v", c.Error(err))
		return
	}
	r.Success(token)
}
