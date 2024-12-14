package validator

import (
	"context"

	"github.com/bufbuild/protovalidate-go"
	"github.com/bufbuild/protovalidate-go/legacy"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	v, err := protovalidate.New(
		legacy.WithLegacySupport(legacy.ModeMerge),
		protovalidate.WithFailFast(true),
	)
	if err != nil {
		panic(err)
	}
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		if msg, ok := req.(proto.Message); ok {
			if err := v.Validate(msg); err != nil {
				return nil, status.Error(codes.InvalidArgument, err.Error())
			}
		}
		return handler(ctx, req)
	}
}
