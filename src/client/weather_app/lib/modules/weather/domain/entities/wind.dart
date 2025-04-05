import 'package:flutter/material.dart';

@immutable
class Wind {
  final double speed;
  final int deg;
  final double gust;

  const Wind({required this.speed, required this.deg, required this.gust});
}
