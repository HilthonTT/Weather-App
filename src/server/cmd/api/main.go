package main

import (
	"context"
	"expvar"
	"runtime"

	"github.com/go-redis/redis/v8"
	"github.com/hilthontt/weather/internal/cache"
	"github.com/hilthontt/weather/internal/config"
	"github.com/hilthontt/weather/internal/ratelimiter"
	"github.com/hilthontt/weather/internal/tracer"
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

// @BasePath					/v1
//
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
// @description				API key authentication for accessing weather data
func main() {
	cfg := config.NewConfig()

	logger := zap.Must(zap.NewProduction()).Sugar()
	defer logger.Sync()

	if err := tracer.SetGlobalTracer(context.TODO(), cfg.JaegerAddr); err != nil {
		logger.Fatal("could not set global tracer", zap.Error(err))
	}

	// Weather Client
	weatherClient := weather.NewClient(cfg.OpenWeather.ApiKey)

	// Ratelimiter
	rateLimiter := ratelimiter.NewFixedWindowLimiter(
		cfg.RateLimiter.RequestsPerTimeFrame,
		cfg.RateLimiter.TimeFrame,
	)

	// Cache
	var rdb *redis.Client
	if cfg.RedisCfg.Enabled {
		rdb = cache.NewRedisClient(cfg.RedisCfg.Addr, cfg.RedisCfg.Pw, cfg.RedisCfg.DB)
		logger.Info("redis cache connection established")

		defer rdb.Close()
	}
	weatherCache := weather.NewWeatherCache(rdb)

	app := &application{
		config:        *cfg,
		weatherClient: *weatherClient,
		logger:        logger,
		rateLimiter:   rateLimiter,
		weatherCache:  weatherCache,
	}

	// Metrics collected
	expvar.Publish("goroutines", expvar.Func(func() any {
		return runtime.NumGoroutine()
	}))

	mux := app.mount()

	logger.Fatal(app.run(mux))
}
