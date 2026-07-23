package api

import (
	"log/slog"
	"net/http"
	"sse-multiplexer/internal/logging"
)

// RequestLogger is HTTP middleware that derives a request-scoped logger from
// base, enriches it with method, path, and remote address (plus the optional
// X-Request-Id header), and stores it in the request context.
func RequestLogger(base *slog.Logger) func(http.Handler) http.Handler {
	if base == nil {
		base = slog.Default()
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			attrs := []any{
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.String("remote_addr", r.RemoteAddr),
			}
			if reqID := r.Header.Get("X-Request-Id"); reqID != "" {
				attrs = append(attrs, slog.String("request_id", reqID))
			}
			reqLogger := base.With(attrs...)
			ctx := logging.WithLogger(r.Context(), reqLogger)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
