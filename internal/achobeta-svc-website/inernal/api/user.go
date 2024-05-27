package api

import (
	"achobeta-svc/internal/achobeta-svc-common/pkg/web"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	r := web.NewResponse(c)

	r.Success(nil)
}
