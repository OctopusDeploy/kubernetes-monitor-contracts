// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: desired_resource_id.proto

package message_contracts

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

type DesiredResourceId struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Value         string                 `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DesiredResourceId) Reset() {
	*x = DesiredResourceId{}
	mi := &file_desired_resource_id_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DesiredResourceId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DesiredResourceId) ProtoMessage() {}

func (x *DesiredResourceId) ProtoReflect() protoreflect.Message {
	mi := &file_desired_resource_id_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DesiredResourceId.ProtoReflect.Descriptor instead.
func (*DesiredResourceId) Descriptor() ([]byte, []int) {
	return file_desired_resource_id_proto_rawDescGZIP(), []int{0}
}

func (x *DesiredResourceId) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_desired_resource_id_proto protoreflect.FileDescriptor

const file_desired_resource_id_proto_rawDesc = "" +
	"\n" +
	"\x19desired_resource_id.proto\")\n" +
	"\x11DesiredResourceId\x12\x14\n" +
	"\x05value\x18\x01 \x01(\tR\x05valueB\x98\x01B\x16DesiredResourceIdProtoP\x01ZNgithub.com/OctopusDeploy/kubernetes-monitor-contracts/go/pkg/message_contracts\xaa\x02+Octopus.Kubernetes.Monitor.MessageContractsb\x06proto3"

var (
	file_desired_resource_id_proto_rawDescOnce sync.Once
	file_desired_resource_id_proto_rawDescData []byte
)

func file_desired_resource_id_proto_rawDescGZIP() []byte {
	file_desired_resource_id_proto_rawDescOnce.Do(func() {
		file_desired_resource_id_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_desired_resource_id_proto_rawDesc), len(file_desired_resource_id_proto_rawDesc)))
	})
	return file_desired_resource_id_proto_rawDescData
}

var file_desired_resource_id_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_desired_resource_id_proto_goTypes = []any{
	(*DesiredResourceId)(nil), // 0: DesiredResourceId
}
var file_desired_resource_id_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_desired_resource_id_proto_init() }
func file_desired_resource_id_proto_init() {
	if File_desired_resource_id_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_desired_resource_id_proto_rawDesc), len(file_desired_resource_id_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_desired_resource_id_proto_goTypes,
		DependencyIndexes: file_desired_resource_id_proto_depIdxs,
		MessageInfos:      file_desired_resource_id_proto_msgTypes,
	}.Build()
	File_desired_resource_id_proto = out.File
	file_desired_resource_id_proto_goTypes = nil
	file_desired_resource_id_proto_depIdxs = nil
}
