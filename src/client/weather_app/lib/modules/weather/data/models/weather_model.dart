import 'package:weather_app/modules/weather/data/models/clouds_model.dart';
import 'package:weather_app/modules/weather/data/models/coord_model.dart';
import 'package:weather_app/modules/weather/data/models/main_model.dart';
import 'package:weather_app/modules/weather/data/models/sys_model.dart';
import 'package:weather_app/modules/weather/data/models/wind_model.dart';
import 'package:weather_app/modules/weather/domain/entities/weather.dart';

final class WeatherResponseModel extends WeatherResponse {
  const WeatherResponseModel({
    required super.coord,
    required super.weather,
    required super.main,
    required super.visibility,
    required super.wind,
    required super.clouds,
    required super.dt,
    required super.sys,
    required super.timezone,
    required super.name,
    required super.cod,
  });

  /// Create a [WeatherResponseModel] from a JSON map
  factory WeatherResponseModel.fromJson(Map<String, dynamic> json) {
    return WeatherResponseModel(
      coord: CoordModel.fromJson(json['coord'] as Map<String, dynamic>),
      weather:
          (json['weather'] as List)
              .map(
                (item) =>
                    WeatherItemModel.fromJson(item as Map<String, dynamic>),
              )
              .toList(),
      main: MainModel.fromJson(json['main'] as Map<String, dynamic>),
      visibility: (json['visibility'] as num).toInt(),
      wind: WindModel.fromJson(json['wind'] as Map<String, dynamic>),
      clouds: CloudsModel.fromJson(json['clouds'] as Map<String, dynamic>),
      dt: (json['dt'] as num).toInt(),
      sys: WeatherSysModel.fromJson(json['sys'] as Map<String, dynamic>),
      timezone: (json['timezone'] as num).toInt(),
      name: json['name'] as String,
      cod: (json['cod'] as num).toInt(),
    );
  }

  /// Convert a [WeatherResponseModel] to a JSON map
  Map<String, dynamic> toJson() {
    return {
      'coord': (coord as CoordModel).toJson(),
      'weather':
          (weather as List<WeatherItemModel>)
              .map((item) => item.toJson())
              .toList(),
      'main': (main as MainModel).toJson(),
      'visibility': visibility,
      'wind': (wind as WindModel).toJson(),
      'clouds': (clouds as CloudsModel).toJson(),
      'dt': dt,
      'sys': (sys as ForecastSysModel).toJson(),
      'timezone': timezone,
      'name': name,
      'cod': cod,
    };
  }
}

final class WeatherItemModel extends WeatherItem {
  const WeatherItemModel({
    required super.id,
    required super.main,
    required super.description,
    required super.icon,
  });

  /// Create a [WeatherItemModel] from a JSON map
  factory WeatherItemModel.fromJson(Map<String, dynamic> json) {
    return WeatherItemModel(
      id: (json['id'] as num).toInt(),
      main: json['main'] as String,
      description: json['description'] as String,
      icon: json['icon'] as String,
    );
  }

  /// Convert a [WeatherItemModel] to a JSON map
  Map<String, dynamic> toJson() {
    return {'id': id, 'main': main, 'description': description, 'icon': icon};
  }
}
