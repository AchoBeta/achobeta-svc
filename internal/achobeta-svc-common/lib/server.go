package server

import (
	"achobeta-svc/internal/achobeta-svc-common/lib/common"
	"achobeta-svc/internal/achobeta-svc-common/lib/interceptor/context"
	"achobeta-svc/internal/achobeta-svc-common/lib/interceptor/validator"
	"achobeta-svc/internal/achobeta-svc-common/lib/tlog"
	"log"
	"net"
	"strconv"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"google.golang.org/grpc"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

type Server interface {
	RegisterService(desc *grpc.ServiceDesc, impl any)
	Start()
	Stop()
}

type server struct {
	grpcServer             *grpc.Server
	lis                    net.Listener
	unaryServerInterceptor []grpc.UnaryServerInterceptor
}

func NewDefaultServer() Server {
	return NewServer(NewConfig())
}

func NewServer(config *Config, ops ...Option) Server {
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(config.port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := server{
		lis: lis,
		unaryServerInterceptor: []grpc.UnaryServerInterceptor{
			// 这里有bug, otelgrpc没有修复
			// otelgrpc.UnaryServerInterceptor(otelgrpc.WithInterceptorFilter(filters.None(filters.HealthCheck()))),
			tlog.UnaryServerInterceptor(),
			validator.UnaryServerInterceptor(),
			context.UnaryServerInterceptor(),
			common.ErrorInterceptor(),
		},
	}
	for _, op := range ops {
		op.apply(&server)
	}
	server.unaryServerInterceptor = append(server.unaryServerInterceptor, grpc_recovery.UnaryServerInterceptor(common.RecoveryInterceptor()))

	streamInterceptors := []grpc.StreamServerInterceptor{
		grpc_validator.StreamServerInterceptor(),
		grpc_recovery.StreamServerInterceptor(common.RecoveryInterceptor()),
	}

	// stats.Handler 无法过滤请求，先使用 WithInterceptorFilter https://github.com/open-telemetry/opentelemetry-go-contrib/issues/4575
	grpcServer := grpc.NewServer(
		// grpc.StatsHandler(otelgrpc.NewServerHandler()),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(server.unaryServerInterceptor...)),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(streamInterceptors...)),
	)
	reflection.Register(grpcServer) // for grpcurl
	healthpb.RegisterHealthServer(grpcServer, NewHealthService(nil))

	server.grpcServer = grpcServer
	return &server
}

func (s *server) RegisterService(desc *grpc.ServiceDesc, impl any) {
	s.grpcServer.RegisterService(desc, impl)
}

func (s *server) Start() {
	go func() {
		if err := s.grpcServer.Serve(s.lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
}

func (s *server) Stop() {
	s.grpcServer.GracefulStop()
	_ = s.lis.Close()
}
