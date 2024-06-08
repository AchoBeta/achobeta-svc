package auth

import (
	"achobeta-svc/internal/achobeta-svc-api/internal/entity"
	"achobeta-svc/internal/achobeta-svc-api/internal/repo/authz"
	"achobeta-svc/internal/achobeta-svc-common/lib/tlog"
	permissionv1 "achobeta-svc/internal/achobeta-svc-proto/gen/go/authz/permission/v1"
	"context"
)

type AuthzLogic struct {
	az authz.AuthzRepo
}

func NewLogic(z authz.AuthzRepo) *AuthzLogic {
	return &AuthzLogic{
		az: z,
	}
}

func (al *AuthzLogic) CreateAccount(ctx context.Context, req *entity.CreateAccountParams) (uint64, error) {
	resp, err := al.az.CreateAccount(ctx, &permissionv1.CreateAccountRequest{
		Username: req.Username,
		Password: req.Password,
		Phone:    req.Phone,
		Email:    req.Email,
	})
	if err != nil {
		tlog.CtxErrorf(ctx, "CreateAccount err: %v", err)
		return 0, err
	}
	return resp.Id, nil
}
