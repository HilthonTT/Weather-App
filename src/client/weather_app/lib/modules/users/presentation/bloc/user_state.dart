part of 'user_bloc.dart';

@immutable
sealed class UserState {
  const UserState();
}

final class UserInitial extends UserState {
  const UserInitial();
}

final class UserFailure extends UserState {
  final String error;

  const UserFailure(this.error);
}

final class UserLoginSuccess extends UserState {
  const UserLoginSuccess();
}

final class UserLoading extends UserState {
  const UserLoading();
}
