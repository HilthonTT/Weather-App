import 'package:flutter/foundation.dart';
import 'package:fpdart/fpdart.dart';
import 'package:weather_app/common/errors/failure.dart';
import 'package:weather_app/common/usecase/usecase.dart';
import 'package:weather_app/modules/users/domain/entities/user.dart';
import 'package:weather_app/modules/users/domain/repositories/user_repository.dart';

@immutable
final class GetCurrentUser implements UseCase<User?, NoParams> {
  final UserRepository userRepository;

  const GetCurrentUser({required this.userRepository});

  @override
  Future<Either<Failure, User?>> call(NoParams params) {
    return userRepository.getCurrent();
  }
}
