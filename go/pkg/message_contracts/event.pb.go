// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: event.proto

package message_contracts

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

type Event struct {
	state               protoimpl.MessageState `protogen:"open.v1"`
	FirstObservedTime   *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=first_observed_time,json=firstObservedTime,proto3" json:"first_observed_time,omitempty"`
	LastObservedTime    *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=last_observed_time,json=lastObservedTime,proto3" json:"last_observed_time,omitempty"`
	Count               int32                  `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	Action              string                 `protobuf:"bytes,4,opt,name=action,proto3" json:"action,omitempty"`
	Reason              string                 `protobuf:"bytes,5,opt,name=reason,proto3" json:"reason,omitempty"`
	Note                string                 `protobuf:"bytes,6,opt,name=note,proto3" json:"note,omitempty"`
	ReportingController string                 `protobuf:"bytes,7,opt,name=reporting_controller,json=reportingController,proto3" json:"reporting_controller,omitempty"`
	ReportingInstance   string                 `protobuf:"bytes,8,opt,name=reporting_instance,json=reportingInstance,proto3" json:"reporting_instance,omitempty"`
	Type                string                 `protobuf:"bytes,9,opt,name=type,proto3" json:"type,omitempty"`
	Manifest            *YamlManifest          `protobuf:"bytes,10,opt,name=manifest,proto3" json:"manifest,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *Event) Reset() {
	*x = Event{}
	mi := &file_event_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_event_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_event_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetFirstObservedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.FirstObservedTime
	}
	return nil
}

func (x *Event) GetLastObservedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LastObservedTime
	}
	return nil
}

func (x *Event) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *Event) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *Event) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *Event) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

func (x *Event) GetReportingController() string {
	if x != nil {
		return x.ReportingController
	}
	return ""
}

func (x *Event) GetReportingInstance() string {
	if x != nil {
		return x.ReportingInstance
	}
	return ""
}

func (x *Event) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Event) GetManifest() *YamlManifest {
	if x != nil {
		return x.Manifest
	}
	return nil
}

var File_event_proto protoreflect.FileDescriptor

const file_event_proto_rawDesc = "" +
	"\n" +
	"\vevent.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x13yaml_manifest.proto\"\x98\x03\n" +
	"\x05Event\x12J\n" +
	"\x13first_observed_time\x18\x01 \x01(\v2\x1a.google.protobuf.TimestampR\x11firstObservedTime\x12H\n" +
	"\x12last_observed_time\x18\x02 \x01(\v2\x1a.google.protobuf.TimestampR\x10lastObservedTime\x12\x14\n" +
	"\x05count\x18\x03 \x01(\x05R\x05count\x12\x16\n" +
	"\x06action\x18\x04 \x01(\tR\x06action\x12\x16\n" +
	"\x06reason\x18\x05 \x01(\tR\x06reason\x12\x12\n" +
	"\x04note\x18\x06 \x01(\tR\x04note\x121\n" +
	"\x14reporting_controller\x18\a \x01(\tR\x13reportingController\x12-\n" +
	"\x12reporting_instance\x18\b \x01(\tR\x11reportingInstance\x12\x12\n" +
	"\x04type\x18\t \x01(\tR\x04type\x12)\n" +
	"\bmanifest\x18\n" +
	" \x01(\v2\r.YamlManifestR\bmanifestB\x8c\x01B\n" +
	"EventProtoP\x01ZNgithub.com/OctopusDeploy/kubernetes-monitor-contracts/go/pkg/message_contracts\xaa\x02+Octopus.Kubernetes.Monitor.MessageContractsb\x06proto3"

var (
	file_event_proto_rawDescOnce sync.Once
	file_event_proto_rawDescData []byte
)

func file_event_proto_rawDescGZIP() []byte {
	file_event_proto_rawDescOnce.Do(func() {
		file_event_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_event_proto_rawDesc), len(file_event_proto_rawDesc)))
	})
	return file_event_proto_rawDescData
}

var file_event_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_event_proto_goTypes = []any{
	(*Event)(nil),                 // 0: Event
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
	(*YamlManifest)(nil),          // 2: YamlManifest
}
var file_event_proto_depIdxs = []int32{
	1, // 0: Event.first_observed_time:type_name -> google.protobuf.Timestamp
	1, // 1: Event.last_observed_time:type_name -> google.protobuf.Timestamp
	2, // 2: Event.manifest:type_name -> YamlManifest
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_event_proto_init() }
func file_event_proto_init() {
	if File_event_proto != nil {
		return
	}
	file_yaml_manifest_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_event_proto_rawDesc), len(file_event_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_event_proto_goTypes,
		DependencyIndexes: file_event_proto_depIdxs,
		MessageInfos:      file_event_proto_msgTypes,
	}.Build()
	File_event_proto = out.File
	file_event_proto_goTypes = nil
	file_event_proto_depIdxs = nil
}
