package login

import (
	"achobeta-svc/internal/achobeta-svc-common/pkg/web"

	"github.com/gin-gonic/gin"
)

func init() {
	web.RouteHandler.RegisterRouter(web.LEVEL_GLOBAL, func(h *gin.RouterGroup) {
		h.POST("/login", Login)
	})
}

func Login(c *gin.Context) {
	r := web.NewResponse(c)
	r.Success(nil)
}
