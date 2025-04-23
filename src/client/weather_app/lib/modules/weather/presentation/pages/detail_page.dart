import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:weather_app/common/constants/app_colors.dart';
import 'package:weather_app/common/constants/text_styles.dart';
import 'package:weather_app/common/extensions/datetime.dart';
import 'package:weather_app/common/extensions/string.dart';
import 'package:weather_app/common/utils/get_weather_icon.dart';
import 'package:weather_app/common/widgets/custom_back_button.dart';
import 'package:weather_app/common/widgets/gradient_container.dart';
import 'package:weather_app/modules/weather/presentation/providers/get_weather_by_city.dart';
import 'package:weather_app/modules/weather/presentation/widgets/weather_info.dart';

final class DetailPage extends ConsumerWidget {
  static route(String city) =>
      MaterialPageRoute(builder: (context) => DetailPage(city: city));

  final String city;

  const DetailPage({super.key, required this.city});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final weatherData = ref.watch(getWeatherByCityProvider(city));

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
                      Text(today.formattedDate, style: TextStyles.subtitleText),

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
                  WeatherInfo(weather: weather),

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
  }
}
