import 'package:flutter/foundation.dart';
import 'package:weather_app/modules/weather/domain/entities/daily_forecast.dart';
import 'package:weather_app/modules/weather/domain/entities/daily_units.dart';
import 'package:weather_app/modules/weather/domain/entities/hourly_forecast.dart';
import 'package:weather_app/modules/weather/domain/entities/hourly_units.dart';

@immutable
class OpenMeteoResponse {
  final double latitude;
  final double longitude;
  final double generationTimeMs;
  final int utcOffsetSeconds;
  final String timezone;
  final String timezoneAbbreviation;
  final double elevation;
  final HourlyUnits hourlyUnits;
  final HourlyForecast hourly;
  final DailyUnits dailyUnits;
  final DailyForecast daily;

  const OpenMeteoResponse({
    required this.latitude,
    required this.longitude,
    required this.generationTimeMs,
    required this.utcOffsetSeconds,
    required this.timezone,
    required this.timezoneAbbreviation,
    required this.elevation,
    required this.hourlyUnits,
    required this.hourly,
    required this.dailyUnits,
    required this.daily,
  });
}
