package user

import (
	"achobeta-svc/internal/achobeta-svc-common/pkg/tlog"
	"achobeta-svc/internal/achobeta-svc-common/pkg/utils"
	"achobeta-svc/internal/achobeta-svc-common/pkg/web"
	"achobeta-svc/internal/achobeta-svc-website/inernal/entity"
	"achobeta-svc/internal/achobeta-svc-website/inernal/service/user"

	"github.com/gin-gonic/gin"
)

func init() {
	web.RouteHandler.RegisterRouter(web.LEVEL_GLOBAL, func(h *gin.RouterGroup) {
		h.POST("/user", Create)
		h.GET("/user", Query)
	})
}

func Create(c *gin.Context) {
	r := web.NewResponse(c)
	// ue := entity.MockUser()
	ue := &entity.User{}
	if err := c.ShouldBindJSON(ue); err != nil {
		tlog.CtxErrorf(c.Request.Context(), "decode user error: %v", err)
		c.Error(err)
		return
	}
	ue = &entity.User{
		Username: ue.Username,
		Password: hashPassword(ue.Password),
		Email:    ue.Email,
		Phone:    ue.Phone,
	}
	if err := user.CreateUser(c.Request.Context(), ue); err != nil {
		tlog.CtxErrorf(c.Request.Context(), "create user error: %v", err)
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
	u := &entity.User{}
	if err := c.ShouldBindJSON(u); err != nil {
		tlog.CtxErrorf(c.Request.Context(), "decode user error: %v", err)
		c.Error(err)
		return
	}
	ue, err := user.QueryUser(c.Request.Context(), u)
	if err != nil {
		tlog.CtxErrorf(c.Request.Context(), "query user error: %v", err)
		c.Error(err)
		return
	}
	r.Success(ue)
}
