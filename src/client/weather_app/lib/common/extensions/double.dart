/// Conversion factor from meters per second to kilometers per hour.
const double _metersPerSecondToKmPerHour = 3.6;

/// Extension to convert wind speed from m/s to km/h.
extension ConvertWindSpeed on double {
  /// Converts the wind speed from meters per second to kilometers per hour
  /// and returns the result as a string with 2 decimal places.
  String get asKmPerHour =>
      (this * _metersPerSecondToKmPerHour).toStringAsFixed(2);
}
