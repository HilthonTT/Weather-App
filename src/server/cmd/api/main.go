package main

import (
	"github.com/hilthontt/weather/internal/config"
	"github.com/hilthontt/weather/internal/ratelimiter"
	"github.com/hilthontt/weather/services/weather"
	_ "github.com/joho/godotenv/autoload"
	"go.uber.org/zap"
)

const version = "1.0.0"

//	@title			Weather API
//	@description	API for fetching weather data
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@BasePath					/v1
//
//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization
//	@description				API key authentication for accessing weather data
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
