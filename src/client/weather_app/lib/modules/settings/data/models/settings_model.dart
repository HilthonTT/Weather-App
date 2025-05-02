import 'package:flutter/foundation.dart';
import 'package:weather_app/modules/settings/domain/entities/settings.dart';

@immutable
final class SettingsModel extends Settings {
  const SettingsModel({
    required super.id,
    required super.userId,
    required super.tempFormat,
    required super.timeFormat,
    required super.speedFormat,
  });

  // Convert SettingsModel to a Map (for saving to Hive)
  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'userId': userId,
      'tempFormat':
          tempFormat.toString().split('.').last, // Convert enum to string
      'timeFormat':
          timeFormat.toString().split('.').last, // Convert enum to string
      'speedFormat':
          speedFormat.toString().split('.').last, // Convert enum to string
    };
  }

  // Convert a Map back to SettingsModel (from JSON)
  factory SettingsModel.fromJson(Map<String, dynamic> json) {
    return SettingsModel(
      id: json['id'],
      userId: json['userId'],
      tempFormat: TempFormat.values.firstWhere(
        (e) => e.toString().split('.').last == json['tempFormat'],
        orElse: () => TempFormat.celsius, // Default to celsius if parsing fails
      ),
      timeFormat: TimeFormat.values.firstWhere(
        (e) => e.toString().split('.').last == json['timeFormat'],
        orElse:
            () =>
                TimeFormat
                    .twentyFourHour, // Default to twentyFourHour if parsing fails
      ),
      speedFormat: SpeedFormat.values.firstWhere(
        (e) => e.toString().split('.').last == json['speedFormat'],
        orElse: () => SpeedFormat.kmph, // Default to kmph if parsing fails
      ),
    );
  }
}
