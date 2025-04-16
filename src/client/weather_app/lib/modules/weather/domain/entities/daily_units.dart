import 'package:flutter/foundation.dart';

@immutable
class DailyUnits {
  final String time;
  final String weatherCode;
  final String apparentTempMax;
  final String apparentTempMin;

  const DailyUnits({
    required this.time,
    required this.weatherCode,
    required this.apparentTempMax,
    required this.apparentTempMin,
  });
}
