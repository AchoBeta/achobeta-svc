//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.package service
package server

import (
	auth_api "achobeta-svc/internal/achobeta-svc-api/internal/api/authz"
	"achobeta-svc/internal/achobeta-svc-api/internal/logic/auth"
	"achobeta-svc/internal/achobeta-svc-api/internal/repo/authz"

	"github.com/google/wire"
)

// 注册的服务都在这里初始化, 更新后执行make wire即可
// InitServices 初始化所有服务
func InitServices() *Apis {
	wire.Build(newServices,
		// services
		auth_api.NewAuthApi,
		// logic
		auth.NewLogic,
		// repo
		authz.New)
	return &Apis{}
}
