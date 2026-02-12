package grpc

import (
	"context"
	jwtpkg "go/helperpkg/auth/jwt"
	ctxmeta "go/helperpkg/ctx_meta"

	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GRPC[T jwt.Claims](verifier *jwtpkg.Verifier) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {

		token := jwtpkg.ExtractFromGRPC(ctx)

		claims, err := jwtpkg.Verify[T](verifier, token)
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, "invalid token")
		}

		ctx = ctxmeta.WithClaims(ctx, claims)

		return handler(ctx, req)
	}
}
