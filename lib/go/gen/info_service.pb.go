// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: info_service.proto

package gen

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetInfoRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetInfoRequest) Reset() {
	*x = GetInfoRequest{}
	mi := &file_info_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInfoRequest) ProtoMessage() {}

func (x *GetInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_info_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInfoRequest.ProtoReflect.Descriptor instead.
func (*GetInfoRequest) Descriptor() ([]byte, []int) {
	return file_info_service_proto_rawDescGZIP(), []int{0}
}

type GetInfoResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Headers       []*HeaderEntry         `protobuf:"bytes,1,rep,name=headers,proto3" json:"headers,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetInfoResponse) Reset() {
	*x = GetInfoResponse{}
	mi := &file_info_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInfoResponse) ProtoMessage() {}

func (x *GetInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_info_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInfoResponse.ProtoReflect.Descriptor instead.
func (*GetInfoResponse) Descriptor() ([]byte, []int) {
	return file_info_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetInfoResponse) GetHeaders() []*HeaderEntry {
	if x != nil {
		return x.Headers
	}
	return nil
}

type HeaderEntry struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           string                 `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
	Value         string                 `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HeaderEntry) Reset() {
	*x = HeaderEntry{}
	mi := &file_info_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HeaderEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeaderEntry) ProtoMessage() {}

func (x *HeaderEntry) ProtoReflect() protoreflect.Message {
	mi := &file_info_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeaderEntry.ProtoReflect.Descriptor instead.
func (*HeaderEntry) Descriptor() ([]byte, []int) {
	return file_info_service_proto_rawDescGZIP(), []int{2}
}

func (x *HeaderEntry) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *HeaderEntry) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_info_service_proto protoreflect.FileDescriptor

const file_info_service_proto_rawDesc = "" +
	"\n" +
	"\x12info_service.proto\"\x10\n" +
	"\x0eGetInfoRequest\"9\n" +
	"\x0fGetInfoResponse\x12&\n" +
	"\aheaders\x18\x01 \x03(\v2\f.HeaderEntryR\aheaders\"5\n" +
	"\vHeaderEntry\x12\x10\n" +
	"\x03Key\x18\x01 \x01(\tR\x03Key\x12\x14\n" +
	"\x05Value\x18\x02 \x01(\tR\x05Value2;\n" +
	"\vInfoService\x12,\n" +
	"\aGetInfo\x12\x0f.GetInfoRequest\x1a\x10.GetInfoResponseB\x84\x01B\x10InfoServiceProtoP\x01Z@github.com/octopusdeploy/kubernetes-monitor-contracts/lib/go/gen\xaa\x02+Octopus.Kubernetes.Monitor.MessageContractsb\x06proto3"

var (
	file_info_service_proto_rawDescOnce sync.Once
	file_info_service_proto_rawDescData []byte
)

func file_info_service_proto_rawDescGZIP() []byte {
	file_info_service_proto_rawDescOnce.Do(func() {
		file_info_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_info_service_proto_rawDesc), len(file_info_service_proto_rawDesc)))
	})
	return file_info_service_proto_rawDescData
}

var file_info_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_info_service_proto_goTypes = []any{
	(*GetInfoRequest)(nil),  // 0: GetInfoRequest
	(*GetInfoResponse)(nil), // 1: GetInfoResponse
	(*HeaderEntry)(nil),     // 2: HeaderEntry
}
var file_info_service_proto_depIdxs = []int32{
	2, // 0: GetInfoResponse.headers:type_name -> HeaderEntry
	0, // 1: InfoService.GetInfo:input_type -> GetInfoRequest
	1, // 2: InfoService.GetInfo:output_type -> GetInfoResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_info_service_proto_init() }
func file_info_service_proto_init() {
	if File_info_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_info_service_proto_rawDesc), len(file_info_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_info_service_proto_goTypes,
		DependencyIndexes: file_info_service_proto_depIdxs,
		MessageInfos:      file_info_service_proto_msgTypes,
	}.Build()
	File_info_service_proto = out.File
	file_info_service_proto_goTypes = nil
	file_info_service_proto_depIdxs = nil
}
