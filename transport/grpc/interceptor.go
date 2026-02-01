package grpc

import (
	"context"

	"go/helperpkg/errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func UnaryErrorInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {

		resp, err := handler(ctx, req)
		if err == nil {
			return resp, nil
		}

		appErr, ok := errors.AsAppError(err)
		if !ok {
			return nil, status.Error(codes.Internal, "internal error")
		}

		return nil, Error(appErr)
	}
}
