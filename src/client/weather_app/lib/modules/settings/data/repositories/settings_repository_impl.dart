import 'package:flutter/foundation.dart';
import 'package:fpdart/fpdart.dart';
import 'package:weather_app/common/errors/failure.dart';
import 'package:weather_app/modules/settings/data/datasources/settings_remote_datasource.dart';
import 'package:weather_app/modules/settings/data/models/settings_model.dart';
import 'package:weather_app/modules/settings/domain/entities/settings.dart';
import 'package:weather_app/modules/settings/domain/repositories/settings_repository.dart';

@immutable
final class SettingsRepositoryImpl implements SettingsRepository {
  final SettingsRemoteDatasource settingsRemoteDatasource;

  const SettingsRepositoryImpl({required this.settingsRemoteDatasource});

  @override
  Either<Failure, Settings> getSettings() {
    try {
      final settings = settingsRemoteDatasource.getSettings();

      return right(settings);
    } catch (e) {
      return left(Failure(e.toString()));
    }
  }

  @override
  Either<Failure, Settings> updateSettings({required Settings settings}) {
    try {
      final updatedSettings = settingsRemoteDatasource.updateSettings(
        settings: SettingsModel(
          id: settings.id,
          userId: settings.userId,
          tempFormat: settings.tempFormat,
          timeFormat: settings.timeFormat,
          speedFormat: settings.speedFormat,
        ),
      );

      return right(updatedSettings);
    } catch (e) {
      return left(Failure(e.toString()));
    }
  }
}
