import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:weather_app/dependency_injection.main.dart';
import 'package:weather_app/modules/weather/domain/entities/forecast.dart';
import 'package:weather_app/modules/weather/domain/usecases/get_forecast_by_city.dart';

final getForecastByCityProvider = FutureProvider.autoDispose
    .family<ForecastResponse, String>((ref, city) async {
      final getForecastByCity = serviceLocator.get<GetForecastByCity>();

      final response = await getForecastByCity(
        GetForecastByCityParams(city: city),
      );

      return response.fold(
        (failure) => throw Exception(failure.message),
        (forecast) => forecast,
      );
    });
