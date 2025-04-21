package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/hilthontt/weather/internal/store"
)

type WeatherStore struct {
	rdb *redis.Client
}

const WeatherExpTime = time.Hour

func (s *WeatherStore) GetWeather(ctx context.Context, lat, lon float64) (*store.WeatherResponse, error) {
	cacheKey := fmt.Sprintf("weather-%f-%f", lat, lon)

	data, err := s.rdb.Get(ctx, cacheKey).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	var weather store.WeatherResponse
	if data != "" {
		err := json.Unmarshal([]byte(data), &weather)
		if err != nil {
			return nil, err
		}
	}

	return &weather, nil
}

func (s *WeatherStore) GetWeatherCity(ctx context.Context, city string) (*store.WeatherResponse, error) {
	cacheKey := fmt.Sprintf("weather-%s", city)

	data, err := s.rdb.Get(ctx, cacheKey).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	var weather store.WeatherResponse
	if data != "" {
		err := json.Unmarshal([]byte(data), &weather)
		if err != nil {
			return nil, err
		}
	}

	return &weather, nil
}

func (s *WeatherStore) SetWeather(ctx context.Context, weather *store.WeatherResponse) error {
	cacheKey1 := fmt.Sprintf("weather-%f-%f", weather.Coord.Lat, weather.Coord.Lon)
	cacheKey2 := fmt.Sprintf("weather-%s", weather.Name)

	json, err := json.Marshal(weather)
	if err != nil {
		return err
	}

	err = s.rdb.SetEX(ctx, cacheKey1, json, WeatherExpTime).Err()
	if err != nil {
		return err
	}

	err = s.rdb.SetEX(ctx, cacheKey2, json, WeatherExpTime).Err()
	if err != nil {
		return err
	}

	return nil
}

func (s *WeatherStore) DeleteWeather(ctx context.Context, lat, lon float64) error {
	cacheKey := fmt.Sprintf("weather-%f-%f", lat, lon)

	return s.rdb.Del(ctx, cacheKey).Err()
}

func (s *WeatherStore) GetForecast(ctx context.Context, lat, lon float64) (*store.ForecastResponse, error) {
	cacheKey := fmt.Sprintf("forecast-%f-%f", lat, lon)

	data, err := s.rdb.Get(ctx, cacheKey).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	var forecast store.ForecastResponse
	if data != "" {
		err := json.Unmarshal([]byte(data), &forecast)
		if err != nil {
			return nil, err
		}
	}

	return &forecast, nil
}

func (s *WeatherStore) GetForecastByCity(ctx context.Context, city string) (*store.ForecastResponse, error) {
	cacheKey := fmt.Sprintf("forecast-%s", city)

	data, err := s.rdb.Get(ctx, cacheKey).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	var forecast store.ForecastResponse
	if data != "" {
		err := json.Unmarshal([]byte(data), &forecast)
		if err != nil {
			return nil, err
		}
	}

	return &forecast, nil
}

func (s *WeatherStore) SetForecast(ctx context.Context, forecast *store.ForecastResponse) error {
	cacheKey1 := fmt.Sprintf("forecast-%f-%f", forecast.City.Coord.Lat, forecast.City.Coord.Lon)
	cacheKey2 := fmt.Sprintf("forecast-%s", forecast.City.Name)

	json, err := json.Marshal(forecast)
	if err != nil {
		return err
	}

	err = s.rdb.SetEX(ctx, cacheKey1, json, WeatherExpTime).Err()
	if err != nil {
		return err
	}

	err = s.rdb.SetEX(ctx, cacheKey2, json, WeatherExpTime).Err()
	if err != nil {
		return err
	}

	return nil
}

func (s *WeatherStore) DeleteForecast(ctx context.Context, lat, lon float64) error {
	cacheKey := fmt.Sprintf("forecast-%f-%f", lat, lon)

	return s.rdb.Del(ctx, cacheKey).Err()
}

func (s *WeatherStore) GetOpenMeteo(ctx context.Context, lat, lon float64) (*store.OpenMeteoResponse, error) {
	cacheKey := fmt.Sprintf("open-meteo-%f-%f", lat, lon)

	data, err := s.rdb.Get(ctx, cacheKey).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	var openMeteo store.OpenMeteoResponse
	if data != "" {
		err := json.Unmarshal([]byte(data), &openMeteo)
		if err != nil {
			return nil, err
		}
	}

	return &openMeteo, nil
}

func (s *WeatherStore) SetOpenMeteo(ctx context.Context, openMeteo *store.OpenMeteoResponse, lat, lon float64) error {
	cacheKey := fmt.Sprintf("open-meteo-%f-%f", lat, lon)

	json, err := json.Marshal(openMeteo)
	if err != nil {
		return err
	}

	err = s.rdb.SetEX(ctx, cacheKey, json, WeatherExpTime).Err()
	if err != nil {
		return err
	}

	return nil
}

func (s *WeatherStore) DeleteOpenMeteo(ctx context.Context, lat, lon float64) error {
	cacheKey := fmt.Sprintf("open-meteo-%f-%f", lat, lon)

	return s.rdb.Del(ctx, cacheKey).Err()
}
