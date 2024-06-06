package common

import (
	"achobeta-svc/internal/achobeta-svc-common/lib/tlog"
	errorsv1 "achobeta-svc/internal/achobeta-svc-proto/gen/go/common/errors/v1"
	"context"
	"runtime/debug"

	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func ErrorInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		resp, err = handler(ctx, req)
		if err == nil {
			return resp, nil
		}
		tlog.Error(ctx, "err_info", zap.Error(err))
		err = errors.Cause(err)
		if st := warpCommonError(ctx, err); st != nil {
			return resp, st.Err()
		}
		if st := wrapGORMError(err); st != nil {
			return resp, st.Err()
		}
		return resp, err
	}
}

func warpCommonError(ctx context.Context, err error) *status.Status {
	var e *CommonError
	ok := errors.As(err, &e)
	if !ok {
		return nil
	}
	st, err1 := status.New(codes.InvalidArgument, e.errMsg).WithDetails(&errorsv1.CommonError{Code: e.errCode, Message: e.errMsg})
	if err1 != nil {
		return nil
	}
	// 将 st.Err() 当做 error 返回
	return st
}

func wrapGORMError(err error) *status.Status {
	switch {
	case errors.Is(err, gorm.ErrDuplicatedKey):
		return status.New(codes.AlreadyExists, err.Error())
	default:
		return nil
	}
}

func RecoveryInterceptor() grpc_recovery.Option {
	return grpc_recovery.WithRecoveryHandlerContext(func(ctx context.Context, p interface{}) (err error) {
		tlog.Error(ctx, "server panic", zap.Any("panic", p), zap.String("stack", string(debug.Stack())))
		return status.Errorf(codes.Internal, "panic: %v", p)
	})
}
