package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/hilthontt/weather/internal/store"
)

func (app *application) RateLimiterMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if app.config.rateLimiter.Enabled {
			if allow, retryAfter := app.rateLimiter.Allow(r.RemoteAddr); !allow {
				app.rateLimitExceededResponse(w, r, retryAfter.String())
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}

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

func (app *application) getUser(ctx context.Context, userID int64) (*store.User, error) {
	if !app.config.redisCfg.enabled {

		return app.store.Users.GetByID(ctx, userID)
	}

	user, err := app.cacheStorage.Users.Get(ctx, userID)
	if err != nil {

		return nil, err
	}

	if user == nil {
		user, err = app.store.Users.GetByID(ctx, userID)
		if err != nil {

			return nil, err
		}

		if err := app.cacheStorage.Users.Set(ctx, user); err != nil {

			return nil, err
		}
	}

	return user, nil
}

type ErrorResponse struct {
	Type      string `json:"type"`
	Title     string `json:"title"`
	Status    int    `json:"status"`
	Instance  string `json:"instance"`
	TraceID   string `json:"traceId"`
	RequestID string `json:"requestId"`
}

type ResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *ResponseWriter) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

func (app *application) statusCodePagesMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rw := &ResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		// Generate traceId and requestId
		traceId := uuid.New().String()
		requestId := uuid.New().String()

		// Add traceId and requestId to the context for use in the error response
		ctx := r.Context()
		ctx = context.WithValue(ctx, "traceId", traceId)
		ctx = context.WithValue(ctx, "requestId", requestId)

		// Call the next handler
		next.ServeHTTP(rw, r.WithContext(ctx))

		// Check the status code and serve the appropriate response for error codes
		if rw.statusCode == http.StatusInternalServerError ||
			rw.statusCode == http.StatusBadRequest || rw.statusCode == http.StatusConflict ||
			rw.statusCode == http.StatusForbidden {
			// Safely retrieve the traceId and requestId from context
			traceID, traceIDExists := r.Context().Value("traceId").(string)
			requestID, requestIDExists := r.Context().Value("requestId").(string)

			// If the context values don't exist, use default empty string
			if !traceIDExists {
				traceID = "unknown"
			}
			if !requestIDExists {
				requestID = "unknown"
			}

			instance := fmt.Sprintf("%s %s", r.Method, r.URL.Path)

			// Determine the error type URL based on the status code
			var errorType string
			switch rw.statusCode {
			case http.StatusConflict:
				errorType = "https://tools.ietf.org/html/rfc7231#section-6.5.8"
			case http.StatusForbidden:
				errorType = "https://tools.ietf.org/html/rfc7231#section-6.5.3"
			case http.StatusBadRequest:
				errorType = "https://tools.ietf.org/html/rfc7231#section-6.5.1"
			default:
				errorType = "https://tools.ietf.org/html/rfc7231#section-6.6.1" // General error type
			}

			// Create the error response
			errResponse := ErrorResponse{
				Type:      errorType,
				Title:     http.StatusText(rw.statusCode),
				Status:    rw.statusCode,
				Instance:  instance,
				TraceID:   traceID,
				RequestID: requestID,
			}

			// Send the error response
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(rw.statusCode)
			json.NewEncoder(w).Encode(errResponse)
		}
	})
}

func (app *application) notFoundHandler(w http.ResponseWriter, r *http.Request) {

	// Generate traceId and requestId
	traceId := uuid.New().String()
	requestId := uuid.New().String()

	// Add traceId and requestId to the context for use in the error response
	ctx := r.Context()
	ctx = context.WithValue(ctx, "traceId", traceId)
	ctx = context.WithValue(ctx, "requestId", requestId)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)

	errResponse := ErrorResponse{
		Type:      "https://tools.ietf.org/html/rfc9110#section-15.5.5",
		Title:     "Not Found",
		Status:    http.StatusNotFound,
		Instance:  fmt.Sprintf("%s %s", r.Method, r.URL.Path),
		TraceID:   traceId,
		RequestID: requestId,
	}

	json.NewEncoder(w).Encode(errResponse)
}
