import 'package:flutter/foundation.dart';
import 'package:hive/hive.dart';

import 'package:weather_app/modules/settings/data/models/settings_model.dart';
import 'package:weather_app/modules/settings/domain/entities/settings.dart';

abstract interface class SettingsLocalDatasource {
  SettingsModel getSettings();

  SettingsModel updateSettings({required SettingsModel settings});
}

@immutable
final class SettingsLocalDatasourceImpl implements SettingsLocalDatasource {
  final Box box;

  const SettingsLocalDatasourceImpl({required this.box});

  static const _tempFormatKey = "tempFormat";
  static const _timeFormatKey = "timeFormat";
  static const _speedFormatKey = "speedFormat";
  static const _userIdKey = "userId";
  static const _idKey = "id";

  @override
  SettingsModel getSettings() {
    final tempFormatIndex = box.get(_tempFormatKey) ?? TempFormat.celsius.index;
    final timeFormatIndex =
        box.get(_timeFormatKey) ?? TimeFormat.twentyFourHour.index;
    final speedFormatIndex = box.get(_speedFormatKey) ?? SpeedFormat.kmph.index;
    final userId = box.get(_userIdKey) ?? 0;
    final id = box.get(_idKey) ?? 0;

    return SettingsModel(
      id: id,
      userId: userId,
      tempFormat: TempFormat.values[tempFormatIndex],
      timeFormat: TimeFormat.values[timeFormatIndex],
      speedFormat: SpeedFormat.values[speedFormatIndex],
    );
  }

  @override
  SettingsModel updateSettings({required SettingsModel settings}) {
    box.put(_tempFormatKey, settings.tempFormat.index);
    box.put(_timeFormatKey, settings.timeFormat.index);
    box.put(_speedFormatKey, settings.speedFormat.index);
    box.put(_userIdKey, settings.userId);
    box.put(_idKey, settings.id);

    return settings;
  }
}
