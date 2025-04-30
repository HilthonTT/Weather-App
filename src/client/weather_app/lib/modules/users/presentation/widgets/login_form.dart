import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:weather_app/common/layouts/home_layout.dart';
import 'package:weather_app/common/utils/show_snackbar.dart';
import 'package:weather_app/common/widgets/rounded_button.dart';
import 'package:weather_app/common/widgets/rounded_text_field.dart';
import 'package:weather_app/modules/users/presentation/bloc/user_bloc.dart';
import 'package:weather_app/modules/users/presentation/pages/register_page.dart';
import 'package:weather_app/modules/users/presentation/providers/get_current_user_provider.dart';

final class LoginForm extends ConsumerStatefulWidget {
  const LoginForm({super.key});

  @override
  ConsumerState<LoginForm> createState() => _LoginFormState();
}

final class _LoginFormState extends ConsumerState<LoginForm> {
  final formKey = GlobalKey<FormState>();

  final _emailController = TextEditingController(text: "test@test.com");
  final _passwordController = TextEditingController(text: "test");

  void _onLogin() {
    final isValid = formKey.currentState!.validate();

    if (!isValid) {
      return;
    }

    context.read<UserBloc>().add(
      UserLoginEvent(
        email: _emailController.text,
        password: _passwordController.text,
      ),
    );
  }

  @override
  void dispose() {
    super.dispose();

    _emailController.dispose();
    _passwordController.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return BlocConsumer<UserBloc, UserState>(
      listener: (context, state) {
        if (state is UserFailure) {
          showSnackbar(context, state.error);
        } else if (state is UserLoginSuccess) {
          showSnackbar(context, "You've logged in!", icon: Icons.verified);
        }
      },
      builder: (context, state) {
        final isLoading = state is UserLoading;

        return Form(
          key: formKey,
          child: Column(
            children: [
              RoundedTextField(
                enabled: !isLoading,
                controller: _emailController,
                hintText: "Email",
              ),

              const SizedBox(height: 16.0),

              RoundedTextField(
                enabled: !isLoading,
                controller: _passwordController,
                hintText: "Password",
              ),

              const SizedBox(height: 16.0),

              RoundedButton(
                disabled: isLoading,
                text: 'Login',
                onPressed: _onLogin,
              ),

              const SizedBox(height: 16.0),

              RoundedButton(
                disabled: isLoading,
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
      },
    );
  }
}
