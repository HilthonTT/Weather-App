package settings

type TemperatureUnit string
type WindSpeedUnit string
type TimeFormat string

const (
	Celsius    TemperatureUnit = "Celsius"
	Fahrenheit TemperatureUnit = "Fahrenheit"

	Kmh     WindSpeedUnit = "Km/h"
	Milesph WindSpeedUnit = "Miles/h"

	H12 TimeFormat = "12h"
	H24 TimeFormat = "24h"
)

type Settings struct {
	ID          int64           `json:"id"`
	UserID      int64           `json:"user_id"`
	Temperature TemperatureUnit `json:"temperature"`
	WindSpeed   WindSpeedUnit   `json:"wind_speed"`
	TimeDisplay TimeFormat      `json:"time_display"`
}
