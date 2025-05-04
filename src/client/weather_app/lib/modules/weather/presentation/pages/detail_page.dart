import 'dart:async';

import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:internet_connection_checker_plus/internet_connection_checker_plus.dart';
import 'package:weather_app/common/constants/app_colors.dart';
import 'package:weather_app/common/constants/text_styles.dart';
import 'package:weather_app/common/extensions/datetime.dart';
import 'package:weather_app/common/extensions/string.dart';
import 'package:weather_app/common/utils/get_weather_icon.dart';
import 'package:weather_app/common/utils/show_snackbar.dart';
import 'package:weather_app/common/widgets/custom_back_button.dart';
import 'package:weather_app/common/widgets/gradient_container.dart';
import 'package:weather_app/dependency_injection.main.dart';
import 'package:weather_app/modules/settings/domain/entities/settings.dart';
import 'package:weather_app/modules/settings/presentation/bloc/settings_bloc.dart';
import 'package:weather_app/modules/weather/presentation/providers/get_weather_by_city.dart';
import 'package:weather_app/modules/weather/presentation/widgets/weather_info.dart';

final class DetailPage extends ConsumerStatefulWidget {
  static route(String city) =>
      MaterialPageRoute(builder: (context) => DetailPage(city: city));

  final String city;

  const DetailPage({super.key, required this.city});

  @override
  ConsumerState<ConsumerStatefulWidget> createState() => _DetailPageState();
}

final class _DetailPageState extends ConsumerState<DetailPage> {
  late final StreamSubscription<InternetStatus> _subscription;
  late final AppLifecycleListener _listener;

  Settings? settings;

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
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    final weatherData = ref.watch(getWeatherByCityProvider(widget.city));

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
          return weatherData.when(
            data: (weather) {
              final today = DateTime.now();

              return Scaffold(
                backgroundColor: AppColors.black,
                body: Stack(
                  children: [
                    GradientContainer(
                      children: [
                        Column(
                          crossAxisAlignment: CrossAxisAlignment.center,
                          children: [
                            const SizedBox(height: 30, width: double.infinity),
                            // Country name text
                            Text(weather.name, style: TextStyles.h1),

                            const SizedBox(height: 20),

                            // Today's date
                            Text(
                              today.formattedDate,
                              style: TextStyles.subtitleText,
                            ),

                            const SizedBox(height: 50),

                            // Weather icon big
                            SizedBox(
                              height: 300,
                              child: Image.asset(
                                getOpenWeatherIcon(weather.weather[0].id),
                                fit: BoxFit.contain,
                              ),
                            ),

                            // Weather description
                            Text(
                              weather.weather[0].description.capitalized,
                              style: TextStyles.h2,
                            ),
                          ],
                        ),

                        const SizedBox(height: 40),

                        // Weather info in a row
                        WeatherInfo(weather: weather, settings: settings),

                        const SizedBox(height: 15),
                      ],
                    ),

                    const CustomBackButton(),
                  ],
                ),
              );
            },
            error: (error, statckTrace) {
              return const Center(child: Text('An error has occurred'));
            },
            loading: () {
              return const Center(child: CircularProgressIndicator());
            },
          );
        },
      ),
    );
  }
}
