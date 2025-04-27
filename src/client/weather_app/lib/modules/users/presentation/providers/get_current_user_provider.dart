import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:weather_app/common/usecase/usecase.dart';
import 'package:weather_app/dependency_injection.dart';
import 'package:weather_app/modules/users/domain/entities/user.dart';
import 'package:weather_app/modules/users/domain/usecases/get_current_user.dart';

final getCurrentUserProvider = FutureProvider.autoDispose<User?>((ref) async {
  final getCurrentUser = serviceLocator.get<GetCurrentUser>();
  final response = await getCurrentUser(const NoParams());

  return response.fold(
    (failure) => throw Exception(failure.message),
    (user) => user,
  );
});
