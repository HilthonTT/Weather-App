import 'package:flutter/foundation.dart';
import 'package:fpdart/fpdart.dart';
import 'package:weather_app/common/errors/failure.dart';
import 'package:weather_app/common/usecase/usecase.dart';
import 'package:weather_app/modules/weather/domain/entities/weather.dart';
import 'package:weather_app/modules/weather/domain/repositories/weather_repository.dart';

@immutable
final class GetWeatherByCity
    implements UseCase<WeatherResponse, GetWeatherByCityParams> {
  final WeatherRepository weatherRepository;

  const GetWeatherByCity({required this.weatherRepository});

  @override
  Future<Either<Failure, WeatherResponse>> call(GetWeatherByCityParams params) {
    return weatherRepository.getWeatherByCity(params.city);
  }
}

@immutable
final class GetWeatherByCityParams {
  final String city;

  const GetWeatherByCityParams({required this.city});
}
