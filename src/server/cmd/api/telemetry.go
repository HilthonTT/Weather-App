package main

import (
	"net/http"
	"strings"

	"go.opentelemetry.io/otel/trace"
)

type TelemetryMiddleware struct {
	tracer trace.Tracer
}

func NewTelemetryMiddleware(tracer trace.Tracer) *TelemetryMiddleware {
	return &TelemetryMiddleware{tracer: tracer}
}

func (tm *TelemetryMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var spanName string
		path := r.URL.Path

		// Define span name based on the path
		switch {
		case strings.Contains(path, "/users/"):
			spanName = "user-endpoints"
		case strings.Contains(path, "/weather/"):
			spanName = "weather-endpoints"
		case strings.Contains(path, "/auth/"):
			spanName = "auth-endpoints"
		default:
			spanName = "unknown-endpoints"
		}

		// Start a new span with context
		ctx, span := tm.tracer.Start(r.Context(), spanName, trace.WithAttributes())
		defer span.End()

		// Replace the request context with the new context that contains the span
		r = r.WithContext(ctx)

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}
