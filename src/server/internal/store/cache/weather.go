package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/hilthontt/weather/internal/store"
)

// WeatherStore caches upstream weather responses in Redis.
type WeatherStore struct {
	rdb *redis.Client
}

// WeatherExpTime is how long weather/forecast responses stay cached.
const WeatherExpTime = time.Hour

// coordKey builds a cache key from coordinates rounded to 2 decimals
// (~1 km). Rounding is required for cache hits: the coordinates a client
// sends (raw GPS) never exactly match the ones OpenWeather echoes back.
func coordKey(prefix string, lat, lon float64) string {
	return fmt.Sprintf("%s-%.2f-%.2f", prefix, lat, lon)
}

// cityKey builds a case-insensitive cache key for a city name, so that
// "london", " London" and "LONDON" share an entry.
func cityKey(prefix, city string) string {
	return fmt.Sprintf("%s-city-%s", prefix, strings.ToLower(strings.TrimSpace(city)))
}

func (s *WeatherStore) GetWeather(ctx context.Context, lat, lon float64) (*store.WeatherResponse, error) {
	return get[store.WeatherResponse](ctx, s.rdb, coordKey("weather", lat, lon))
}

func (s *WeatherStore) GetWeatherCity(ctx context.Context, city string) (*store.WeatherResponse, error) {
	return get[store.WeatherResponse](ctx, s.rdb, cityKey("weather", city))
}

// SetWeather caches the response under both its coordinates and city name.
func (s *WeatherStore) SetWeather(ctx context.Context, weather *store.WeatherResponse) error {
	return set(ctx, s.rdb, weather,
		coordKey("weather", weather.Coord.Lat, weather.Coord.Lon),
		cityKey("weather", weather.Name),
	)
}

func (s *WeatherStore) DeleteWeather(ctx context.Context, lat, lon float64) error {
	return s.rdb.Del(ctx, coordKey("weather", lat, lon)).Err()
}

func (s *WeatherStore) GetForecast(ctx context.Context, lat, lon float64) (*store.ForecastResponse, error) {
	return get[store.ForecastResponse](ctx, s.rdb, coordKey("forecast", lat, lon))
}

func (s *WeatherStore) GetForecastByCity(ctx context.Context, city string) (*store.ForecastResponse, error) {
	return get[store.ForecastResponse](ctx, s.rdb, cityKey("forecast", city))
}

// SetForecast caches the response under both its coordinates and city name.
func (s *WeatherStore) SetForecast(ctx context.Context, forecast *store.ForecastResponse) error {
	return set(ctx, s.rdb, forecast,
		coordKey("forecast", forecast.City.Coord.Lat, forecast.City.Coord.Lon),
		cityKey("forecast", forecast.City.Name),
	)
}

func (s *WeatherStore) DeleteForecast(ctx context.Context, lat, lon float64) error {
	return s.rdb.Del(ctx, coordKey("forecast", lat, lon)).Err()
}

func (s *WeatherStore) GetOpenMeteo(ctx context.Context, lat, lon float64) (*store.OpenMeteoResponse, error) {
	return get[store.OpenMeteoResponse](ctx, s.rdb, coordKey("open-meteo", lat, lon))
}

func (s *WeatherStore) SetOpenMeteo(ctx context.Context, openMeteo *store.OpenMeteoResponse, lat, lon float64) error {
	return set(ctx, s.rdb, openMeteo, coordKey("open-meteo", lat, lon))
}

func (s *WeatherStore) DeleteOpenMeteo(ctx context.Context, lat, lon float64) error {
	return s.rdb.Del(ctx, coordKey("open-meteo", lat, lon)).Err()
}

// get returns (nil, nil) on a cache miss.
func get[T any](ctx context.Context, rdb *redis.Client, key string) (*T, error) {
	data, err := rdb.Get(ctx, key).Bytes()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	var value T
	if err := json.Unmarshal(data, &value); err != nil {
		return nil, err
	}

	return &value, nil
}

func set(ctx context.Context, rdb *redis.Client, value any, keys ...string) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	pipe := rdb.TxPipeline()
	for _, key := range keys {
		pipe.SetEX(ctx, key, data, WeatherExpTime)
	}

	_, err = pipe.Exec(ctx)
	return err
}
