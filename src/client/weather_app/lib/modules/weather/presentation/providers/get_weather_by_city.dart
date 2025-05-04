import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:weather_app/dependency_injection.main.dart';
import 'package:weather_app/modules/weather/domain/entities/weather.dart';
import 'package:weather_app/modules/weather/domain/usecases/get_weather_by_city.dart';

final getWeatherByCityProvider = FutureProvider.autoDispose
    .family<WeatherResponse, String>((ref, city) async {
      final getWeatherByCity = serviceLocator.get<GetWeatherByCity>();

      final response = await getWeatherByCity(
        GetWeatherByCityParams(city: city),
      );

      return response.fold(
        (failure) => throw Exception(failure.message),
        (weather) => weather,
      );
    });
