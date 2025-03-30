package weather

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DbName           = "weather"
	ForecastCollName = "Forecasts"
	WeatherCollName  = "Weathers"
)

type Store struct {
	db *mongo.Client
}

func NewStore(db *mongo.Client) *Store {
	return &Store{db}
}

func (s *Store) CreateWeather(ctx context.Context, w WeatherModel) (primitive.ObjectID, error) {
	col := s.db.Database(DbName).Collection(WeatherCollName)

	newWeather, err := col.InsertOne(ctx, w)

	id := newWeather.InsertedID.(primitive.ObjectID)

	return id, err
}

func (s *Store) GetWeather(ctx context.Context, id string) (*WeatherModel, error) {
	col := s.db.Database(DbName).Collection(WeatherCollName)

	oID, _ := primitive.ObjectIDFromHex(id)

	var w WeatherModel
	err := col.FindOne(ctx, bson.M{
		"_id": oID,
	}).Decode(&w)

	return &w, err
}

func (s *Store) GetWeatherByCoords(ctx context.Context, lat, lon float64) (*WeatherModel, error) {
	col := s.db.Database(DbName).Collection(WeatherCollName)

	var w WeatherModel
	err := col.FindOne(ctx, bson.M{
		"coord.lat": lat,
		"coord.lon": lon,
	}).Decode(&w)

	return &w, err
}

func (s *Store) GetWeatherByCityName(ctx context.Context, cityName string) (*WeatherModel, error) {
	col := s.db.Database(DbName).Collection(WeatherCollName)

	var w WeatherModel
	err := col.FindOne(ctx, bson.M{
		"name": cityName,
	}).Decode(&w)

	return &w, err
}

func (s *Store) CreateForecast(ctx context.Context, w ForecastModel) (primitive.ObjectID, error) {
	col := s.db.Database(DbName).Collection(WeatherCollName)

	newWeather, err := col.InsertOne(ctx, w)

	id := newWeather.InsertedID.(primitive.ObjectID)

	return id, err
}

func (s *Store) GetForecast(ctx context.Context, id string) (*ForecastModel, error) {
	col := s.db.Database(DbName).Collection(WeatherCollName)

	oID, _ := primitive.ObjectIDFromHex(id)

	var w ForecastModel
	err := col.FindOne(ctx, bson.M{
		"_id": oID,
	}).Decode(&w)

	return &w, err
}

func (s *Store) GetForecastByCoords(ctx context.Context, lat, lon float64) (*ForecastModel, error) {
	col := s.db.Database(DbName).Collection(WeatherCollName)

	var f ForecastModel
	err := col.FindOne(ctx, bson.M{
		"city.coord.lat": lat,
		"city.coord.lon": lon,
	}).Decode(&f)

	return &f, err
}

func (s *Store) GetForecastByCityName(ctx context.Context, cityName string) (*ForecastModel, error) {
	col := s.db.Database(DbName).Collection(WeatherCollName)

	var f ForecastModel
	err := col.FindOne(ctx, bson.M{
		"city.name": cityName,
	}).Decode(&f)

	return &f, err
}
