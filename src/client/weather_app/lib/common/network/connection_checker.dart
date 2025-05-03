import 'package:flutter/foundation.dart';
import 'package:internet_connection_checker_plus/internet_connection_checker_plus.dart';

abstract interface class ConnectionChecker {
  Future<bool> get isConnected;
}

@immutable
final class ConnectionCheckerImpl implements ConnectionChecker {
  final InternetConnection internetConnection;

  const ConnectionCheckerImpl({required this.internetConnection});

  @override
  Future<bool> get isConnected => internetConnection.hasInternetAccess;
}
