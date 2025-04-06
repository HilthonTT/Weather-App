import 'package:flutter/foundation.dart';
import 'package:fpdart/fpdart.dart';
import 'package:weather_app/common/errors/failure.dart';
import 'package:weather_app/common/usecase/usecase.dart';
import 'package:weather_app/modules/weather/domain/entities/forecast.dart';
import 'package:weather_app/modules/weather/domain/repositories/weather_repository.dart';

@immutable
final class GetForecast implements UseCase<ForecastResponse, NoParams> {
  final WeatherRepository weatherRepository;

  const GetForecast({required this.weatherRepository});

  @override
  Future<Either<Failure, ForecastResponse>> call(NoParams params) {
    return weatherRepository.getForecast();
  }
}
