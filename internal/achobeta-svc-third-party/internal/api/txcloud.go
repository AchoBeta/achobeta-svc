package api

import (
	"achobeta-svc/internal/achobeta-svc-common/pkg/constant"
	"achobeta-svc/internal/achobeta-svc-common/pkg/tlog"
	"achobeta-svc/internal/achobeta-svc-common/pkg/utils"
	"achobeta-svc/internal/achobeta-svc-common/pkg/web"

	"achobeta-svc/internal/achobeta-svc-third-party/internal/service/txcloud"

	"github.com/gin-gonic/gin"
)

func init() {
	web.RouteHandler.RegisterRouter(web.LEVEL_GLOBAL, func(r *gin.RouterGroup) {
		r.GET("/txcloud", Ping)
		r.PUT("/txcloud/upload", Upload)
	})
}

func Ping(c *gin.Context) {
	r := web.NewResponse(c)
	tlog.CtxInfof(c.Request.Context(), "ctx test log zzz...")
	tlog.Infof("test log zzz...")
	r.Success(nil)
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
	if err = txcloud.PutObject(f, fileName); err != nil {
		tlog.CtxErrorf(ctx, "put object error: %v", err)
		r.ErrorMsg(constant.COMMON_FAIL, err.Error())
		return
	}
	r.Success(nil)
}
