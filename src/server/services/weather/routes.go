package weather

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/hilthontt/weather/internal/utils"
)

type Handler struct {
	client Client
}

func NewHandler(client Client) *Handler {
	return &Handler{client}
}

func (h *Handler) RegisterRoutes(r chi.Router) {
	r.Get("/weather/{city}", h.handleGetWeatherByCity)
	r.Get("/weather/coords/{latitude}/{longitude}", h.handleGetWeatherByCoordinates)

	r.Get("/forecast/{city}", h.handleGetForecast)
	r.Get("/forecast/coords/{latitude}/{longitude}", h.handleGetForecastByCoordinates)
}

func (h *Handler) handleGetWeatherByCity(w http.ResponseWriter, r *http.Request) {
	city := chi.URLParam(r, "city")

	weather, err := h.client.GetWeather(city)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	utils.JsonResponse(w, http.StatusOK, weather)
}

func (h *Handler) handleGetWeatherByCoordinates(w http.ResponseWriter, r *http.Request) {
	latitudeStr := chi.URLParam(r, "latitude")
	longitudeStr := chi.URLParam(r, "longitude")

	latitude, err := strconv.ParseFloat(latitudeStr, 64)
	if err != nil {
		utils.BadRequestResponse(w, r, fmt.Errorf("invalid latitude: %s", latitudeStr))
		return
	}

	longitude, err := strconv.ParseFloat(longitudeStr, 64)
	if err != nil {
		utils.BadRequestResponse(w, r, fmt.Errorf("invalid longitude: %s", longitudeStr))
		return
	}

	weather, err := h.client.GetWeatherByCoords(latitude, longitude)
	if err != nil {
		utils.InternalServerError(w, r, err)
		return
	}

	utils.JsonResponse(w, http.StatusOK, weather)
}

func (h *Handler) handleGetForecast(w http.ResponseWriter, r *http.Request) {
	city := chi.URLParam(r, "city")

	weather, err := h.client.GetForecast(city)
	if err != nil {
		utils.InternalServerError(w, r, err)
		return
	}

	utils.JsonResponse(w, http.StatusOK, weather)
}

func (h *Handler) handleGetForecastByCoordinates(w http.ResponseWriter, r *http.Request) {
	latitudeStr := chi.URLParam(r, "latitude")
	longitudeStr := chi.URLParam(r, "longitude")

	latitude, err := strconv.ParseFloat(latitudeStr, 64)
	if err != nil {
		utils.BadRequestResponse(w, r, fmt.Errorf("invalid latitude: %s", latitudeStr))
		return
	}

	longitude, err := strconv.ParseFloat(longitudeStr, 64)
	if err != nil {
		utils.BadRequestResponse(w, r, fmt.Errorf("invalid longitude: %s", longitudeStr))
		return
	}

	forecast, err := h.client.GetForecastByCoords(latitude, longitude)
	if err != nil {
		utils.InternalServerError(w, r, err)
		return
	}

	utils.JsonResponse(w, http.StatusOK, forecast)
}
