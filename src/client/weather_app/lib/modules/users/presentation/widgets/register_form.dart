import 'package:flutter/material.dart';
import 'package:weather_app/common/layouts/home_layout.dart';
import 'package:weather_app/common/widgets/rounded_button.dart';
import 'package:weather_app/common/widgets/rounded_text_field.dart';
import 'package:weather_app/modules/users/presentation/pages/login_page.dart';

final class RegisterForm extends StatefulWidget {
  const RegisterForm({super.key});

  @override
  State<RegisterForm> createState() => _RegisterFormState();
}

final class _RegisterFormState extends State<RegisterForm> {
  final _emailController = TextEditingController();
  final _passwordController = TextEditingController();
  final _usernameController = TextEditingController();

  @override
  void dispose() {
    super.dispose();

    _emailController.dispose();
    _passwordController.dispose();
    _usernameController.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Form(
      child: Column(
        children: [
          RoundedTextField(controller: _emailController, hintText: "Email"),

          const SizedBox(height: 16.0),

          RoundedTextField(
            controller: _usernameController,
            hintText: "Username",
          ),

          const SizedBox(height: 16.0),

          RoundedTextField(
            controller: _passwordController,
            hintText: "Password",
          ),

          const SizedBox(height: 16.0),

          RoundedButton(
            text: 'Register',
            onPressed: () {
              // Handle register logic here
              print('Register pressed');
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
                Navigator.push(context, LoginPage.route());
              },
              child: const Text(
                "Already have an account?",
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
