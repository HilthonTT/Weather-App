import 'package:get_it/get_it.dart';
import 'package:weather_app/modules/weather/data/datasources/weather_remote_datasource.dart';
import 'package:weather_app/modules/weather/data/repositories/weather_repository_impl.dart';
import 'package:weather_app/modules/weather/domain/repositories/weather_repository.dart';
import 'package:weather_app/modules/weather/domain/usecases/get_forecast.dart';
import 'package:weather_app/modules/weather/domain/usecases/get_weather.dart';

final serviceLocator = GetIt.instance;

void initDependencies() {
  _initWeather();
}

void _initWeather() {
  serviceLocator
    ..registerFactory<WeatherRemoteDatasource>(
      () => WeatherRemoteDatasourceImpl(),
    )
    ..registerFactory<WeatherRepository>(
      () => WeatherRepositoryImpl(weatherRemoteDatasource: serviceLocator()),
    )
    ..registerFactory(() => GetForecast(weatherRepository: serviceLocator()))
    ..registerFactory(() => GetWeather(weatherRepository: serviceLocator()));
}
