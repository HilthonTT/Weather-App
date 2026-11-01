package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/hilthontt/weather/internal/store"
)

// All weather handlers follow the same read-through cache pattern:
//  1. try the cache (errors are logged, never fatal),
//  2. on a miss, call the upstream weather provider,
//  3. store the result in the cache and return it.

// GetWeatherByCity godoc
//
//	@Summary		Fetches current weather by city
//	@Description	Fetches the current weather data (metric units) for a specified city
//	@Tags			weather
//	@Accept			json
//	@Produce		json
//	@Param			city	path		string	true	"City Name"
//	@Success		200		{object}	store.WeatherResponse
//	@Failure		400		{object}	error
//	@Failure		404		{object}	error
//	@Failure		500		{object}	error
//	@Router			/weather/{city} [get]
func (app *application) handleGetWeatherByCity(w http.ResponseWriter, r *http.Request) {
	city := chi.URLParam(r, "city")
	ctx := r.Context()

	weather, err := app.cacheStorage.Weathers.GetWeatherCity(ctx, city)
	if err != nil {
		app.logger.Warnw("error fetching weather from cache", "city", city, "error", err)
	}

	if weather == nil {
		weather, err = app.store.Weathers.GetWeather(ctx, city)
		if err != nil {
			app.upstreamErrorResponse(w, r, err)
			return
		}

		if err := app.cacheStorage.Weathers.SetWeather(ctx, weather); err != nil {
			app.logger.Warnw("error setting weather in cache", "city", city, "error", err)
		}
	}

	if err := app.jsonResponse(w, http.StatusOK, weather); err != nil {
		app.internalServerError(w, r, err)
	}
}

// GetWeatherByCoordinates godoc
//
//	@Summary		Fetches current weather by coordinates (latitude and longitude)
//	@Description	Fetches the current weather data (metric units) for a specified latitude and longitude
//	@Tags			weather
//	@Accept			json
//	@Produce		json
//	@Param			latitude	path		string	true	"Latitude"
//	@Param			longitude	path		string	true	"Longitude"
//	@Success		200			{object}	store.WeatherResponse
//	@Failure		400			{object}	error
//	@Failure		404			{object}	error
//	@Failure		500			{object}	error
//	@Router			/weather/coords/{latitude}/{longitude} [get]
func (app *application) handleGetWeatherByCoordinates(w http.ResponseWriter, r *http.Request) {
	latitude, longitude, err := parseCoordinates(r)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()

	weather, err := app.cacheStorage.Weathers.GetWeather(ctx, latitude, longitude)
	if err != nil {
		app.logger.Warnw("error fetching weather from cache", "error", err)
	}

	if weather == nil {
		weather, err = app.store.Weathers.GetWeatherByCoords(ctx, latitude, longitude)
		if err != nil {
			app.upstreamErrorResponse(w, r, err)
			return
		}

		if err := app.cacheStorage.Weathers.SetWeather(ctx, weather); err != nil {
			app.logger.Warnw("error setting weather in cache", "error", err)
		}
	}

	if err := app.jsonResponse(w, http.StatusOK, weather); err != nil {
		app.internalServerError(w, r, err)
	}
}

// GetForecastByCity godoc
//
//	@Summary		Fetches weather forecast by city
//	@Description	Fetches the 5 day / 3 hour forecast (metric units) for a specified city
//	@Tags			weather
//	@Accept			json
//	@Produce		json
//	@Param			city	path		string	true	"City Name"
//	@Success		200		{object}	store.ForecastResponse
//	@Failure		400		{object}	error
//	@Failure		404		{object}	error
//	@Failure		500		{object}	error
//	@Router			/weather/forecast/{city} [get]
func (app *application) handleGetForecast(w http.ResponseWriter, r *http.Request) {
	city := chi.URLParam(r, "city")
	ctx := r.Context()

	forecast, err := app.cacheStorage.Weathers.GetForecastByCity(ctx, city)
	if err != nil {
		app.logger.Warnw("error fetching forecast from cache", "city", city, "error", err)
	}

	if forecast == nil {
		forecast, err = app.store.Weathers.GetForecast(ctx, city)
		if err != nil {
			app.upstreamErrorResponse(w, r, err)
			return
		}

		if err := app.cacheStorage.Weathers.SetForecast(ctx, forecast); err != nil {
			app.logger.Warnw("error setting forecast in cache", "city", city, "error", err)
		}
	}

	if err := app.jsonResponse(w, http.StatusOK, forecast); err != nil {
		app.internalServerError(w, r, err)
	}
}

// GetForecastByCoordinates godoc
//
//	@Summary		Fetches weather forecast by coordinates (latitude and longitude)
//	@Description	Fetches the 5 day / 3 hour forecast (metric units) for a specified latitude and longitude
//	@Tags			weather
//	@Accept			json
//	@Produce		json
//	@Param			latitude	path		string	true	"Latitude"
//	@Param			longitude	path		string	true	"Longitude"
//	@Success		200			{object}	store.ForecastResponse
//	@Failure		400			{object}	error
//	@Failure		404			{object}	error
//	@Failure		500			{object}	error
//	@Router			/weather/forecast/coords/{latitude}/{longitude} [get]
func (app *application) handleGetForecastByCoordinates(w http.ResponseWriter, r *http.Request) {
	latitude, longitude, err := parseCoordinates(r)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()

	forecast, err := app.cacheStorage.Weathers.GetForecast(ctx, latitude, longitude)
	if err != nil {
		app.logger.Warnw("error fetching forecast from cache", "error", err)
	}

	if forecast == nil {
		forecast, err = app.store.Weathers.GetForecastByCoords(ctx, latitude, longitude)
		if err != nil {
			app.upstreamErrorResponse(w, r, err)
			return
		}

		if err := app.cacheStorage.Weathers.SetForecast(ctx, forecast); err != nil {
			app.logger.Warnw("error setting forecast in cache", "error", err)
		}
	}

	if err := app.jsonResponse(w, http.StatusOK, forecast); err != nil {
		app.internalServerError(w, r, err)
	}
}

// GetOpenMeteoByCoordinates godoc
//
//	@Summary		Fetches open meteo by coordinates (latitude and longitude)
//	@Description	Fetches the Open-Meteo daily + hourly forecast for a specified latitude and longitude
//	@Tags			weather
//	@Accept			json
//	@Produce		json
//	@Param			latitude	path		string	true	"Latitude"
//	@Param			longitude	path		string	true	"Longitude"
//	@Success		200			{object}	store.OpenMeteoResponse
//	@Failure		400			{object}	error
//	@Failure		404			{object}	error
//	@Failure		500			{object}	error
//	@Router			/weather/open-meteo/coords/{latitude}/{longitude} [get]
func (app *application) handleGetOpenMeteoByCoordinates(w http.ResponseWriter, r *http.Request) {
	latitude, longitude, err := parseCoordinates(r)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()

	openMeteo, err := app.cacheStorage.Weathers.GetOpenMeteo(ctx, latitude, longitude)
	if err != nil {
		app.logger.Warnw("error fetching open meteo from cache", "error", err)
	}

	if openMeteo == nil {
		openMeteo, err = app.store.Weathers.GetOpenMeteoByCoords(ctx, latitude, longitude)
		if err != nil {
			app.upstreamErrorResponse(w, r, err)
			return
		}

		if err := app.cacheStorage.Weathers.SetOpenMeteo(ctx, openMeteo, latitude, longitude); err != nil {
			app.logger.Warnw("error setting open meteo in cache", "error", err)
		}
	}

	if err := app.jsonResponse(w, http.StatusOK, openMeteo); err != nil {
		app.internalServerError(w, r, err)
	}
}

// parseCoordinates reads and validates the {latitude}/{longitude} URL params.
func parseCoordinates(r *http.Request) (float64, float64, error) {
	latitudeStr := chi.URLParam(r, "latitude")
	longitudeStr := chi.URLParam(r, "longitude")

	latitude, err := strconv.ParseFloat(latitudeStr, 64)
	if err != nil || latitude < -90 || latitude > 90 {
		return 0, 0, fmt.Errorf("invalid latitude: %s", latitudeStr)
	}

	longitude, err := strconv.ParseFloat(longitudeStr, 64)
	if err != nil || longitude < -180 || longitude > 180 {
		return 0, 0, fmt.Errorf("invalid longitude: %s", longitudeStr)
	}

	return latitude, longitude, nil
}

// upstreamErrorResponse maps weather-provider errors to HTTP responses.
func (app *application) upstreamErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	if errors.Is(err, store.ErrNotFound) {
		app.notFoundResponse(w, r, err)
		return
	}

	app.internalServerError(w, r, err)
}
