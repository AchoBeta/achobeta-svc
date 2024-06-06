package authz

import (
	permissionv1 "achobeta-svc/internal/achobeta-svc-proto/gen/go/authz/permission/v1"
	"context"
)

func (z *impl) CreateAccount(ctx context.Context, req *permissionv1.CreateAccountRequest) (*permissionv1.CreateAccountResponse, error) {
	return z.pms.CreateAccount(ctx, req)
}
