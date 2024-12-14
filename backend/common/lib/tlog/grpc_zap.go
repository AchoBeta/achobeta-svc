package tlog

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		var fields []zap.Field
		if md, ok := metadata.FromIncomingContext(ctx); ok {
			ids := md.Get("x-trace-id")
			if len(ids) > 0 {
				fields = append(fields, zap.String("x-trace-id", ids[0]))
			}
		}
		newCtx := NewContext(ctx, fields...)
		resp, err = handler(newCtx, req)
		if err != nil {
			newCtx = NewContext(newCtx, zap.Error(err))
		}
		return
	}
}
