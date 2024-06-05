//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.package service
package service

import (
	"achobeta-svc/internal/achobeta-svc-authz/internal/logic/account"
	"achobeta-svc/internal/achobeta-svc-authz/internal/logic/user"
	"achobeta-svc/internal/achobeta-svc-authz/internal/repo/cache"
	"achobeta-svc/internal/achobeta-svc-authz/internal/repo/casbin"
	"achobeta-svc/internal/achobeta-svc-authz/internal/repo/database"
	"achobeta-svc/internal/achobeta-svc-authz/internal/service/permission"

	"github.com/google/wire"
)

// InitServices 初始化所有服务
func InitServices() *Services {
	wire.Build(newServices,
		// services
		permission.NewPermissionService,
		// logic
		account.New, user.New,
		// repo
		database.New, cache.New, casbin.New)
	return &Services{}
}
