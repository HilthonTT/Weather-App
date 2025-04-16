import 'package:flutter/foundation.dart';

@immutable
class HourlyUnits {
  final String time;
  final String temperature2m;

  const HourlyUnits({required this.time, required this.temperature2m});
}
