package main

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"github.com/hilthontt/weather/internal/store"
)

// The TestAuthenticator always authenticates as user 1 (see auth/mocks.go).
func TestSettings(t *testing.T) {
	app := newTestApplication(t, config{})
	mux := app.mount()

	token, err := app.authenticator.GenerateToken(nil)
	if err != nil {
		t.Fatal(err)
	}

	do := func(method, path, body string) (int, map[string]any) {
		t.Helper()

		var req *http.Request
		if body == "" {
			req, err = http.NewRequest(method, path, nil)
		} else {
			req, err = http.NewRequest(method, path, strings.NewReader(body))
		}
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Authorization", "Bearer "+token)

		rr := executeRequest(req, mux)

		var envelope map[string]any
		_ = json.NewDecoder(rr.Body).Decode(&envelope)
		return rr.Code, envelope
	}

	t.Run("creates default settings on first read", func(t *testing.T) {
		code, body := do(http.MethodGet, "/v1/users/1/settings/", "")
		checkResponseCode(t, http.StatusOK, code)

		data := body["data"].(map[string]any)
		if data["temp_format"] != string(store.Celsius) {
			t.Errorf("expected default celsius, got %v", data["temp_format"])
		}
	})

	t.Run("forbids reading another user's settings", func(t *testing.T) {
		code, _ := do(http.MethodGet, "/v1/users/2/settings/", "")
		checkResponseCode(t, http.StatusForbidden, code)
	})

	t.Run("forbids updating another user's settings", func(t *testing.T) {
		code, _ := do(http.MethodPatch, "/v1/users/2/settings/", `{"temp_format":"fahrenheit"}`)
		checkResponseCode(t, http.StatusForbidden, code)
	})

	t.Run("rejects a non-numeric user id", func(t *testing.T) {
		code, _ := do(http.MethodGet, "/v1/users/abc/settings/", "")
		checkResponseCode(t, http.StatusBadRequest, code)
	})

	t.Run("partial update keeps omitted fields", func(t *testing.T) {
		code, _ := do(http.MethodPatch, "/v1/users/1/settings/", `{"time_format":"12h"}`)
		checkResponseCode(t, http.StatusOK, code)

		code, body := do(http.MethodPatch, "/v1/users/1/settings/", `{"temp_format":"fahrenheit"}`)
		checkResponseCode(t, http.StatusOK, code)

		data := body["data"].(map[string]any)
		if data["temp_format"] != "fahrenheit" || data["time_format"] != "12h" || data["speed_format"] != "kmph" {
			t.Errorf("unexpected settings after partial updates: %v", data)
		}
	})

	t.Run("rejects invalid values", func(t *testing.T) {
		code, _ := do(http.MethodPatch, "/v1/users/1/settings/", `{"temp_format":"kelvin"}`)
		checkResponseCode(t, http.StatusBadRequest, code)
	})
}

func TestClientIP(t *testing.T) {
	cases := map[string]string{
		"192.168.1.1:54321": "192.168.1.1",
		"[::1]:8080":        "::1",
		"10.0.0.1":          "10.0.0.1", // already stripped by middleware.RealIP
	}

	for remoteAddr, want := range cases {
		r := &http.Request{RemoteAddr: remoteAddr}
		if got := clientIP(r); got != want {
			t.Errorf("clientIP(%q) = %q, want %q", remoteAddr, got, want)
		}
	}
}
