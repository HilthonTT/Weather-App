import 'package:flutter/material.dart';
import 'package:weather_app/common/constants/app_colors.dart';

final class RoundedButton extends StatelessWidget {
  final String text;
  final VoidCallback onPressed;

  const RoundedButton({super.key, required this.text, required this.onPressed});

  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onTap: onPressed,
      child: Container(
        height: 55,
        decoration: BoxDecoration(
          color: AppColors.accentBlue,
          borderRadius: BorderRadius.circular(10),
        ),
        child: Center(
          child: Text(
            text.toUpperCase(),
            style: const TextStyle(
              color: Colors.white,
              fontWeight: FontWeight.bold,
              fontSize: 16,
            ),
          ),
        ),
      ),
    );
  }
}
