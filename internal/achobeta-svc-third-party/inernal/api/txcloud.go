package api

import (
	"achobeta-svc/internal/achobeta-svc-common/pkg/constant"
	"achobeta-svc/internal/achobeta-svc-common/pkg/tlog"
	"achobeta-svc/internal/achobeta-svc-common/pkg/utils"
	"achobeta-svc/internal/achobeta-svc-common/pkg/web"
	"achobeta-svc/internal/achobeta-svc-third-party/inernal/router/manager"
	"achobeta-svc/internal/achobeta-svc-third-party/inernal/service"

	"github.com/gin-gonic/gin"
)

func init() {
	manager.RouteHandler.RegisterRouter(manager.LEVEL_GLOBAL, func(r *gin.RouterGroup) {
		r.GET("/txcloud", Ping)
		r.PUT("/txcloud/upload", Upload)
	})
}

func Ping(c *gin.Context) {
	r := web.NewResponse(c)
	r.Success("pong")
}

func Upload(c *gin.Context) {
	var err error
	ctx := c.Request.Context()
	r := web.NewResponse(c)
	file, err := c.FormFile("file")
	if err != nil {
		tlog.CtxErrorf(ctx, "get file error: %v", err)
		r.ErrorMsg(constant.COMMON_FAIL, err.Error())
		return
	}
	fileName := utils.GetSnowflakeUUID()
	f, err := file.Open()
	if err != nil {
		tlog.CtxErrorf(ctx, "open file error: %v", err)
		r.ErrorMsg(constant.COMMON_FAIL, err.Error())
		return
	}
	defer f.Close()
	if err = service.PutObject(f, fileName); err != nil {
		tlog.CtxErrorf(ctx, "put object error: %v", err)
		r.ErrorMsg(constant.COMMON_FAIL, err.Error())
		return
	}
	r.Success(nil)
}
