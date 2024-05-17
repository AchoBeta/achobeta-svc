package api

import (
	"achobeta-svc/internal/achobeta-svc-common/pkg/constant"
	"achobeta-svc/internal/achobeta-svc-common/pkg/tlog"
	"achobeta-svc/internal/achobeta-svc-common/pkg/web"
	"achobeta-svc/internal/achobeta-svc-third-party/inernal/router/manager"
	"achobeta-svc/internal/achobeta-svc-third-party/inernal/service"
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/route"
)

func init() {
	manager.RouteHandler.RegisterRouter(manager.LEVEL_GLOBAL, func(r *route.RouterGroup) {
		r.GET("/txcloud", Ping)
		r.PUT("/txcloud/upload", Upload)
	})
}

func Ping(ctx context.Context, c *app.RequestContext) {
	r := web.NewResponse(c)
	r.Success("pong")
}

func Upload(ctx context.Context, c *app.RequestContext) {
	var err error
	r := web.NewResponse(c)
	file, err := c.FormFile("file")
	if err != nil {
		tlog.CtxErrorf(ctx, "get file error: %v", err)
		r.ErrorMsg(constant.COMMON_FAIL, err.Error())
		return
	}
	fileName := file.Filename
	f, err := file.Open()
	if err != nil {
		tlog.CtxErrorf(ctx, "open file error: %v", err)
		r.ErrorMsg(constant.COMMON_FAIL, err.Error())
		return
	}
	defer f.Close()
	tlog.Infof("upload file: %s", fileName)
	err = service.PutObject(f, fileName)
	if err != nil {
		tlog.CtxErrorf(ctx, "put object error: %v", err)
		r.ErrorMsg(constant.COMMON_FAIL, err.Error())
		return
	}
	r.Success(nil)
}
