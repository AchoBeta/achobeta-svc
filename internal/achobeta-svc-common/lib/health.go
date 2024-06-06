package server

import (
	"context"

	"google.golang.org/grpc/codes"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
)

type Checker = func(ctx context.Context, in *healthpb.HealthCheckRequest) (healthpb.HealthCheckResponse_ServingStatus, error)

type healthService struct {
	check Checker
}

func (s *healthService) Check(ctx context.Context, in *healthpb.HealthCheckRequest) (*healthpb.HealthCheckResponse, error) {
	code := healthpb.HealthCheckResponse_SERVING
	var err error
	if s.check != nil {
		code, err = s.check(ctx, in)
	}
	// yeah, right, open 24x7, like 7-11
	return &healthpb.HealthCheckResponse{Status: code}, err
}

func (s *healthService) Watch(in *healthpb.HealthCheckRequest, srv healthpb.Health_WatchServer) error {
	return status.Error(codes.Unimplemented, "Watch is not implemented")
}

func NewHealthService(check Checker) healthpb.HealthServer {
	return &healthService{
		check: check,
	}
}
