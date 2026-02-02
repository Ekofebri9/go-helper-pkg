package logger

import "log/slog"

func RequestID(id string) slog.Attr {
	return slog.String("request_id", id)
}

func UserID(id string) slog.Attr {
	return slog.String("user_id", id)
}

func ErrorAttr(err error) slog.Attr {
	return slog.Any("error", err)
}
