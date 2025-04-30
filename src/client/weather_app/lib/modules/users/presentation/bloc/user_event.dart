part of 'user_bloc.dart';

@immutable
sealed class UserEvent {
  const UserEvent();
}

final class UserLoginEvent extends UserEvent {
  final String email;
  final String password;

  const UserLoginEvent({required this.email, required this.password});
}
