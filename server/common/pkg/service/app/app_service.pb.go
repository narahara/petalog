// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: service/app/app_service.proto

package app

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_service_app_app_service_proto protoreflect.FileDescriptor

var file_service_app_app_service_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x61, 0x70,
	0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x13, 0x70, 0x65, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x61, 0x70, 0x70, 0x1a, 0x21, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x70,
	0x70, 0x2f, 0x61, 0x70, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x64, 0x74,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x5b, 0x0a, 0x0a, 0x41, 0x70, 0x70, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x20, 0x2e,
	0x70, 0x65, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x61, 0x70, 0x70, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x21, 0x2e, 0x70, 0x65, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x61, 0x70, 0x70, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_service_app_app_service_proto_goTypes = []any{
	(*PingRequest)(nil),  // 0: petalog.service.app.PingRequest
	(*PingResponse)(nil), // 1: petalog.service.app.PingResponse
}
var file_service_app_app_service_proto_depIdxs = []int32{
	0, // 0: petalog.service.app.AppService.Ping:input_type -> petalog.service.app.PingRequest
	1, // 1: petalog.service.app.AppService.Ping:output_type -> petalog.service.app.PingResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_service_app_app_service_proto_init() }
func file_service_app_app_service_proto_init() {
	if File_service_app_app_service_proto != nil {
		return
	}
	file_service_app_app_service_dto_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_app_app_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_app_app_service_proto_goTypes,
		DependencyIndexes: file_service_app_app_service_proto_depIdxs,
	}.Build()
	File_service_app_app_service_proto = out.File
	file_service_app_app_service_proto_rawDesc = nil
	file_service_app_app_service_proto_goTypes = nil
	file_service_app_app_service_proto_depIdxs = nil
}
