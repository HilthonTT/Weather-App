import 'package:flutter/material.dart';
import 'package:weather_app/modules/weather/domain/entities/clouds.dart';
import 'package:weather_app/modules/weather/domain/entities/coord.dart';
import 'package:weather_app/modules/weather/domain/entities/main.dart';
import 'package:weather_app/modules/weather/domain/entities/sys.dart';
import 'package:weather_app/modules/weather/domain/entities/wind.dart';

@immutable
class WeatherResponse {
  final Coord coord;
  final List<WeatherItem> weather;
  final Main main;
  final int visibility;
  final Wind wind;
  final Clouds clouds;
  final int dt;
  final ForecastSys sys;
  final int timezone;
  final String name;
  final int cod;

  const WeatherResponse({
    required this.coord,
    required this.weather,
    required this.main,
    required this.visibility,
    required this.wind,
    required this.clouds,
    required this.dt,
    required this.sys,
    required this.timezone,
    required this.name,
    required this.cod,
  });
}

@immutable
class WeatherItem {
  final int id;
  final String main;
  final String description;
  final String icon;

  const WeatherItem({
    required this.id,
    required this.main,
    required this.description,
    required this.icon,
  });
}
