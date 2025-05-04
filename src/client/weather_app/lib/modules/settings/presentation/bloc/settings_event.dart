part of 'settings_bloc.dart';

@immutable
sealed class SettingsEvent {
  const SettingsEvent();
}

@immutable
final class UpdateSettingsEvent extends SettingsEvent {
  final Settings settings;

  const UpdateSettingsEvent(this.settings);
}

@immutable
final class GetSettingsEvent extends SettingsEvent {
  const GetSettingsEvent();
}
