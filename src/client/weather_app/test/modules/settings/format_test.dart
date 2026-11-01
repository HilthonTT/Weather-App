import 'package:flutter_test/flutter_test.dart';
import 'package:weather_app/modules/settings/domain/entities/settings.dart';
import 'package:weather_app/modules/settings/presentation/utils/format.dart';

const _celsius = Settings(
  id: 1,
  userId: 1,
  tempFormat: TempFormat.celsius,
  timeFormat: TimeFormat.twentyFourHour,
  speedFormat: SpeedFormat.kmph,
);

void main() {
  group('formatTemperature', () {
    test('defaults to Celsius when settings are not loaded', () {
      expect(formatTemperature(20, null), '20.0°C');
      expect(temperatureUnit(null), '°C');
    });

    test('converts to Fahrenheit', () {
      final settings = _celsius.copyWith(tempFormat: TempFormat.fahrenheit);

      expect(formatTemperature(20, settings), '68.0°F');
      expect(temperatureValue(-40, settings), '-40.0');
      expect(temperatureUnit(settings), '°F');
    });
  });

  group('formatWindSpeed', () {
    test('converts m/s to km/h by default', () {
      expect(formatWindSpeed(5, null), '18.00 km/h');
    });

    test('converts m/s to mph', () {
      final settings = _celsius.copyWith(speedFormat: SpeedFormat.mph);

      expect(formatWindSpeed(10, settings), '22.37 mph');
    });
  });

  test('dropdown labels', () {
    expect(formatTempFormat(TempFormat.fahrenheit), 'Fahrenheit');
    expect(formatTimeFormat(TimeFormat.twelveHour), '12h');
    expect(formatSpeedFormat(SpeedFormat.mph), 'Mi/h');
  });
}
