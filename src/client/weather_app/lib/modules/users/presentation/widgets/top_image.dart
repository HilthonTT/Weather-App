import 'package:flutter/material.dart';
import 'package:weather_app/common/constants/text_styles.dart';

final class TopImage extends StatelessWidget {
  final String title;
  final String image;

  const TopImage({super.key, required this.title, required this.image});

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        Text(title, style: TextStyles.h1),
        const SizedBox(height: 32.0),
        Row(
          children: [
            const Spacer(),
            Expanded(
              child: Image.asset(
                image,
                fit: BoxFit.contain,
                width: 200,
                height: 200,
              ),
            ),
            const Spacer(),
          ],
        ),

        const SizedBox(height: 32.0),
      ],
    );
  }
}
