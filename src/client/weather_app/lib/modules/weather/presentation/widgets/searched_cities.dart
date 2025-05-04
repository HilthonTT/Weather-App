import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:weather_app/common/constants/app_colors.dart';
import 'package:weather_app/common/constants/text_styles.dart';
import 'package:weather_app/common/extensions/double.dart';
import 'package:weather_app/common/extensions/string.dart';
import 'package:weather_app/common/models/city.dart';
import 'package:weather_app/common/utils/get_weather_icon.dart';
import 'package:weather_app/modules/settings/domain/entities/settings.dart';
import 'package:weather_app/modules/weather/presentation/pages/detail_page.dart';
import 'package:weather_app/modules/weather/presentation/providers/get_weather_by_city.dart';

final class SearchedCities extends StatelessWidget {
  final List<City> cities;
  final Settings? settings;

  const SearchedCities({super.key, required this.cities, this.settings});

  @override
  Widget build(BuildContext context) {
    return GridView.builder(
      gridDelegate: const SliverGridDelegateWithFixedCrossAxisCount(
        crossAxisCount: 2,
        crossAxisSpacing: 20,
        mainAxisSpacing: 20,
      ),
      physics: const NeverScrollableScrollPhysics(),
      shrinkWrap: true,
      itemCount: cities.length,
      itemBuilder: (context, index) {
        final city = cities[index];

        return InkWell(
          onTap: () {
            Navigator.of(context).push(DetailPage.route(city.name));
          },
          child: CityTile(index: index, city: city.name, settings: settings),
        );
      },
    );
  }
}

final class CityTile extends ConsumerWidget {
  final String city;
  final int index;
  final Settings? settings;

  const CityTile({
    super.key,
    required this.index,
    required this.city,
    this.settings,
  });

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final weatherData = ref.watch(getWeatherByCityProvider(city));

    return weatherData.when(
      data: (weather) {
        return Material(
          color: index == 0 ? AppColors.lightBlue : AppColors.accentBlue,
          elevation: index == 0 ? 8 : 0,
          borderRadius: BorderRadius.circular(25),
          child: Padding(
            padding: const EdgeInsets.symmetric(horizontal: 18, vertical: 24),
            child: Column(
              mainAxisAlignment: MainAxisAlignment.spaceBetween,
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Row(
                  children: [
                    Expanded(
                      child: Column(
                        crossAxisAlignment: CrossAxisAlignment.start,
                        children: [
                          Text(
                            _formatTemperature(weather.main.temp, settings),
                            style: TextStyles.h2,
                          ),
                          const SizedBox(height: 10),
                          Text(
                            weather.weather[0].description.capitalized,
                            style: TextStyles.subtitleText,
                            overflow: TextOverflow.ellipsis,
                            maxLines: 2,
                          ),
                        ],
                      ),
                    ),
                    Image.asset(
                      getOpenWeatherIcon(weather.weather[0].id),
                      width: 50,
                    ),
                  ],
                ),
                Text(
                  weather.name,
                  style: TextStyle(
                    fontSize: 18,
                    color: Colors.white.withValues(alpha: 0.8),
                    fontWeight: FontWeight.w400,
                  ),
                ),
              ],
            ),
          ),
        );
      },
      error: (error, stackTrace) {
        return Center(
          child: Text(error.toString(), style: TextStyles.subtitleText),
        );
      },
      loading: () {
        return const Center(child: CircularProgressIndicator());
      },
    );
  }

  String _formatTemperature(double temp, Settings? settings) {
    final temperature = switch (settings?.tempFormat) {
      TempFormat.fahrenheit => '${temp.asFahrenheit}°F',
      _ => '${temp.asCelsius}°C',
    };

    return temperature;
  }
}
