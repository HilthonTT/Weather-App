import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:weather_app/common/constants/app_colors.dart';
import 'package:weather_app/common/constants/text_styles.dart';
import 'package:weather_app/common/widgets/custom_back_button.dart';
import 'package:weather_app/common/widgets/gradient_container.dart';
import 'package:weather_app/modules/users/presentation/pages/login_page.dart';
import 'package:weather_app/modules/users/presentation/providers/get_current_user_provider.dart';
import 'package:weather_app/modules/users/presentation/widgets/icon_item_row.dart';

final class SettingsPage extends ConsumerStatefulWidget {
  const SettingsPage({super.key});

  @override
  ConsumerState<ConsumerStatefulWidget> createState() => _SettingsPageState();
}

final class _SettingsPageState extends ConsumerState<SettingsPage> {
  @override
  Widget build(BuildContext context) {
    final getCurrentUser = ref.watch(getCurrentUserProvider);

    return getCurrentUser.when(
      data: (user) {
        if (user == null) {
          WidgetsBinding.instance.addPostFrameCallback((_) {
            Navigator.push(context, LoginPage.route());
          });
          return const GradientContainer(
            children: [
              Align(
                alignment: Alignment.center,
                child: CircularProgressIndicator(),
              ),
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
                      InkWell(
                        borderRadius: BorderRadius.circular(15),
                        onTap: () {},
                        child: Container(
                          padding: const EdgeInsets.all(6),
                          decoration: BoxDecoration(
                            border: Border.all(
                              color: AppColors.border.withValues(alpha: .15),
                            ),
                            color: AppColors.gray60.withValues(alpha: .2),
                            borderRadius: BorderRadius.circular(15),
                          ),
                          child: Text(
                            "Edit profile",
                            style: TextStyle(
                              color: AppColors.white,
                              fontSize: 12,
                              fontWeight: FontWeight.w600,
                            ),
                          ),
                        ),
                      ),
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
                                top: 20,
                                bottom: 8,
                              ),
                              child: const Text(
                                "Formatting",
                                style: TextStyles.h3,
                              ),
                            ),
                            Container(
                              padding: const EdgeInsets.symmetric(vertical: 8),
                              decoration: BoxDecoration(
                                border: Border.all(
                                  color: AppColors.border.withValues(alpha: .1),
                                ),
                                color: AppColors.gray60.withValues(alpha: .2),
                                borderRadius: BorderRadius.circular(16),
                              ),
                              child: const Column(
                                children: [
                                  IconItemRow(
                                    title: "Time Format",
                                    icon: "assets/icons/time.png",
                                    value: "24h",
                                  ),
                                  SizedBox(height: 10),
                                  IconItemRow(
                                    title: "Speed Format",
                                    icon: "assets/icons/speed.png",
                                    value: "Km/h",
                                  ),
                                  SizedBox(height: 10),
                                  IconItemRow(
                                    title: "Temperature Format",
                                    icon: "assets/icons/temp.png",
                                    value: "Celcius",
                                  ),
                                ],
                              ),
                            ),
                          ],
                        ),
                      ),
                    ],
                  ),
                ],
              ),
              const CustomBackButton(),
            ],
          ),
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
