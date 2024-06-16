package authz

import (
	"achobeta-svc/internal/achobeta-svc-api/internal/entity"
	"achobeta-svc/internal/achobeta-svc-api/internal/repo/authz"
	"achobeta-svc/internal/achobeta-svc-common/lib/tlog"
	permissionv1 "achobeta-svc/internal/achobeta-svc-proto/gen/go/authz/permission/v1"
	"context"
)

type Logic struct {
	az authz.Repo
}

func NewLogic(z authz.Repo) *Logic {
	return &Logic{
		az: z,
	}
}

func (al *Logic) CreateAccount(ctx context.Context, req *entity.CreateAccountParams) (uint64, error) {
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

func (al *Logic) Login(ctx context.Context, req *entity.LoginAccountParams) (string, error) {
	tlog.CtxInfof(ctx, "%+v", req)
	resp, err := al.az.Login(ctx, &permissionv1.LoginRequest{
		Username:  req.Username,
		Password:  req.Password,
		Email:     req.Email,
		Phone:     req.Phone,
		LoginType: req.LoginType,
	})
	if err != nil {
		tlog.CtxErrorf(ctx, "Login err: %v", err)
		return "", err
	}
	return resp.Token, nil
}
