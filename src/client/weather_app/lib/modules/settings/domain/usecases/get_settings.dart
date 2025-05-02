import 'package:flutter/foundation.dart';
import 'package:fpdart/fpdart.dart';
import 'package:weather_app/common/errors/failure.dart';
import 'package:weather_app/common/usecase/usecase.dart';
import 'package:weather_app/modules/settings/domain/entities/settings.dart';
import 'package:weather_app/modules/settings/domain/repositories/settings_repository.dart';

@immutable
final class GetSettings implements UseCase<Settings, NoParams> {
  final SettingsRepository settingsRepository;

  const GetSettings({required this.settingsRepository});

  @override
  Future<Either<Failure, Settings>> call(NoParams params) {
    return Future.value(settingsRepository.getSettings());
  }
}
