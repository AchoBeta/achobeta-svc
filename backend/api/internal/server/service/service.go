package service

import (
	"achobeta-svc/backend/api/internal/api/authz"
	"achobeta-svc/backend/api/internal/api/health"
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
