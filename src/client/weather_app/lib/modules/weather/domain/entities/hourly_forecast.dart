import 'package:flutter/foundation.dart';

@immutable
class HourlyForecast {
  final List<String> time;
  final List<double> temperature2m;

  const HourlyForecast({required this.time, required this.temperature2m});
}
