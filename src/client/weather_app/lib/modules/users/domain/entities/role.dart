import 'package:flutter/foundation.dart';

@immutable
class Role {
  final int id;
  final String name;
  final String description;
  final int level;

  const Role({
    required this.id,
    required this.name,
    required this.description,
    required this.level,
  });
}
