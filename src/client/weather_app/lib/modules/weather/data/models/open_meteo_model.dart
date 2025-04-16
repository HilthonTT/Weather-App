import 'package:flutter/foundation.dart';
import 'package:weather_app/modules/weather/data/models/daily_forecast_model.dart';
import 'package:weather_app/modules/weather/data/models/daily_units_model.dart';
import 'package:weather_app/modules/weather/data/models/hourly_forecast_model.dart';
import 'package:weather_app/modules/weather/data/models/hourly_units_model.dart';
import 'package:weather_app/modules/weather/domain/entities/open_meteo.dart';

@immutable
final class OpenMeteoResponseModel extends OpenMeteoResponse {
  const OpenMeteoResponseModel({
    required super.latitude,
    required super.longitude,
    required super.generationTimeMs,
    required super.utcOffsetSeconds,
    required super.timezone,
    required super.timezoneAbbreviation,
    required super.elevation,
    required super.hourlyUnits,
    required super.hourly,
    required super.dailyUnits,
    required super.daily,
  });

  factory OpenMeteoResponseModel.fromJson(Map<String, dynamic> json) {
    return OpenMeteoResponseModel(
      latitude: (json['latitude'] as num).toDouble(),
      longitude: (json['longitude'] as num).toDouble(),
      generationTimeMs: (json['generationtime_ms'] as num).toDouble(),
      utcOffsetSeconds: (json['utc_offset_seconds'] as num).toInt(),
      timezone: json['timezone'],
      timezoneAbbreviation: json['timezone_abbreviation'],
      elevation: (json['elevation'] as num).toDouble(),
      hourlyUnits: HourlyUnitsModel.fromJson(json['hourly_units']),
      hourly: HourlyForecastModel.fromJson(json['hourly']),
      dailyUnits: DailyUnitsModel.fromJson(json['daily_units']),
      daily: DailyForecastModel.fromJson(json['daily']),
    );
  }
}
