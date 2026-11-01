import 'package:flutter/foundation.dart';

@immutable
final class Constants {
  /// Base URL of the Go backend, without a trailing slash.
  ///
  /// Override at build/run time:
  /// ```sh
  /// flutter run --dart-define=API_URL=http://192.168.1.11:8080
  /// ```
  /// The default `10.0.2.2` is how the Android emulator reaches the host
  /// machine's `localhost`. Physical devices need your computer's LAN IP.
  static const String apiUrl = String.fromEnvironment(
    'API_URL',
    defaultValue: 'http://10.0.2.2:8080',
  );
}
