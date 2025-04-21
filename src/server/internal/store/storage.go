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
		GetWeather(string) (*WeatherResponse, error)
		GetWeatherByCoords(float64, float64) (*WeatherResponse, error)
		GetForecast(string) (*ForecastResponse, error)
		GetForecastByCoords(float64, float64) (*ForecastResponse, error)
		GetOpenMeteoByCoords(float64, float64) (*OpenMeteoResponse, error)
	}
}

func NewStorage(db *sql.DB, openWeatherApiKey string) Storage {
	return Storage{
		Users: &UserStore{db},
		Roles: &RoleStore{db},
		Weathers: &WeatherStore{
			openWeatherApiKey,
			&http.Client{
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
