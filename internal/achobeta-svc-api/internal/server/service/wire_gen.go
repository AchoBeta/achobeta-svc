// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package service

import (
	auth_api "achobeta-svc/internal/achobeta-svc-api/internal/api/authz"
	"achobeta-svc/internal/achobeta-svc-api/internal/logic/auth"
	"achobeta-svc/internal/achobeta-svc-api/internal/repo/authz"
)

// Injectors from wire.go:

// 注册的服务都在这里初始化, 更新后执行make wire即可
// InitServices 初始化所有服务
func InitServices() *Apis {
	authzRepo := authz.New()
	authzLogic := auth.NewLogic(authzRepo)
	authzApi := auth_api.NewAuthApi(authzLogic)
	apis := newApiService(authzApi)
	return apis
}
