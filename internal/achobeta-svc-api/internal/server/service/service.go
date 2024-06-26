package service

import (
	"achobeta-svc/internal/achobeta-svc-api/internal/api/authz"
	"achobeta-svc/internal/achobeta-svc-api/internal/api/health"
)

type Apis struct {
	authApi   *authz.Api
	healthApi *health.Api
}

func newApiService(a *authz.Api, h *health.Api) *Apis {
	return &Apis{
		authApi:   a,
		healthApi: h,
	}
}
