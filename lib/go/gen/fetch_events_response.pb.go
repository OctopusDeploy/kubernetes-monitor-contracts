// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: fetch_events_response.proto

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

type FetchEventsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	SessionId     *SessionId             `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	Events        []*Event               `protobuf:"bytes,2,rep,name=events,proto3" json:"events,omitempty"`
	Error         *Error                 `protobuf:"bytes,3,opt,name=error,proto3,oneof" json:"error,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FetchEventsResponse) Reset() {
	*x = FetchEventsResponse{}
	mi := &file_fetch_events_response_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FetchEventsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchEventsResponse) ProtoMessage() {}

func (x *FetchEventsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_fetch_events_response_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchEventsResponse.ProtoReflect.Descriptor instead.
func (*FetchEventsResponse) Descriptor() ([]byte, []int) {
	return file_fetch_events_response_proto_rawDescGZIP(), []int{0}
}

func (x *FetchEventsResponse) GetSessionId() *SessionId {
	if x != nil {
		return x.SessionId
	}
	return nil
}

func (x *FetchEventsResponse) GetEvents() []*Event {
	if x != nil {
		return x.Events
	}
	return nil
}

func (x *FetchEventsResponse) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_fetch_events_response_proto protoreflect.FileDescriptor

const file_fetch_events_response_proto_rawDesc = "" +
	"\n" +
	"\x1bfetch_events_response.proto\x1a\verror.proto\x1a\vevent.proto\x1a\x10session_id.proto\"\x8d\x01\n" +
	"\x13FetchEventsResponse\x12)\n" +
	"\n" +
	"session_id\x18\x01 \x01(\v2\n" +
	".SessionIdR\tsessionId\x12\x1e\n" +
	"\x06events\x18\x02 \x03(\v2\x06.EventR\x06events\x12!\n" +
	"\x05error\x18\x03 \x01(\v2\x06.ErrorH\x00R\x05error\x88\x01\x01B\b\n" +
	"\x06_errorB\x85\x01B\x18FetchEventsResponseProtoP\x01Z9github.com/octopusdeploy/kubernetes-monitor-contracts/gen\xaa\x02+Octopus.Kubernetes.Monitor.MessageContractsb\x06proto3"

var (
	file_fetch_events_response_proto_rawDescOnce sync.Once
	file_fetch_events_response_proto_rawDescData []byte
)

func file_fetch_events_response_proto_rawDescGZIP() []byte {
	file_fetch_events_response_proto_rawDescOnce.Do(func() {
		file_fetch_events_response_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_fetch_events_response_proto_rawDesc), len(file_fetch_events_response_proto_rawDesc)))
	})
	return file_fetch_events_response_proto_rawDescData
}

var file_fetch_events_response_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_fetch_events_response_proto_goTypes = []any{
	(*FetchEventsResponse)(nil), // 0: FetchEventsResponse
	(*SessionId)(nil),           // 1: SessionId
	(*Event)(nil),               // 2: Event
	(*Error)(nil),               // 3: Error
}
var file_fetch_events_response_proto_depIdxs = []int32{
	1, // 0: FetchEventsResponse.session_id:type_name -> SessionId
	2, // 1: FetchEventsResponse.events:type_name -> Event
	3, // 2: FetchEventsResponse.error:type_name -> Error
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_fetch_events_response_proto_init() }
func file_fetch_events_response_proto_init() {
	if File_fetch_events_response_proto != nil {
		return
	}
	file_error_proto_init()
	file_event_proto_init()
	file_session_id_proto_init()
	file_fetch_events_response_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_fetch_events_response_proto_rawDesc), len(file_fetch_events_response_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_fetch_events_response_proto_goTypes,
		DependencyIndexes: file_fetch_events_response_proto_depIdxs,
		MessageInfos:      file_fetch_events_response_proto_msgTypes,
	}.Build()
	File_fetch_events_response_proto = out.File
	file_fetch_events_response_proto_goTypes = nil
	file_fetch_events_response_proto_depIdxs = nil
}
