import 'package:flutter/material.dart';
import 'package:weather_app/common/constants/app_colors.dart';
import 'package:weather_app/common/constants/text_styles.dart';
import 'package:weather_app/common/extensions/datetime.dart';
import 'package:weather_app/common/widgets/gradient_container.dart';
import 'package:weather_app/modules/weather/presentation/widgets/hourly_forecast.dart';
import 'package:weather_app/modules/weather/presentation/widgets/weekly_forecast.dart';

final class ForecastPage extends StatelessWidget {
  static const int currentIndex = 2;

  const ForecastPage({super.key});

  @override
  Widget build(BuildContext context) {
    final today = DateTime.now().formattedDate;

    return Scaffold(
      backgroundColor: AppColors.black,
      body: GradientContainer(
        children: [
          const Align(
            alignment: Alignment.center,
            child: Text('Forecast report', style: TextStyles.h1),
          ),
          const SizedBox(height: 40),
          Row(
            mainAxisAlignment: MainAxisAlignment.spaceBetween,
            children: [
              const Text('Today', style: TextStyles.h2),
              Text(today, style: TextStyles.subtitleText),
            ],
          ),
          const SizedBox(height: 20),

          const HourlyForecast(),

          const SizedBox(height: 20),

          Row(
            mainAxisAlignment: MainAxisAlignment.spaceBetween,
            children: [
              const Text('Next forecast', style: TextStyles.h1),
              IconButton(
                onPressed: () {},
                icon: const Icon(
                  Icons.calendar_month_outlined,
                  color: AppColors.white,
                ),
              ),
            ],
          ),

          const SizedBox(height: 20),

          const WeeklyForecast(),
        ],
      ),
    );
  }
}
