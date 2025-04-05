import 'package:flutter/foundation.dart';
import 'package:weather_app/modules/weather/domain/entities/city.dart';
import 'package:weather_app/modules/weather/domain/entities/clouds.dart';
import 'package:weather_app/modules/weather/domain/entities/main.dart';
import 'package:weather_app/modules/weather/domain/entities/sys.dart';
import 'package:weather_app/modules/weather/domain/entities/weather.dart';
import 'package:weather_app/modules/weather/domain/entities/wind.dart';

@immutable
class ForecastResponse {
  final String cod;
  final int message;
  final int cnt;
  final List<ForecastItem> list;
  final City city;

  const ForecastResponse({
    required this.cod,
    required this.message,
    required this.cnt,
    required this.list,
    required this.city,
  });
}

@immutable
class ForecastItem {
  final int dt;
  final Main main;
  final List<WeatherItem> weather;
  final Clouds clouds;
  final Wind wind;
  final int visibility;
  final double pop;
  final ForecastSys sys;
  final String dtTxt;

  const ForecastItem({
    required this.dt,
    required this.main,
    required this.weather,
    required this.clouds,
    required this.wind,
    required this.visibility,
    required this.pop,
    required this.sys,
    required this.dtTxt,
  });
}
