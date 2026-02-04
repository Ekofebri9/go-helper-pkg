package grpc

import (
	"context"
	ctxmeta "go/helperpkg/ctx_meta"
	"go/helperpkg/middleware"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func RequestIDUnary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {

		id := ""

		if md, ok := metadata.FromIncomingContext(ctx); ok {
			if v := md.Get(middleware.RequestIDHeader); len(v) > 0 {
				id = v[0]
			}
		}

		if id == "" {
			id = uuid.NewString()
		}

		ctx = ctxmeta.WithRequestID(ctx, id)
		_ = grpc.SetHeader(ctx, metadata.Pairs(middleware.RequestIDHeader, id))

		return handler(ctx, req)
	}
}
