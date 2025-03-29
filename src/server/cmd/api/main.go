package main

import (
	"log"

	"github.com/hilthontt/weather/internal/env"
	"github.com/hilthontt/weather/services/weather"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		openWeather: openWeatherConfig{
			apiKey: env.GetString("OPEN_WEATHER_API_KEY", ""),
		},
	}

	// Weather Client
	weatherClient := weather.NewClient(cfg.openWeather.apiKey)

	app := &application{
		config:        cfg,
		weatherClient: *weatherClient,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
