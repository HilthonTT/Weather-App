import 'package:flutter/foundation.dart';
import 'package:fpdart/fpdart.dart';
import 'package:weather_app/common/errors/failure.dart';
import 'package:weather_app/common/usecase/usecase.dart';
import 'package:weather_app/modules/users/domain/repositories/user_repository.dart';

@immutable
final class LoginUser implements UseCase<String, LoginUserParams> {
  final UserRepository userRepository;

  const LoginUser({required this.userRepository});

  @override
  Future<Either<Failure, String>> call(LoginUserParams params) {
    return userRepository.login(email: params.email, password: params.password);
  }
}

@immutable
final class LoginUserParams {
  final String email;
  final String password;

  const LoginUserParams({required this.email, required this.password});
}
