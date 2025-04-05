package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hilthontt/weather/internal/auth"
	"github.com/hilthontt/weather/internal/config"
	"github.com/hilthontt/weather/internal/ratelimiter"
	"go.uber.org/zap"
)

func newTestApplication(t *testing.T, cfg config.Config) *application {
	t.Helper()

	logger := zap.NewNop().Sugar()
	// Uncomment to enable logs
	// logger := zap.Must(zap.NewProduction()).Sugar()

	testAuth := &auth.TestAuthenticator{}

	rateLimiter := ratelimiter.NewFixedWindowLimiter(
		cfg.RateLimiter.RequestsPerTimeFrame,
		cfg.RateLimiter.TimeFrame,
	)

	return &application{
		logger:        logger,
		authenticator: testAuth,
		config:        cfg,
		rateLimiter:   rateLimiter,
	}
}
func executeRequest(req *http.Request, mux http.Handler) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d", expected, actual)
	}
}
