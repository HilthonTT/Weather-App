package weather

import "go.mongodb.org/mongo-driver/bson/primitive"

type WeatherModel struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Coord   Coord              `bson:"coord,omit_empty"`
	Weather []Weather          `bson:"weather,omit_empty"`
	Main    struct {
		Temp      float64 `bson:"temp,omit_empty"`
		FeelsLike float64 `bson:"feels_like,omit_empty"`
		TempMin   float64 `bson:"temp_min,omit_empty"`
		TempMax   float64 `bson:"temp_max,omit_empty"`
		Pressure  int     `bson:"pressure,omit_empty"`
		Humidity  int     `bson:"humidity,omit_empty"`
		SeaLevel  int     `bson:"sea_level,omit_empty"`
		GrndLevel int     `bson:"grnd_level,omit_empty"`
	} `bson:"main,omit_empty"`
	Visibility int    `bson:"visibility,omit_empty"`
	Wind       Wind   `bson:"wind,omit_empty"`
	Clouds     Clouds `bson:"clouds,omit_empty"`
	Dt         int64  `bson:"dt,omit_empty"` // Timestamp
	Sys        struct {
		Country string `bson:"country,omit_empty"`
		Sunrise int64  `bson:"sunrise,omit_empty"`
		Sunset  int64  `bson:"sunset,omit_empty"`
	} `bson:"sys,omit_empty"`
	Timezone int    `bson:"timezone,omit_empty"`
	Name     string `bson:"name,omit_empty"`
	Cod      int    `bson:"cod,omit_empty"`
}

type ForecastModel struct {
	Cod     string         `bson:"cod,omit_empty"`
	Message int            `bson:"message,omit_empty"`
	Cnt     int            `bson:"cnt,omit_empty"`
	List    []ForecastItem `bson:"list,omit_empty"`
	City    City           `bson:"city,omit_empty"`
}

type ForecastItem struct {
	Dt         int64     `bson:"dt,omit_empty"`
	Main       Main      `bson:"main,omit_empty"`
	Weather    []Weather `bson:"weather,omit_empty"`
	Clouds     Clouds    `bson:"clouds,omit_empty"`
	Wind       Wind      `bson:"wind,omit_empty"`
	Visibility int       `bson:"visibility,omit_empty"`
	Pop        float64   `bson:"pop,omit_empty"`
	Sys        Sys       `bson:"sys,omit_empty"`
	DtTxt      string    `bson:"dt_txt,omit_empty"`
}

type Main struct {
	Temp      float64 `bson:"temp,omit_empty"`
	FeelsLike float64 `bson:"feels_like,omit_empty"`
	TempMin   float64 `bson:"temp_min,omit_empty"`
	TempMax   float64 `bson:"temp_max,omit_empty"`
	Pressure  int     `bson:"pressure,omit_empty"`
	SeaLevel  int     `bson:"sea_level,omit_empty"`
	GrndLevel int     `bson:"grnd_level,omit_empty"`
	Humidity  int     `bson:"humidity,omit_empty"`
	TempKf    float64 `bson:"temp_kf,omit_empty"`
}

type Weather struct {
	ID          int    `bson:"id,omit_empty"`
	Main        string `bson:"main,omit_empty"`
	Description string `bson:"description,omit_empty"`
	Icon        string `bson:"icon,omit_empty"`
}

type Clouds struct {
	All int `bson:"all,omit_empty"`
}

type Wind struct {
	Speed float64 `bson:"speed,omit_empty"`
	Deg   int     `bson:"deg,omit_empty"`
	Gust  float64 `bson:"gust,omit_empty"`
}

type Sys struct {
	Pod string `bson:"pod,omit_empty"`
}

type City struct {
	ID         int    `bson:"id,omit_empty"`
	Name       string `bson:"name,omit_empty"`
	Coord      Coord  `bson:"coord,omit_empty"`
	Country    string `bson:"country,omit_empty"`
	Population int    `bson:"population,omit_empty"`
	Timezone   int    `bson:"timezone,omit_empty"`
	Sunrise    int64  `bson:"sunrise,omit_empty"`
	Sunset     int64  `bson:"sunset,omit_empty"`
}

type Coord struct {
	Lat float64 `bson:"lat,omit_empty"`
	Lon float64 `bson:"lon,omit_empty"`
}
