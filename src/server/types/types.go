package types

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
