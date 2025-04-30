import 'package:flutter/material.dart';
import 'package:weather_app/common/constants/app_colors.dart';

final class RoundedTextField extends StatelessWidget {
  final Function(String)? onSubmitted;
  final TextEditingController controller;
  final String hintText;
  final bool? enabled;

  const RoundedTextField({
    super.key,
    required this.controller,
    this.hintText = "Search",
    this.onSubmitted,
    this.enabled = true,
  });

  @override
  Widget build(BuildContext context) {
    final effectiveColor =
        enabled == true ? AppColors.accentBlue : Colors.grey[600];
    final effectiveTextColor =
        enabled == true ? Colors.white : Colors.grey[400];

    return Container(
      height: 55,
      decoration: BoxDecoration(
        borderRadius: BorderRadius.circular(10),
        color: effectiveColor,
      ),
      child: TextField(
        enabled: enabled,
        onSubmitted: onSubmitted,
        controller: controller,
        style: TextStyle(color: effectiveTextColor),
        decoration: InputDecoration(
          contentPadding: const EdgeInsets.only(left: 20, top: 5),
          border: InputBorder.none,
          hintText: hintText,
          hintStyle: TextStyle(
            color: effectiveTextColor,
            fontWeight: FontWeight.w400,
          ),
        ),
      ),
    );
  }
}
