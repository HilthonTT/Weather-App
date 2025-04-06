import 'package:jiffy/jiffy.dart';

extension ConvertTimestampToTime on int {
  /// Returns the time in `HH:mm` format (e.g., 14:30) from a UNIX timestamp.
  String get time =>
      Jiffy.parseFromMillisecondsSinceEpoch(this * 1000).format(pattern: 'Hm');

  /// Returns the full date in `yMMMMd` format (e.g., April 6, 2025) from a UNIX timestamp.
  String get formattedDate => Jiffy.parseFromMillisecondsSinceEpoch(
    this * 1000,
  ).format(pattern: 'yMMMMd');
}
