import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:weather_app/common/usecase/usecase.dart';
import 'package:weather_app/modules/settings/domain/entities/settings.dart';
import 'package:weather_app/modules/settings/domain/usecases/get_settings.dart';
import 'package:weather_app/modules/settings/domain/usecases/update_settings.dart';

part 'settings_event.dart';
part 'settings_state.dart';

/// Single source of truth for the user's display settings.
///
/// Every page listens for [SettingsDisplaySuccess] to format values. After a
/// successful update the bloc emits [SettingsUpdateSuccess] (for feedback)
/// followed by [SettingsDisplaySuccess] with the new settings, so all pages
/// refresh without re-fetching.
final class SettingsBloc extends Bloc<SettingsEvent, SettingsState> {
  final GetSettings _getSettings;
  final UpdateSettings _updateSettings;

  SettingsBloc({
    required UpdateSettings updateSettings,
    required GetSettings getSettings,
  }) : _updateSettings = updateSettings,
       _getSettings = getSettings,
       super(const SettingsInitial()) {
    on<UpdateSettingsEvent>(_onUpdateSettings);
    on<GetSettingsEvent>(_onGetSettings);
  }

  /// The most recently loaded settings, or `null` before the first load.
  Settings? get currentSettings => _currentSettings;
  Settings? _currentSettings;

  Future<void> _onUpdateSettings(
    UpdateSettingsEvent event,
    Emitter<SettingsState> emit,
  ) async {
    emit(const SettingsLoading());

    final response = await _updateSettings(
      UpdateSettingsParams(settings: event.settings),
    );

    response.fold((failure) => emit(SettingsFailure(failure.message)), (
      settings,
    ) {
      _currentSettings = settings;
      emit(const SettingsUpdateSuccess());
      emit(SettingsDisplaySuccess(settings));
    });
  }

  Future<void> _onGetSettings(
    GetSettingsEvent event,
    Emitter<SettingsState> emit,
  ) async {
    emit(const SettingsLoading());

    final response = await _getSettings(const NoParams());

    response.fold((failure) => emit(SettingsFailure(failure.message)), (
      settings,
    ) {
      _currentSettings = settings;
      emit(SettingsDisplaySuccess(settings));
    });
  }
}
