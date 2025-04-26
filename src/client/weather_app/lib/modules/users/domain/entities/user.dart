import 'package:flutter/foundation.dart';
import 'package:weather_app/modules/users/domain/entities/role.dart';

@immutable
class User {
  final int id;
  final String username;
  final String email;
  final DateTime createdAt;
  final bool emailVerified;
  final int roleId;
  final Role role;

  const User({
    required this.id,
    required this.username,
    required this.email,
    required this.createdAt,
    required this.emailVerified,
    required this.roleId,
    required this.role,
  });
}
