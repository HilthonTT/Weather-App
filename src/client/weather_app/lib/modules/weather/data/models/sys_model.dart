import 'package:weather_app/modules/weather/domain/entities/sys.dart';

final class ForecastSysModel extends ForecastSys {
  const ForecastSysModel({required super.pod});

  /// Create a [ForecastSysModel] from a JSON map
  factory ForecastSysModel.fromJson(Map<String, dynamic> json) {
    return ForecastSysModel(pod: json['pod'] as String);
  }

  /// Convert a [ForecastSysModel] to a JSON map
  Map<String, dynamic> toJson() {
    return {'pod': pod};
  }
}

final class WeatherSysModel extends WeatherSys {
  const WeatherSysModel({
    required super.country,
    required super.sunrise,
    required super.sunset,
  });

  /// Create a [WeatherSysModel] from a JSON map
  factory WeatherSysModel.fromJson(Map<String, dynamic> json) {
    return WeatherSysModel(
      country: json['country'] as String,
      sunrise: json['sunrise'] as int,
      sunset: json['sunset'] as int,
    );
  }

  /// Convert a [WeatherSysModel] to a JSON map
  Map<String, dynamic> toJson() {
    return {'country': country, 'sunrise': sunrise, 'sunset': sunset};
  }
}
