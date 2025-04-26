import 'package:flutter/foundation.dart';
import 'package:weather_app/modules/users/domain/entities/role.dart';

@immutable
final class RoleModel extends Role {
  const RoleModel({
    required super.id,
    required super.name,
    required super.description,
    required super.level,
  });

  factory RoleModel.fromJson(Map<String, dynamic> json) {
    return RoleModel(
      id: (json['id'] as num).toInt(),
      name: json['name'] as String,
      description: json['description'] as String,
      level: (json['level'] as num).toInt(),
    );
  }

  Map<String, dynamic> toJson() {
    return {'id': id, 'name': name, 'description': description, 'level': level};
  }
}
