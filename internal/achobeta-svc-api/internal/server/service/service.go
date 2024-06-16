package service

import (
	"achobeta-svc/internal/achobeta-svc-api/internal/api/authz"
)

type Apis struct {
	authApi *authz.Api
}

func newApiService(a *authz.Api) *Apis {
	return &Apis{
		authApi: a,
	}
}
