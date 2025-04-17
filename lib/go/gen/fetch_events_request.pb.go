// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: fetch_events_request.proto

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

type FetchEventsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	SessionId     *SessionId             `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	ClusterId     *ClusterId             `protobuf:"bytes,2,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	Namespace     string                 `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Kind          string                 `protobuf:"bytes,4,opt,name=kind,proto3" json:"kind,omitempty"`
	Name          string                 `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FetchEventsRequest) Reset() {
	*x = FetchEventsRequest{}
	mi := &file_fetch_events_request_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FetchEventsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchEventsRequest) ProtoMessage() {}

func (x *FetchEventsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fetch_events_request_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchEventsRequest.ProtoReflect.Descriptor instead.
func (*FetchEventsRequest) Descriptor() ([]byte, []int) {
	return file_fetch_events_request_proto_rawDescGZIP(), []int{0}
}

func (x *FetchEventsRequest) GetSessionId() *SessionId {
	if x != nil {
		return x.SessionId
	}
	return nil
}

func (x *FetchEventsRequest) GetClusterId() *ClusterId {
	if x != nil {
		return x.ClusterId
	}
	return nil
}

func (x *FetchEventsRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *FetchEventsRequest) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *FetchEventsRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_fetch_events_request_proto protoreflect.FileDescriptor

const file_fetch_events_request_proto_rawDesc = "" +
	"\n" +
	"\x1afetch_events_request.proto\x1a\x10session_id.proto\x1a\x10cluster_id.proto\"\xb0\x01\n" +
	"\x12FetchEventsRequest\x12)\n" +
	"\n" +
	"session_id\x18\x01 \x01(\v2\n" +
	".SessionIdR\tsessionId\x12)\n" +
	"\n" +
	"cluster_id\x18\x02 \x01(\v2\n" +
	".ClusterIdR\tclusterId\x12\x1c\n" +
	"\tnamespace\x18\x03 \x01(\tR\tnamespace\x12\x12\n" +
	"\x04kind\x18\x04 \x01(\tR\x04kind\x12\x12\n" +
	"\x04name\x18\x05 \x01(\tR\x04nameB\x84\x01B\x17FetchEventsRequestProtoP\x01Z9github.com/octopusdeploy/kubernetes-monitor-contracts/gen\xaa\x02+Octopus.Kubernetes.Monitor.MessageContractsb\x06proto3"

var (
	file_fetch_events_request_proto_rawDescOnce sync.Once
	file_fetch_events_request_proto_rawDescData []byte
)

func file_fetch_events_request_proto_rawDescGZIP() []byte {
	file_fetch_events_request_proto_rawDescOnce.Do(func() {
		file_fetch_events_request_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_fetch_events_request_proto_rawDesc), len(file_fetch_events_request_proto_rawDesc)))
	})
	return file_fetch_events_request_proto_rawDescData
}

var file_fetch_events_request_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_fetch_events_request_proto_goTypes = []any{
	(*FetchEventsRequest)(nil), // 0: FetchEventsRequest
	(*SessionId)(nil),          // 1: SessionId
	(*ClusterId)(nil),          // 2: ClusterId
}
var file_fetch_events_request_proto_depIdxs = []int32{
	1, // 0: FetchEventsRequest.session_id:type_name -> SessionId
	2, // 1: FetchEventsRequest.cluster_id:type_name -> ClusterId
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_fetch_events_request_proto_init() }
func file_fetch_events_request_proto_init() {
	if File_fetch_events_request_proto != nil {
		return
	}
	file_session_id_proto_init()
	file_cluster_id_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_fetch_events_request_proto_rawDesc), len(file_fetch_events_request_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_fetch_events_request_proto_goTypes,
		DependencyIndexes: file_fetch_events_request_proto_depIdxs,
		MessageInfos:      file_fetch_events_request_proto_msgTypes,
	}.Build()
	File_fetch_events_request_proto = out.File
	file_fetch_events_request_proto_goTypes = nil
	file_fetch_events_request_proto_depIdxs = nil
}
