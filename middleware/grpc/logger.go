package grpc

import (
	"context"
	"log/slog"

	"google.golang.org/grpc"

	ctxmeta "go/helperpkg/ctx_meta"
	"go/helperpkg/logger"
)

func LoggerUnary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {

		reqID := ctxmeta.RequestID(ctx)

		l := logger.Base().With(
			slog.String("request_id", reqID),
			slog.String("grpc_method", info.FullMethod),
		)

		ctx = logger.WithContext(ctx, l)
		return handler(ctx, req)
	}
}
