import 'package:flutter/material.dart';

@immutable
class Settings {
  final int id;
  final int userId;
  final TempFormat tempFormat;
  final TimeFormat timeFormat;
  final SpeedFormat speedFormat;

  const Settings({
    required this.id,
    required this.userId,
    required this.tempFormat,
    required this.timeFormat,
    required this.speedFormat,
  });

  Settings copyWith({
    int? id,
    int? userId,
    TempFormat? tempFormat,
    TimeFormat? timeFormat,
    SpeedFormat? speedFormat,
  }) {
    return Settings(
      id: id ?? this.id,
      userId: userId ?? this.userId,
      tempFormat: tempFormat ?? this.tempFormat,
      timeFormat: timeFormat ?? this.timeFormat,
      speedFormat: speedFormat ?? this.speedFormat,
    );
  }
}

enum TempFormat { celsius, fahrenheit }

enum TimeFormat { twentyFourHour, twelveHour }

enum SpeedFormat { kmph, mph }
