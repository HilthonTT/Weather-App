/// Conversion factor from meters per second to kilometers per hour.
const double _metersPerSecondToKmPerHour = 3.6;
const double _metersPerSecondToMilesPerHour = 2.23694;

/// Extension to convert wind speed from m/s to km/h.
extension ConvertWindSpeed on double {
  /// Converts the wind speed from meters per second to kilometers per hour
  /// and returns the result as a string with 2 decimal places.
  String get asKmPerHour =>
      (this * _metersPerSecondToKmPerHour).toStringAsFixed(2);

  /// Converts the wind speed from meters per second to miles per hour
  /// and returns the result as a string with 2 decimal places.
  String get asMilesPerHour =>
      (this * _metersPerSecondToMilesPerHour).toStringAsFixed(2);
}

extension ConvertTemperature on double {
  /// Converts Celsius to Fahrenheit and returns a string with 1 decimal.
  String get asFahrenheit => ((this * 9 / 5) + 32).toStringAsFixed(1);

  /// Returns Celsius with 1 decimal.
  String get asCelsius => toStringAsFixed(1);
}
