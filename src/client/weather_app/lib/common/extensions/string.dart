extension Capitalize on String {
  /// Capitalizes the first letter of the string.
  String get capitalized {
    final trimmed = trim();

    return trimmed.isNotEmpty
        ? '${trimmed[0].toUpperCase()}${trimmed.substring(1)}'
        : '';
  }
}
