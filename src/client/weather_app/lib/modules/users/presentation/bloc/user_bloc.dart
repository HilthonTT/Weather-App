import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:weather_app/modules/users/domain/usecases/get_current_user.dart';
import 'package:weather_app/modules/users/domain/usecases/login_user.dart';
import 'package:weather_app/modules/users/domain/usecases/register_user.dart';

part 'user_event.dart';
part 'user_state.dart';

final class UserBloc extends Bloc<UserEvent, UserState> {
  final LoginUser _loginUser;
  final RegisterUser _registerUser;
  final GetCurrentUser _getCurrentUser;

  UserBloc({
    required LoginUser loginUser,
    required RegisterUser registerUser,
    required GetCurrentUser getCurrentUser,
  }) : _loginUser = loginUser,
       _registerUser = registerUser,
       _getCurrentUser = getCurrentUser,
       super(const UserInitial()) {
    on<UserEvent>((event, emit) {
      emit(const UserLoading());
    });

    on<UserLoginEvent>(_onUserLogin);
  }

  Future<void> _onUserLogin(
    UserLoginEvent event,
    Emitter<UserState> emit,
  ) async {
    final params = LoginUserParams(
      email: event.email,
      password: event.password,
    );

    final response = await _loginUser(params);

    response.fold((failure) => emit(UserFailure(failure.message)), (jwtToken) {
      debugPrint(jwtToken);
      emit(const UserLoginSuccess());
    });
  }
}
