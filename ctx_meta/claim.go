package ctxmeta

import "context"

func ClaimsFrom[T any](ctx context.Context) (*T, bool) {
	v, ok := ctx.Value(claimsKey{}).(*T)
	return v, ok
}

func WithClaims[T any](ctx context.Context, claims *T) context.Context {
	return context.WithValue(ctx, claimsKey{}, claims)
}
