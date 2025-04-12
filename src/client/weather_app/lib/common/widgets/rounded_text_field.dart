import 'package:flutter/material.dart';
import 'package:weather_app/common/constants/app_colors.dart';

final class RoundedTextField extends StatelessWidget {
  final Function(String)? onSubmitted;
  final TextEditingController controller;
  final String hintText;

  const RoundedTextField({
    super.key,
    required this.controller,
    this.hintText = "Search",
    this.onSubmitted,
  });

  @override
  Widget build(BuildContext context) {
    return Container(
      height: 55,
      decoration: BoxDecoration(
        borderRadius: BorderRadius.circular(10),
        color: AppColors.accentBlue,
      ),
      child: TextField(
        onSubmitted: onSubmitted,
        controller: controller,
        style: const TextStyle(color: Colors.white),
        decoration: InputDecoration(
          contentPadding: EdgeInsets.only(left: 20, top: 5),
          border: InputBorder.none,
          fillColor: Colors.white,
          focusColor: Colors.white,
          hintText: hintText,
          hintStyle: TextStyle(color: Colors.grey, fontWeight: FontWeight.w400),
        ),
      ),
    );
  }
}
