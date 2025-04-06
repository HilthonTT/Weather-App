import 'package:jiffy/jiffy.dart';

extension DateTimeFormatting on DateTime {
  /// Returns the date in `yMMMMd` format (e.g., April 6, 2025).
  String get formattedDate =>
      Jiffy.parseFromDateTime(this).format(pattern: 'yMMMMd');

  /// Returns the full day of the week (e.g., Sunday).
  String get dayOfWeek => Jiffy.parseFromDateTime(this).format(pattern: 'EEEE');
}
