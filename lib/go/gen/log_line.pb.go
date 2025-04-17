// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: log_line.proto

package gen

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type LogLine struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Timestamp     *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LogLine) Reset() {
	*x = LogLine{}
	mi := &file_log_line_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogLine) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogLine) ProtoMessage() {}

func (x *LogLine) ProtoReflect() protoreflect.Message {
	mi := &file_log_line_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogLine.ProtoReflect.Descriptor instead.
func (*LogLine) Descriptor() ([]byte, []int) {
	return file_log_line_proto_rawDescGZIP(), []int{0}
}

func (x *LogLine) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *LogLine) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

var File_log_line_proto protoreflect.FileDescriptor

const file_log_line_proto_rawDesc = "" +
	"\n" +
	"\x0elog_line.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"]\n" +
	"\aLogLine\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage\x128\n" +
	"\ttimestamp\x18\x02 \x01(\v2\x1a.google.protobuf.TimestampR\ttimestampB\x80\x01B\fLogLineProtoP\x01Z@github.com/octopusdeploy/kubernetes-monitor-contracts/lib/go/gen\xaa\x02+Octopus.Kubernetes.Monitor.MessageContractsb\x06proto3"

var (
	file_log_line_proto_rawDescOnce sync.Once
	file_log_line_proto_rawDescData []byte
)

func file_log_line_proto_rawDescGZIP() []byte {
	file_log_line_proto_rawDescOnce.Do(func() {
		file_log_line_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_log_line_proto_rawDesc), len(file_log_line_proto_rawDesc)))
	})
	return file_log_line_proto_rawDescData
}

var file_log_line_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_log_line_proto_goTypes = []any{
	(*LogLine)(nil),               // 0: LogLine
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_log_line_proto_depIdxs = []int32{
	1, // 0: LogLine.timestamp:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_log_line_proto_init() }
func file_log_line_proto_init() {
	if File_log_line_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_log_line_proto_rawDesc), len(file_log_line_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_log_line_proto_goTypes,
		DependencyIndexes: file_log_line_proto_depIdxs,
		MessageInfos:      file_log_line_proto_msgTypes,
	}.Build()
	File_log_line_proto = out.File
	file_log_line_proto_goTypes = nil
	file_log_line_proto_depIdxs = nil
}
