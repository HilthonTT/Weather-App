import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:weather_app/common/constants/app_colors.dart';
import 'package:weather_app/common/constants/text_styles.dart';
import 'package:weather_app/common/extensions/int.dart';
import 'package:weather_app/common/utils/get_weather_icon.dart';
import 'package:weather_app/modules/weather/domain/entities/forecast.dart';
import 'package:weather_app/modules/weather/presentation/providers/get_forecast_provider.dart';

final class HourlyForecast extends ConsumerWidget {
  const HourlyForecast({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final forecastData = ref.watch(getForecastProvider);

    return forecastData.when(
      data: (response) {
        return SizedBox(
          height: 100,
          child: ListView.builder(
            itemCount: response.cnt,
            scrollDirection: Axis.horizontal,
            shrinkWrap: true,
            itemBuilder: (context, index) {
              final forecast = response.list[index];

              return HourlyForecastTile(
                forecast: forecast,
                isActive: index == 0,
              );
            },
          ),
        );
      },
      error: (error, stackTrace) {
        return const SizedBox();
      },
      loading: () {
        return const Center(child: CircularProgressIndicator());
      },
    );
  }
}

final class HourlyForecastTile extends StatelessWidget {
  final ForecastItem forecast;
  final bool isActive;

  const HourlyForecastTile({
    super.key,
    required this.forecast,
    required this.isActive,
  });

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.only(right: 16, top: 12, bottom: 12),
      child: Material(
        color: isActive ? AppColors.lightBlue : AppColors.accentBlue,
        borderRadius: BorderRadius.circular(15.0),
        elevation: isActive ? 8 : 0,
        child: Padding(
          padding: const EdgeInsets.symmetric(horizontal: 16, vertical: 10),
          child: Row(
            crossAxisAlignment: CrossAxisAlignment.center,
            children: [
              Image.asset(
                getOpenWeatherIcon(forecast.weather[0].id),
                width: 50,
              ),
              const SizedBox(width: 10),
              Column(
                mainAxisAlignment: MainAxisAlignment.spaceBetween,
                mainAxisSize: MainAxisSize.min,
                children: [
                  Text(
                    forecast.dt.time,
                    style: const TextStyle(color: AppColors.white),
                  ),
                  const SizedBox(height: 5),
                  Text('${forecast.main.temp.round()}°', style: TextStyles.h3),
                ],
              ),
            ],
          ),
        ),
      ),
    );
  }
}
