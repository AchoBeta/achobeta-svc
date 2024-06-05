package service

// Services 是所有服务的集合
type Services struct {
	ReportService    *reportService
	DashboardService *dashboardService
}

func newServices(reportService *reportService, dashboardService *dashboardService) *Services {
	return &Services{
		ReportService:    reportService,
		DashboardService: dashboardService,
	}
}
