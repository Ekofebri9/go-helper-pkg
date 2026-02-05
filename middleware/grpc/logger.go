package grpc

import (
	"context"
	"log/slog"

	"go.opentelemetry.io/otel/trace"
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

		span := trace.SpanFromContext(ctx)
		if span.SpanContext().IsValid() {
			traceID := span.SpanContext().TraceID().String()
			l = l.With(slog.String("trace_id", traceID))
		}

		ctx = logger.WithContext(ctx, l)
		return handler(ctx, req)
	}
}
