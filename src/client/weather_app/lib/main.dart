import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:weather_app/common/layouts/home_layout.dart';
import 'package:weather_app/dependency_injection.dart';
import 'package:weather_app/modules/users/presentation/bloc/user_bloc.dart';

Future<void> main() async {
  WidgetsFlutterBinding.ensureInitialized();

  await initDependencies();

  runApp(
    ProviderScope(
      child: MultiBlocProvider(
        providers: [BlocProvider(create: (_) => serviceLocator<UserBloc>())],
        child: const MyApp(),
      ),
    ),
  );
}

final class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      debugShowCheckedModeBanner: false,
      title: 'Weather App',
      theme: ThemeData(
        colorScheme: ColorScheme.fromSeed(seedColor: Colors.deepPurple),
      ),
      home: const HomeLayout(),
    );
  }
}
