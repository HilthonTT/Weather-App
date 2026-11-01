import 'package:flutter_test/flutter_test.dart';
import 'package:fpdart/fpdart.dart';
import 'package:weather_app/common/errors/failure.dart';
import 'package:weather_app/modules/settings/domain/entities/settings.dart';
import 'package:weather_app/modules/settings/domain/repositories/settings_repository.dart';
import 'package:weather_app/modules/settings/domain/usecases/get_settings.dart';
import 'package:weather_app/modules/settings/domain/usecases/update_settings.dart';
import 'package:weather_app/modules/settings/presentation/bloc/settings_bloc.dart';

/// In-memory repository used instead of Hive.
final class _FakeSettingsRepository implements SettingsRepository {
  Settings stored = const Settings(
    id: 1,
    userId: 1,
    tempFormat: TempFormat.celsius,
    timeFormat: TimeFormat.twentyFourHour,
    speedFormat: SpeedFormat.kmph,
  );

  @override
  Either<Failure, Settings> getSettings() => right(stored);

  @override
  Either<Failure, Settings> updateSettings({required Settings settings}) {
    stored = settings;
    return right(settings);
  }
}

void main() {
  late _FakeSettingsRepository repository;
  late SettingsBloc bloc;

  setUp(() {
    repository = _FakeSettingsRepository();
    bloc = SettingsBloc(
      getSettings: GetSettings(settingsRepository: repository),
      updateSettings: UpdateSettings(settingsRepository: repository),
    );
  });

  tearDown(() => bloc.close());

  test('loads settings and exposes them as currentSettings', () async {
    bloc.add(const GetSettingsEvent());

    await expectLater(
      bloc.stream,
      emitsInOrder([isA<SettingsLoading>(), isA<SettingsDisplaySuccess>()]),
    );
    expect(bloc.currentSettings?.tempFormat, TempFormat.celsius);
  });

  test('update emits success then the new settings', () async {
    bloc.add(const GetSettingsEvent());
    await bloc.stream.firstWhere((s) => s is SettingsDisplaySuccess);

    final updated = bloc.currentSettings!.copyWith(
      tempFormat: TempFormat.fahrenheit,
    );
    bloc.add(UpdateSettingsEvent(updated));

    await expectLater(
      bloc.stream,
      emitsInOrder([
        isA<SettingsLoading>(),
        isA<SettingsUpdateSuccess>(),
        isA<SettingsDisplaySuccess>().having(
          (s) => s.settings.tempFormat,
          'tempFormat',
          TempFormat.fahrenheit,
        ),
      ]),
    );
  });

  test('successive updates to different fields are all kept', () async {
    bloc.add(const GetSettingsEvent());
    await bloc.stream.firstWhere((s) => s is SettingsDisplaySuccess);

    // Mirrors SettingsPage: each change is applied to bloc.currentSettings.
    bloc.add(
      UpdateSettingsEvent(
        bloc.currentSettings!.copyWith(timeFormat: TimeFormat.twelveHour),
      ),
    );
    await bloc.stream.firstWhere((s) => s is SettingsDisplaySuccess);

    bloc.add(
      UpdateSettingsEvent(
        bloc.currentSettings!.copyWith(speedFormat: SpeedFormat.mph),
      ),
    );
    await bloc.stream.firstWhere((s) => s is SettingsDisplaySuccess);

    expect(repository.stored.timeFormat, TimeFormat.twelveHour);
    expect(repository.stored.speedFormat, SpeedFormat.mph);
  });
}
