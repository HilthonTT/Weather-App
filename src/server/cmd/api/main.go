// Command api runs the Weather HTTP API.
//
// It proxies OpenWeather / Open-Meteo requests (with an optional Redis cache),
// and manages users, authentication and per-user display settings backed by
// PostgreSQL. Configuration is read from the environment; see config.go.
package main

import (
	"context"
	"expvar"
	"runtime"

	"github.com/hilthontt/weather/internal/auth"
	"github.com/hilthontt/weather/internal/db"
	"github.com/hilthontt/weather/internal/mailer"
	"github.com/hilthontt/weather/internal/ratelimiter"
	"github.com/hilthontt/weather/internal/store"
	"github.com/hilthontt/weather/internal/store/cache"
	"github.com/hilthontt/weather/internal/tracing"
	_ "github.com/joho/godotenv/autoload"
	"go.uber.org/zap"
)

const version = "1.0.0"

//	@title			Weather API
//	@description	This is the Weather API for a weather app system.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	MIT
//	@license.url	https://opensource.org/licenses/MIT

//	@BasePath					/v1
//
//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization
//	@description				Use your API key in the Authorization header

func main() {
	cfg := loadConfig()

	// Logger
	logger := zap.Must(zap.NewProduction()).Sugar()
	defer logger.Sync()

	if err := tracing.SetGlobalTracer(context.Background(), cfg.serviceName, cfg.jaegerAddr); err != nil {
		logger.Fatalw("failed to set global tracer", "error", err)
	}

	// Main Database
	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)
	if err != nil {
		logger.Fatal(err)
	}

	defer db.Close()
	logger.Info("database connection pool established")

	// Cache. When Redis is disabled a no-op cache is used so handlers never
	// have to special-case a missing client.
	cacheStorage := cache.NewNoopStorage()
	if cfg.redisCfg.enabled {
		rdb := cache.NewRedisClient(cfg.redisCfg.addr, cfg.redisCfg.pw, cfg.redisCfg.db)
		logger.Info("redis cache connection established")
		defer rdb.Close()

		cacheStorage = cache.NewRedisStorage(rdb)
	}

	// Rate limiter
	rateLimiter := ratelimiter.NewFixedWindowLimiter(
		cfg.rateLimiter.RequestsPerTimeFrame,
		cfg.rateLimiter.TimeFrame,
	)

	// Mailer
	mailtrap, err := mailer.NewMailTrapClient(
		cfg.mail.mailTrap.username,
		cfg.mail.mailTrap.password,
		cfg.mail.fromEmail,
	)
	if err != nil {
		logger.Fatal(err)
	}

	// Authenticator (audience and issuer are both the service name)
	jwtAuthenticator := auth.NewJWTAuthenticator(
		cfg.auth.token.secret,
		cfg.auth.token.iss,
		cfg.auth.token.iss,
	)

	app := &application{
		config:        cfg,
		store:         store.NewStorage(db, cfg.openWeather.apiKey),
		cacheStorage:  cacheStorage,
		logger:        logger,
		authenticator: jwtAuthenticator,
		rateLimiter:   rateLimiter,
		mailer:        mailtrap,
	}

	// Metrics collected
	expvar.NewString("version").Set(version)
	expvar.Publish("database", expvar.Func(func() any {
		return db.Stats()
	}))
	expvar.Publish("goroutines", expvar.Func(func() any {
		return runtime.NumGoroutine()
	}))

	mux := app.mount()

	logger.Fatal(app.run(mux))
}
