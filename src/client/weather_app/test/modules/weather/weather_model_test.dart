import 'package:flutter_test/flutter_test.dart';
import 'package:weather_app/modules/weather/data/models/weather_model.dart';

/// Shape of `GET /v1/weather/{city}` → `data` (see server store.WeatherResponse).
const _json = <String, dynamic>{
  'coord': {'lat': 51.5085, 'lon': -0.1257},
  'weather': [
    {'id': 800, 'main': 'Clear', 'description': 'clear sky', 'icon': '01d'},
  ],
  'main': {
    'temp': 12, // integers must be accepted for doubles
    'feels_like': 11.2,
    'temp_min': 10.0,
    'temp_max': 14.0,
    'pressure': 1012,
    'humidity': 70,
    'sea_level': 1012,
    'grnd_level': 1008,
  },
  'visibility': 10000,
  'wind': {'speed': 3, 'deg': 200, 'gust': 0},
  'clouds': {'all': 0},
  'dt': 1700000000,
  'sys': {'country': 'GB', 'sunrise': 1699990000, 'sunset': 1700020000},
  'timezone': 0,
  'name': 'London',
  'cod': 200,
};

void main() {
  test('parses a backend weather response', () {
    final weather = WeatherResponseModel.fromJson(_json);

    expect(weather.name, 'London');
    expect(weather.main.temp, 12.0);
    expect(weather.wind.speed, 3.0);
    expect(weather.weather.single.id, 800);
    expect(weather.sys.country, 'GB');
  });

  test('toJson does not throw and round-trips', () {
    final weather = WeatherResponseModel.fromJson(_json);

    final again = WeatherResponseModel.fromJson(weather.toJson());

    expect(again.name, weather.name);
    expect(again.sys.sunset, weather.sys.sunset);
  });
}
