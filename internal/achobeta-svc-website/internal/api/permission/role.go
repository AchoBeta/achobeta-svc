package permission

import (
	"achobeta-svc/internal/achobeta-svc-common/pkg/tlog"
	"achobeta-svc/internal/achobeta-svc-common/pkg/web"
	"achobeta-svc/internal/achobeta-svc-website/internal/entity"

	"github.com/gin-gonic/gin"
)

func init() {
	web.RouteHandler.RegisterRouter(web.LEVEL_GLOBAL, func(h *gin.RouterGroup) {
		h.POST("/permission/addPolicy", AddPolicy)
	})
}

func AddPolicy(c *gin.Context) {
	r := web.NewResponse(c)
	req := &entity.AddPolicyRequest{}
	if err := c.BindJSON(req); err != nil {
		tlog.CtxErrorf(c.Request.Context(), "bind json error: %v", err)
		c.Error(err)
		return
	}
	r.Success(nil)
}
