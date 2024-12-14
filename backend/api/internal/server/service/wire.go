//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.package service
//
//go:generate go run -mod=mod github.com/google/wire/cmd/wire
package service

import (
	authapi "achobeta-svc/backend/api/internal/api/authz"
	"achobeta-svc/backend/api/internal/api/health"
	authz2 "achobeta-svc/backend/api/internal/logic/authz"
	"achobeta-svc/backend/api/internal/repo/authz"

	"github.com/google/wire"
)

// 注册的服务都在这里初始化, 更新后执行make wire即可
// InitServices 初始化所有服务
func InitServices() *Apis {
	wire.Build(newApiService,
		// services
		authapi.NewAuthApi, health.NewHealthApi,
		// logic
		authz2.NewLogic,
		// repo
		authz.New)
	return &Apis{}
}
