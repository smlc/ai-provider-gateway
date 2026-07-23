package logging

import (
	"context"
	"log/slog"
)

type contextKey struct{}

// WithLogger stores l in ctx and returns the derived context.
func WithLogger(ctx context.Context, l *slog.Logger) context.Context {
	return context.WithValue(ctx, contextKey{}, l)
}

// FromContext returns the logger stored in ctx by WithLogger.
// If none is present it returns slog.Default() as a safe fallback.
func FromContext(ctx context.Context) *slog.Logger {
	if l, ok := ctx.Value(contextKey{}).(*slog.Logger); ok && l != nil {
		return l
	}
	return slog.Default()
}

// FromContextOr returns the logger stored in ctx by WithLogger.
// If none is present it returns fallback instead.
func FromContextOr(ctx context.Context, fallback *slog.Logger) *slog.Logger {
	if l, ok := ctx.Value(contextKey{}).(*slog.Logger); ok && l != nil {
		return l
	}
	return fallback
}
