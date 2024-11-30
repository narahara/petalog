import 'package:flutter_dotenv/flutter_dotenv.dart';

class Config {
  late String host;
  late int port;
  static final Config _config = Config._internal();
  factory Config() {
    return _config;
  }
  Config._internal();

  Future<void> setup() async {
    await dotenv.load(fileName: ".env");
    host = dotenv.env['APP_SERVER_HOST'] ?? '';
    port = int.parse(dotenv.env['APP_SERVER_PORT'] ?? '0');
  }
}
