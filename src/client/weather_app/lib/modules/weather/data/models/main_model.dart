import 'package:weather_app/modules/weather/domain/entities/main.dart';

final class MainModel extends Main {
  const MainModel({
    required super.temp,
    required super.feelsLike,
    required super.tempMin,
    required super.tempMax,
    required super.pressure,
    required super.seaLevel,
    required super.grndLevel,
    required super.humidity,
  });

  /// Create a [MainModel] from a JSON map
  factory MainModel.fromJson(Map<String, dynamic> json) {
    return MainModel(
      temp: (json['temp'] as double).toDouble(),
      feelsLike: (json['feels_like'] as double).toDouble(),
      tempMin: (json['temp_min'] as double).toDouble(),
      tempMax: (json['temp_max'] as double).toDouble(),
      pressure: json['pressure'] as int,
      seaLevel: json['sea_level'] as int,
      grndLevel: json['grnd_level'] as int,
      humidity: json['humidity'] as int,
    );
  }

  /// Convert a [MainModel] to a JSON map
  Map<String, dynamic> toJson() {
    return {
      'temp': temp,
      'feels_like': feelsLike,
      'temp_min': tempMin,
      'temp_max': tempMax,
      'pressure': pressure,
      'sea_level': seaLevel,
      'grnd_level': grndLevel,
      'humidity': humidity,
    };
  }
}
