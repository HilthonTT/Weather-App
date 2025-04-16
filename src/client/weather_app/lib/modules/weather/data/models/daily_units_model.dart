import 'package:weather_app/modules/weather/domain/entities/daily_units.dart';

final class DailyUnitsModel extends DailyUnits {
  const DailyUnitsModel({
    required super.time,
    required super.weatherCode,
    required super.apparentTempMax,
    required super.apparentTempMin,
  });

  factory DailyUnitsModel.fromJson(Map<String, dynamic> json) {
    return DailyUnitsModel(
      time: json['time'],
      weatherCode: json['weather_code'],
      apparentTempMax: json['apparent_temperature_max'],
      apparentTempMin: json['apparent_temperature_min'],
    );
  }
}
