import 'package:flutter/foundation.dart';
import 'package:hive/hive.dart';

import 'package:weather_app/modules/settings/data/models/settings_model.dart';
import 'package:weather_app/modules/settings/domain/entities/settings.dart';

abstract interface class SettingsRemoteDatasource {
  SettingsModel getSettings();

  SettingsModel updateSettings({required SettingsModel settings});
}

@immutable
final class SettingsRemoteDatasourceImpl implements SettingsRemoteDatasource {
  static const _settingsKey = "settings";

  final Box box;

  const SettingsRemoteDatasourceImpl({required this.box});

  @override
  SettingsModel getSettings() {
    final data = box.get(_settingsKey);

    print(data);

    if (data != null && data is Map<String, dynamic>) {
      print("returning data here");
      return SettingsModel.fromJson(data);
    }

    final defaultSettings = SettingsModel(
      id: 0,
      userId: 0,
      tempFormat: TempFormat.celsius,
      timeFormat: TimeFormat.twentyFourHour,
      speedFormat: SpeedFormat.kmph,
    );

    box.put(_settingsKey, defaultSettings.toJson());

    return defaultSettings;
  }

  @override
  SettingsModel updateSettings({required SettingsModel settings}) {
    final json = settings.toJson();

    box.put(_settingsKey, json);

    print("Updated settings: $json");

    return settings;
  }
}
