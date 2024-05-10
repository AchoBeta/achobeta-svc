package router

import (
	"achobeta-svc/internal/achobeta-svc-common/pkg/tlog"
	"achobeta-svc/internal/achobeta-svc-third-party/config"
	"achobeta-svc/internal/achobeta-svc-third-party/inernal/router/manager"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func RunServer() {
	h, err := listen()
	if err != nil {
		tlog.Errorf("Listen error: %v", err)
		panic(err.Error())
	}
	h.Spin()
}

func listen() (*server.Hertz, error) {
	c := config.Get()
	h := server.Default(server.WithHostPorts(fmt.Sprintf("%s:%d", c.Host, c.Port)))
	manager.RouteHandler.Register(h)
	return h, nil
}
