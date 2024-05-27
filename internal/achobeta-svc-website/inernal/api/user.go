package api

import (
	"achobeta-svc/internal/achobeta-svc-common/pkg/tlog"
	"achobeta-svc/internal/achobeta-svc-common/pkg/utils"
	"achobeta-svc/internal/achobeta-svc-common/pkg/web"
	"achobeta-svc/internal/achobeta-svc-website/inernal/entity"
	"achobeta-svc/internal/achobeta-svc-website/inernal/service/user"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func init() {
	web.RouteHandler.RegisterRouter(web.LEVEL_GLOBAL, func(h *gin.RouterGroup) {
		h.POST("/user", Create)
	})
}

func Create(c *gin.Context) {
	r := web.NewResponse(c)
	uuid := uuid.New().String()
	email := fmt.Sprintf("%s@qq.com", uuid[len(uuid)-8:])
	password, err := utils.HashPassword("abc")
	if err != nil {
		tlog.CtxErrorf(c.Request.Context(), "xxx: %v", err)
		c.Error(err)
		return
	}
	ue := &entity.User{
		Username: uuid[:8],
		Password: string(password),
		Phone:    uuid[len(uuid)-11:],
		Email:    &email,
	}
	if err := user.CreateUser(c.Request.Context(), ue); err != nil {
		tlog.CtxErrorf(c.Request.Context(), "create user error: %v", err)
		c.Error(err)
		return
	}
	r.Success(nil)
}
