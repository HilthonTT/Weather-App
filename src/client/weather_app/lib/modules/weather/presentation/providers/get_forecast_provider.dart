import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:weather_app/common/usecase/usecase.dart';
import 'package:weather_app/dependency_injection.main.dart';
import 'package:weather_app/modules/weather/domain/entities/forecast.dart';
import 'package:weather_app/modules/weather/domain/usecases/get_forecast.dart';

final getForecastProvider = FutureProvider.autoDispose<ForecastResponse>((
  ref,
) async {
  final getForecast = serviceLocator.get<GetForecast>();
  final response = await getForecast(const NoParams());

  return response.fold(
    (failure) => throw Exception(failure.message),
    (forecast) => forecast,
  );
});
