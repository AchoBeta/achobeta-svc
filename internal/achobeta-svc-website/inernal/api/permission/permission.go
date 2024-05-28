package permission

import (
	"achobeta-svc/internal/achobeta-svc-common/pkg/constant"
	"achobeta-svc/internal/achobeta-svc-common/pkg/tlog"
	"achobeta-svc/internal/achobeta-svc-common/pkg/web"
	"achobeta-svc/internal/achobeta-svc-website/inernal/entity"
	"achobeta-svc/internal/achobeta-svc-website/inernal/service/login"

	"github.com/gin-gonic/gin"
)

func init() {
	web.RouteHandler.RegisterRouter(web.LEVEL_GLOBAL, func(h *gin.RouterGroup) {
		h.POST("/permission/login", Login)
	})
}

func Login(c *gin.Context) {
	r := web.NewResponse(c)
	req := &entity.LoginRequest{}
	if err := c.BindJSON(req); err != nil {
		tlog.CtxErrorf(c.Request.Context(), "bind json error: %v", err)
		c.Error(err)
		return
	}
	if !checkParams(req) {
		r.ErrorCode(constant.USER_CREDENTIALS_ERROR)
		return
	}
	// 登录
	token, err := login.Login(c.Request.Context(), req)
	if err != nil {
		r.ErrorMsg(constant.USER_NOT_LOGIN, err.Error())
		return
	}
	// todo 缓存token
	r.Success(token)
}

func checkParams(req *entity.LoginRequest) bool {
	if req.Username == "" && req.Email == "" && req.Phone == "" {
		return false
	}
	if req.Username == "" && req.Password == "" {
		return false
	}
	return true
}