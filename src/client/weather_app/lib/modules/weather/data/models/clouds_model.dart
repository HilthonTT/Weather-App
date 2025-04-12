import 'package:weather_app/modules/weather/domain/entities/clouds.dart';

final class CloudsModel extends Clouds {
  const CloudsModel({required super.all});

  factory CloudsModel.fromJson(Map<String, dynamic> json) {
    return CloudsModel(all: (json['all'] as num).toInt());
  }

  Map<String, dynamic> toJson() {
    return {'all': all};
  }
}
