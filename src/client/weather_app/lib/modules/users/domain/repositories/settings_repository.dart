import 'package:fpdart/fpdart.dart';
import 'package:weather_app/common/errors/failure.dart';
import 'package:weather_app/modules/users/domain/entities/settings.dart';

abstract interface class SettingsRepository {
  Future<Either<Failure, Settings>> getSettings();

  Future<Either<Failure, Settings>> updateSettings({
    required Settings settings,
  });
}
