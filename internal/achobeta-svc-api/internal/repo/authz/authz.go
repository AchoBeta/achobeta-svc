package authz

import (
	permissionv1 "achobeta-svc/internal/achobeta-svc-proto/gen/go/authz/permission/v1"
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthzRepo interface {
	CreateAccount(ctx context.Context, req *permissionv1.CreateAccountRequest) (*permissionv1.CreateAccountResponse, error)
	VerifyToken(ctx context.Context, req *permissionv1.VerifyTokenRequest) (*permissionv1.VerifyTokenResponse, error)
}

type impl struct {
	pms permissionv1.AuthzServiceClient
}

// New creates a new impl repo
func New() AuthzRepo {
	conn, err := grpc.NewClient("achobeta-svc-authz:4396",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	return &impl{
		pms: permissionv1.NewAuthzServiceClient(conn),
	}
}
