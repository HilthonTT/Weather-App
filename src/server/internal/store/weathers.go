package store

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type WeatherResponse struct {
	Coord   Coord     `json:"coord"`
	Weather []Weather `json:"weather"`
	Main    struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
		SeaLevel  int     `json:"sea_level"`
		GrndLevel int     `json:"grnd_level"`
	} `json:"main"`
	Visibility int    `json:"visibility"`
	Wind       Wind   `json:"wind"`
	Clouds     Clouds `json:"clouds"`
	Dt         int64  `json:"dt"` // Timestamp
	Sys        struct {
		Country string `json:"country"`
		Sunrise int64  `json:"sunrise"`
		Sunset  int64  `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}

type ForecastResponse struct {
	Cod     string         `json:"cod"`
	Message int            `json:"message"`
	Cnt     int            `json:"cnt"`
	List    []ForecastItem `json:"list"`
	City    City           `json:"city"`
}

type ForecastItem struct {
	Dt         int64     `json:"dt"`
	Main       Main      `json:"main"`
	Weather    []Weather `json:"weather"`
	Clouds     Clouds    `json:"clouds"`
	Wind       Wind      `json:"wind"`
	Visibility int       `json:"visibility"`
	Pop        float64   `json:"pop"`
	Sys        Sys       `json:"sys"`
	DtTxt      string    `json:"dt_txt"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int     `json:"pressure"`
	SeaLevel  int     `json:"sea_level"`
	GrndLevel int     `json:"grnd_level"`
	Humidity  int     `json:"humidity"`
	TempKf    float64 `json:"temp_kf"`
}

type Weather struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Clouds struct {
	All int `json:"all"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int     `json:"deg"`
	Gust  float64 `json:"gust"`
}

type Sys struct {
	Pod string `json:"pod"`
}

type City struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Coord      Coord  `json:"coord"`
	Country    string `json:"country"`
	Population int    `json:"population"`
	Timezone   int    `json:"timezone"`
	Sunrise    int64  `json:"sunrise"`
	Sunset     int64  `json:"sunset"`
}

type Coord struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type OpenMeteoResponse struct {
	Latitude             float64        `json:"latitude"`
	Longitude            float64        `json:"longitude"`
	GenerationTimeMs     float64        `json:"generationtime_ms"`
	UTCOffsetSeconds     int            `json:"utc_offset_seconds"`
	Timezone             string         `json:"timezone"`
	TimezoneAbbreviation string         `json:"timezone_abbreviation"`
	Elevation            float64        `json:"elevation"`
	HourlyUnits          HourlyUnits    `json:"hourly_units"`
	Hourly               HourlyForecast `json:"hourly"`
	DailyUnits           DailyUnits     `json:"daily_units"`
	Daily                DailyForecast  `json:"daily"`
}

type HourlyUnits struct {
	Time          string `json:"time"`
	Temperature2m string `json:"temperature_2m"`
}

type HourlyForecast struct {
	Time          []string  `json:"time"`
	Temperature2m []float64 `json:"temperature_2m"`
}

type DailyUnits struct {
	Time            string `json:"time"`
	WeatherCode     string `json:"weather_code"`
	ApparentTempMax string `json:"apparent_temperature_max"`
	ApparentTempMin string `json:"apparent_temperature_min"`
}

type DailyForecast struct {
	Time            []string  `json:"time"`
	WeatherCode     []int     `json:"weather_code"`
	ApparentTempMax []float64 `json:"apparent_temperature_max"`
	ApparentTempMin []float64 `json:"apparent_temperature_min"`
}

// WeatherStore fetches data from the OpenWeather and Open-Meteo HTTP APIs.
// All temperatures are returned in Celsius and wind speeds in m/s.
type WeatherStore struct {
	APIKey string
	HTTP   *http.Client
}

const (
	openWeatherURL         = "https://api.openweathermap.org/data/2.5/weather"
	openWeatherForecastURL = "https://api.openweathermap.org/data/2.5/forecast"
	openMeteoURL           = "https://api.open-meteo.com/v1/forecast"

	// fetchRetries is the number of attempts made for transient upstream failures.
	fetchRetries = 3
)

// ErrMissingAPIKey is returned when OPEN_WEATHER_API_KEY is not configured.
var ErrMissingAPIKey = errors.New("OpenWeather API key is required")

// retryBackoff is the base delay between attempts; a variable so tests can
// shorten it.
var retryBackoff = time.Second

func (s *WeatherStore) GetWeather(ctx context.Context, city string) (*WeatherResponse, error) {
	if s.APIKey == "" {
		return nil, ErrMissingAPIKey
	}

	url := buildURL(openWeatherURL, map[string]string{
		"q":     city,
		"appid": s.APIKey,
		"units": "metric",
	})

	return fetchJSON[WeatherResponse](ctx, s.HTTP, url, fetchRetries)
}

func (s *WeatherStore) GetWeatherByCoords(ctx context.Context, lat, lon float64) (*WeatherResponse, error) {
	if s.APIKey == "" {
		return nil, ErrMissingAPIKey
	}

	url := buildURL(openWeatherURL, map[string]string{
		"lat":   formatCoord(lat),
		"lon":   formatCoord(lon),
		"appid": s.APIKey,
		"units": "metric",
	})

	return fetchJSON[WeatherResponse](ctx, s.HTTP, url, fetchRetries)
}

func (s *WeatherStore) GetForecast(ctx context.Context, city string) (*ForecastResponse, error) {
	if s.APIKey == "" {
		return nil, ErrMissingAPIKey
	}

	url := buildURL(openWeatherForecastURL, map[string]string{
		"q":     city,
		"appid": s.APIKey,
		"units": "metric",
	})

	return fetchJSON[ForecastResponse](ctx, s.HTTP, url, fetchRetries)
}

func (s *WeatherStore) GetForecastByCoords(ctx context.Context, lat, lon float64) (*ForecastResponse, error) {
	if s.APIKey == "" {
		return nil, ErrMissingAPIKey
	}

	url := buildURL(openWeatherForecastURL, map[string]string{
		"lat":   formatCoord(lat),
		"lon":   formatCoord(lon),
		"appid": s.APIKey,
		"units": "metric",
	})

	return fetchJSON[ForecastResponse](ctx, s.HTTP, url, fetchRetries)
}

// GetOpenMeteoByCoords does not need an API key.
func (s *WeatherStore) GetOpenMeteoByCoords(ctx context.Context, lat, lon float64) (*OpenMeteoResponse, error) {
	url := buildURL(openMeteoURL, map[string]string{
		"latitude":  formatCoord(lat),
		"longitude": formatCoord(lon),
		"daily":     "weather_code,apparent_temperature_max,apparent_temperature_min",
		"hourly":    "temperature_2m",
	})

	return fetchJSON[OpenMeteoResponse](ctx, s.HTTP, url, fetchRetries)
}

// fetchJSON performs a GET request and decodes a JSON body into T.
//
// 5xx responses and network errors are retried with a linear backoff.
// A 404 is mapped to ErrNotFound and other 4xx responses fail immediately,
// since retrying them cannot succeed.
func fetchJSON[T any](ctx context.Context, client *http.Client, url string, retries int) (*T, error) {
	var lastErr error

	for attempt := range retries {
		if attempt > 0 {
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			case <-time.After(time.Duration(attempt) * retryBackoff):
			}
		}

		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			return nil, err
		}
		req.Header.Set("Accept", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			lastErr = err
			continue
		}

		data, err := decodeResponse[T](resp)
		if err == nil {
			return data, nil
		}

		lastErr = err
		if resp.StatusCode >= 400 && resp.StatusCode < 500 {
			return nil, err
		}
	}

	return nil, fmt.Errorf("request failed after %d attempts: %w", retries, lastErr)
}

func decodeResponse[T any](resp *http.Response) (*T, error) {
	defer resp.Body.Close()

	switch {
	case resp.StatusCode == http.StatusNotFound:
		return nil, ErrNotFound
	case resp.StatusCode != http.StatusOK:
		return nil, fmt.Errorf("upstream returned status %d", resp.StatusCode)
	}

	var data T
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("failed to decode JSON: %w", err)
	}

	return &data, nil
}

func formatCoord(v float64) string {
	return strconv.FormatFloat(v, 'f', -1, 64)
}

func buildURL(base string, params map[string]string) string {
	u, _ := url.Parse(base)
	q := u.Query()
	for k, v := range params {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()
	return u.String()
}
