import 'package:flutter/foundation.dart';

@immutable
final class ServerException implements Exception {
  final String message;

  const ServerException(this.message);
}
