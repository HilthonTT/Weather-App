package weather

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hilthontt/weather/types"
)

type Client struct {
	APIKey string
	HTTP   *http.Client
}

func NewClient(apiKey string) *Client {
	return &Client{
		APIKey: apiKey,
		HTTP:   &http.Client{},
	}
}

func (c *Client) GetWeather(city string) (*types.WeatherResponse, error) {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, c.APIKey)

	return c.fetchWeather(url)
}

func (c *Client) GetWeatherByCoords(lat, lon float64) (*types.WeatherResponse, error) {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&appid=%s&units=metric", lat, lon, c.APIKey)

	return c.fetchWeather(url)
}

func (c *Client) GetForecast(city string) (*types.ForecastResponse, error) {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/forecast?q=%s&appid=%s&units=metric", city, c.APIKey)

	return c.fetchForecast(url)
}

func (c *Client) GetForecastByCoords(lat, lon float64) (*types.ForecastResponse, error) {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/forecast?lat=%f&lon=%f&appid=%s&units=metric", lat, lon, c.APIKey)

	return c.fetchForecast(url)
}

func (c *Client) fetchWeather(url string) (*types.WeatherResponse, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.HTTP.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch weather: %s", resp.Status)
	}

	var weatherData types.WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weatherData); err != nil {
		return nil, err
	}

	return &weatherData, nil
}

func (c *Client) fetchForecast(url string) (*types.ForecastResponse, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.HTTP.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch weather: %s", resp.Status)
	}

	var forecastData types.ForecastResponse
	if err := json.NewDecoder(resp.Body).Decode(&forecastData); err != nil {
		return nil, err
	}

	return &forecastData, nil
}
