import 'package:get_it/get_it.dart';
import 'package:hive/hive.dart';
import 'package:path_provider/path_provider.dart';
import 'package:weather_app/modules/users/data/datasources/settings_remote_datasource.dart';
import 'package:weather_app/modules/users/data/datasources/user_remote_datasource.dart';
import 'package:weather_app/modules/users/data/repositories/settings_repository_impl.dart';
import 'package:weather_app/modules/users/data/repositories/user_repository_impl.dart';
import 'package:weather_app/modules/users/domain/repositories/settings_repository.dart';
import 'package:weather_app/modules/users/domain/repositories/user_repository.dart';
import 'package:weather_app/modules/users/domain/usecases/get_current_user.dart';
import 'package:weather_app/modules/users/domain/usecases/login_user.dart';
import 'package:weather_app/modules/users/domain/usecases/register_user.dart';
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
  _initUser();
}

Future<void> _initServices() async {
  final appDocumentsDirectory = await getApplicationDocumentsDirectory();

  Hive.defaultDirectory = appDocumentsDirectory.path;

  serviceLocator.registerLazySingleton(() => Hive.box(name: "weathers"));
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
      () => WeatherRepositoryImpl(weatherRemoteDatasource: serviceLocator()),
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

void _initUser() {
  serviceLocator
    ..registerFactory<UserRemoteDatasource>(
      () => UserRemoteDatasourceImpl(box: serviceLocator()),
    )
    ..registerFactory<SettingsRemoteDatasource>(
      () => SettingsRemoteDatasourceImpl(box: serviceLocator()),
    )
    ..registerFactory<UserRepository>(
      () => UserRepositoryImpl(userRemoteDatasource: serviceLocator()),
    )
    ..registerFactory<SettingsRepository>(
      () => SettingsRepositoryImpl(settingsRemoteDatasource: serviceLocator()),
    )
    ..registerFactory(() => GetCurrentUser(userRepository: serviceLocator()))
    ..registerFactory(() => LoginUser(userRepository: serviceLocator()))
    ..registerFactory(() => RegisterUser(userRepository: serviceLocator()));
}
