//
//  Generated code. Do not modify.
//  source: service/app/app_service.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:async' as $async;
import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'package:protobuf/protobuf.dart' as $pb;

import 'app_service_dto.pb.dart' as $0;

export 'app_service.pb.dart';

@$pb.GrpcServiceName('petalog.service.app.AppService')
class AppServiceClient extends $grpc.Client {
  static final _$ping = $grpc.ClientMethod<$0.PingRequest, $0.PingResponse>(
      '/petalog.service.app.AppService/Ping',
      ($0.PingRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.PingResponse.fromBuffer(value));

  AppServiceClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options,
        interceptors: interceptors);

  $grpc.ResponseFuture<$0.PingResponse> ping($0.PingRequest request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$ping, request, options: options);
  }
}

@$pb.GrpcServiceName('petalog.service.app.AppService')
abstract class AppServiceBase extends $grpc.Service {
  $core.String get $name => 'petalog.service.app.AppService';

  AppServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.PingRequest, $0.PingResponse>(
        'Ping',
        ping_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.PingRequest.fromBuffer(value),
        ($0.PingResponse value) => value.writeToBuffer()));
  }

  $async.Future<$0.PingResponse> ping_Pre($grpc.ServiceCall call, $async.Future<$0.PingRequest> request) async {
    return ping(call, await request);
  }

  $async.Future<$0.PingResponse> ping($grpc.ServiceCall call, $0.PingRequest request);
}
