import 'package:dio/dio.dart';
import 'package:flutter/foundation.dart';
import 'package:weather_app/common/constants/constants.dart';
import 'package:weather_app/common/errors/exceptions.dart';
import 'package:weather_app/modules/weather/data/datasources/location_local_data_source.dart';
import 'package:weather_app/modules/weather/data/models/forecast_model.dart';
import 'package:weather_app/modules/weather/data/models/open_meteo_model.dart';
import 'package:weather_app/modules/weather/data/models/weather_model.dart';

abstract interface class WeatherRemoteDatasource {
  Future<WeatherResponseModel> getWeather();
  Future<WeatherResponseModel> getWeatherByCity(String city);

  Future<ForecastResponseModel> getForecast();
  Future<ForecastResponseModel> getForecastByCity(String city);

  Future<OpenMeteoResponseModel> getOpenMeteo();
}

@immutable
final class WeatherRemoteDatasourceImpl implements WeatherRemoteDatasource {
  final LocationLocalDataSource locationLocalDataSource;

  const WeatherRemoteDatasourceImpl({required this.locationLocalDataSource});

  static final _dio = Dio();

  @override
  Future<ForecastResponseModel> getForecast() async {
    try {
      final position = await locationLocalDataSource.getPosition();

      final url =
          '${Constants.apiUrl}/v1/weather/forecast/coords/${position.latitude}/${position.longitude}';

      final response = await _dio.get(url);

      if (response.statusCode != 200) {
        throw ServerException('Failed to load forecast data');
      }

      final forecast = ForecastResponseModel.fromJson(response.data["data"]);

      return forecast;
    } catch (e) {
      debugPrint(e.toString());

      throw ServerException('Failed to load forecast data');
    }
  }

  @override
  Future<WeatherResponseModel> getWeather() async {
    try {
      final position = await locationLocalDataSource.getPosition();

      final url =
          '${Constants.apiUrl}/v1/weather/coords/${position.latitude}/${position.longitude}';

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

  @override
  Future<ForecastResponseModel> getForecastByCity(String city) async {
    try {
      final url = '${Constants.apiUrl}/v1/forecast/$city';

      final response = await _dio.get(url);

      if (response.statusCode != 200) {
        throw ServerException('Failed to load forecast data');
      }

      final forecast = ForecastResponseModel.fromJson(response.data["data"]);

      return forecast;
    } catch (e) {
      debugPrint(e.toString());

      throw ServerException('Failed to load forecast data');
    }
  }

  @override
  Future<WeatherResponseModel> getWeatherByCity(String city) async {
    try {
      final url = '${Constants.apiUrl}/v1/weather/$city';

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

  @override
  Future<OpenMeteoResponseModel> getOpenMeteo() async {
    try {
      final position = await locationLocalDataSource.getPosition();

      final url =
          '${Constants.apiUrl}/v1/weather/open-meteo/coords/${position.latitude}/${position.longitude}';

      final response = await _dio.get(url);

      if (response.statusCode != 200) {
        throw ServerException('Failed to load open meteo data');
      }

      final openMeteo = OpenMeteoResponseModel.fromJson(response.data["data"]);

      return openMeteo;
    } catch (e) {
      debugPrint(e.toString());

      throw ServerException('Failed to load open meteo data');
    }
  }
}
