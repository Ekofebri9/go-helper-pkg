package otel

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
)

type TracerConfig struct {
	ServiceName string
	Exporter    sdktrace.SpanExporter
}

func InitTracer(cfg TracerConfig) (*sdktrace.TracerProvider, error) {
	res, err := resource.New(
		context.Background(),
		resource.WithAttributes(
			semconv.ServiceName(cfg.ServiceName),
		),
	)
	if err != nil {
		return nil, err
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithResource(res),
		sdktrace.WithBatcher(cfg.Exporter),
	)

	otel.SetTracerProvider(tp)
	return tp, nil
}

func Shutdown(ctx context.Context, fns ...func(context.Context) error) {
	for _, fn := range fns {
		_ = fn(ctx)
	}
}

func Meter(name string) metric.Meter {
	return otel.Meter(name)
}
