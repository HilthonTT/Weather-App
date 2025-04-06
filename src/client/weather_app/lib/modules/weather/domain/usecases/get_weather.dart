import 'package:flutter/foundation.dart';
import 'package:fpdart/fpdart.dart';
import 'package:weather_app/common/errors/failure.dart';
import 'package:weather_app/common/usecase/usecase.dart';
import 'package:weather_app/modules/weather/domain/entities/weather.dart';
import 'package:weather_app/modules/weather/domain/repositories/weather_repository.dart';

@immutable
final class GetWeather implements UseCase<WeatherResponse, NoParams> {
  final WeatherRepository weatherRepository;

  const GetWeather({required this.weatherRepository});

  @override
  Future<Either<Failure, WeatherResponse>> call(NoParams params) {
    return weatherRepository.getWeather();
  }
}
