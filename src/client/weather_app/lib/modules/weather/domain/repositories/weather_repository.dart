import 'package:fpdart/fpdart.dart';
import 'package:weather_app/common/errors/failure.dart';
import 'package:weather_app/modules/weather/domain/entities/forecast.dart';
import 'package:weather_app/modules/weather/domain/entities/weather.dart';

abstract interface class WeatherRepository {
  Future<Either<Failure, WeatherResponse>> getWeather();

  Future<Either<Failure, WeatherResponse>> getWeatherByCity(String city);

  Future<Either<Failure, ForecastResponse>> getForecast();

  Future<Either<Failure, ForecastResponse>> getForecastByCity(String city);
}
