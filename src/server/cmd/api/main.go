package main

import (
	"context"
	"fmt"
	"time"

	"github.com/hilthontt/weather/internal/config"
	"github.com/hilthontt/weather/internal/ratelimiter"
	"github.com/hilthontt/weather/internal/tracer"
	"github.com/hilthontt/weather/services/weather"
	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
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

	// MongoDB connection
	uri := fmt.Sprintf("mongodb://%s:%s@%s", cfg.Db.MongoUser, cfg.Db.MongoPassword, cfg.Db.MongoAddr)
	mongoClient, err := connectToMongoDB(uri)
	if err != nil {
		logger.Fatal("failed to connect to mongo db", zap.Error(err))
	}

	weatherStore := weather.NewStore(mongoClient)

	app := &application{
		config:        *cfg,
		weatherClient: *weatherClient,
		logger:        logger,
		rateLimiter:   rateLimiter,
		weatherStore:  weatherStore,
	}

	mux := app.mount()

	logger.Fatal(app.run(mux))
}

func connectToMongoDB(uri string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, readpref.Primary())

	return client, err
}
