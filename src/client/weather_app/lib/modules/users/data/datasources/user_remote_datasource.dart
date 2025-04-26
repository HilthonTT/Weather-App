import 'package:dio/dio.dart';
import 'package:flutter/foundation.dart';
import 'package:hive/hive.dart';
import 'package:weather_app/common/constants/constants.dart';
import 'package:weather_app/common/errors/exceptions.dart';
import 'package:weather_app/modules/users/data/models/user_model.dart';
import 'package:weather_app/modules/users/data/user_constants.dart';

abstract interface class UserRemoteDatasource {
  Future<UserModel?> getCurrent();

  Future<String> login({required String email, required String password});

  Future<void> register({
    required String username,
    required String email,
    required String password,
  });

  void logout();
}

@immutable
final class UserRemoteDatasourceImpl implements UserRemoteDatasource {
  static final _dio = Dio();

  final Box box;

  const UserRemoteDatasourceImpl({required this.box});

  @override
  Future<UserModel?> getCurrent() async {
    final jwtToken = box.get(UserConstants.jwtTokenKey) as String?;
    if (jwtToken == null || jwtToken.isEmpty) {
      return null;
    }

    const url = "${Constants.apiUrl}/v1/users/me";

    final response = await _dio.get(
      url,
      options: Options(headers: {'Authorization': 'Bearer $jwtToken'}),
    );

    final user = UserModel.fromJson(response.data['data']);

    return user;
  }

  @override
  Future<String> login({
    required String email,
    required String password,
  }) async {
    try {
      const url = "${Constants.apiUrl}/v1/authentication/login";

      final response = await _dio.post(
        url,
        data: {'email': email, 'password': password},
      );

      final jwtToken = response.data['data'] as String;

      box.put(UserConstants.jwtTokenKey, jwtToken);

      return jwtToken;
    } catch (e) {
      debugPrint(e.toString());

      throw ServerException('Failed to login');
    }
  }

  @override
  Future<void> register({
    required String username,
    required String email,
    required String password,
  }) async {
    try {
      const url = "${Constants.apiUrl}/v1/authentication/login";

      await _dio.post(
        url,
        data: {'email': email, 'password': password, 'username': username},
      );
    } catch (e) {
      debugPrint(e.toString());

      throw ServerException('Failed to login');
    }
  }

  @override
  void logout() {
    box.delete(UserConstants.jwtTokenKey);
  }
}
