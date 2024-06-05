package service

import (
	"achobeta-svc/internal/achobeta-svc-authz/internal/logic/service/permission"
	permissionv1 "achobeta-svc/internal/achobeta-svc-proto/gen/go/authz/permission/v1"

	"google.golang.org/grpc"
)

func New(s *grpc.Server) {
	permissionv1.RegisterAuthzServiceServer(s, &permission.AuthzServiceServer{})
}
