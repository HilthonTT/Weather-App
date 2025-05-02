import 'package:flutter/foundation.dart';
import 'package:fpdart/fpdart.dart';
import 'package:weather_app/common/errors/failure.dart';
import 'package:weather_app/common/usecase/usecase.dart';
import 'package:weather_app/modules/settings/domain/entities/settings.dart';
import 'package:weather_app/modules/settings/domain/repositories/settings_repository.dart';

@immutable
final class UpdateSettings implements UseCase<Settings, UpdateSettingsParams> {
  final SettingsRepository settingsRepository;

  const UpdateSettings({required this.settingsRepository});

  @override
  Future<Either<Failure, Settings>> call(UpdateSettingsParams params) {
    return Future.value(
      settingsRepository.updateSettings(settings: params.settings),
    );
  }
}

@immutable
final class UpdateSettingsParams {
  final Settings settings;

  const UpdateSettingsParams({required this.settings});
}
