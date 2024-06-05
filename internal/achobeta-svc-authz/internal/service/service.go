package service

import "achobeta-svc/internal/achobeta-svc-authz/internal/service/permission"

// Services 是所有服务的集合

type Services struct {
	PermissionService *permission.PermissionServiceServer
}

func newServices(p *permission.PermissionServiceServer) *Services {
	return &Services{
		PermissionService: p,
	}
}
