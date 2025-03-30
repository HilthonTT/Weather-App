package config

import (
	"time"

	"github.com/hilthontt/weather/internal/env"
	"github.com/hilthontt/weather/internal/ratelimiter"
)

type Config struct {
	Addr        string
	ApiURL      string
	OpenWeather OpenWeatherConfig
	RateLimiter ratelimiter.Config
	Env         string
	Db          DbConfig
}

type OpenWeatherConfig struct {
	ApiKey string
}

type DbConfig struct {
	MongoUser     string
	MongoPassword string
	MongoAddr     string
}

func NewConfig() *Config {
	return &Config{
		Addr:   env.GetString("ADDR", ":8080"),
		ApiURL: env.GetString("EXTERNAL_URL", "localhost:8080"),
		OpenWeather: OpenWeatherConfig{
			ApiKey: env.GetString("OPEN_WEATHER_API_KEY", ""),
		},
		RateLimiter: ratelimiter.Config{
			RequestsPerTimeFrame: env.GetInt("RATELIMITER_REQUESTS_COUNT", 20),
			TimeFrame:            time.Second * 5,
			Enabled:              env.GetBool("RATE_LIMITER_ENABLED", true),
		},
		Db: DbConfig{
			MongoUser:     env.GetString("MONGO_DB_USER", "root"),
			MongoPassword: env.GetString("MONGO_DB_PASS", "example"),
			MongoAddr:     env.GetString("MONGO_DB_HOST", "localhost:27017"),
		},
	}
}
