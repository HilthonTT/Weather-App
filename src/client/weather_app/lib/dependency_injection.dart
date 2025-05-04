part of "dependency_injection.main.dart";

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
      () => SettingsBloc(
        updatedSettings: serviceLocator(),
        getSettings: serviceLocator(),
      ),
    );
}
