import 'package:weather_app/modules/weather/data/models/coord_model.dart';
import 'package:weather_app/modules/weather/domain/entities/city.dart';

final class CityModel extends City {
  const CityModel({
    required super.id,
    required super.name,
    required super.coord,
    required super.country,
    required super.population,
    required super.timezone,
    required super.sunrise,
    required super.sunset,
  });

  /// Create a [CoordModel] from a JSON map
  factory CityModel.fromJson(Map<String, dynamic> json) {
    return CityModel(
      id: (json['id'] as num).toInt(),
      name: json['name'] as String,
      coord: CoordModel.fromJson(json['coord'] as Map<String, dynamic>),
      country: json['country'] as String,
      population: (json['population'] as num).toInt(),
      timezone: (json['timezone'] as num).toInt(),
      sunrise: (json['sunrise'] as num).toInt(),
      sunset: (json['sunset'] as num).toInt(),
    );
  }

  /// Convert a [CityModel] to a JSON map
  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'name': name,
      'coord': (coord as CoordModel).toJson(),
      'country': country,
      'population': population,
      'timezone': timezone,
      'sunrise': sunrise,
      'sunset': sunset,
    };
  }
}
