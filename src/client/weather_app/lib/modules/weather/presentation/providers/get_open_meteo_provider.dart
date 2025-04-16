import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:weather_app/common/usecase/usecase.dart';
import 'package:weather_app/dependency_injection.dart';
import 'package:weather_app/modules/weather/domain/entities/open_meteo.dart';
import 'package:weather_app/modules/weather/domain/usecases/get_open_meteo.dart';

final getOpenMeteoProvider = FutureProvider.autoDispose<OpenMeteoResponse>((
  ref,
) async {
  final getOpenMeteo = serviceLocator.get<GetOpenMeteo>();

  final response = await getOpenMeteo(const NoParams());

  return response.fold(
    (failure) => throw Exception(failure.message),
    (openMeteo) => openMeteo,
  );
});
