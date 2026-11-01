package main

import (
	"time"

	"github.com/hilthontt/weather/internal/env"
	"github.com/hilthontt/weather/internal/ratelimiter"
)

// config holds every runtime setting of the API. Values are read from the
// environment (and from a local .env file, via godotenv/autoload) in
// loadConfig. See .env.example for the full list of variables.
type config struct {
	serviceName string
	jaegerAddr  string
	addr        string
	db          dbConfig
	env         string
	apiURL      string
	frontendURL string
	auth        authConfig
	redisCfg    redisConfig
	rateLimiter ratelimiter.Config
	openWeather openWeatherConfig
	mail        mailConfig
}

type redisConfig struct {
	addr    string
	pw      string
	db      int
	enabled bool
}

type authConfig struct {
	basic basicConfig
	token tokenConfig
}

type tokenConfig struct {
	secret string
	exp    time.Duration
	iss    string
}

type basicConfig struct {
	user string
	pass string
}

type dbConfig struct {
	addr         string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}

type openWeatherConfig struct {
	apiKey string
}

type mailConfig struct {
	mailTrap  mailTrapConfig
	fromEmail string
	exp       time.Duration
}

type mailTrapConfig struct {
	username string
	password string
}

// loadConfig builds the application configuration from environment
// variables, falling back to development-friendly defaults.
func loadConfig() config {
	return config{
		serviceName: "weather",
		jaegerAddr:  env.GetString("JAEGER_ADDR", "localhost:4318"),
		addr:        env.GetString("ADDR", ":8080"),
		apiURL:      env.GetString("EXTERNAL_URL", "localhost:8080"),
		frontendURL: env.GetString("FRONTEND_URL", "http://localhost:5173"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/weather?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		redisCfg: redisConfig{
			addr:    env.GetString("REDIS_ADDR", "localhost:6379"),
			pw:      env.GetString("REDIS_PW", ""),
			db:      env.GetInt("REDIS_DB", 0),
			enabled: env.GetBool("REDIS_ENABLED", true),
		},
		env: env.GetString("ENV", "development"),
		auth: authConfig{
			basic: basicConfig{
				user: env.GetString("AUTH_BASIC_USER", "admin"),
				pass: env.GetString("AUTH_BASIC_PASS", "admin"),
			},
			token: tokenConfig{
				secret: env.GetString("AUTH_TOKEN_SECRET", "example"),
				exp:    time.Hour * 24 * 3, // 3 days
				iss:    "weather",
			},
		},
		rateLimiter: ratelimiter.Config{
			RequestsPerTimeFrame: env.GetInt("RATELIMITER_REQUESTS_COUNT", 20),
			TimeFrame:            time.Second * 5,
			Enabled:              env.GetBool("RATE_LIMITER_ENABLED", true),
		},
		openWeather: openWeatherConfig{
			apiKey: env.GetString("OPEN_WEATHER_API_KEY", ""),
		},
		mail: mailConfig{
			exp:       time.Hour * 24 * 3, // 3 days
			fromEmail: env.GetString("FROM_EMAIL", "from@example.com"),
			mailTrap: mailTrapConfig{
				username: env.GetString("MAILTRAP_USERNAME", ""),
				password: env.GetString("MAILTRAP_PASSWORD", ""),
			},
		},
	}
}
