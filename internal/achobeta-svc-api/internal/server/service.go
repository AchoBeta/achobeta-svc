package server

import (
	auth_api "achobeta-svc/internal/achobeta-svc-api/internal/api/authz"
)

type Apis struct {
	authApi *auth_api.AuthzApi
}

func newServices(a *auth_api.AuthzApi) *Apis {
	return &Apis{
		authApi: a,
	}
}
