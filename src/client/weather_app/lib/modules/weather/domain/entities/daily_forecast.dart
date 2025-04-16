import 'package:flutter/foundation.dart';

@immutable
class DailyForecast {
  final List<String> time;
  final List<int> weatherCode;
  final List<double> apparentTempMax;
  final List<double> apparentTempMin;

  const DailyForecast({
    required this.time,
    required this.weatherCode,
    required this.apparentTempMax,
    required this.apparentTempMin,
  });
}
