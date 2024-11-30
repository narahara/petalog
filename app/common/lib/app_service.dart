import 'package:grpc/grpc.dart';

import 'config.dart';
import 'service/app/app_service.pbgrpc.dart';
import 'service/app/app_service_dto.pb.dart';

class AppService {
  AppServiceClient? client;

  Future<void> init() async {
    final config = Config();
    await config.setup();
    const ChannelOptions opt =
        ChannelOptions(credentials: ChannelCredentials.insecure());
    var channel = ClientChannel(config.host, port: config.port, options: opt);
    client = AppServiceClient(channel,
        options: CallOptions(timeout: const Duration(seconds: 5)));
  }

  Future<void> Ping() async {
    await checkInit();
    final req = PingRequest(
      message: 'Hello:${DateTime.now()}',
      name: "tom"
    );
    final res = await client!.ping(req);
    print(res);
  }

  Future<void> checkInit() async {
    if (client == null) {
      await init();
    }
    return;
  }
}
