// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: resource_details.proto

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

type DesiredResourceDetails struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	Name             string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	AssumedNamespace string                 `protobuf:"bytes,2,opt,name=assumed_namespace,json=assumedNamespace,proto3" json:"assumed_namespace,omitempty"`
	GroupVersionKind *GroupVersionKind      `protobuf:"bytes,3,opt,name=group_version_kind,json=groupVersionKind,proto3" json:"group_version_kind,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *DesiredResourceDetails) Reset() {
	*x = DesiredResourceDetails{}
	mi := &file_resource_details_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DesiredResourceDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DesiredResourceDetails) ProtoMessage() {}

func (x *DesiredResourceDetails) ProtoReflect() protoreflect.Message {
	mi := &file_resource_details_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DesiredResourceDetails.ProtoReflect.Descriptor instead.
func (*DesiredResourceDetails) Descriptor() ([]byte, []int) {
	return file_resource_details_proto_rawDescGZIP(), []int{0}
}

func (x *DesiredResourceDetails) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DesiredResourceDetails) GetAssumedNamespace() string {
	if x != nil {
		return x.AssumedNamespace
	}
	return ""
}

func (x *DesiredResourceDetails) GetGroupVersionKind() *GroupVersionKind {
	if x != nil {
		return x.GroupVersionKind
	}
	return nil
}

type ResourceDetails struct {
	state               protoimpl.MessageState `protogen:"open.v1"`
	Name                string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	KubernetesNamespace string                 `protobuf:"bytes,2,opt,name=kubernetes_namespace,json=kubernetesNamespace,proto3" json:"kubernetes_namespace,omitempty"`
	GroupVersionKind    *GroupVersionKind      `protobuf:"bytes,3,opt,name=group_version_kind,json=groupVersionKind,proto3" json:"group_version_kind,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *ResourceDetails) Reset() {
	*x = ResourceDetails{}
	mi := &file_resource_details_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResourceDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceDetails) ProtoMessage() {}

func (x *ResourceDetails) ProtoReflect() protoreflect.Message {
	mi := &file_resource_details_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceDetails.ProtoReflect.Descriptor instead.
func (*ResourceDetails) Descriptor() ([]byte, []int) {
	return file_resource_details_proto_rawDescGZIP(), []int{1}
}

func (x *ResourceDetails) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ResourceDetails) GetKubernetesNamespace() string {
	if x != nil {
		return x.KubernetesNamespace
	}
	return ""
}

func (x *ResourceDetails) GetGroupVersionKind() *GroupVersionKind {
	if x != nil {
		return x.GroupVersionKind
	}
	return nil
}

type GroupVersionKind struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Group         string                 `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
	Version       string                 `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Kind          string                 `protobuf:"bytes,3,opt,name=kind,proto3" json:"kind,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GroupVersionKind) Reset() {
	*x = GroupVersionKind{}
	mi := &file_resource_details_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GroupVersionKind) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupVersionKind) ProtoMessage() {}

func (x *GroupVersionKind) ProtoReflect() protoreflect.Message {
	mi := &file_resource_details_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupVersionKind.ProtoReflect.Descriptor instead.
func (*GroupVersionKind) Descriptor() ([]byte, []int) {
	return file_resource_details_proto_rawDescGZIP(), []int{2}
}

func (x *GroupVersionKind) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *GroupVersionKind) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *GroupVersionKind) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

var File_resource_details_proto protoreflect.FileDescriptor

const file_resource_details_proto_rawDesc = "" +
	"\n" +
	"\x16resource_details.proto\"\x9a\x01\n" +
	"\x16DesiredResourceDetails\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12+\n" +
	"\x11assumed_namespace\x18\x02 \x01(\tR\x10assumedNamespace\x12?\n" +
	"\x12group_version_kind\x18\x03 \x01(\v2\x11.GroupVersionKindR\x10groupVersionKind\"\x99\x01\n" +
	"\x0fResourceDetails\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x121\n" +
	"\x14kubernetes_namespace\x18\x02 \x01(\tR\x13kubernetesNamespace\x12?\n" +
	"\x12group_version_kind\x18\x03 \x01(\v2\x11.GroupVersionKindR\x10groupVersionKind\"V\n" +
	"\x10GroupVersionKind\x12\x14\n" +
	"\x05group\x18\x01 \x01(\tR\x05group\x12\x18\n" +
	"\aversion\x18\x02 \x01(\tR\aversion\x12\x12\n" +
	"\x04kind\x18\x03 \x01(\tR\x04kindB\x96\x01B\x14ResourceDetailsProtoP\x01ZNgithub.com/OctopusDeploy/kubernetes-monitor-contracts/go/pkg/message_contracts\xaa\x02+Octopus.Kubernetes.Monitor.MessageContractsb\x06proto3"

var (
	file_resource_details_proto_rawDescOnce sync.Once
	file_resource_details_proto_rawDescData []byte
)

func file_resource_details_proto_rawDescGZIP() []byte {
	file_resource_details_proto_rawDescOnce.Do(func() {
		file_resource_details_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_resource_details_proto_rawDesc), len(file_resource_details_proto_rawDesc)))
	})
	return file_resource_details_proto_rawDescData
}

var file_resource_details_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_resource_details_proto_goTypes = []any{
	(*DesiredResourceDetails)(nil), // 0: DesiredResourceDetails
	(*ResourceDetails)(nil),        // 1: ResourceDetails
	(*GroupVersionKind)(nil),       // 2: GroupVersionKind
}
var file_resource_details_proto_depIdxs = []int32{
	2, // 0: DesiredResourceDetails.group_version_kind:type_name -> GroupVersionKind
	2, // 1: ResourceDetails.group_version_kind:type_name -> GroupVersionKind
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_resource_details_proto_init() }
func file_resource_details_proto_init() {
	if File_resource_details_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_resource_details_proto_rawDesc), len(file_resource_details_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_resource_details_proto_goTypes,
		DependencyIndexes: file_resource_details_proto_depIdxs,
		MessageInfos:      file_resource_details_proto_msgTypes,
	}.Build()
	File_resource_details_proto = out.File
	file_resource_details_proto_goTypes = nil
	file_resource_details_proto_depIdxs = nil
}
