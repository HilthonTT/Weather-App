// Package store contains the persistence layer: PostgreSQL-backed users,
// roles and settings, plus the HTTP clients for the upstream weather APIs.
package store

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"time"
)

var (
	ErrNotFound          = errors.New("resource not found")
	ErrConflict          = errors.New("resource already exists")
	QueryTimeoutDuration = time.Second * 5
)

// Storage groups every repository used by the API. Each field is an
// interface so handlers can be tested with mocks (see mocks.go).
type Storage struct {
	Users interface {
		GetByID(context.Context, int64) (*User, error)
		GetByEmail(context.Context, string) (*User, error)
		Create(context.Context, *sql.Tx, *User) error
		CreateAndSendEmailVerification(context.Context, *User, string, time.Duration) error
		Verify(context.Context, string) error
		Delete(context.Context, int64) error
	}
	Roles interface {
		GetByName(context.Context, string) (*Role, error)
	}
	Weathers interface {
		GetWeather(context.Context, string) (*WeatherResponse, error)
		GetWeatherByCoords(context.Context, float64, float64) (*WeatherResponse, error)
		GetForecast(context.Context, string) (*ForecastResponse, error)
		GetForecastByCoords(context.Context, float64, float64) (*ForecastResponse, error)
		GetOpenMeteoByCoords(context.Context, float64, float64) (*OpenMeteoResponse, error)
	}
	Settings interface {
		Create(ctx context.Context, settings *Settings) error
		GetByUserID(ctx context.Context, userID int64) (*Settings, error)
		Update(ctx context.Context, settings *Settings) error
	}
}

// NewStorage returns the production Storage implementation.
func NewStorage(db *sql.DB, openWeatherApiKey string) Storage {
	return Storage{
		Users:    &UserStore{db},
		Roles:    &RoleStore{db},
		Settings: &SettingsStore{db},
		Weathers: &WeatherStore{
			APIKey: openWeatherApiKey,
			HTTP: &http.Client{
				Timeout: 10 * time.Second,
			},
		},
	}
}

func withTx(db *sql.DB, ctx context.Context, fn func(*sql.Tx) error) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	if err := fn(tx); err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}
