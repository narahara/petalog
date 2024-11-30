import 'package:common/app_service.dart';
import 'package:flutter_test/flutter_test.dart';

void main() {
  test('Service test', () async {
    final service = AppService();
    await service.Ping();
  });
}