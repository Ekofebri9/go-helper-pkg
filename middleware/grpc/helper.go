package grpc

import (
	"context"

	"google.golang.org/grpc"
)

func ChainUnary(interceptors ...grpc.UnaryServerInterceptor) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {

		chained := handler
		for i := len(interceptors) - 1; i >= 0; i-- {
			current := interceptors[i]
			next := chained

			chained = func(ctx context.Context, req any) (any, error) {
				return current(ctx, req, info, next)
			}
		}

		return chained(ctx, req)
	}
}
