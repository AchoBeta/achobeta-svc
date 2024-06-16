package service

import "achobeta-svc/internal/achobeta-svc-authz/internal/service/permission"

// Services 是所有服务的集合

type Services struct {
	PermissionService *permission.ServiceServer
}

func newServices(p *permission.ServiceServer) *Services {
	return &Services{
		PermissionService: p,
	}
}
