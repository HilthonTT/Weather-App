import 'package:flutter/foundation.dart';
import 'package:fpdart/fpdart.dart';
import 'package:weather_app/common/errors/exceptions.dart';
import 'package:weather_app/common/errors/failure.dart';
import 'package:weather_app/modules/users/data/datasources/settings_remote_datasource.dart';
import 'package:weather_app/modules/users/data/models/settings_model.dart';
import 'package:weather_app/modules/users/domain/entities/settings.dart';
import 'package:weather_app/modules/users/domain/repositories/settings_repository.dart';

@immutable
final class SettingsRepositoryImpl implements SettingsRepository {
  final SettingsRemoteDatasource settingsRemoteDatasource;

  const SettingsRepositoryImpl({required this.settingsRemoteDatasource});

  @override
  Future<Either<Failure, Settings>> getSettings() async {
    try {
      final settings = await settingsRemoteDatasource.getSettings();

      return right(settings);
    } on ServerException catch (e) {
      return left(Failure(e.message));
    }
  }

  @override
  Future<Either<Failure, Settings>> updateSettings({
    required Settings settings,
  }) async {
    try {
      final updatedSettings = await settingsRemoteDatasource.updateSettings(
        settings: SettingsModel(
          id: settings.id,
          userId: settings.userId,
          tempFormat: settings.tempFormat,
          timeFormat: settings.timeFormat,
          speedFormat: settings.speedFormat,
        ),
      );

      return right(updatedSettings);
    } on ServerException catch (e) {
      return left(Failure(e.message));
    }
  }
}
