package permission

import (
	"achobeta-svc/internal/achobeta-svc-authz/internal/entity"
	"achobeta-svc/internal/achobeta-svc-authz/internal/logic/account"
	"achobeta-svc/internal/achobeta-svc-authz/internal/logic/user"
	"achobeta-svc/internal/achobeta-svc-common/lib/tlog"
	permissionv1 "achobeta-svc/internal/achobeta-svc-proto/gen/go/authz/permission/v1"
	"context"
	"fmt"
)

type ServiceServer struct {
	// UnimplementedAuthzServiceServer这个结构体是必须要内嵌进来的
	// 嵌入之后，我们就已经实现了GRPC这个服务的接口，但是实现之后我们什么都没做，没有写自己的业务逻辑，
	// 我们要重写实现的这个接口里的函数，这样才能提供一个真正的rpc的能力。
	permissionv1.UnimplementedAuthzServiceServer
	// pms 是logic 层的部分, 用于处理业务逻辑
	pms *account.Permission
	ul  *user.Logic
}

func NewPermissionService(p *account.Permission, u *user.Logic) *ServiceServer {
	return &ServiceServer{
		pms: p,
		ul:  u,
	}
}

// CreateAccount 创建账号接口
func (p *ServiceServer) CreateAccount(ctx context.Context, req *permissionv1.CreateAccountRequest) (*permissionv1.CreateAccountResponse, error) {
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

func (p *ServiceServer) VerifyToken(ctx context.Context, req *permissionv1.VerifyTokenRequest) (*permissionv1.VerifyTokenResponse, error) {
	tlog.CtxInfof(ctx, "VerifyToken request: %s", req.GetToken())
	vaild, err := p.pms.CheckToken(ctx, req.Token)
	if err != nil {
		tlog.CtxErrorf(ctx, "CheckToken err: %v", err)
		return nil, err
	}
	return &permissionv1.VerifyTokenResponse{
		Valid: vaild,
	}, nil
}

func (p *ServiceServer) Login(ctx context.Context, req *permissionv1.LoginRequest) (*permissionv1.LoginResponse, error) {
	// 构建登录请求
	loginReq, err := buildLoginRequest(req)
	if err != nil {
		tlog.CtxErrorf(ctx, "buildLoginRequest err: %v", err)
		return nil, err
	}
	// 登录逻辑
	token, err := p.pms.Login(ctx, loginReq)
	if err != nil {
		tlog.CtxErrorf(ctx, "Login err: %v", err)
		return nil, err
	}
	return &permissionv1.LoginResponse{
		Token: token,
	}, nil
}

func buildLoginRequest(req *permissionv1.LoginRequest) (*entity.LoginRequest, error) {
	if checkLoginParams(req) {
		return nil, fmt.Errorf("login params error")
	}
	var loginKey string
	if req.LoginType == permissionv1.LoginType_LOGIN_TYPE_USERNAME {
		loginKey = req.GetUsername()
	} else if req.LoginType == permissionv1.LoginType_LOGIN_TYPE_PHONE {
		loginKey = req.GetPhone()
	} else if req.LoginType == permissionv1.LoginType_LOGIN_TYPE_EMAIL {
		loginKey = req.GetEmail()
	}
	return &entity.LoginRequest{
		LoginKey: loginKey,
		LoginPwd: req.GetPassword(),
		Type:     converLoginType(req.GetLoginType()),
	}, nil
}

func checkLoginParams(req *permissionv1.LoginRequest) bool {
	if req.Username == nil && req.Phone == nil && req.Email == nil {
		return false
	}
	if req.LoginType == permissionv1.LoginType_LOGIN_TYPE_USERNAME && req.Password == nil {
		return false
	}
	return true
}
func converLoginType(t permissionv1.LoginType) entity.LoginType {
	switch t {
	case permissionv1.LoginType_LOGIN_TYPE_USERNAME:
		return entity.LoginTypeUsername
	case permissionv1.LoginType_LOGIN_TYPE_PHONE:
		return entity.LoginTypePhone
	case permissionv1.LoginType_LOGIN_TYPE_EMAIL:
		return entity.LoginTypeEmail
	}
	return entity.LoginTypeUsername
}
