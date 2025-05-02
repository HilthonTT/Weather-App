import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:weather_app/common/usecase/usecase.dart';
import 'package:weather_app/dependency_injection.dart';
import 'package:weather_app/modules/settings/domain/entities/settings.dart';
import 'package:weather_app/modules/settings/domain/usecases/get_settings.dart';

final getSettingsProvider = FutureProvider.autoDispose<Settings>((ref) async {
  final getSettings = serviceLocator.get<GetSettings>();
  final response = await getSettings(const NoParams());

  return response.fold(
    (failure) => throw Exception(failure.message),
    (forecast) => forecast,
  );
});
