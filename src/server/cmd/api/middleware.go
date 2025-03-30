package main

import (
	"net/http"

	"github.com/hilthontt/weather/internal/utils"
)

func (app *application) RateLimiterMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if app.config.RateLimiter.Enabled {
			if allow, retryAfter := app.rateLimiter.Allow(r.RemoteAddr); !allow {
				utils.RateLimitExceededResponse(w, r, retryAfter.String(), app.logger)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}
