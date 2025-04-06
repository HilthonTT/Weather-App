import 'package:weather_app/modules/weather/data/models/city_model.dart';
import 'package:weather_app/modules/weather/data/models/clouds_model.dart';
import 'package:weather_app/modules/weather/data/models/main_model.dart';
import 'package:weather_app/modules/weather/data/models/sys_model.dart';
import 'package:weather_app/modules/weather/data/models/weather_model.dart';
import 'package:weather_app/modules/weather/data/models/wind_model.dart';
import 'package:weather_app/modules/weather/domain/entities/forecast.dart';

final class ForecastResponseModel extends ForecastResponse {
  const ForecastResponseModel({
    required super.cod,
    required super.message,
    required super.cnt,
    required super.list,
    required super.city,
  });

  /// Create a [ForecastResponseModel] from a JSON map
  factory ForecastResponseModel.fromJson(Map<String, dynamic> json) {
    return ForecastResponseModel(
      cod: json['cod'] as String,
      message: json['message'] as int,
      cnt: json['cnt'] as int,
      list:
          (json['list'] as List)
              .map(
                (item) =>
                    ForecastItemModel.fromJson(item as Map<String, dynamic>),
              )
              .toList(),
      city: CityModel.fromJson(json['city'] as Map<String, dynamic>),
    );
  }

  /// Convert a [ForecastResponseModel] to a JSON map
  Map<String, dynamic> toJson() {
    return {
      'cod': cod,
      'message': message,
      'cnt': cnt,
      'list': list.map((item) => (item as ForecastItemModel).toJson()).toList(),
      'city': (city as CityModel).toJson(),
    };
  }
}

final class ForecastItemModel extends ForecastItem {
  const ForecastItemModel({
    required super.dt,
    required super.main,
    required super.weather,
    required super.clouds,
    required super.wind,
    required super.visibility,
    required super.pop,
    required super.sys,
    required super.dtTxt,
  });

  /// Create a [ForecastItemModel] from a JSON map
  factory ForecastItemModel.fromJson(Map<String, dynamic> json) {
    return ForecastItemModel(
      dt: json['dt'] as int,
      main: MainModel.fromJson(json['main'] as Map<String, dynamic>),
      weather:
          (json['weather'] as List)
              .map(
                (item) =>
                    WeatherItemModel.fromJson(item as Map<String, dynamic>),
              )
              .toList(),
      clouds: CloudsModel.fromJson(json['clouds'] as Map<String, dynamic>),
      wind: WindModel.fromJson(json['wind'] as Map<String, dynamic>),
      visibility: json['visibility'] as int,
      pop: (json['pop'] as num).toDouble(),
      sys: ForecastSysModel.fromJson(json['sys'] as Map<String, dynamic>),
      dtTxt: json['dt_txt'] as String,
    );
  }

  /// Convert a [ForecastItemModel] to a JSON map
  Map<String, dynamic> toJson() {
    return {
      'dt': dt,
      'main': (main as MainModel).toJson(),
      'weather':
          weather.map((item) => (item as WeatherItemModel).toJson()).toList(),
      'clouds': (clouds as CloudsModel).toJson(),
      'wind': (wind as WindModel).toJson(),
      'visibility': visibility,
      'pop': pop,
      'sys': (sys as ForecastSysModel).toJson(),
      'dt_txt': dtTxt,
    };
  }
}
