import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:weather_app/common/usecase/usecase.dart';
import 'package:weather_app/dependency_injection.dart';
import 'package:weather_app/modules/weather/domain/entities/weather.dart';
import 'package:weather_app/modules/weather/domain/usecases/get_weather.dart';

final getWeatherProvider = FutureProvider.autoDispose<WeatherResponse>((
  ref,
) async {
  final getWeather = serviceLocator.get<GetWeather>();
  final response = await getWeather(const NoParams());

  return response.fold(
    (failure) => throw Exception(failure.message),
    (weather) => weather,
  );
});
