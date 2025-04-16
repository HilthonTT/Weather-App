import 'package:weather_app/modules/weather/domain/entities/daily_forecast.dart';

final class DailyForecastModel extends DailyForecast {
  const DailyForecastModel({
    required super.time,
    required super.weatherCode,
    required super.apparentTempMax,
    required super.apparentTempMin,
  });

  factory DailyForecastModel.fromJson(Map<String, dynamic> json) {
    return DailyForecastModel(
      time: List<String>.from(json['time']),
      weatherCode: List<int>.from(json['weather_code']),
      apparentTempMax: List<double>.from(
        json['apparent_temperature_max'].map(
          (e) => e is num ? e.toDouble() : e,
        ),
      ),
      apparentTempMin: List<double>.from(
        json['apparent_temperature_min'].map(
          (e) => e is num ? e.toDouble() : e,
        ),
      ),
    );
  }
}
