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
	JaegerAddr  string
	RedisCfg    RedisConfig
	Auth        AuthConfig
}

type OpenWeatherConfig struct {
	ApiKey string
}

type DbConfig struct {
	Addr               string
	MaxOpenConnections int
	MaxIdleConnections int
	MaxIdleTime        string
}

type RedisConfig struct {
	Addr    string
	Pw      string
	DB      int
	Enabled bool
}

type AuthConfig struct {
	Basic BasicConfig
	Token TokenConfig
}

type BasicConfig struct {
	User string
	Pass string
}

type TokenConfig struct {
	Secret string
	Exp    time.Duration
	Iss    string
}

func NewConfig() *Config {
	return &Config{
		Addr:       env.GetString("ADDR", ":8080"),
		ApiURL:     env.GetString("EXTERNAL_URL", "localhost:8080"),
		JaegerAddr: env.GetString("JAEGER_ADDR", "localhost:4318"),
		Env:        env.GetString("ENV", "development"),
		OpenWeather: OpenWeatherConfig{
			ApiKey: env.GetString("OPEN_WEATHER_API_KEY", ""),
		},
		RateLimiter: ratelimiter.Config{
			RequestsPerTimeFrame: env.GetInt("RATELIMITER_REQUESTS_COUNT", 20),
			TimeFrame:            time.Second * 5,
			Enabled:              env.GetBool("RATE_LIMITER_ENABLED", true),
		},
		Db: DbConfig{
			Addr:               env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/weather?sslmode=disable"),
			MaxOpenConnections: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			MaxIdleConnections: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			MaxIdleTime:        env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		RedisCfg: RedisConfig{
			Addr:    env.GetString("REDIS_ADDR", "localhost:6379"),
			Pw:      env.GetString("REDIS_PW", ""),
			DB:      env.GetInt("REDIS_DB", 0),
			Enabled: env.GetBool("REDIS_ENABLED", true),
		},
		Auth: AuthConfig{
			Basic: BasicConfig{
				User: env.GetString("AUTH_BASIC_USER", "admin"),
				Pass: env.GetString("AUTH_BASIC_PASS", "admin"),
			},
			Token: TokenConfig{
				Secret: env.GetString("AUTH_TOKEN_SECRET", "very-secret-key-do-not-share-pls-pls"),
				Exp:    time.Hour * 24 * 3, // 3 days
				Iss:    "weather",
			},
		},
	}
}

func NewTestConfig() *Config {
	return &Config{
		RateLimiter: ratelimiter.Config{
			RequestsPerTimeFrame: 20,
			TimeFrame:            time.Second * 5,
			Enabled:              true,
		},
		Addr: ":8080",
	}
}
