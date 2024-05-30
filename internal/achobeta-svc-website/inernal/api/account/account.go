package account

import (
	"achobeta-svc/internal/achobeta-svc-common/pkg/constant"
	"achobeta-svc/internal/achobeta-svc-common/pkg/tlog"
	"achobeta-svc/internal/achobeta-svc-common/pkg/utils"
	"achobeta-svc/internal/achobeta-svc-common/pkg/web"
	"achobeta-svc/internal/achobeta-svc-website/inernal/entity"
	"achobeta-svc/internal/achobeta-svc-website/inernal/service/account"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func init() {
	web.RouteHandler.RegisterRouter(web.LEVEL_GLOBAL, func(h *gin.RouterGroup) {
		h.POST("/account", Create)
	})
	web.RouteHandler.RegisterRouter(web.LEVEL_V1, func(h *gin.RouterGroup) {
		h.GET("/account", Query)
		h.GET("/account/self", QuerySelf)
	})
}

func QuerySelf(c *gin.Context) {
	r := web.NewResponse(c)
	id, ok := c.Request.Context().Value((constant.RequestHeaderKeyAccountId)).(uint)
	if !ok {
		c.Error(fmt.Errorf("failed to retrieve account id from context"))
		return
	}
	ue, err := account.QueryAccount(c.Request.Context(), &entity.Account{
		Model: gorm.Model{
			ID: id,
		},
	})
	if err != nil {
		tlog.CtxErrorf(c.Request.Context(), "query account error: %v", err)
		c.Error(err)
		return
	}
	r.Success(ue)
}

func Create(c *gin.Context) {
	r := web.NewResponse(c)
	ue := &entity.Account{}
	if err := c.ShouldBindJSON(ue); err != nil {
		tlog.CtxErrorf(c.Request.Context(), "decode account error: %v", err)
		c.Error(err)
		return
	}
	ue = &entity.Account{
		Username: ue.Username,
		Password: hashPassword(ue.Password),
		Email:    ue.Email,
		Phone:    ue.Phone,
	}
	if err := account.CreateAccount(c.Request.Context(), ue); err != nil {
		tlog.CtxErrorf(c.Request.Context(), "create account error: %v", err)
		c.Error(err)
		return
	}
	r.Success(nil)
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
	ue, err := account.QueryAccount(c.Request.Context(), u)
	if err != nil {
		tlog.CtxErrorf(c.Request.Context(), "query account error: %v", err)
		c.Error(err)
		return
	}
	r.Success(ue)
}
