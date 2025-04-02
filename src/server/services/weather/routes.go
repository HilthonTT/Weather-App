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
}

func NewHandler(client Client) *Handler {
	return &Handler{client}
}

func (h *Handler) RegisterRoutes(r chi.Router, logger *zap.SugaredLogger, weatherCache *WeatherCache) {
	r.Get("/weather/{city}", func(w http.ResponseWriter, r *http.Request) {
		h.handleGetWeatherByCity(w, r, logger, weatherCache)
	})
	r.Get("/weather/coords/{latitude}/{longitude}", func(w http.ResponseWriter, r *http.Request) {
		h.handleGetWeatherByCoordinates(w, r, logger, weatherCache)
	})

	r.Get("/weather/forecast/{city}", func(w http.ResponseWriter, r *http.Request) {
		h.handleGetForecast(w, r, logger, weatherCache)
	})
	r.Get("/weather/forecast/coords/{latitude}/{longitude}", func(w http.ResponseWriter, r *http.Request) {
		h.handleGetForecastByCoordinates(w, r, logger, weatherCache)
	})
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
func (h *Handler) handleGetWeatherByCity(
	w http.ResponseWriter,
	r *http.Request,
	logger *zap.SugaredLogger,
	weatherCache *WeatherCache,
) {
	city := chi.URLParam(r, "city")

	ctx := r.Context()
	weather, err := weatherCache.GetWeatherCity(ctx, city)
	if err != nil {
		logger.Errorf("Error fetching weather from cache: %v", err)
	}

	if weather != nil {
		utils.JsonResponse(w, http.StatusOK, weather)
		return
	}

	weather, err = h.client.GetWeather(city)
	if err != nil {
		utils.BadRequestResponse(w, r, err, logger)
		return
	}

	if err := weatherCache.SetWeather(ctx, weather); err != nil {
		logger.Errorf("Error setting weather in cache: %v", err)
	}

	utils.JsonResponse(w, http.StatusOK, weather)
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
func (h *Handler) handleGetWeatherByCoordinates(
	w http.ResponseWriter,
	r *http.Request,
	logger *zap.SugaredLogger,
	weatherCache *WeatherCache,
) {
	latitudeStr := chi.URLParam(r, "latitude")
	longitudeStr := chi.URLParam(r, "longitude")

	latitude, err := strconv.ParseFloat(latitudeStr, 64)
	if err != nil {
		utils.BadRequestResponse(w, r, fmt.Errorf("invalid latitude: %s", latitudeStr), logger)
		return
	}

	longitude, err := strconv.ParseFloat(longitudeStr, 64)
	if err != nil {
		utils.BadRequestResponse(w, r, fmt.Errorf("invalid longitude: %s", longitudeStr), logger)
		return
	}

	ctx := r.Context()
	weather, err := weatherCache.GetWeather(ctx, latitude, longitude)
	if err != nil {
		logger.Errorf("Error fetching weather from cache: %v", err)
	}

	if weather != nil {
		utils.JsonResponse(w, http.StatusOK, weather)
		return
	}

	weather, err = h.client.GetWeatherByCoords(latitude, longitude)
	if err != nil {
		utils.InternalServerError(w, r, err, logger)
		return
	}

	if err := weatherCache.SetWeather(ctx, weather); err != nil {
		logger.Errorf("Error setting weather in cache: %v", err)
	}

	utils.JsonResponse(w, http.StatusOK, weather)
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
func (h *Handler) handleGetForecast(
	w http.ResponseWriter,
	r *http.Request,
	logger *zap.SugaredLogger,
	weatherCache *WeatherCache,
) {
	city := chi.URLParam(r, "city")

	ctx := r.Context()
	forecast, err := weatherCache.GetForecastByCity(ctx, city)
	if err != nil {
		logger.Errorf("Error fetching forecast from cache: %v", err)
	}

	if forecast != nil {
		utils.JsonResponse(w, http.StatusOK, forecast)
		return
	}

	forecast, err = h.client.GetForecast(city)
	if err != nil {
		utils.InternalServerError(w, r, err, logger)
		return
	}

	if err := weatherCache.SetForecast(ctx, forecast); err != nil {
		logger.Errorf("Error setting forecast in cache: %v", err)
	}

	utils.JsonResponse(w, http.StatusOK, forecast)
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
func (h *Handler) handleGetForecastByCoordinates(
	w http.ResponseWriter,
	r *http.Request,
	logger *zap.SugaredLogger,
	weatherCache *WeatherCache,
) {
	latitudeStr := chi.URLParam(r, "latitude")
	longitudeStr := chi.URLParam(r, "longitude")

	latitude, err := strconv.ParseFloat(latitudeStr, 64)
	if err != nil {
		utils.BadRequestResponse(w, r, fmt.Errorf("invalid latitude: %s", latitudeStr), logger)
		return
	}

	longitude, err := strconv.ParseFloat(longitudeStr, 64)
	if err != nil {
		utils.BadRequestResponse(w, r, fmt.Errorf("invalid longitude: %s", longitudeStr), logger)
		return
	}

	ctx := r.Context()

	forecast, err := weatherCache.GetForecast(ctx, latitude, longitude)
	if err != nil {
		logger.Errorf("Error fetching forecast from cache: %v", err)
	}
	if forecast != nil {

		utils.JsonResponse(w, http.StatusOK, forecast)
		return
	}

	forecast, err = h.client.GetForecastByCoords(latitude, longitude)
	if err != nil {
		utils.InternalServerError(w, r, err, logger)
		return
	}

	if err := weatherCache.SetForecast(ctx, forecast); err != nil {
		logger.Errorf("Error setting forecast in cache: %v", err)
	}

	utils.JsonResponse(w, http.StatusOK, forecast)
}
