import 'package:flutter/material.dart';
import 'package:weather_app/common/constants/app_colors.dart';

final class GradientContainer extends StatelessWidget {
  final List<Widget> children;

  const GradientContainer({super.key, required this.children});

  @override
  Widget build(BuildContext context) {
    final screenSize = MediaQuery.of(context).size;

    return SafeArea(
      child: Container(
        height: screenSize.height,
        width: screenSize.width,
        decoration: BoxDecoration(
          gradient: LinearGradient(
            begin: Alignment.bottomLeft,
            end: Alignment.topRight,
            colors: [
              AppColors.black,
              AppColors.secondaryBlack,
              AppColors.secondaryBlack.withValues(alpha: 0.99),
              AppColors.secondaryBlack.withValues(alpha: 0.98),
              AppColors.secondaryBlack.withValues(alpha: 0.97),
              AppColors.secondaryBlack.withValues(alpha: 0.96),
              AppColors.secondaryBlack.withValues(alpha: 0.95),
              AppColors.secondaryBlack.withValues(alpha: 0.94),
              AppColors.secondaryBlack.withValues(alpha: 0.93),
              AppColors.secondaryBlack.withValues(alpha: 0.92),
              AppColors.secondaryBlack.withValues(alpha: 0.91),
              AppColors.secondaryBlack.withValues(alpha: 0.90),
              AppColors.darkBlue,
              AppColors.accentBlue,
              AppColors.lightBlue,
            ],
          ),
        ),
        child: Padding(
          padding: const EdgeInsets.only(left: 24.0, right: 24.0, top: 36.0),
          child: CustomScrollView(
            slivers: [SliverList(delegate: SliverChildListDelegate(children))],
          ),
        ),
      ),
    );
  }
}
