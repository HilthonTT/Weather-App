import 'package:flutter/foundation.dart';
import 'package:geolocator/geolocator.dart';
import 'package:hive/hive.dart';
import 'package:weather_app/common/services/geolocator.dart';

abstract interface class LocationLocalDataSource {
  Future<Position> getPosition();
}

@immutable
final class LocationLocalDataSourceImpl implements LocationLocalDataSource {
  final Box box;

  const LocationLocalDataSourceImpl({required this.box});

  static const String _latitudeKey = "position-latitude";
  static const String _longitudeKey = "position-longitude";
  static const String _timestampKey = "position-time-saved";

  @override
  Future<Position> getPosition() async {
    final latitude = box.get(_latitudeKey) as double?;
    final longitude = box.get(_longitudeKey) as double?;
    final timeSavedString = box.get(_timestampKey) as String?;
    final timeSaved =
        timeSavedString != null ? DateTime.tryParse(timeSavedString) : null;

    final isExpired =
        timeSaved == null ||
        DateTime.now().difference(timeSaved) > const Duration(days: 1);

    if (latitude == null || longitude == null || isExpired) {
      final newPosition = await getLocation();

      box.put(_latitudeKey, newPosition.latitude);
      box.put(_longitudeKey, newPosition.longitude);
      box.put(_timestampKey, DateTime.now().toIso8601String());

      return newPosition;
    }

    return Position(
      latitude: latitude,
      longitude: longitude,
      timestamp: timeSaved,
      accuracy: 0,
      altitude: 0,
      heading: 0,
      speed: 0,
      speedAccuracy: 0,
      altitudeAccuracy: 0,
      headingAccuracy: 0,
    );
  }
}
