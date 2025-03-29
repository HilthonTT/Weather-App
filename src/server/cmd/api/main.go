package main

import (
	"log"
	"time"

	"github.com/hilthontt/weather/internal/env"
	"github.com/hilthontt/weather/internal/ratelimiter"
	"github.com/hilthontt/weather/services/weather"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		openWeather: openWeatherConfig{
			apiKey: env.GetString("OPEN_WEATHER_API_KEY", ""),
		},
		rateLimiter: ratelimiter.Config{
			RequestsPerTimeFrame: env.GetInt("RATELIMITER_REQUESTS_COUNT", 20),
			TimeFrame:            time.Second * 5,
			Enabled:              env.GetBool("RATE_LIMITER_ENABLED", true),
		},
	}

	// Weather Client
	weatherClient := weather.NewClient(cfg.openWeather.apiKey)

	rateLimiter := ratelimiter.NewFixedWindowLimiter(
		cfg.rateLimiter.RequestsPerTimeFrame,
		cfg.rateLimiter.TimeFrame,
	)

	app := &application{
		config:        cfg,
		weatherClient: *weatherClient,
		rateLimiter:   rateLimiter,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
