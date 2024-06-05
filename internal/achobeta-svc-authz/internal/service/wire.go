package service

import (
	"achobeta-svc/internal/achobeta-svc-authz/internal/repo/cache"
	"achobeta-svc/internal/achobeta-svc-authz/internal/repo/database"

	"github.com/google/wire"
)

// InitServices 初始化所有服务
func InitServices() *Services {
	wire.Build(newServices,
		// services
		// newDashboardService, newReportService,
		// logic
		// report.New, dashboard.New,
		// repo
		cache.New, database.New)
	return &Services{}
}
