package cache

import (
	"context"

	"github.com/hilthontt/weather/internal/store"
)

// NewNoopStorage returns a cache that never stores anything. It is used when
// Redis is disabled (REDIS_ENABLED=false) so callers never have to nil-check.
func NewNoopStorage() Storage {
	return Storage{
		Users:    noopUserStore{},
		Weathers: noopWeatherStore{},
	}
}

type noopUserStore struct{}

func (noopUserStore) Get(context.Context, int64) (*store.User, error) { return nil, nil }
func (noopUserStore) Set(context.Context, *store.User) error          { return nil }
func (noopUserStore) Delete(context.Context, int64)                   {}

type noopWeatherStore struct{}

func (noopWeatherStore) GetWeather(context.Context, float64, float64) (*store.WeatherResponse, error) {
	return nil, nil
}
func (noopWeatherStore) GetWeatherCity(context.Context, string) (*store.WeatherResponse, error) {
	return nil, nil
}
func (noopWeatherStore) SetWeather(context.Context, *store.WeatherResponse) error { return nil }
func (noopWeatherStore) DeleteWeather(context.Context, float64, float64) error    { return nil }

func (noopWeatherStore) GetForecast(context.Context, float64, float64) (*store.ForecastResponse, error) {
	return nil, nil
}
func (noopWeatherStore) GetForecastByCity(context.Context, string) (*store.ForecastResponse, error) {
	return nil, nil
}
func (noopWeatherStore) SetForecast(context.Context, *store.ForecastResponse) error { return nil }
func (noopWeatherStore) DeleteForecast(context.Context, float64, float64) error     { return nil }

func (noopWeatherStore) GetOpenMeteo(context.Context, float64, float64) (*store.OpenMeteoResponse, error) {
	return nil, nil
}
func (noopWeatherStore) SetOpenMeteo(context.Context, *store.OpenMeteoResponse, float64, float64) error {
	return nil
}
func (noopWeatherStore) DeleteOpenMeteo(context.Context, float64, float64) error { return nil }
