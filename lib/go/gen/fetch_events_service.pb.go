// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: fetch_events_service.proto

package gen

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_fetch_events_service_proto protoreflect.FileDescriptor

const file_fetch_events_service_proto_rawDesc = "" +
	"\n" +
	"\x1afetch_events_service.proto\x1a\x1afetch_events_request.proto\x1a\x1bfetch_events_response.proto2R\n" +
	"\x12FetchEventsService\x12<\n" +
	"\vFetchEvents\x12\x14.FetchEventsResponse\x1a\x13.FetchEventsRequest(\x010\x01B\x84\x01B\x17FetchEventsServiceProtoP\x01Z9github.com/octopusdeploy/kubernetes-monitor-contracts/gen\xaa\x02+Octopus.Kubernetes.Monitor.MessageContractsb\x06proto3"

var file_fetch_events_service_proto_goTypes = []any{
	(*FetchEventsResponse)(nil), // 0: FetchEventsResponse
	(*FetchEventsRequest)(nil),  // 1: FetchEventsRequest
}
var file_fetch_events_service_proto_depIdxs = []int32{
	0, // 0: FetchEventsService.FetchEvents:input_type -> FetchEventsResponse
	1, // 1: FetchEventsService.FetchEvents:output_type -> FetchEventsRequest
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_fetch_events_service_proto_init() }
func file_fetch_events_service_proto_init() {
	if File_fetch_events_service_proto != nil {
		return
	}
	file_fetch_events_request_proto_init()
	file_fetch_events_response_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_fetch_events_service_proto_rawDesc), len(file_fetch_events_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_fetch_events_service_proto_goTypes,
		DependencyIndexes: file_fetch_events_service_proto_depIdxs,
	}.Build()
	File_fetch_events_service_proto = out.File
	file_fetch_events_service_proto_goTypes = nil
	file_fetch_events_service_proto_depIdxs = nil
}
