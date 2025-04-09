import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:weather_app/common/constants/app_colors.dart';
import 'package:weather_app/common/constants/text_styles.dart';
import 'package:weather_app/common/extensions/datetime.dart';
import 'package:weather_app/common/extensions/string.dart';
import 'package:weather_app/common/widgets/gradient_container.dart';
import 'package:weather_app/modules/weather/presentation/providers/get_weather_provider.dart';
import 'package:weather_app/modules/weather/presentation/widgets/hourly_forecast.dart';
import 'package:weather_app/modules/weather/presentation/widgets/weather_info.dart';

final class WeatherPage extends ConsumerWidget {
  const WeatherPage({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final openWeather = ref.watch(getWeatherProvider);

    final today = DateTime.now().formattedDate;

    return openWeather.when(
      data: (weather) {
        return GradientContainer(
          children: [
            Column(
              crossAxisAlignment: CrossAxisAlignment.center,
              children: [
                const SizedBox(width: double.infinity),
                Text(weather.name, style: TextStyles.h1),

                const SizedBox(height: 20),

                // Today's date
                Text(today, style: TextStyles.subtitleText),

                const SizedBox(height: 30),

                SizedBox(
                  height: 260,
                  child: Image.asset(
                    'assets/icons/${weather.weather[0].icon.replaceAll('n', 'd')}.png',
                    fit: BoxFit.contain,
                  ),
                ),

                const SizedBox(height: 30),

                Text(
                  weather.weather[0].description.capitalized,
                  style: TextStyles.h2,
                ),
              ],
            ),

            const SizedBox(height: 40),

            WeatherInfo(weather: weather),

            const SizedBox(height: 40),

            const Row(
              mainAxisAlignment: MainAxisAlignment.spaceBetween,
              children: [
                Text(
                  'Today',
                  style: TextStyle(fontSize: 20, color: AppColors.white),
                ),
                InkWell(
                  child: Text(
                    'View full report',
                    style: TextStyle(color: AppColors.lightBlue),
                  ),
                ),
              ],
            ),

            const SizedBox(height: 15),

            const HourlyForecast(),
          ],
        );
      },
      error: (error, stackTrace) {
        return const GradientContainer(
          children: [Center(child: Text('An error has occurred'))],
        );
      },
      loading: () {
        return const GradientContainer(
          children: [
            Align(
              alignment: Alignment.center,
              child: CircularProgressIndicator(),
            ),
          ],
        );
      },
    );
  }
}
