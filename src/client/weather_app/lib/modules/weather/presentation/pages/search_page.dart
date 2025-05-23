import 'dart:async';

import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:internet_connection_checker_plus/internet_connection_checker_plus.dart';
import 'package:weather_app/common/constants/app_colors.dart';
import 'package:weather_app/common/constants/text_styles.dart';
import 'package:weather_app/common/icons/location_icon.dart';
import 'package:weather_app/common/models/city.dart';
import 'package:weather_app/common/utils/show_snackbar.dart';
import 'package:weather_app/common/widgets/gradient_container.dart';
import 'package:weather_app/common/widgets/rounded_text_field.dart';
import 'package:weather_app/dependency_injection.main.dart';
import 'package:weather_app/modules/settings/domain/entities/settings.dart';
import 'package:weather_app/modules/settings/presentation/bloc/settings_bloc.dart';
import 'package:weather_app/modules/weather/presentation/widgets/searched_cities.dart';

final class SearchPage extends StatefulWidget {
  static const int currentIndex = 1;

  const SearchPage({super.key});

  @override
  State<SearchPage> createState() => _SearchPageState();
}

final class _SearchPageState extends State<SearchPage> {
  late final StreamSubscription<InternetStatus> _subscription;
  late final AppLifecycleListener _listener;

  final _controller = TextEditingController();

  final _cities = [
    const City(name: 'Tokyo', lat: 35.6833, lon: 139.7667),
    const City(name: 'New Delhi', lat: 28.5833, lon: 77.2),
    const City(name: 'Paris', lat: 48.85, lon: 2.3333),
    const City(name: 'London', lat: 51.4833, lon: -0.0833),
    const City(name: 'New York', lat: 40.7167, lon: -74.0167),
    const City(name: 'Tehran', lat: 35.6833, lon: 51.4167),
  ];

  Settings? settings;

  void _addCity(String cityName) {
    final name = cityName.trim();

    if (name.isEmpty) {
      return;
    }

    final city = City(name: name, lat: 0, lon: 0);

    setState(() {
      _cities.add(city);
    });

    _controller.clear();
  }

  @override
  void initState() {
    super.initState();

    _subscription = serviceLocator<InternetConnection>().onStatusChange.listen(
      (status) {},
    );

    _listener = AppLifecycleListener(
      onResume: _subscription.resume,
      onHide: _subscription.pause,
      onPause: _subscription.pause,
    );

    context.read<SettingsBloc>().add(const GetSettingsEvent());
  }

  @override
  void dispose() {
    _subscription.cancel();
    _listener.dispose();
    _controller.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: AppColors.black,
      body: BlocConsumer<SettingsBloc, SettingsState>(
        listener: (context, state) {
          if (state is SettingsDisplaySuccess) {
            setState(() {
              settings = state.settings;
            });
          } else if (state is SettingsFailure) {
            showSnackbar(context, state.error, icon: Icons.error);
          }
        },
        builder: (context, state) {
          return GradientContainer(
            children: [
              const Column(
                crossAxisAlignment: CrossAxisAlignment.center,
                children: [
                  Text(
                    'Pick location',
                    style: TextStyles.h1,
                    textAlign: TextAlign.center,
                  ),
                  SizedBox(height: 30),
                  Text(
                    'Find the area or city that you want to know the detailed weather info at this time',
                    style: TextStyles.subtitleText,
                    textAlign: TextAlign.center,
                  ),
                ],
              ),

              const SizedBox(height: 40),

              Row(
                children: [
                  Expanded(
                    child: RoundedTextField(
                      controller: _controller,
                      onSubmitted: _addCity,
                    ),
                  ),
                  const SizedBox(width: 15),
                  const LocationIcon(),
                ],
              ),

              const SizedBox(height: 30),

              SearchedCities(cities: _cities, settings: settings),
            ],
          );
        },
      ),
    );
  }
}
