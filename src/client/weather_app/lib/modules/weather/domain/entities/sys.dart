import 'package:flutter/material.dart';

@immutable
class ForecastSys {
  final String pod;

  const ForecastSys({required this.pod});
}

@immutable
class WeatherSys {
  final String country;
  final int sunrise;
  final int sunset;

  const WeatherSys({
    required this.country,
    required this.sunrise,
    required this.sunset,
  });
}
