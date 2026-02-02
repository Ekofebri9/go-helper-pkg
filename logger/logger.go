package logger

import (
	"context"
	"log/slog"
	"os"
)

var base *slog.Logger

type Config struct {
	Level string // debug | info | warn | error
}

func Init(cfg Config) {
	level := slog.LevelInfo

	switch cfg.Level {
	case "debug":
		level = slog.LevelDebug
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	}

	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	})

	base = slog.New(handler)
}

// Base returns base logger (optional, tapi berguna)
func Base() *slog.Logger {
	if base != nil {
		return base
	}
	return slog.Default()
}

func Debug(ctx context.Context, msg string, attrs ...slog.Attr) {
	fromContext(ctx).Debug(msg, attrsToAny(attrs)...)
}

func Info(ctx context.Context, msg string, attrs ...slog.Attr) {
	fromContext(ctx).Info(msg, attrsToAny(attrs)...)
}

func Warn(ctx context.Context, msg string, attrs ...slog.Attr) {
	fromContext(ctx).Warn(msg, attrsToAny(attrs)...)
}

func Error(ctx context.Context, msg string, attrs ...slog.Attr) {
	fromContext(ctx).Error(msg, attrsToAny(attrs)...)
}
