package grpc

import (
	"google.golang.org/grpc"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
)

func UnaryServer() grpc.ServerOption {
	return grpc.StatsHandler(
		otelgrpc.NewServerHandler(),
	)
}
