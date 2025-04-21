package cache

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/hilthontt/weather/internal/store"
)

type Storage struct {
	Users interface {
		Get(context.Context, int64) (*store.User, error)
		Set(context.Context, *store.User) error
		Delete(context.Context, int64)
	}
	Weathers interface {
		GetWeather(context.Context, float64, float64) (*store.WeatherResponse, error)
		GetWeatherCity(context.Context, string) (*store.WeatherResponse, error)
		SetWeather(context.Context, *store.WeatherResponse) error
		DeleteWeather(context.Context, float64, float64) error

		GetForecast(context.Context, float64, float64) (*store.ForecastResponse, error)
		GetForecastByCity(context.Context, string) (*store.ForecastResponse, error)
		SetForecast(context.Context, *store.ForecastResponse) error
		DeleteForecast(context.Context, float64, float64) error

		GetOpenMeteo(context.Context, float64, float64) (*store.OpenMeteoResponse, error)
		SetOpenMeteo(context.Context, *store.OpenMeteoResponse, float64, float64) error
		DeleteOpenMeteo(context.Context, float64, float64) error
	}
}

func NewRedisStorage(rdb *redis.Client) Storage {
	return Storage{
		Users:    &UserStore{rdb: rdb},
		Weathers: &WeatherStore{rdb: rdb},
	}
}
