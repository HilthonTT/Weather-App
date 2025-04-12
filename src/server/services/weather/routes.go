package weather

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/hilthontt/weather/internal/utils"
	_ "github.com/hilthontt/weather/types"
	"go.uber.org/zap"
)

type Handler struct {
	client Client
	logger *zap.SugaredLogger
	cache  *WeatherCache
}

func NewHandler(client Client, logger *zap.SugaredLogger, cache *WeatherCache) *Handler {
	return &Handler{client, logger, cache}
}

func (h *Handler) RegisterRoutes(r chi.Router) {
	r.Get("/weather/{city}", h.handleGetWeatherByCity)
	r.Get("/weather/coords/{latitude}/{longitude}", h.handleGetWeatherByCoordinates)

	r.Get("/weather/forecast/{city}", h.handleGetForecast)
	r.Get("/weather/forecast/coords/{latitude}/{longitude}", h.handleGetForecastByCoordinates)

	r.Get("/weather/open-meteo/coords/{latitude}/{longitude}", h.handleGetOpenMeteoByCoordinates)
}

// GetWeatherByCity godoc
//
//	@Summary		Fetches current weather by city
//	@Description	Fetches the current weather data for a specified city
//	@Tags			weather
//	@Accept			json
//	@Produce		json
//	@Param			city	path		string	true	"City Name"
//	@Success		200		{object}	types.WeatherResponse
//	@Failure		400		{object}	error
//	@Failure		404		{object}	error
//	@Failure		500		{object}	error
//	@Router			/weather/{city} [get]
func (h *Handler) handleGetWeatherByCity(w http.ResponseWriter, r *http.Request) {
	city := chi.URLParam(r, "city")

	ctx := r.Context()
	weather, err := h.cache.GetWeatherCity(ctx, city)
	if err != nil {
		h.logger.Errorf("Error fetching weather from cache: %v", err)
	}

	if weather != nil {
		utils.JsonResponse(w, http.StatusOK, weather)
		return
	}

	weather, err = h.client.GetWeather(city)
	if err != nil {
		utils.BadRequestResponse(w, r, err, h.logger)
		return
	}

	if err := h.cache.SetWeather(ctx, weather); err != nil {
		h.logger.Errorf("Error setting weather in cache: %v", err)
	}

	if err := utils.JsonResponse(w, http.StatusOK, weather); err != nil {
		utils.InternalServerError(w, r, err, h.logger)
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
//	@Success		200			{object}	types.WeatherResponse
//	@Failure		400			{object}	error
//	@Failure		404			{object}	error
//	@Failure		500			{object}	error
//	@Router			/weather/coords/{latitude}/{longitude} [get]
func (h *Handler) handleGetWeatherByCoordinates(w http.ResponseWriter, r *http.Request) {
	latitudeStr := chi.URLParam(r, "latitude")
	longitudeStr := chi.URLParam(r, "longitude")

	latitude, err := strconv.ParseFloat(latitudeStr, 64)
	if err != nil {
		utils.BadRequestResponse(w, r, fmt.Errorf("invalid latitude: %s", latitudeStr), h.logger)
		return
	}

	longitude, err := strconv.ParseFloat(longitudeStr, 64)
	if err != nil {
		utils.BadRequestResponse(w, r, fmt.Errorf("invalid longitude: %s", longitudeStr), h.logger)
		return
	}

	ctx := r.Context()
	weather, err := h.cache.GetWeather(ctx, latitude, longitude)
	if err != nil {
		h.logger.Errorf("Error fetching weather from cache: %v", err)
	}

	if weather != nil {
		utils.JsonResponse(w, http.StatusOK, weather)
		return
	}

	weather, err = h.client.GetWeatherByCoords(latitude, longitude)
	if err != nil {
		utils.InternalServerError(w, r, err, h.logger)
		return
	}

	if err := h.cache.SetWeather(ctx, weather); err != nil {
		h.logger.Errorf("Error setting weather in cache: %v", err)
	}

	if err := utils.JsonResponse(w, http.StatusOK, weather); err != nil {
		utils.InternalServerError(w, r, err, h.logger)
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
//	@Success		200		{object}	types.ForecastResponse
//	@Failure		400		{object}	error
//	@Failure		404		{object}	error
//	@Failure		500		{object}	error
//	@Router			/weather/forecast/{city} [get]
func (h *Handler) handleGetForecast(w http.ResponseWriter, r *http.Request) {
	city := chi.URLParam(r, "city")

	ctx := r.Context()
	forecast, err := h.cache.GetForecastByCity(ctx, city)
	if err != nil {
		h.logger.Errorf("Error fetching forecast from cache: %v", err)
	}

	if forecast != nil {
		utils.JsonResponse(w, http.StatusOK, forecast)
		return
	}

	forecast, err = h.client.GetForecast(city)
	if err != nil {
		utils.InternalServerError(w, r, err, h.logger)
		return
	}

	if err := h.cache.SetForecast(ctx, forecast); err != nil {
		h.logger.Errorf("Error setting forecast in cache: %v", err)
	}

	if err := utils.JsonResponse(w, http.StatusOK, forecast); err != nil {
		utils.InternalServerError(w, r, err, h.logger)
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
//	@Success		200			{object}	types.ForecastResponse
//	@Failure		400			{object}	error
//	@Failure		404			{object}	error
//	@Failure		500			{object}	error
//	@Router			/weather/forecast/coords/{latitude}/{longitude} [get]
func (h *Handler) handleGetForecastByCoordinates(w http.ResponseWriter, r *http.Request) {
	latitudeStr := chi.URLParam(r, "latitude")
	longitudeStr := chi.URLParam(r, "longitude")

	latitude, err := strconv.ParseFloat(latitudeStr, 64)
	if err != nil {
		utils.BadRequestResponse(w, r, fmt.Errorf("invalid latitude: %s", latitudeStr), h.logger)
		return
	}

	longitude, err := strconv.ParseFloat(longitudeStr, 64)
	if err != nil {
		utils.BadRequestResponse(w, r, fmt.Errorf("invalid longitude: %s", longitudeStr), h.logger)
		return
	}

	ctx := r.Context()

	forecast, err := h.cache.GetForecast(ctx, latitude, longitude)
	if err != nil {
		h.logger.Errorf("Error fetching forecast from cache: %v", err)
	}

	if forecast != nil {
		utils.JsonResponse(w, http.StatusOK, forecast)
		return
	}

	forecast, err = h.client.GetForecastByCoords(latitude, longitude)
	if err != nil {
		utils.InternalServerError(w, r, err, h.logger)
		return
	}

	if err := h.cache.SetForecast(ctx, forecast); err != nil {
		h.logger.Errorf("Error setting forecast in cache: %v", err)
	}

	if err := utils.JsonResponse(w, http.StatusOK, forecast); err != nil {
		utils.InternalServerError(w, r, err, h.logger)
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
//	@Success		200			{object}	types.OpenMeteoResponse
//	@Failure		400			{object}	error
//	@Failure		404			{object}	error
//	@Failure		500			{object}	error
//	@Router			/weather/open-meteo/coords/{latitude}/{longitude} [get]
func (h *Handler) handleGetOpenMeteoByCoordinates(w http.ResponseWriter, r *http.Request) {
	latitudeStr := chi.URLParam(r, "latitude")
	longitudeStr := chi.URLParam(r, "longitude")

	latitude, err := strconv.ParseFloat(latitudeStr, 64)
	if err != nil {
		utils.BadRequestResponse(w, r, fmt.Errorf("invalid latitude: %s", latitudeStr), h.logger)
		return
	}

	longitude, err := strconv.ParseFloat(longitudeStr, 64)
	if err != nil {
		utils.BadRequestResponse(w, r, fmt.Errorf("invalid longitude: %s", longitudeStr), h.logger)
		return
	}

	ctx := r.Context()

	openMeteo, err := h.cache.GetOpenMeteo(ctx, latitude, longitude)
	if err != nil {
		h.logger.Errorf("Error fetching open meteo from cache: %v", err)
	}

	if openMeteo != nil {
		utils.JsonResponse(w, http.StatusOK, openMeteo)
		return
	}

	openMeteo, err = h.client.GetOpenMeteoByCoords(latitude, longitude)
	if err != nil {
		utils.InternalServerError(w, r, err, h.logger)
		return
	}

	if err := h.cache.SetOpenMeteo(ctx, openMeteo, latitude, longitude); err != nil {
		h.logger.Errorf("Error setting forecast in cache: %v", err)
	}

	if err := utils.JsonResponse(w, http.StatusOK, openMeteo); err != nil {
		utils.InternalServerError(w, r, err, h.logger)
	}
}
