import 'package:flutter/material.dart';
import 'package:weather_app/common/constants/app_colors.dart';

final class RoundedButton extends StatelessWidget {
  final String text;
  final VoidCallback onPressed;
  final bool disabled;

  const RoundedButton({
    super.key,
    required this.text,
    required this.onPressed,
    this.disabled = false,
  });

  @override
  Widget build(BuildContext context) {
    final effectiveOnTap = disabled ? null : onPressed;
    final buttonColor = disabled ? Colors.grey[400] : AppColors.accentBlue;
    final textColor = disabled ? Colors.grey[600] : Colors.white;

    return GestureDetector(
      onTap: effectiveOnTap,
      child: Container(
        height: 55,
        decoration: BoxDecoration(
          color: buttonColor,
          borderRadius: BorderRadius.circular(10),
        ),
        child: Center(
          child: Text(
            text.toUpperCase(),
            style: TextStyle(
              color: textColor,
              fontWeight: FontWeight.bold,
              fontSize: 16,
            ),
          ),
        ),
      ),
    );
  }
}
