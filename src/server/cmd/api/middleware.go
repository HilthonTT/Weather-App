package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"math"
	"net"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/hilthontt/weather/internal/store"
)

// RateLimiterMiddleware throttles requests per client IP.
func (app *application) RateLimiterMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if app.config.rateLimiter.Enabled {
			if allow, retryAfter := app.rateLimiter.Allow(clientIP(r)); !allow {
				seconds := int(math.Ceil(retryAfter.Seconds()))
				app.rateLimitExceededResponse(w, r, strconv.Itoa(seconds))
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}

// clientIP returns the caller's IP without the port. Without this, every new
// TCP connection (new ephemeral port) would get a fresh rate-limit bucket.
// middleware.RealIP has already replaced RemoteAddr with X-Forwarded-For /
// X-Real-IP when present.
func clientIP(r *http.Request) string {
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}

	return host
}

// BasicAuthMiddleware protects operational endpoints (e.g. /debug/vars)
// with the AUTH_BASIC_USER / AUTH_BASIC_PASS credentials.
func (app *application) BasicAuthMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// read the auth header
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				app.unauthorizedBasicErrorResponse(w, r, fmt.Errorf("authorization header is missing"))
				return
			}

			// parse it -> get the base64
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Basic" {
				app.unauthorizedBasicErrorResponse(w, r, fmt.Errorf("authorization header is malformed"))
				return
			}

			// decode it
			decoded, err := base64.StdEncoding.DecodeString(parts[1])
			if err != nil {
				app.unauthorizedBasicErrorResponse(w, r, err)
				return
			}

			// check the credentials
			username := app.config.auth.basic.user
			pass := app.config.auth.basic.pass

			creds := strings.SplitN(string(decoded), ":", 2)
			if len(creds) != 2 || creds[0] != username || creds[1] != pass {
				app.unauthorizedBasicErrorResponse(w, r, fmt.Errorf("invalid credentials"))
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

// AuthTokenMiddleware validates the "Authorization: Bearer <jwt>" header and
// stores the authenticated *store.User in the request context.
func (app *application) AuthTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			app.unauthorizedErrorResponse(w, r, fmt.Errorf("authorization header is missing"))
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			app.unauthorizedErrorResponse(w, r, fmt.Errorf("authorization header is malformed"))
			return
		}

		token := parts[1]
		jwtToken, err := app.authenticator.ValidateToken(token)
		if err != nil {
			app.unauthorizedErrorResponse(w, r, err)
			return
		}

		claims, _ := jwtToken.Claims.(jwt.MapClaims)

		userID, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["sub"]), 10, 64)
		if err != nil {
			app.unauthorizedErrorResponse(w, r, err)
			return
		}

		ctx := r.Context()

		user, err := app.getUser(ctx, userID)

		if err != nil {
			app.unauthorizedErrorResponse(w, r, err)
			return
		}

		ctx = context.WithValue(ctx, userCtx, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// getUser loads a user, going through the Redis cache when it is enabled.
// Cache failures are logged and fall back to the database rather than
// failing the request.
func (app *application) getUser(ctx context.Context, userID int64) (*store.User, error) {
	if !app.config.redisCfg.enabled {
		return app.store.Users.GetByID(ctx, userID)
	}

	user, err := app.cacheStorage.Users.Get(ctx, userID)
	if err != nil {
		app.logger.Warnw("error reading user from cache", "userID", userID, "error", err)
	}

	if user != nil {
		return user, nil
	}

	user, err = app.store.Users.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	if err := app.cacheStorage.Users.Set(ctx, user); err != nil {
		app.logger.Warnw("error writing user to cache", "userID", userID, "error", err)
	}

	return user, nil
}
