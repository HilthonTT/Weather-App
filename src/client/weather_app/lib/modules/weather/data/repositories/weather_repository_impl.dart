import 'package:flutter/foundation.dart';
import 'package:fpdart/fpdart.dart';
import 'package:weather_app/common/errors/exceptions.dart';
import 'package:weather_app/common/errors/failure.dart';
import 'package:weather_app/modules/weather/data/datasources/weather_remote_datasource.dart';
import 'package:weather_app/modules/weather/domain/entities/forecast.dart';
import 'package:weather_app/modules/weather/domain/entities/open_meteo.dart';
import 'package:weather_app/modules/weather/domain/entities/weather.dart';
import 'package:weather_app/modules/weather/domain/repositories/weather_repository.dart';

@immutable
final class WeatherRepositoryImpl implements WeatherRepository {
  final WeatherRemoteDatasource weatherRemoteDatasource;

  const WeatherRepositoryImpl({required this.weatherRemoteDatasource});

  @override
  Future<Either<Failure, ForecastResponse>> getForecast() async {
    try {
      final forecast = await weatherRemoteDatasource.getForecast();

      return right(forecast);
    } on ServerException catch (e) {
      return left(Failure(e.message));
    }
  }

  @override
  Future<Either<Failure, WeatherResponse>> getWeather() async {
    try {
      final weather = await weatherRemoteDatasource.getWeather();

      return right(weather);
    } on ServerException catch (e) {
      return left(Failure(e.message));
    }
  }

  @override
  Future<Either<Failure, ForecastResponse>> getForecastByCity(
    String city,
  ) async {
    try {
      final forecast = await weatherRemoteDatasource.getForecastByCity(city);

      return right(forecast);
    } on ServerException catch (e) {
      return left(Failure(e.message));
    }
  }

  @override
  Future<Either<Failure, WeatherResponse>> getWeatherByCity(String city) async {
    try {
      final weather = await weatherRemoteDatasource.getWeatherByCity(city);

      return right(weather);
    } on ServerException catch (e) {
      return left(Failure(e.message));
    }
  }

  @override
  Future<Either<Failure, OpenMeteoResponse>> getOpenMeteo() async {
    try {
      final openMeteo = await weatherRemoteDatasource.getOpenMeteo();

      return right(openMeteo);
    } on ServerException catch (e) {
      return left(Failure(e.message));
    }
  }
}
