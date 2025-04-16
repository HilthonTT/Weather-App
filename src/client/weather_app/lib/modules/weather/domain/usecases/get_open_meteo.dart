import 'package:flutter/foundation.dart';
import 'package:fpdart/fpdart.dart';
import 'package:weather_app/common/errors/failure.dart';
import 'package:weather_app/common/usecase/usecase.dart';
import 'package:weather_app/modules/weather/domain/entities/open_meteo.dart';
import 'package:weather_app/modules/weather/domain/repositories/weather_repository.dart';

@immutable
final class GetOpenMeteo implements UseCase<OpenMeteoResponse, NoParams> {
  final WeatherRepository weatherRepository;

  const GetOpenMeteo({required this.weatherRepository});

  @override
  Future<Either<Failure, OpenMeteoResponse>> call(NoParams params) {
    return weatherRepository.getOpenMeteo();
  }
}
