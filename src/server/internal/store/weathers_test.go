package store

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"
)

func init() {
	retryBackoff = time.Millisecond
}

func TestFetchJSON_NotFoundIsNotRetried(t *testing.T) {
	var calls atomic.Int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls.Add(1)
		http.Error(w, `{"cod":"404","message":"city not found"}`, http.StatusNotFound)
	}))
	defer srv.Close()

	_, err := fetchJSON[WeatherResponse](context.Background(), srv.Client(), srv.URL, 3)
	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("err = %v, want ErrNotFound", err)
	}
	if calls.Load() != 1 {
		t.Errorf("upstream called %d times, want 1", calls.Load())
	}
}

func TestFetchJSON_RetriesServerErrors(t *testing.T) {
	var calls atomic.Int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if calls.Add(1) < 3 {
			w.WriteHeader(http.StatusBadGateway)
			return
		}
		_, _ = w.Write([]byte(`{"name":"London","main":{"temp":12.5}}`))
	}))
	defer srv.Close()

	got, err := fetchJSON[WeatherResponse](context.Background(), srv.Client(), srv.URL, 3)
	if err != nil {
		t.Fatal(err)
	}
	if got.Name != "London" || got.Main.Temp != 12.5 {
		t.Errorf("unexpected response: %+v", got)
	}
}

func TestFetchJSON_ReportsLastError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusServiceUnavailable)
	}))
	defer srv.Close()

	_, err := fetchJSON[WeatherResponse](context.Background(), srv.Client(), srv.URL, 2)
	if err == nil || err.Error() != "request failed after 2 attempts: upstream returned status 503" {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestGetWeather_UsesMetricUnits(t *testing.T) {
	var units string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		units = r.URL.Query().Get("units")
		_, _ = w.Write([]byte(`{}`))
	}))
	defer srv.Close()

	// Point the client at the test server by rewriting the request URL.
	client := srv.Client()
	client.Transport = rewriteTransport{target: srv.URL, base: client.Transport}

	s := &WeatherStore{APIKey: "key", HTTP: client}
	if _, err := s.GetWeather(context.Background(), "London"); err != nil {
		t.Fatal(err)
	}
	if units != "metric" {
		t.Errorf("units = %q, want metric", units)
	}
}

func TestSettingsUpdate_IsPartial(t *testing.T) {
	s := DefaultSettings(1)
	s.Update("", TwelveHour, "")

	if s.TempFormat != Celsius || s.TimeFormat != TwelveHour || s.SpeedFormat != Kmph {
		t.Errorf("unexpected settings: %+v", s)
	}
}

type rewriteTransport struct {
	target string
	base   http.RoundTripper
}

func (t rewriteTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	u, _ := r.URL.Parse(t.target)
	r.URL.Scheme, r.URL.Host = u.Scheme, u.Host
	return t.base.RoundTrip(r)
}
