package authz

import (
	permissionv1 "achobeta-svc/internal/achobeta-svc-proto/gen/go/authz/permission/v1"
	"context"
)

func (z *impl) VerifyToken(ctx context.Context, req *permissionv1.VerifyTokenRequest) (*permissionv1.VerifyTokenResponse, error) {
	return z.pms.VerifyToken(ctx, req)
}

func (z *impl) Login(ctx context.Context, req *permissionv1.LoginRequest) (*permissionv1.LoginResponse, error) {
	return z.pms.Login(ctx, req)
}
