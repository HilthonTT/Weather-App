package weather

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/hilthontt/weather/internal/utils"
	"go.uber.org/zap"
)

type Handler struct {
	client Client
}

func NewHandler(client Client) *Handler {
	return &Handler{client}
}

func (h *Handler) RegisterRoutes(r chi.Router, logger *zap.SugaredLogger) {
	r.Get("/weather/{city}", func(w http.ResponseWriter, r *http.Request) {
		h.handleGetWeatherByCity(w, r, logger)
	})
	r.Get("/weather/coords/{latitude}/{longitude}", func(w http.ResponseWriter, r *http.Request) {
		h.handleGetWeatherByCoordinates(w, r, logger)
	})

	r.Get("/forecast/{city}", func(w http.ResponseWriter, r *http.Request) {
		h.handleGetForecast(w, r, logger)
	})
	r.Get("/forecast/coords/{latitude}/{longitude}", func(w http.ResponseWriter, r *http.Request) {
		h.handleGetForecastByCoordinates(w, r, logger)
	})
}

func (h *Handler) handleGetWeatherByCity(w http.ResponseWriter, r *http.Request, logger *zap.SugaredLogger) {
	city := chi.URLParam(r, "city")

	weather, err := h.client.GetWeather(city)
	if err != nil {
		utils.BadRequestResponse(w, r, err, logger)
		return
	}

	utils.JsonResponse(w, http.StatusOK, weather)
}

func (h *Handler) handleGetWeatherByCoordinates(w http.ResponseWriter, r *http.Request, logger *zap.SugaredLogger) {
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

	weather, err := h.client.GetWeatherByCoords(latitude, longitude)
	if err != nil {
		utils.InternalServerError(w, r, err, logger)
		return
	}

	utils.JsonResponse(w, http.StatusOK, weather)
}

func (h *Handler) handleGetForecast(w http.ResponseWriter, r *http.Request, logger *zap.SugaredLogger) {
	city := chi.URLParam(r, "city")

	weather, err := h.client.GetForecast(city)
	if err != nil {
		utils.InternalServerError(w, r, err, logger)
		return
	}

	utils.JsonResponse(w, http.StatusOK, weather)
}

func (h *Handler) handleGetForecastByCoordinates(w http.ResponseWriter, r *http.Request, logger *zap.SugaredLogger) {
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

	forecast, err := h.client.GetForecastByCoords(latitude, longitude)
	if err != nil {
		utils.InternalServerError(w, r, err, logger)
		return
	}

	utils.JsonResponse(w, http.StatusOK, forecast)
}
