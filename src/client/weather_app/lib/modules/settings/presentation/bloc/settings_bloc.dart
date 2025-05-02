import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:weather_app/modules/settings/domain/entities/settings.dart';
import 'package:weather_app/modules/settings/domain/usecases/update_settings.dart';

part 'settings_event.dart';
part 'settings_state.dart';

final class SettingsBloc extends Bloc<SettingsEvent, SettingsState> {
  final UpdateSettings _updateSettings;

  SettingsBloc({required UpdateSettings updatedSettings})
    : _updateSettings = updatedSettings,
      super(const SettingsInitial()) {
    on<SettingsEvent>((event, emit) {
      emit(const SettingsLoading());
    });

    on<UpdateSettingsEvent>(_onUpdateSettings);
  }

  Future<void> _onUpdateSettings(
    UpdateSettingsEvent event,
    Emitter<SettingsState> emit,
  ) async {
    final response = await _updateSettings(
      UpdateSettingsParams(settings: event.settings),
    );

    response.fold(
      (failure) => emit(SettingsFailure(failure.message)),
      (blog) => emit(const SettingsUpdateSuccess()),
    );
  }
}
