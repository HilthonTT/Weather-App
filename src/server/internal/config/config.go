package config

import (
	"time"

	"github.com/hilthontt/weather/internal/env"
	"github.com/hilthontt/weather/internal/ratelimiter"
)

type Config struct {
	Addr        string
	OpenWeather OpenWeatherConfig
	RateLimiter ratelimiter.Config
	Env         string
}

type OpenWeatherConfig struct {
	ApiKey string
}

func NewConfig() *Config {
	return &Config{
		Addr: env.GetString("ADDR", ":8080"),
		OpenWeather: OpenWeatherConfig{
			ApiKey: env.GetString("OPEN_WEATHER_API_KEY", ""),
		},
		RateLimiter: ratelimiter.Config{
			RequestsPerTimeFrame: env.GetInt("RATELIMITER_REQUESTS_COUNT", 20),
			TimeFrame:            time.Second * 5,
			Enabled:              env.GetBool("RATE_LIMITER_ENABLED", true),
		},
	}
}
