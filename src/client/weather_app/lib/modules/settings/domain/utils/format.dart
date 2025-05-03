import 'package:weather_app/modules/settings/domain/entities/settings.dart';

String formatTempFormat(TempFormat tempFormat) {
  switch (tempFormat) {
    case TempFormat.celsius:
      return "Celsius";
    case TempFormat.fahrenheit:
      return "Fahrenheit";
  }
}

String formatTimeFormat(TimeFormat timeFormat) {
  switch (timeFormat) {
    case TimeFormat.twelveHour:
      return "12h";
    case TimeFormat.twentyFourHour:
      return "24h";
  }
}

String formatSpeedFormat(SpeedFormat speedFormat) {
  switch (speedFormat) {
    case SpeedFormat.kmph:
      return "Km/h";
    case SpeedFormat.mph:
      return "Mi/h";
  }
}
