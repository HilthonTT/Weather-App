import 'dart:convert';

import 'package:dio/dio.dart';
import 'package:flutter/foundation.dart';
import 'package:hive/hive.dart';
import 'package:weather_app/common/constants/constants.dart';
import 'package:weather_app/common/errors/exceptions.dart';
import 'package:weather_app/modules/users/data/models/settings_model.dart';
import 'package:weather_app/modules/users/data/user_constants.dart';

abstract interface class SettingsRemoteDatasource {
  Future<SettingsModel> getSettings();

  Future<SettingsModel> updateSettings({required SettingsModel settings});
}

@immutable
final class SettingsRemoteDatasourceImpl implements SettingsRemoteDatasource {
  static final _dio = Dio();

  final Box box;

  const SettingsRemoteDatasourceImpl({required this.box});

  @override
  Future<SettingsModel> getSettings() async {
    try {
      final jwtToken = box.get(UserConstants.jwtTokenKey) as String?;
      if (jwtToken == null || jwtToken.isEmpty) {
        throw ServerException('Unauthorized');
      }

      final userId = getUserIdFromToken(jwtToken);

      final url = "${Constants.apiUrl}/v1/users/$userId/settings";

      final response = await _dio.get(
        url,
        options: Options(headers: {'Authorization': 'Bearer $jwtToken'}),
      );

      final settings = SettingsModel.fromJson(response.data['data']);

      return settings;
    } catch (e) {
      debugPrint(e.toString());

      throw ServerException('Failed to fetch settings');
    }
  }

  @override
  Future<SettingsModel> updateSettings({
    required SettingsModel settings,
  }) async {
    try {
      final jwtToken = box.get(UserConstants.jwtTokenKey) as String?;
      if (jwtToken == null || jwtToken.isEmpty) {
        throw ServerException('Unauthorized');
      }

      final userId = getUserIdFromToken(jwtToken);

      final url = "${Constants.apiUrl}/v1/users/$userId/settings";

      final response = await _dio.patch(
        url,
        data: {
          'speed_format': settings.speedFormat,
          'temp_format': settings.tempFormat,
          'time_format': settings.timeFormat,
        },
        options: Options(headers: {'Authorization': 'Bearer $jwtToken'}),
      );

      final updatedSettings = SettingsModel.fromJson(response.data['data']);

      return updatedSettings;
    } catch (e) {
      debugPrint(e.toString());

      throw ServerException('Failed to fetch settings');
    }
  }

  int getUserIdFromToken(String jwtToken) {
    final parts = jwtToken.split('.');
    if (parts.length != 3) {
      throw FormatException('Invalid JWT token');
    }

    final payload = parts[1];

    // Base64 decode
    final normalized = base64.normalize(payload);
    final decoded = utf8.decode(base64Url.decode(normalized));

    final payloadMap = json.decode(decoded) as Map<String, dynamic>;

    if (!payloadMap.containsKey('sub')) {
      throw FormatException('JWT does not contain sub');
    }

    return int.parse(payloadMap['sub'].toString());
  }
}
