import 'package:weather_app/modules/weather/domain/entities/coord.dart';

final class CoordModel extends Coord {
  const CoordModel({required super.lat, required super.lon});

  /// Create a [CoordModel] from a JSON map
  factory CoordModel.fromJson(Map<String, dynamic> json) {
    return CoordModel(lat: _toDouble(json['lat']), lon: _toDouble(json['lon']));
  }

  /// Convert a [CoordModel] to a JSON map
  Map<String, dynamic> toJson() {
    return {'lat': lat, 'lon': lon};
  }

  static double _toDouble(dynamic value) {
    if (value is int) {
      return value.toDouble();
    }
    if (value is double) {
      return value;
    }

    throw Exception('Invalid type for coordinate value: $value');
  }
}
