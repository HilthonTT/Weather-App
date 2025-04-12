package weather

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/hilthontt/weather/types"
)

type WeatherCache struct {
	rdb *redis.Client
}

func NewWeatherCache(rdb *redis.Client) *WeatherCache {
	return &WeatherCache{rdb}
}

const WeatherExpTime = time.Hour * 24
const ForecastExpTime = time.Hour * 24
const OpenMeteoExpTime = time.Hour * 24

func (c *WeatherCache) GetWeather(ctx context.Context, lat, lon float64) (*types.WeatherResponse, error) {
	cacheKey := fmt.Sprintf("weather-%f-%f", lat, lon)

	data, err := c.rdb.Get(ctx, cacheKey).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	var weather types.WeatherResponse
	if data != "" {
		err := json.Unmarshal([]byte(data), &weather)
		if err != nil {
			return nil, err
		}
	}

	return &weather, nil
}

func (c *WeatherCache) GetWeatherCity(ctx context.Context, city string) (*types.WeatherResponse, error) {
	cacheKey := fmt.Sprintf("weather-%s", city)

	data, err := c.rdb.Get(ctx, cacheKey).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	var weather types.WeatherResponse
	if data != "" {
		err := json.Unmarshal([]byte(data), &weather)
		if err != nil {
			return nil, err
		}
	}

	return &weather, nil
}

func (c *WeatherCache) SetWeather(ctx context.Context, weather *types.WeatherResponse) error {
	cacheKey1 := fmt.Sprintf("weather-%f-%f", weather.Coord.Lat, weather.Coord.Lon)
	cacheKey2 := fmt.Sprintf("weather-%s", weather.Name)

	json, err := json.Marshal(weather)
	if err != nil {
		return err
	}

	err = c.rdb.SetEX(ctx, cacheKey1, json, WeatherExpTime).Err()
	if err != nil {
		return err
	}

	err = c.rdb.SetEX(ctx, cacheKey2, json, WeatherExpTime).Err()
	if err != nil {
		return err
	}

	return nil
}

func (c *WeatherCache) DeleteWeather(ctx context.Context, lat, lon float64) error {
	cacheKey := fmt.Sprintf("weather-%f-%f", lat, lon)

	return c.rdb.Del(ctx, cacheKey).Err()
}

func (c *WeatherCache) GetForecast(ctx context.Context, lat, lon float64) (*types.ForecastResponse, error) {
	cacheKey := fmt.Sprintf("forecast-%f-%f", lat, lon)

	data, err := c.rdb.Get(ctx, cacheKey).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	var forecast types.ForecastResponse
	if data != "" {
		err := json.Unmarshal([]byte(data), &forecast)
		if err != nil {
			return nil, err
		}
	}

	return &forecast, nil
}

func (c *WeatherCache) GetForecastByCity(ctx context.Context, city string) (*types.ForecastResponse, error) {
	cacheKey := fmt.Sprintf("forecast-%s", city)

	data, err := c.rdb.Get(ctx, cacheKey).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	var forecast types.ForecastResponse
	if data != "" {
		err := json.Unmarshal([]byte(data), &forecast)
		if err != nil {
			return nil, err
		}
	}

	return &forecast, nil
}

func (c *WeatherCache) SetForecast(ctx context.Context, forecast *types.ForecastResponse) error {
	cacheKey1 := fmt.Sprintf("forecast-%f-%f", forecast.City.Coord.Lat, forecast.City.Coord.Lon)
	cacheKey2 := fmt.Sprintf("forecast-%s", forecast.City.Name)

	json, err := json.Marshal(forecast)
	if err != nil {
		return err
	}

	err = c.rdb.SetEX(ctx, cacheKey1, json, WeatherExpTime).Err()
	if err != nil {
		return err
	}

	err = c.rdb.SetEX(ctx, cacheKey2, json, WeatherExpTime).Err()
	if err != nil {
		return err
	}

	return nil
}

func (c *WeatherCache) DeleteForecast(ctx context.Context, lat, lon float64) error {
	cacheKey := fmt.Sprintf("forecast-%f-%f", lat, lon)

	return c.rdb.Del(ctx, cacheKey).Err()
}

func (c *WeatherCache) GetOpenMeteo(ctx context.Context, lat, lon float64) (*types.OpenMeteoResponse, error) {
	cacheKey := fmt.Sprintf("open-meteo-%f-%f", lat, lon)

	data, err := c.rdb.Get(ctx, cacheKey).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	var openMeteo types.OpenMeteoResponse
	if data != "" {
		err := json.Unmarshal([]byte(data), &openMeteo)
		if err != nil {
			return nil, err
		}
	}

	return &openMeteo, nil
}

func (c *WeatherCache) SetOpenMeteo(ctx context.Context, openMeteo *types.OpenMeteoResponse, lat, lon float64) error {
	cacheKey := fmt.Sprintf("open-meteo-%f-%f", lat, lon)

	json, err := json.Marshal(openMeteo)
	if err != nil {
		return err
	}

	err = c.rdb.SetEX(ctx, cacheKey, json, WeatherExpTime).Err()
	if err != nil {
		return err
	}

	return nil
}
