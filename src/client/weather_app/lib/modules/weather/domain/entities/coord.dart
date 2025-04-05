import 'package:flutter/material.dart';

@immutable
class Coord {
  final double lat;
  final double lon;

  const Coord({required this.lat, required this.lon});
}
