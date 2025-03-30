package main

import (
	"github.com/hilthontt/weather/internal/config"
	"github.com/hilthontt/weather/internal/ratelimiter"
	"github.com/hilthontt/weather/services/weather"
	_ "github.com/joho/godotenv/autoload"
	"go.uber.org/zap"
)

func main() {
	cfg := config.NewConfig()

	logger := zap.Must(zap.NewProduction()).Sugar()
	defer logger.Sync()

	// Weather Client
	weatherClient := weather.NewClient(cfg.OpenWeather.ApiKey)

	// Ratelimiter
	rateLimiter := ratelimiter.NewFixedWindowLimiter(
		cfg.RateLimiter.RequestsPerTimeFrame,
		cfg.RateLimiter.TimeFrame,
	)

	app := &application{
		config:        *cfg,
		weatherClient: *weatherClient,
		logger:        logger,
		rateLimiter:   rateLimiter,
	}

	mux := app.mount()

	logger.Fatal(app.run(mux))
}
