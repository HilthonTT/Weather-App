import 'package:flutter/foundation.dart';
import 'package:weather_app/modules/users/data/models/role_model.dart';
import 'package:weather_app/modules/users/domain/entities/user.dart';

@immutable
final class UserModel extends User {
  const UserModel({
    required super.id,
    required super.username,
    required super.email,
    required super.createdAt,
    required super.emailVerified,
    required super.roleId,
    required super.role,
  });

  factory UserModel.fromJson(Map<String, dynamic> json) {
    return UserModel(
      id: (json['id'] as num).toInt(),
      username: json['username'] as String,
      email: json['email'] as String,
      createdAt: DateTime.parse(json['created_at'] as String),
      emailVerified: json['email_verified'] as bool,
      roleId: (json['role_id'] as num).toInt(),
      role: RoleModel.fromJson(json['role'] as Map<String, dynamic>),
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'username': username,
      'email': email,
      'created_at': createdAt.toIso8601String(),
      'email_verified': emailVerified,
      'role_id': roleId,
      'role': role,
    };
  }
}
