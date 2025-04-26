import 'package:flutter/foundation.dart';
import 'package:fpdart/fpdart.dart';
import 'package:weather_app/common/errors/exceptions.dart';
import 'package:weather_app/common/errors/failure.dart';
import 'package:weather_app/modules/users/data/datasources/user_remote_datasource.dart';
import 'package:weather_app/modules/users/domain/entities/user.dart';
import 'package:weather_app/modules/users/domain/repositories/user_repository.dart';

@immutable
final class UserRepositoryImpl implements UserRepository {
  final UserRemoteDatasource userRemoteDatasource;

  const UserRepositoryImpl({required this.userRemoteDatasource});

  @override
  Future<Either<Failure, User?>> getCurrent() async {
    try {
      final user = await userRemoteDatasource.getCurrent();

      return right(user);
    } on ServerException catch (e) {
      return left(Failure(e.message));
    }
  }

  @override
  Future<Either<Failure, String>> login({
    required String email,
    required String password,
  }) async {
    try {
      final jwtToken = await userRemoteDatasource.login(
        email: email,
        password: password,
      );

      return right(jwtToken);
    } on ServerException catch (e) {
      return left(Failure(e.message));
    }
  }

  @override
  Future<Either<Failure, String>> register({
    required String username,
    required String email,
    required String password,
  }) async {
    try {
      await userRemoteDatasource.login(email: email, password: password);

      return right("Success");
    } on ServerException catch (e) {
      return left(Failure(e.message));
    }
  }
}
