package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// GetWeatherByCity godoc
//
//	@Summary		Fetches current weather by city
//	@Description	Fetches the current weather data for a specified city
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
		app.logger.Errorf("Error fetching weather from cache: %v", err)
	}

	if weather != nil {
		app.jsonResponse(w, http.StatusOK, weather)
		return
	}

	weather, err = app.store.Weathers.GetWeather(city)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := app.cacheStorage.Weathers.SetWeather(ctx, weather); err != nil {
		app.logger.Errorf("Error setting weather in cache: %v", err)
	}

	if err := app.jsonResponse(w, http.StatusOK, weather); err != nil {
		app.internalServerError(w, r, err)
	}
}

// GetWeatherByCoordinates godoc
//
//	@Summary		Fetches current weather by coordinates (latitude and longitude)
//	@Description	Fetches the current weather data for a specified latitude and longitude
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
	latitudeStr := chi.URLParam(r, "latitude")
	longitudeStr := chi.URLParam(r, "longitude")

	latitude, err := strconv.ParseFloat(latitudeStr, 64)
	if err != nil {
		app.badRequestResponse(w, r, fmt.Errorf("invalid latitude: %s", latitudeStr))
		return
	}

	longitude, err := strconv.ParseFloat(longitudeStr, 64)
	if err != nil {
		app.badRequestResponse(w, r, fmt.Errorf("invalid longitude: %s", longitudeStr))
		return
	}

	ctx := r.Context()
	weather, err := app.cacheStorage.Weathers.GetWeather(ctx, latitude, longitude)
	if err != nil {
		app.logger.Errorf("Error fetching weather from cache: %v", err)
	}

	if weather != nil {
		app.jsonResponse(w, http.StatusOK, weather)
		return
	}

	weather, err = app.store.Weathers.GetWeatherByCoords(latitude, longitude)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.cacheStorage.Weathers.SetWeather(ctx, weather); err != nil {
		app.logger.Errorf("Error setting weather in cache: %v", err)
	}

	if err := app.jsonResponse(w, http.StatusOK, weather); err != nil {
		app.internalServerError(w, r, err)
	}
}

// GetForecastByCity godoc
//
//	@Summary		Fetches weather forecast by city
//	@Description	Fetches the weather forecast data for a specified city
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
		app.logger.Errorf("Error fetching forecast from cache: %v", err)
	}

	if forecast != nil {
		app.jsonResponse(w, http.StatusOK, forecast)
		return
	}

	forecast, err = app.store.Weathers.GetForecast(city)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.cacheStorage.Weathers.SetForecast(ctx, forecast); err != nil {
		app.logger.Errorf("Error setting forecast in cache: %v", err)
	}

	if err := app.jsonResponse(w, http.StatusOK, forecast); err != nil {
		app.internalServerError(w, r, err)
	}
}

// GetForecastByCoordinates godoc
//
//	@Summary		Fetches weather forecast by coordinates (latitude and longitude)
//	@Description	Fetches the weather forecast data for a specified latitude and longitude
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
	latitudeStr := chi.URLParam(r, "latitude")
	longitudeStr := chi.URLParam(r, "longitude")

	latitude, err := strconv.ParseFloat(latitudeStr, 64)
	if err != nil {
		app.badRequestResponse(w, r, fmt.Errorf("invalid latitude: %s", latitudeStr))
		return
	}

	longitude, err := strconv.ParseFloat(longitudeStr, 64)
	if err != nil {
		app.badRequestResponse(w, r, fmt.Errorf("invalid longitude: %s", longitudeStr))
		return
	}

	ctx := r.Context()

	forecast, err := app.cacheStorage.Weathers.GetForecast(ctx, latitude, longitude)
	if err != nil {
		app.logger.Errorf("Error fetching forecast from cache: %v", err)
	}

	if forecast != nil {
		app.jsonResponse(w, http.StatusOK, forecast)
		return
	}

	forecast, err = app.store.Weathers.GetForecastByCoords(latitude, longitude)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.cacheStorage.Weathers.SetForecast(ctx, forecast); err != nil {
		app.logger.Errorf("Error setting forecast in cache: %v", err)
	}

	if err := app.jsonResponse(w, http.StatusOK, forecast); err != nil {
		app.internalServerError(w, r, err)
	}
}

// GetOpenMeteoByCoordinates godoc
//
//	@Summary		Fetches open meteo by coordinates (latitude and longitude)
//	@Description	Fetches the open meteo data for a specified latitude and longitude
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
	latitudeStr := chi.URLParam(r, "latitude")
	longitudeStr := chi.URLParam(r, "longitude")

	latitude, err := strconv.ParseFloat(latitudeStr, 64)
	if err != nil {
		app.badRequestResponse(w, r, fmt.Errorf("invalid latitude: %s", latitudeStr))
		return
	}

	longitude, err := strconv.ParseFloat(longitudeStr, 64)
	if err != nil {
		app.badRequestResponse(w, r, fmt.Errorf("invalid longitude: %s", longitudeStr))
		return
	}

	ctx := r.Context()

	openMeteo, err := app.cacheStorage.Weathers.GetOpenMeteo(ctx, latitude, longitude)
	if err != nil {
		app.logger.Errorf("Error fetching open meteo from cache: %v", err)
	}

	if openMeteo != nil {
		app.jsonResponse(w, http.StatusOK, openMeteo)
		return
	}

	openMeteo, err = app.store.Weathers.GetOpenMeteoByCoords(latitude, longitude)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.cacheStorage.Weathers.SetOpenMeteo(ctx, openMeteo, latitude, longitude); err != nil {
		app.logger.Errorf("Error setting forecast in cache: %v", err)
	}

	if err := app.jsonResponse(w, http.StatusOK, openMeteo); err != nil {
		app.internalServerError(w, r, err)
	}
}
