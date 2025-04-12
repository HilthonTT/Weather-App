import 'package:weather_app/modules/weather/domain/entities/wind.dart';

final class WindModel extends Wind {
  const WindModel({
    required super.speed,
    required super.deg,
    required super.gust,
  });

  /// Create a [WindModel] from a JSON map
  factory WindModel.fromJson(Map<String, dynamic> json) {
    return WindModel(
      speed: (json['speed'] as num).toDouble(),
      deg: json['deg'] as int,
      gust: (json['gust'] as num).toDouble(),
    );
  }

  /// Convert a [WindModel] to a JSON map
  Map<String, dynamic> toJson() {
    return {'speed': speed, 'deg': deg, 'gust': gust};
  }
}
