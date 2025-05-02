import 'package:fpdart/fpdart.dart';
import 'package:weather_app/common/errors/failure.dart';
import 'package:weather_app/modules/settings/domain/entities/settings.dart';

abstract interface class SettingsRepository {
  Either<Failure, Settings> getSettings();

  Either<Failure, Settings> updateSettings({required Settings settings});
}
