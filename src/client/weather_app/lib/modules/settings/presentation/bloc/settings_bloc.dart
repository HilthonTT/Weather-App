import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:weather_app/common/usecase/usecase.dart';
import 'package:weather_app/modules/settings/domain/entities/settings.dart';
import 'package:weather_app/modules/settings/domain/usecases/get_settings.dart';
import 'package:weather_app/modules/settings/domain/usecases/update_settings.dart';

part 'settings_event.dart';
part 'settings_state.dart';

final class SettingsBloc extends Bloc<SettingsEvent, SettingsState> {
  final GetSettings _getSettings;
  final UpdateSettings _updateSettings;

  SettingsBloc({
    required UpdateSettings updatedSettings,
    required GetSettings getSettings,
  }) : _updateSettings = updatedSettings,
       _getSettings = getSettings,
       super(const SettingsInitial()) {
    on<SettingsEvent>((event, emit) {
      emit(const SettingsLoading());
    });

    on<UpdateSettingsEvent>(_onUpdateSettings);
    on<GetSettingsEvent>(_onGetSettings);
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

  Future<void> _onGetSettings(
    GetSettingsEvent event,
    Emitter<SettingsState> emit,
  ) async {
    final response = await _getSettings(const NoParams());

    response.fold(
      (failure) => emit(SettingsFailure(failure.message)),
      (settings) => emit(SettingsDisplaySuccess(settings)),
    );
  }
}
