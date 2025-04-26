import 'package:fpdart/fpdart.dart';
import 'package:weather_app/common/errors/failure.dart';
import 'package:weather_app/modules/users/domain/entities/user.dart';

abstract interface class UserRepository {
  Future<Either<Failure, User?>> getCurrent();

  Future<Either<Failure, String>> login({
    required String email,
    required String password,
  });

  Future<Either<Failure, String>> register({
    required String username,
    required String email,
    required String password,
  });
}
