import 'package:flutter_test/flutter_test.dart';
import 'package:weather_app/modules/settings/data/models/settings_model.dart';
import 'package:weather_app/modules/settings/domain/entities/settings.dart';

void main() {
  test('SettingsModel round-trips through JSON', () {
    const model = SettingsModel(
      id: 3,
      userId: 7,
      tempFormat: TempFormat.fahrenheit,
      timeFormat: TimeFormat.twelveHour,
      speedFormat: SpeedFormat.mph,
    );

    final decoded = SettingsModel.fromJson(model.toJson());

    expect(decoded.id, 3);
    expect(decoded.userId, 7);
    expect(decoded.tempFormat, TempFormat.fahrenheit);
    expect(decoded.timeFormat, TimeFormat.twelveHour);
    expect(decoded.speedFormat, SpeedFormat.mph);
  });

  test('unknown enum values fall back to defaults', () {
    final decoded = SettingsModel.fromJson({
      'id': 1,
      'userId': 1,
      'tempFormat': 'kelvin',
      'timeFormat': '?',
      'speedFormat': '?',
    });

    expect(decoded.tempFormat, TempFormat.celsius);
    expect(decoded.timeFormat, TimeFormat.twentyFourHour);
    expect(decoded.speedFormat, SpeedFormat.kmph);
  });
}
