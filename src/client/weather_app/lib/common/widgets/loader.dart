import 'package:flutter/material.dart';
import 'package:weather_app/common/widgets/gradient_container.dart';

final class LoaderWithBackground extends StatelessWidget {
  const LoaderWithBackground({super.key});

  @override
  Widget build(BuildContext context) {
    return const GradientContainer(
      children: [
        Align(alignment: Alignment.center, child: CircularProgressIndicator()),
      ],
    );
  }
}

final class Loader extends StatelessWidget {
  const Loader({super.key});

  @override
  Widget build(BuildContext context) {
    return Align(
      alignment: Alignment.center,
      child: CircularProgressIndicator(),
    );
  }
}
