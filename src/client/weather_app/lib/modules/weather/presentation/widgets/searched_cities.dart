import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:weather_app/common/constants/app_colors.dart';
import 'package:weather_app/common/models/city.dart';

final class SearchedCities extends StatelessWidget {
  final List<City> cities;

  const SearchedCities({super.key, required this.cities});

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

        return InkWell(onTap: () {}, child: CityTile(index: index));
      },
    );
  }
}

final class CityTile extends ConsumerWidget {
  final int index;

  const CityTile({super.key, required this.index});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    return Material(
      color: index == 0 ? AppColors.lightBlue : AppColors.accentBlue,
      elevation: index == 0 ? 8 : 0,
      borderRadius: BorderRadius.circular(25),
    );
  }
}
