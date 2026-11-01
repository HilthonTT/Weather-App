import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:weather_app/common/constants/app_colors.dart';
import 'package:weather_app/common/constants/text_styles.dart';
import 'package:weather_app/common/utils/show_snackbar.dart';
import 'package:weather_app/common/widgets/gradient_container.dart';
import 'package:weather_app/common/widgets/loader.dart';
import 'package:weather_app/modules/settings/domain/entities/settings.dart';
import 'package:weather_app/modules/settings/presentation/utils/format.dart';
import 'package:weather_app/modules/settings/presentation/bloc/settings_bloc.dart';
import 'package:weather_app/modules/settings/presentation/widgets/icon_item_row.dart';

/// Lets the user pick temperature, time and wind-speed units.
///
/// Reads the current settings from [SettingsBloc] (never from a local copy),
/// so successive changes to different fields are all preserved.
final class SettingsPage extends StatelessWidget {
  static const int currentIndex = 3;

  const SettingsPage({super.key});

  void _onUpdateSettings(BuildContext context, Settings settings) {
    context.read<SettingsBloc>().add(UpdateSettingsEvent(settings));
  }

  @override
  Widget build(BuildContext context) {
    return BlocConsumer<SettingsBloc, SettingsState>(
          listener: (context, state) {
            if (state is SettingsFailure) {
              showSnackbar(context, state.error, icon: Icons.error);
            } else if (state is SettingsUpdateSuccess) {
              showSnackbar(context, "Settings updated!");
            }
          },
          builder: (context, state) {
            final settings = context.read<SettingsBloc>().currentSettings;

            if (settings == null) {
              return GradientContainer(
                children: [
                  if (state is SettingsFailure)
                    const Center(
                      child: Text(
                        'An error has occurred',
                        style: TextStyles.h1,
                      ),
                    )
                  else
                    const Loader(),
                ],
              );
            }

            return Scaffold(
              backgroundColor: AppColors.black,
              body: Stack(
                children: [
                  GradientContainer(
                    children: [
                      Column(
                        children: [
                          Row(
                            mainAxisAlignment: MainAxisAlignment.center,
                            children: [
                              const Text("Settings", style: TextStyles.h1),
                            ],
                          ),

                          const SizedBox(height: 20),

                          Row(
                            mainAxisAlignment: MainAxisAlignment.center,
                            children: [
                              Image.asset(
                                "assets/icons/user.png",
                                width: 70,
                                height: 70,
                              ),
                            ],
                          ),

                          const SizedBox(height: 10),

                          Padding(
                            padding: const EdgeInsets.symmetric(
                              vertical: 20,
                              horizontal: 20,
                            ),
                            child: Column(
                              crossAxisAlignment: CrossAxisAlignment.start,
                              children: [
                                Padding(
                                  padding: const EdgeInsets.only(
                                    top: 10,
                                    bottom: 8,
                                  ),
                                  child: const Text(
                                    "Formatting",
                                    style: TextStyles.h3,
                                  ),
                                ),
                                Container(
                                  padding: const EdgeInsets.symmetric(
                                    vertical: 8,
                                  ),
                                  decoration: BoxDecoration(
                                    border: Border.all(
                                      color: AppColors.border.withValues(
                                        alpha: .1,
                                      ),
                                    ),
                                    color: AppColors.gray60.withValues(
                                      alpha: .2,
                                    ),
                                    borderRadius: BorderRadius.circular(16),
                                  ),
                                  child: Column(
                                    children: [
                                      IconItemDropdownRow(
                                        title: "Time Format",
                                        icon: "assets/icons/time.png",
                                        options: ["12h", "24h"],
                                        selectedValue: formatTimeFormat(
                                          settings.timeFormat,
                                        ),
                                        onChanged: (val) {
                                          final timeFormat =
                                              val == "12h"
                                                  ? TimeFormat.twelveHour
                                                  : TimeFormat.twentyFourHour;

                                          final updatedSettings = settings
                                              .copyWith(timeFormat: timeFormat);
                                          _onUpdateSettings(context, updatedSettings);
                                        },
                                      ),

                                      IconItemDropdownRow(
                                        title: "Speed Format",
                                        icon: "assets/icons/speed.png",
                                        options: ["Km/h", "Mi/h"],
                                        selectedValue: formatSpeedFormat(
                                          settings.speedFormat,
                                        ),
                                        onChanged: (val) {
                                          final speedFormat =
                                              val == "Km/h"
                                                  ? SpeedFormat.kmph
                                                  : SpeedFormat.mph;

                                          final updatedSettings = settings
                                              .copyWith(
                                                speedFormat: speedFormat,
                                              );
                                          _onUpdateSettings(context, updatedSettings);
                                        },
                                      ),

                                      IconItemDropdownRow(
                                        title: "Temperature Format",
                                        icon: "assets/icons/temp.png",
                                        options: ["Celsius", "Fahrenheit"],
                                        selectedValue: formatTempFormat(
                                          settings.tempFormat,
                                        ),
                                        onChanged: (val) {
                                          final tempFormat =
                                              val == "Celsius"
                                                  ? TempFormat.celsius
                                                  : TempFormat.fahrenheit;

                                          final updatedSettings = settings
                                              .copyWith(tempFormat: tempFormat);
                                          _onUpdateSettings(context, updatedSettings);
                                        },
                                      ),
                                    ],
                                  ),
                                ),

                                const SizedBox(height: 20),
                              ],
                            ),
                          ),
                        ],
                      ),
                    ],
                  ),
                ],
              ),
            );
          },
        );
  }
}
