# Weather App (Flutter)

Mobile client for the Weather API. State management uses
[Riverpod](https://riverpod.dev/) for data fetching and
[flutter_bloc](https://bloclibrary.dev/) for user settings; dependencies are
wired with [get_it](https://pub.dev/packages/get_it) and settings/location are
persisted with [Hive](https://pub.dev/packages/hive).

## Running

```bash
flutter pub get
flutter run --dart-define=API_URL=http://<backend-host>:8080
```

| `API_URL` | When to use |
| --------- | ----------- |
| `http://10.0.2.2:8080` (default) | Android emulator → backend on your machine |
| `http://localhost:8080` | iOS simulator / desktop |
| `http://<LAN-IP>:8080` | Physical device on the same network |

The backend must be running—see [`src/server/README.md`](../../server/README.md).
Location permission is required for the home and forecast screens.

## Project structure

Each feature follows clean architecture: **data** (API/Hive access and JSON
models) → **domain** (entities, repository interfaces, use cases) →
**presentation** (pages, widgets, providers/blocs).

```
lib/
  main.dart                    # bootstraps DI, loads settings, runs WeatherApp
  app/
    app.dart                   # MaterialApp
    home_layout.dart           # bottom navigation + IndexedStack of pages
    current_page_index_provider.dart
  common/
    constants/                 # API URL, colors, text styles
    di/service_locator.dart    # get_it registrations for every module
    errors/                    # Failure / exceptions
    extensions/                # unit conversions, date formatting
    network/                   # internet connectivity checker
    services/geolocator.dart   # device location with permission handling
    usecase/                   # UseCase base interface
    utils/ widgets/            # shared helpers and widgets
  modules/
    settings/                  # unit preferences (stored locally in Hive)
      data/ domain/ presentation/
        presentation/utils/format.dart   # temperature / wind formatting
    weather/                   # current weather, forecasts, city search
      data/ domain/ presentation/
test/                          # unit tests mirroring lib/modules
```

### Data flow

```
Page/Widget ──watch──▶ FutureProvider ──▶ UseCase ──▶ Repository ──▶ RemoteDatasource ──▶ Go API
     │                                                    │
     └──── BlocConsumer<SettingsBloc> ◀── Hive (settings, cached location)
```

- The API always returns **metric** values; widgets format them through
  `formatTemperature` / `formatWindSpeed` with the current `Settings`.
- `SettingsBloc` is loaded once at startup. After an update it emits
  `SettingsUpdateSuccess` followed by `SettingsDisplaySuccess`, so every page
  re-renders with the new units.
- The device location is cached in Hive for 24 hours.

## Tests

```bash
flutter analyze
flutter test
```
