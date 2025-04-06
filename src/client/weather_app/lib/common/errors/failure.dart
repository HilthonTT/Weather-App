import 'package:flutter/foundation.dart';

@immutable
final class Failure {
  final String message;

  const Failure([this.message = "An unexpected error occured"]);
}
