import 'package:flutter/foundation.dart';

@immutable
final class SavedCity {
  final String name;
  final double lat;
  final double lon;

  const SavedCity({required this.name, required this.lat, required this.lon});
}
