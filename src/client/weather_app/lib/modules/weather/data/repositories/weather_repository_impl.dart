import 'package:flutter/foundation.dart';
import 'package:fpdart/fpdart.dart';
import 'package:weather_app/common/errors/exceptions.dart';
import 'package:weather_app/common/errors/failure.dart';
import 'package:weather_app/common/errors/general_errors.dart';
import 'package:weather_app/common/network/connection_checker.dart';
import 'package:weather_app/modules/weather/data/datasources/weather_remote_datasource.dart';
import 'package:weather_app/modules/weather/domain/entities/forecast.dart';
import 'package:weather_app/modules/weather/domain/entities/open_meteo.dart';
import 'package:weather_app/modules/weather/domain/entities/weather.dart';
import 'package:weather_app/modules/weather/domain/repositories/weather_repository.dart';

@immutable
final class WeatherRepositoryImpl implements WeatherRepository {
  final WeatherRemoteDatasource weatherRemoteDatasource;
  final ConnectionChecker connectionChecker;

  const WeatherRepositoryImpl({
    required this.weatherRemoteDatasource,
    required this.connectionChecker,
  });

  @override
  Future<Either<Failure, ForecastResponse>> getForecast() async {
    try {
      if (!await connectionChecker.isConnected) {
        return left(Failure(GeneralErrors.noInternetConnection));
      }

      final forecast = await weatherRemoteDatasource.getForecast();

      return right(forecast);
    } on ServerException catch (e) {
      return left(Failure(e.message));
    }
  }

  @override
  Future<Either<Failure, WeatherResponse>> getWeather() async {
    try {
      if (!await connectionChecker.isConnected) {
        return left(Failure(GeneralErrors.noInternetConnection));
      }

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
      if (!await connectionChecker.isConnected) {
        return left(Failure(GeneralErrors.noInternetConnection));
      }

      final forecast = await weatherRemoteDatasource.getForecastByCity(city);

      return right(forecast);
    } on ServerException catch (e) {
      return left(Failure(e.message));
    }
  }

  @override
  Future<Either<Failure, WeatherResponse>> getWeatherByCity(String city) async {
    try {
      if (!await connectionChecker.isConnected) {
        return left(Failure(GeneralErrors.noInternetConnection));
      }

      final weather = await weatherRemoteDatasource.getWeatherByCity(city);

      return right(weather);
    } on ServerException catch (e) {
      return left(Failure(e.message));
    }
  }

  @override
  Future<Either<Failure, OpenMeteoResponse>> getOpenMeteo() async {
    try {
      if (!await connectionChecker.isConnected) {
        return left(Failure(GeneralErrors.noInternetConnection));
      }

      final openMeteo = await weatherRemoteDatasource.getOpenMeteo();

      return right(openMeteo);
    } on ServerException catch (e) {
      return left(Failure(e.message));
    }
  }
}
