package auth_api

import "achobeta-svc/internal/achobeta-svc-api/internal/logic/auth"

type AuthzApi struct {
	// UnimplementedAuthzServiceServer这个结构体是必须要内嵌进来的
	// 嵌入之后，我们就已经实现了GRPC这个服务的接口，但是实现之后我们什么都没做，没有写自己的业务逻辑，
	// 我们要重写实现的这个接口里的函数，这样才能提供一个真正的rpc的能力。
	// pms 是logic 层的部分, 用于处理业务逻辑
	authLogic *auth.AuthzLogic
}

func NewAuthApi(al *auth.AuthzLogic) *AuthzApi {
	return &AuthzApi{
		authLogic: al,
	}
}
