package weather

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/hilthontt/weather/types"
)

type Client struct {
	APIKey string
	HTTP   *http.Client
}

func NewClient(apiKey string) *Client {

	return &Client{
		APIKey: apiKey,
		HTTP: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *Client) GetWeather(city string) (*types.WeatherResponse, error) {
	if c.APIKey == "" {
		return nil, errors.New("OpenWeather API key is required")
	}

	url := buildURL("https://api.openweathermap.org/data/2.5/weather", map[string]string{
		"q":     city,
		"appid": c.APIKey,
		"units": "metric",
	})

	return c.fetchWeather(url)
}

func (c *Client) GetWeatherByCoords(lat, lon float64) (*types.WeatherResponse, error) {
	if c.APIKey == "" {
		return nil, errors.New("OpenWeather API key is required")
	}

	url := buildURL("https://api.openweathermap.org/data/2.5/weather", map[string]string{
		"lat":   fmt.Sprintf("%f", lat),
		"lon":   fmt.Sprintf("%f", lon),
		"appid": c.APIKey,
		"units": "metric",
	})

	return c.fetchWeather(url)
}

func (c *Client) GetForecast(city string) (*types.ForecastResponse, error) {
	url := buildURL("https://api.openweathermap.org/data/2.5/forecast", map[string]string{
		"q":     city,
		"appid": c.APIKey,
		"units": "metric",
	})

	return c.fetchForecast(url)
}

func (c *Client) GetForecastByCoords(lat, lon float64) (*types.ForecastResponse, error) {
	url := buildURL("https://api.openweathermap.org/data/2.5/forecast", map[string]string{
		"lat":   fmt.Sprintf("%f", lat),
		"lon":   fmt.Sprintf("%f", lon),
		"appid": c.APIKey,
		"units": "metric",
	})

	return c.fetchForecast(url)
}

func (c *Client) GetOpenMeteoByCoords(lat, lon float64) (*types.OpenMeteoResponse, error) {
	url := fmt.Sprintf(
		"https://api.open-meteo.com/v1/forecast?latitude=%f&longitude=%f&daily=weather_code,apparent_temperature_max,apparent_temperature_min&hourly=temperature_2m",
		lat,
		lon,
	)

	return c.fetchOpenMeteo(url)
}

func (c *Client) fetchWeather(url string) (*types.WeatherResponse, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")

	return fetchJSON[types.WeatherResponse](c.HTTP, req, 3)
}

func (c *Client) fetchForecast(url string) (*types.ForecastResponse, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")

	return fetchJSON[types.ForecastResponse](c.HTTP, req, 3)
}

func (c *Client) fetchOpenMeteo(url string) (*types.OpenMeteoResponse, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")

	return fetchJSON[types.OpenMeteoResponse](c.HTTP, req, 3)
}

func fetchJSON[T any](client *http.Client, req *http.Request, retries int) (*T, error) {
	var resp *http.Response
	var err error

	for attempt := range retries {
		resp, err = client.Do(req)
		if err == nil && resp.StatusCode == http.StatusOK {
			defer resp.Body.Close()

			var data T
			if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
				return nil, fmt.Errorf("failed to decode JSON: %w", err)
			}
			return &data, nil
		}

		if resp != nil {
			resp.Body.Close()
		}

		time.Sleep(time.Duration(attempt+1) * time.Second) // basic backoff
	}

	return nil, fmt.Errorf("request failed after %d attempts: %w", retries, err)
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
