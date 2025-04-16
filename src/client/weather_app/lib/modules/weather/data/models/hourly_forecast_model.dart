import 'package:flutter/foundation.dart';
import 'package:weather_app/modules/weather/domain/entities/hourly_forecast.dart';

@immutable
final class HourlyForecastModel extends HourlyForecast {
  const HourlyForecastModel({
    required super.time,
    required super.temperature2m,
  });

  factory HourlyForecastModel.fromJson(Map<String, dynamic> json) {
    return HourlyForecastModel(
      time: List<String>.from(json['time']),
      temperature2m: List<double>.from(
        json['temperature_2m'].map((e) => e is int ? e.toDouble() : e),
      ),
    );
  }
}
