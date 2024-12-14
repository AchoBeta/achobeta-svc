package server

import "google.golang.org/grpc"

type Option struct {
	apply func(*server)
}

func WithUnaryServerInterceptor(interceptors ...grpc.UnaryServerInterceptor) Option {
	return Option{
		apply: func(s *server) {
			s.unaryServerInterceptor = append(s.unaryServerInterceptor, interceptors...)
		},
	}
}
