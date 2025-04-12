import 'package:flutter/foundation.dart';

@immutable
final class City {
  final String name;
  final double lat;
  final double lon;

  const City({required this.name, required this.lat, required this.lon});
}
