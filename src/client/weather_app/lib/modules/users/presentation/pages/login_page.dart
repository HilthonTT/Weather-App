import 'package:flutter/material.dart';
import 'package:weather_app/common/constants/app_colors.dart';
import 'package:weather_app/common/widgets/gradient_container.dart';
import 'package:weather_app/modules/users/presentation/widgets/login_form.dart';
import 'package:weather_app/modules/users/presentation/widgets/top_image.dart';

final class LoginPage extends StatelessWidget {
  static route() => MaterialPageRoute(builder: (context) => const LoginPage());

  const LoginPage({super.key});

  @override
  Widget build(BuildContext context) {
    return Material(
      color: AppColors.black,
      child: GradientContainer(
        children: [
          const Column(
            mainAxisAlignment: MainAxisAlignment.center,
            crossAxisAlignment: CrossAxisAlignment.center,
            children: [
              TopImage(title: "Login", image: "assets/icons/01d.png"),
              Row(
                children: [
                  Spacer(),
                  Expanded(flex: 50, child: LoginForm()),
                  Spacer(),
                ],
              ),
            ],
          ),
        ],
      ),
    );
  }
}
