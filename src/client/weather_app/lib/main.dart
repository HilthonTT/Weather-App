import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:weather_app/app/app.dart';
import 'package:weather_app/common/di/service_locator.dart';
import 'package:weather_app/modules/settings/presentation/bloc/settings_bloc.dart';

Future<void> main() async {
  WidgetsFlutterBinding.ensureInitialized();

  await initDependencies();

  runApp(
    ProviderScope(
      child: MultiBlocProvider(
        providers: [
          // Load the user's settings once at startup; every page reads them
          // from this bloc.
          BlocProvider(
            create:
                (_) =>
                    serviceLocator<SettingsBloc>()
                      ..add(const GetSettingsEvent()),
          ),
        ],
        child: const WeatherApp(),
      ),
    ),
  );
}
