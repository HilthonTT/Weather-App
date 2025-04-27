import 'package:flutter/foundation.dart';
import 'package:weather_app/modules/users/domain/entities/settings.dart';

@immutable
final class SettingsModel extends Settings {
  const SettingsModel({
    required super.id,
    required super.userId,
    required super.tempFormat,
    required super.timeFormat,
    required super.speedFormat,
  });

  factory SettingsModel.fromJson(Map<String, dynamic> json) {
    return SettingsModel(
      id: (json['id'] as num).toInt(),
      userId: (json['userId'] as num).toInt(),
      tempFormat: TempFormat.values.firstWhere(
        (e) => e.name == json['tempFormat'],
      ),
      timeFormat: TimeFormat.values.firstWhere(
        (e) => e.name == json['timeFormat'],
      ),
      speedFormat: SpeedFormat.values.firstWhere(
        (e) => e.name == json['speedFormat'],
      ),
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'userId': userId,
      'tempFormat': tempFormat.name,
      'timeFormat': timeFormat.name,
      'speedFormat': speedFormat.name,
    };
  }
}
