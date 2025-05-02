import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:weather_app/common/constants/app_colors.dart';
import 'package:weather_app/common/providers/current_page_index_provider.dart';
import 'package:weather_app/modules/settings/presentation/pages/settings_page.dart';
import 'package:weather_app/modules/weather/presentation/pages/forecast_page.dart';
import 'package:weather_app/modules/weather/presentation/pages/search_page.dart';
import 'package:weather_app/modules/weather/presentation/pages/weather_page.dart';

final class HomeLayout extends ConsumerStatefulWidget {
  const HomeLayout({super.key});

  @override
  ConsumerState<HomeLayout> createState() => _HomeLayoutState();
}

final class _HomeLayoutState extends ConsumerState<HomeLayout> {
  final _destinations = const [
    NavigationDestination(
      icon: Icon(Icons.home_outlined, color: Colors.white),
      selectedIcon: Icon(Icons.home, color: Colors.white),
      label: '',
    ),
    NavigationDestination(
      icon: Icon(Icons.search_outlined, color: Colors.white),
      selectedIcon: Icon(Icons.search, color: Colors.white),
      label: '',
    ),
    NavigationDestination(
      icon: Icon(Icons.wb_sunny_outlined, color: Colors.white),
      selectedIcon: Icon(Icons.wb_sunny, color: Colors.white),
      label: '',
    ),
    NavigationDestination(
      icon: Icon(Icons.settings_outlined, color: Colors.white),
      selectedIcon: Icon(Icons.settings, color: Colors.white),
      label: '',
    ),
  ];

  final _pages = const [
    WeatherPage(),
    SearchPage(),
    ForecastPage(),
    SettingsPage(),
  ];

  @override
  Widget build(BuildContext context) {
    final currentPageIndex = ref.watch(currentPageIndexProvider);

    return Scaffold(
      backgroundColor: AppColors.black,
      bottomNavigationBar: NavigationBarTheme(
        data: NavigationBarThemeData(backgroundColor: AppColors.secondaryBlack),
        child: NavigationBar(
          destinations: _destinations,
          labelBehavior: NavigationDestinationLabelBehavior.alwaysHide,
          selectedIndex: currentPageIndex,
          indicatorColor: Colors.transparent,
          onDestinationSelected: (index) {
            ref.read(currentPageIndexProvider.notifier).state = index;
          },
        ),
      ),
      body: IndexedStack(index: currentPageIndex, children: _pages),
    );
  }
}
