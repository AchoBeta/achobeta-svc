package permission

import (
	"achobeta-svc/internal/achobeta-svc-authz/internal/entity"
	"achobeta-svc/internal/achobeta-svc-authz/internal/logic/account"
	"achobeta-svc/internal/achobeta-svc-authz/internal/logic/user"
	"achobeta-svc/internal/achobeta-svc-common/pkg/tlog"
	permissionv1 "achobeta-svc/internal/achobeta-svc-proto/gen/go/authz/permission/v1"
	"context"
)

type PermissionServiceServer struct {
	// UnimplementedAuthzServiceServer这个结构体是必须要内嵌进来的
	// 嵌入之后，我们就已经实现了GRPC这个服务的接口，但是实现之后我们什么都没做，没有写自己的业务逻辑，
	// 我们要重写实现的这个接口里的函数，这样才能提供一个真正的rpc的能力。
	permissionv1.UnimplementedAuthzServiceServer
	// pms 是logic 层的部分, 用于处理业务逻辑
	pms *account.Permission
	ul  *user.UserLogic
}

func NewPermissionService(p *account.Permission, u *user.UserLogic) *PermissionServiceServer {
	return &PermissionServiceServer{
		pms: p,
		ul:  u,
	}
}

// CreateAccount 创建账号接口
func (p *PermissionServiceServer) CreateAccount(ctx context.Context, req *permissionv1.CreateAccountRequest) (*permissionv1.CreateAccountResponse, error) {
	// 创建账号前, 先创建角色, 这里是mock一个虚拟的角色, 后续更新
	uid, err := p.ul.CreateUser(ctx, entity.MockUser())
	if err != nil {
		tlog.CtxErrorf(ctx, "CreateUser err: %v", err)
		return nil, err
	}
	ue := &entity.Account{
		Username: req.Username,
		UserId:   uid,
		Password: req.Password,
		Phone:    req.Phone,
		Email:    req.Email,
	}
	if err := p.pms.CreateAccount(ctx, ue); err != nil {
		tlog.CtxErrorf(ctx, "CreateAccount err: %v", err)
		return nil, err
	}
	resp := &permissionv1.CreateAccountResponse{
		Id: uint64(ue.ID),
	}
	return resp, nil
}