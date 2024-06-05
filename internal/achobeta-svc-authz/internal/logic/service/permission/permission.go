package permission

import (
	"achobeta-svc/internal/achobeta-svc-authz/internal/entity"
	"achobeta-svc/internal/achobeta-svc-authz/internal/logic/handler/account"
	permissionv1 "achobeta-svc/internal/achobeta-svc-proto/gen/go/authz/permission/v1"
	"context"
)

type AuthzServiceServer struct {
	// UnimplementedHelloServiceServer这个结构体是必须要内嵌进来的
	// 也就是说我们定义的这个结构体对象必须继承UnimplementedHelloServiceServer。
	// 嵌入之后，我们就已经实现了GRPC这个服务的接口，但是实现之后我们什么都没做，没有写自己的业务逻辑，
	// 我们要重写实现的这个接口里的函数，这样才能提供一个真正的rpc的能力。
	permissionv1.UnimplementedAuthzServiceServer
}

// CreateAccount 创建账号接口
func (p *AuthzServiceServer) CreateAccount(ctx context.Context, req *permissionv1.CreateAccountRequest) (*permissionv1.CreateAccountResponse, error) {
	ue := &entity.Account{
		Username: req.Username,
		Password: req.Password,
		Phone:    req.Phone,
		Email:    req.Email,
	}
	account.CreateAccount(ctx, ue)
	resp := &permissionv1.CreateAccountResponse{
		Id: uint64(ue.ID),
	}
	return resp, nil
}
