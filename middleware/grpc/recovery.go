package grpc

import (
	"context"
	"fmt"
	"runtime/debug"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go/helperpkg/logger"
)

func RecoveryUnary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp any, err error) {

		defer func() {
			if r := recover(); r != nil {
				logger.Error(
					ctx,
					"grpc panic recovered",
					logger.ErrorAttr(anyToError(r)),
				)

				err = status.Error(codes.Internal, "internal server error")
			}
		}()

		return handler(ctx, req)
	}
}

func anyToError(v any) error {
	if err, ok := v.(error); ok {
		return err
	}
	return fmt.Errorf("%v\n%s", v, debug.Stack())
}
