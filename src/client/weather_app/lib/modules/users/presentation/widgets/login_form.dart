import 'package:flutter/material.dart';
import 'package:weather_app/common/layouts/home_layout.dart';
import 'package:weather_app/common/widgets/rounded_button.dart';
import 'package:weather_app/common/widgets/rounded_text_field.dart';
import 'package:weather_app/modules/users/presentation/pages/register_page.dart';

final class LoginForm extends StatefulWidget {
  const LoginForm({super.key});

  @override
  State<LoginForm> createState() => _LoginFormState();
}

final class _LoginFormState extends State<LoginForm> {
  final _emailController = TextEditingController();
  final _passwordController = TextEditingController();

  @override
  void dispose() {
    super.dispose();

    _emailController.dispose();
    _passwordController.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Form(
      child: Column(
        children: [
          RoundedTextField(controller: _emailController, hintText: "Email"),

          const SizedBox(height: 16.0),

          RoundedTextField(
            controller: _passwordController,
            hintText: "Password",
          ),

          const SizedBox(height: 16.0),

          RoundedButton(
            text: 'Login',
            onPressed: () {
              // Handle login logic here
              print('Login pressed');
            },
          ),

          const SizedBox(height: 16.0),

          RoundedButton(
            text: 'Nevermind',
            onPressed: () {
              Navigator.push(context, HomeLayout.route());
            },
          ),
          const SizedBox(height: 16.0),

          Align(
            alignment: Alignment.centerLeft,
            child: TextButton(
              onPressed: () {
                Navigator.push(context, RegisterPage.route());
              },
              child: const Text(
                "Don't have an account?",
                style: TextStyle(
                  color: Colors.white,
                  fontWeight: FontWeight.w400,
                ),
              ),
            ),
          ),
        ],
      ),
    );
  }
}
