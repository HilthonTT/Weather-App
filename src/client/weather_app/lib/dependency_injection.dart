import 'package:get_it/get_it.dart';
import 'package:hive/hive.dart';
import 'package:internet_connection_checker_plus/internet_connection_checker_plus.dart';
import 'package:path_provider/path_provider.dart';
import 'package:weather_app/common/network/connection_checker.dart';
import 'package:weather_app/modules/settings/data/datasources/settings_local_datasource.dart';
import 'package:weather_app/modules/settings/data/repositories/settings_repository_impl.dart';
import 'package:weather_app/modules/settings/domain/repositories/settings_repository.dart';
import 'package:weather_app/modules/settings/domain/usecases/get_settings.dart';
import 'package:weather_app/modules/settings/domain/usecases/update_settings.dart';
import 'package:weather_app/modules/settings/presentation/bloc/settings_bloc.dart';
import 'package:weather_app/modules/weather/data/datasources/location_local_data_source.dart';
import 'package:weather_app/modules/weather/data/datasources/weather_remote_datasource.dart';
import 'package:weather_app/modules/weather/data/repositories/weather_repository_impl.dart';
import 'package:weather_app/modules/weather/domain/repositories/weather_repository.dart';
import 'package:weather_app/modules/weather/domain/usecases/get_forecast.dart';
import 'package:weather_app/modules/weather/domain/usecases/get_forecast_by_city.dart';
import 'package:weather_app/modules/weather/domain/usecases/get_open_meteo.dart';
import 'package:weather_app/modules/weather/domain/usecases/get_weather.dart';
import 'package:weather_app/modules/weather/domain/usecases/get_weather_by_city.dart';

final serviceLocator = GetIt.instance;

Future<void> initDependencies() async {
  await _initServices();

  _initWeather();
  _initSettings();
}

Future<void> _initServices() async {
  final appDocumentsDirectory = await getApplicationDocumentsDirectory();

  Hive.defaultDirectory = appDocumentsDirectory.path;

  serviceLocator.registerLazySingleton(() => Hive.box(name: "weathers"));

  serviceLocator.registerFactory(() => InternetConnection());

  serviceLocator.registerFactory<ConnectionChecker>(
    () => ConnectionCheckerImpl(internetConnection: serviceLocator()),
  );
}

void _initWeather() {
  serviceLocator
    ..registerFactory<LocationLocalDataSource>(
      () => LocationLocalDataSourceImpl(box: serviceLocator()),
    )
    ..registerFactory<WeatherRemoteDatasource>(
      () => WeatherRemoteDatasourceImpl(
        locationLocalDataSource: serviceLocator(),
      ),
    )
    ..registerFactory<WeatherRepository>(
      () => WeatherRepositoryImpl(
        weatherRemoteDatasource: serviceLocator(),
        connectionChecker: serviceLocator(),
      ),
    )
    ..registerFactory(() => GetForecast(weatherRepository: serviceLocator()))
    ..registerFactory(() => GetWeather(weatherRepository: serviceLocator()))
    ..registerFactory(
      () => GetWeatherByCity(weatherRepository: serviceLocator()),
    )
    ..registerFactory(
      () => GetForecastByCity(weatherRepository: serviceLocator()),
    )
    ..registerFactory(() => GetOpenMeteo(weatherRepository: serviceLocator()));
}

void _initSettings() {
  serviceLocator
    ..registerFactory<SettingsLocalDatasource>(
      () => SettingsLocalDatasourceImpl(box: serviceLocator()),
    )
    ..registerFactory<SettingsRepository>(
      () => SettingsRepositoryImpl(settingsRemoteDatasource: serviceLocator()),
    )
    ..registerFactory(() => GetSettings(settingsRepository: serviceLocator()))
    ..registerFactory(
      () => UpdateSettings(settingsRepository: serviceLocator()),
    )
    ..registerLazySingleton(
      () => SettingsBloc(updatedSettings: serviceLocator()),
    );
}
