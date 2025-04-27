import 'package:flutter/foundation.dart';
import 'package:fpdart/fpdart.dart';
import 'package:weather_app/common/errors/failure.dart';
import 'package:weather_app/common/usecase/usecase.dart';
import 'package:weather_app/modules/users/domain/repositories/user_repository.dart';

@immutable
final class RegisterUser implements UseCase<String, RegisterUserParams> {
  final UserRepository userRepository;

  const RegisterUser({required this.userRepository});

  @override
  Future<Either<Failure, String>> call(RegisterUserParams params) {
    return userRepository.register(
      username: params.username,
      email: params.email,
      password: params.password,
    );
  }
}

@immutable
final class RegisterUserParams {
  final String username;
  final String email;
  final String password;

  const RegisterUserParams({
    required this.username,
    required this.email,
    required this.password,
  });
}
