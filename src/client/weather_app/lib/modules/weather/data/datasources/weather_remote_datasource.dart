import 'package:dio/dio.dart';
import 'package:flutter/foundation.dart';
import 'package:weather_app/common/constants/constants.dart';
import 'package:weather_app/common/errors/exceptions.dart';
import 'package:weather_app/modules/weather/data/models/forecast_model.dart';
import 'package:weather_app/modules/weather/data/models/weather_model.dart';

abstract interface class WeatherRemoteDatasource {
  Future<WeatherResponseModel> getWeather();

  Future<ForecastResponseModel> getForecast();
}

@immutable
final class WeatherRemoteDatasourceImpl implements WeatherRemoteDatasource {
  const WeatherRemoteDatasourceImpl();

  static final _dio = Dio();

  @override
  Future<ForecastResponseModel> getForecast() async {
    try {
      final latitude = 50;
      final longitude = 50;

      final url = '${Constants.apiUrl}/v1/forecast/coords/$latitude/$longitude';

      final response = await _dio.get(url);

      if (response.statusCode != 200) {
        throw ServerException('Failed to load forecast data');
      }

      final forecast = ForecastResponseModel.fromJson(response.data["data"]);

      return forecast;
    } catch (e) {
      debugPrint(e.toString());

      throw ServerException('Failed to load weather data');
    }
  }

  @override
  Future<WeatherResponseModel> getWeather() async {
    try {
      final latitude = 50;
      final longitude = 50;

      final url = '${Constants.apiUrl}/v1/weather/coords/$latitude/$longitude';

      final response = await _dio.get(url);

      if (response.statusCode != 200) {
        throw ServerException('Failed to load weather data');
      }

      final weather = WeatherResponseModel.fromJson(response.data["data"]);

      return weather;
    } catch (e) {
      debugPrint(e.toString());

      throw ServerException('Failed to load weather data');
    }
  }
}
