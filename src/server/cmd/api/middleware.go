package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/hilthontt/weather/internal/utils"

	"github.com/hilthontt/weather/services/users"
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
func (app *application) AuthTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.UnauthorizedErrorResponse(w, r, fmt.Errorf("authorization header is missing"), app.logger)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.UnauthorizedErrorResponse(w, r, fmt.Errorf("authorization header is malformed"), app.logger)
			return
		}

		token := parts[1]
		jwtToken, err := app.authenticator.ValidateToken(token)
		if err != nil {
			utils.UnauthorizedErrorResponse(w, r, err, app.logger)
			return
		}

		claims, _ := jwtToken.Claims.(jwt.MapClaims)

		userID, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["sub"]), 10, 64)
		if err != nil {
			utils.UnauthorizedErrorResponse(w, r, err, app.logger)
			return
		}

		ctx := r.Context()

		user, err := app.getUser(ctx, userID)
		if err != nil {
			utils.UnauthorizedErrorResponse(w, r, err, app.logger)
			return
		}

		ctx = context.WithValue(ctx, users.UserCtx, user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (app *application) BasicAuthMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// read the auth header
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				utils.UnauthorizedBasicErrorResponse(w, r, fmt.Errorf("authorization header is missing"), app.logger)
				return
			}

			// parse it -> get the base64
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Basic" {
				utils.UnauthorizedBasicErrorResponse(w, r, fmt.Errorf("authorization header is malformed"), app.logger)
				return
			}

			// decode it
			decoded, err := base64.StdEncoding.DecodeString(parts[1])
			if err != nil {
				utils.UnauthorizedBasicErrorResponse(w, r, err, app.logger)
				return
			}

			// check the credentials
			username := app.config.Auth.Basic.User
			pass := app.config.Auth.Basic.Pass

			creds := strings.SplitN(string(decoded), ":", 2)
			if len(creds) != 2 || creds[0] != username || creds[1] != pass {
				utils.UnauthorizedBasicErrorResponse(w, r, fmt.Errorf("invalid credentials"), app.logger)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func (app *application) getUser(ctx context.Context, userID int64) (*users.User, error) {
	if !app.config.RedisCfg.Enabled {
		return app.userStore.GetByID(ctx, userID)
	}

	user, err := app.userCache.Get(ctx, userID)
	if err != nil {
		return nil, err
	}

	if user == nil {
		user, err = app.userStore.GetByID(ctx, userID)
		if err != nil {
			return nil, err
		}

		if err := app.userCache.Set(ctx, user); err != nil {
			return nil, err
		}
	}

	return user, nil
}
