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
      throw ServerException('No JWT token found, user not logged in');
    }

    const url = "${Constants.apiUrl}/v1/users/me";

    final response = await _dio.get(
      url,
      options: Options(
        headers: {'Authorization': 'Bearer $jwtToken'},
        validateStatus: (status) => true,
      ),
    );

    final statusCode = response.statusCode ?? 0;
    final responseData = response.data;

    if (statusCode == 200 &&
        responseData is Map &&
        responseData.containsKey('data')) {
      final user = UserModel.fromJson(responseData['data']);
      return user;
    }

    // Handling different status codes
    if (statusCode == 401) {
      throw ServerException('Unauthorized: Invalid token or expired session');
    }

    if (statusCode == 404) {
      throw ServerException('User not found');
    }

    if (statusCode == 500) {
      throw ServerException('Server error, please try again later');
    }

    throw ServerException('Unexpected response: $responseData');
  }

  @override
  Future<String> login({
    required String email,
    required String password,
  }) async {
    const url = "${Constants.apiUrl}/v1/authentication/login";

    final response = await _dio.post(
      url,
      data: {'email': email, 'password': password},
      options: Options(validateStatus: (status) => true),
    );

    final statusCode = response.statusCode ?? 0;
    final responseData = response.data;

    if (statusCode == 201 &&
        responseData is Map &&
        responseData.containsKey('data')) {
      final jwtToken = responseData['data'] as String;
      box.put(UserConstants.jwtTokenKey, jwtToken);

      return jwtToken;
    }

    if (statusCode == 401) {
      throw ServerException('Invalid email or password');
    }

    if (statusCode == 500) {
      throw ServerException('Server error, please try again later');
    }

    throw ServerException('Unexpected response: $responseData');
  }

  @override
  Future<void> register({
    required String username,
    required String email,
    required String password,
  }) async {
    const url = "${Constants.apiUrl}/v1/authentication/register";

    final response = await _dio.post(
      url,
      data: {'email': email, 'password': password, 'username': username},
      options: Options(validateStatus: (status) => true),
    );

    final statusCode = response.statusCode ?? 0;
    final data = response.data;

    if (statusCode == 201) {
      return;
    }

    if (statusCode == 400) {
      throw ServerException(
        'Invalid input: ${data['message'] ?? 'Bad request'}',
      );
    }

    if (statusCode == 409) {
      throw ServerException('Email is already registered');
    }

    if (statusCode == 500) {
      throw ServerException('Server error, please try again later');
    }

    throw ServerException('Unexpected response: $data');
  }

  @override
  void logout() {
    box.delete(UserConstants.jwtTokenKey);
  }
}
