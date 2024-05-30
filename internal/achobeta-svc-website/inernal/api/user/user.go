package user

import (
	"achobeta-svc/internal/achobeta-svc-common/pkg/constant"
	"achobeta-svc/internal/achobeta-svc-common/pkg/tlog"
	"achobeta-svc/internal/achobeta-svc-common/pkg/utils"
	"achobeta-svc/internal/achobeta-svc-common/pkg/web"
	"achobeta-svc/internal/achobeta-svc-website/inernal/entity"
	"achobeta-svc/internal/achobeta-svc-website/inernal/service/account"
	"achobeta-svc/internal/achobeta-svc-website/inernal/service/user"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func init() {
	web.RouteHandler.RegisterRouter(web.LEVEL_GLOBAL, func(h *gin.RouterGroup) {
		h.GET("/user", Query)
	})
	web.RouteHandler.RegisterRouter(web.LEVEL_V1, func(h *gin.RouterGroup) {
		h.GET("/user/self", QuerySelf)
		h.PUT("/user", Modify)
	})
}

func Query(c *gin.Context) {
	r := web.NewResponse(c)
	pid := c.DefaultQuery("id", "0")
	qid, err := utils.ConverStr2Uint(pid)
	if err != nil {
		tlog.CtxErrorf(c.Request.Context(), "convert %s to uint id error: %v", pid, err)
		c.Error(err)
		return
	}
	ue := &entity.User{
		Model: gorm.Model{
			ID: qid,
		},
	}
	if err := user.Query(c.Request.Context(), ue); err != nil {
		tlog.CtxErrorf(c.Request.Context(), "query user error: %v", err)
		c.Error(err)
		return
	}

	r.Success(ue)
}

func QuerySelf(c *gin.Context) {
	r := web.NewResponse(c)
	var err error
	id, ok := c.Request.Context().Value((constant.RequestHeaderKeyAccountId)).(uint)
	if !ok {
		c.Error(fmt.Errorf("failed to retrieve account id from context"))
		return
	}
	acct, err := account.QueryAccount(c.Request.Context(), &entity.Account{
		Model: gorm.Model{
			ID: id,
		},
	})
	if err != nil {
		tlog.CtxErrorf(c.Request.Context(), "query account error: %v", err)
		c.Error(err)
		return
	}
	ue := &entity.User{
		Model: gorm.Model{
			ID: acct.UserId,
		},
	}
	err = user.Query(c.Request.Context(), ue)
	if err != nil {
		tlog.CtxErrorf(c.Request.Context(), "query user error: %v", err)
		c.Error(err)
		return
	}

	r.Success(&entity.UserInfoEntity{
		Id:       acct.ID,
		Username: acct.Username,
		Email:    acct.Email,
		Phone:    acct.Phone,
		User:     ue,
	})
}

func Modify(c *gin.Context) {
	r := web.NewResponse(c)
	ue := &entity.User{}
	if err := c.ShouldBindJSON(ue); err != nil {
		tlog.CtxErrorf(c.Request.Context(), "decode account error: %v", err)
		c.Error(err)
		return
	}
	if err := user.Modify(c.Request.Context(), ue); err != nil {
		tlog.CtxErrorf(c.Request.Context(), err.Error())
		c.Error(err)
		return
	}
	r.Success(ue.ID)
}
