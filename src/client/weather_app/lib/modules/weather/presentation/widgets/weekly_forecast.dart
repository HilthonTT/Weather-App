import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:weather_app/common/constants/app_colors.dart';
import 'package:weather_app/common/constants/text_styles.dart';
import 'package:weather_app/common/extensions/datetime.dart';
import 'package:weather_app/common/utils/get_weather_icon.dart';
import 'package:weather_app/common/widgets/superscript_text.dart';
import 'package:weather_app/modules/weather/presentation/providers/get_open_meteo_provider.dart';

final class WeeklyForecast extends ConsumerWidget {
  const WeeklyForecast({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final openMeteoData = ref.watch(getOpenMeteoProvider);

    return openMeteoData.when(
      data: (meteoData) {
        return ListView.builder(
          itemCount: meteoData.daily.weatherCode.length,
          shrinkWrap: true,
          physics: const NeverScrollableScrollPhysics(),
          itemBuilder: (context, index) {
            final dayOfWeek =
                DateTime.parse(meteoData.daily.time[index]).dayOfWeek;
            final date = meteoData.daily.time[index];
            final temp = meteoData.daily.apparentTempMax[index];
            final icon = meteoData.daily.weatherCode[index];

            return WeeklyForecastTitle(
              date: date,
              day: dayOfWeek,
              temp: temp,
              icon: getOpenMeteoIcon(icon),
            );
          },
        );
      },
      error: (error, stackTrace) {
        return Center(child: Text(error.toString()));
      },
      loading: () {
        return const Center(child: CircularProgressIndicator());
      },
    );
  }
}

final class WeeklyForecastTitle extends StatelessWidget {
  final String day;
  final String date;
  final double temp;
  final String icon;

  const WeeklyForecastTitle({
    super.key,
    required this.day,
    required this.date,
    required this.temp,
    required this.icon,
  });

  @override
  Widget build(BuildContext context) {
    return Container(
      padding: const EdgeInsets.symmetric(horizontal: 16, vertical: 20),
      margin: const EdgeInsets.symmetric(vertical: 12),
      decoration: BoxDecoration(
        borderRadius: BorderRadius.circular(15),
        color: AppColors.accentBlue,
      ),
      child: Row(
        mainAxisAlignment: MainAxisAlignment.spaceBetween,
        children: [
          Column(
            children: [
              Text(day, style: TextStyles.h3),
              const SizedBox(height: 5),
              Text(date, style: TextStyles.subtitleText),
            ],
          ),

          // Temperature
          SuperscriptText(
            text: '$temp',
            color: AppColors.white,
            superScript: '°C',
            superscriptColor: AppColors.white,
          ),

          // weather icon
          Image.asset(icon, width: 60),
        ],
      ),
    );
  }
}
