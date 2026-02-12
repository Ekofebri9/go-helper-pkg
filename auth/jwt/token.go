package jwt

import (
	"context"
	"go/helperpkg/middleware"
	"net/http"
	"strings"

	"google.golang.org/grpc/metadata"
)

func ExtractFromGRPC(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}

	values := md.Get(middleware.AuthHeader)
	if len(values) == 0 {
		return ""
	}

	return extract(values[0])
}

func ExtractFromHTTP(r *http.Request) string {
	authHeader := r.Header.Get(middleware.AuthHeader)
	if authHeader == "" {
		return ""
	}
	return extract(authHeader)
}

func extract(authHeader string) string {
	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 {
		return ""
	}

	if !strings.EqualFold(parts[0], "Bearer") {
		return ""
	}

	return strings.TrimSpace(parts[1])
}
