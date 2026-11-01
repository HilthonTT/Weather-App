/// Display helpers that turn raw API values into strings according to the
/// user's [Settings].
///
/// The backend always returns metric values: temperatures in °C and wind
/// speed in m/s. When [Settings] is `null` (not loaded yet) the metric
/// defaults (°C, km/h) are used.
library;

import 'package:weather_app/common/extensions/double.dart';
import 'package:weather_app/modules/settings/domain/entities/settings.dart';

/// Label shown in the settings dropdown for [tempFormat].
String formatTempFormat(TempFormat tempFormat) {
  switch (tempFormat) {
    case TempFormat.celsius:
      return "Celsius";
    case TempFormat.fahrenheit:
      return "Fahrenheit";
  }
}

/// Label shown in the settings dropdown for [timeFormat].
String formatTimeFormat(TimeFormat timeFormat) {
  switch (timeFormat) {
    case TimeFormat.twelveHour:
      return "12h";
    case TimeFormat.twentyFourHour:
      return "24h";
  }
}

/// Label shown in the settings dropdown for [speedFormat].
String formatSpeedFormat(SpeedFormat speedFormat) {
  switch (speedFormat) {
    case SpeedFormat.kmph:
      return "Km/h";
    case SpeedFormat.mph:
      return "Mi/h";
  }
}

/// The temperature unit symbol for [settings], e.g. `°C`.
String temperatureUnit(Settings? settings) =>
    settings?.tempFormat == TempFormat.fahrenheit ? '°F' : '°C';

/// The numeric part of [celsius] converted to the user's unit, e.g. `12.5`.
String temperatureValue(double celsius, Settings? settings) =>
    settings?.tempFormat == TempFormat.fahrenheit
        ? celsius.asFahrenheit
        : celsius.asCelsius;

/// [celsius] formatted with its unit, e.g. `12.5°C` or `54.5°F`.
String formatTemperature(double celsius, Settings? settings) =>
    '${temperatureValue(celsius, settings)}${temperatureUnit(settings)}';

/// [metersPerSecond] formatted in the user's unit, e.g. `18.00 km/h`.
String formatWindSpeed(double metersPerSecond, Settings? settings) =>
    settings?.speedFormat == SpeedFormat.mph
        ? '${metersPerSecond.asMilesPerHour} mph'
        : '${metersPerSecond.asKmPerHour} km/h';
