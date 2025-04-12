import 'package:flutter/foundation.dart';
import 'package:fpdart/fpdart.dart';
import 'package:weather_app/common/errors/failure.dart';
import 'package:weather_app/common/usecase/usecase.dart';
import 'package:weather_app/modules/weather/domain/entities/forecast.dart';
import 'package:weather_app/modules/weather/domain/repositories/weather_repository.dart';

@immutable
final class GetForecastByCity
    implements UseCase<ForecastResponse, GetForecastByCityParams> {
  final WeatherRepository weatherRepository;

  const GetForecastByCity({required this.weatherRepository});

  @override
  Future<Either<Failure, ForecastResponse>> call(
    GetForecastByCityParams params,
  ) {
    return weatherRepository.getForecastByCity(params.city);
  }
}

@immutable
final class GetForecastByCityParams {
  final String city;

  const GetForecastByCityParams({required this.city});
}
