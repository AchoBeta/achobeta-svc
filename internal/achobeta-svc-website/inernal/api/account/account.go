package account

import (
	"achobeta-svc/internal/achobeta-svc-common/pkg/tlog"
	"achobeta-svc/internal/achobeta-svc-common/pkg/utils"
	"achobeta-svc/internal/achobeta-svc-common/pkg/web"
	"achobeta-svc/internal/achobeta-svc-website/inernal/entity"
	"achobeta-svc/internal/achobeta-svc-website/inernal/service/account"
	"achobeta-svc/internal/achobeta-svc-website/inernal/service/user"

	"github.com/gin-gonic/gin"
)

func init() {
	web.RouteHandler.RegisterRouter(web.LEVEL_GLOBAL, func(h *gin.RouterGroup) {
		h.POST("/account", Create)
	})
	web.RouteHandler.RegisterRouter(web.LEVEL_V1, func(h *gin.RouterGroup) {
		h.GET("/account", Query)
	})
}

func Create(c *gin.Context) {
	r := web.NewResponse(c)
	acct := &entity.Account{}
	if err := c.ShouldBindJSON(acct); err != nil {
		tlog.CtxErrorf(c.Request.Context(), "decode account error: %v", err)
		c.Error(err)
		return
	}
	userId, err := user.Create(c.Request.Context(), entity.MockUser())
	if err != nil {
		tlog.CtxErrorf(c.Request.Context(), "create user error: %v", err)
		c.Error(err)
		return
	}
	acct.UserId = userId
	acct.Password = hashPassword(acct.Password)
	if err := account.CreateAccount(c.Request.Context(), acct); err != nil {
		tlog.CtxErrorf(c.Request.Context(), "create account error: %v", err)
		c.Error(err)
		return
	}
	r.Success(acct.ID)
}

func hashPassword(pwd string) string {
	hashedPwd, err := utils.HashPassword(pwd)
	if err != nil {
		tlog.Errorf("hash password error: %v", err)
		return pwd
	}
	return string(hashedPwd)
}

func Query(c *gin.Context) {
	r := web.NewResponse(c)
	u := &entity.Account{}
	if err := c.ShouldBindJSON(u); err != nil {
		tlog.CtxErrorf(c.Request.Context(), "decode account error: %v", err)
		c.Error(err)
		return
	}
	account, err := account.QueryAccount(c.Request.Context(), u)
	if err != nil {
		tlog.CtxErrorf(c.Request.Context(), "query account error: %v", err)
		c.Error(err)
		return
	}
	r.Success(account)
}
