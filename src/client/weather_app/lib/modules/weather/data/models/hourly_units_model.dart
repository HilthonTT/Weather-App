import 'package:flutter/foundation.dart';
import 'package:weather_app/modules/weather/domain/entities/hourly_units.dart';

@immutable
final class HourlyUnitsModel extends HourlyUnits {
  const HourlyUnitsModel({required super.time, required super.temperature2m});

  factory HourlyUnitsModel.fromJson(Map<String, dynamic> json) {
    return HourlyUnitsModel(
      time: json['time'],
      temperature2m: json['temperature_2m'],
    );
  }
}
