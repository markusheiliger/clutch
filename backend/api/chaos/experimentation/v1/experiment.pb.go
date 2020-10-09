// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: chaos/experimentation/v1/experiment.proto

package experimentationv1

import (
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Experiment_Status int32

const (
	// Unspecifiedf Status.
	Experiment_UNSPECIFIED Experiment_Status = 0
	// An experiment is specified to be run at a future date.
	Experiment_SCHEDULED Experiment_Status = 1
	// An experiment is currently running.
	Experiment_RUNNING Experiment_Status = 2
	// An experiment has completed.
	Experiment_COMPLETED Experiment_Status = 3
	// An experiment has been canceled before it started.
	Experiment_CANCELED Experiment_Status = 4
	// An experiment was running and was stopped before its originally scheduled end time.
	Experiment_STOPPED Experiment_Status = 5
)

// Enum value maps for Experiment_Status.
var (
	Experiment_Status_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "SCHEDULED",
		2: "RUNNING",
		3: "COMPLETED",
		4: "CANCELED",
		5: "STOPPED",
	}
	Experiment_Status_value = map[string]int32{
		"UNSPECIFIED": 0,
		"SCHEDULED":   1,
		"RUNNING":     2,
		"COMPLETED":   3,
		"CANCELED":    4,
		"STOPPED":     5,
	}
)

func (x Experiment_Status) Enum() *Experiment_Status {
	p := new(Experiment_Status)
	*p = x
	return p
}

func (x Experiment_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Experiment_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_chaos_experimentation_v1_experiment_proto_enumTypes[0].Descriptor()
}

func (Experiment_Status) Type() protoreflect.EnumType {
	return &file_chaos_experimentation_v1_experiment_proto_enumTypes[0]
}

func (x Experiment_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Experiment_Status.Descriptor instead.
func (Experiment_Status) EnumDescriptor() ([]byte, []int) {
	return file_chaos_experimentation_v1_experiment_proto_rawDescGZIP(), []int{0, 0}
}

type Experiment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// The experiment configuration specific to the type of experiment
	Config *any.Any `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	// The time when the experiment should start. If unspecified, defaults to 'now'
	StartTime *timestamp.Timestamp `protobuf:"bytes,3,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// The time when the experiment should end, If unspecified, the experiment runs indefinitely unless it is manually
	// stopped by a user.
	EndTime *timestamp.Timestamp `protobuf:"bytes,4,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
}

func (x *Experiment) Reset() {
	*x = Experiment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chaos_experimentation_v1_experiment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Experiment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Experiment) ProtoMessage() {}

func (x *Experiment) ProtoReflect() protoreflect.Message {
	mi := &file_chaos_experimentation_v1_experiment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Experiment.ProtoReflect.Descriptor instead.
func (*Experiment) Descriptor() ([]byte, []int) {
	return file_chaos_experimentation_v1_experiment_proto_rawDescGZIP(), []int{0}
}

func (x *Experiment) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Experiment) GetConfig() *any.Any {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *Experiment) GetStartTime() *timestamp.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *Experiment) GetEndTime() *timestamp.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

var File_chaos_experimentation_v1_experiment_proto protoreflect.FileDescriptor

var file_chaos_experimentation_v1_experiment_proto_rawDesc = []byte{
	0x0a, 0x29, 0x63, 0x68, 0x61, 0x6f, 0x73, 0x2f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65,
	0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x70, 0x65, 0x72,
	0x69, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x63, 0x6c, 0x75,
	0x74, 0x63, 0x68, 0x2e, 0x63, 0x68, 0x61, 0x6f, 0x73, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69,
	0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9d, 0x02, 0x0a, 0x0a, 0x45, 0x78, 0x70,
	0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2c, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x06, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07,
	0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x5f, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x43, 0x48, 0x45, 0x44, 0x55, 0x4c, 0x45, 0x44, 0x10,
	0x01, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x0d,
	0x0a, 0x09, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0c, 0x0a,
	0x08, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x45, 0x44, 0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x53,
	0x54, 0x4f, 0x50, 0x50, 0x45, 0x44, 0x10, 0x05, 0x42, 0x13, 0x5a, 0x11, 0x65, 0x78, 0x70, 0x65,
	0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chaos_experimentation_v1_experiment_proto_rawDescOnce sync.Once
	file_chaos_experimentation_v1_experiment_proto_rawDescData = file_chaos_experimentation_v1_experiment_proto_rawDesc
)

func file_chaos_experimentation_v1_experiment_proto_rawDescGZIP() []byte {
	file_chaos_experimentation_v1_experiment_proto_rawDescOnce.Do(func() {
		file_chaos_experimentation_v1_experiment_proto_rawDescData = protoimpl.X.CompressGZIP(file_chaos_experimentation_v1_experiment_proto_rawDescData)
	})
	return file_chaos_experimentation_v1_experiment_proto_rawDescData
}

var file_chaos_experimentation_v1_experiment_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_chaos_experimentation_v1_experiment_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_chaos_experimentation_v1_experiment_proto_goTypes = []interface{}{
	(Experiment_Status)(0),      // 0: clutch.chaos.experimentation.v1.Experiment.Status
	(*Experiment)(nil),          // 1: clutch.chaos.experimentation.v1.Experiment
	(*any.Any)(nil),             // 2: google.protobuf.Any
	(*timestamp.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_chaos_experimentation_v1_experiment_proto_depIdxs = []int32{
	2, // 0: clutch.chaos.experimentation.v1.Experiment.config:type_name -> google.protobuf.Any
	3, // 1: clutch.chaos.experimentation.v1.Experiment.start_time:type_name -> google.protobuf.Timestamp
	3, // 2: clutch.chaos.experimentation.v1.Experiment.end_time:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_chaos_experimentation_v1_experiment_proto_init() }
func file_chaos_experimentation_v1_experiment_proto_init() {
	if File_chaos_experimentation_v1_experiment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chaos_experimentation_v1_experiment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Experiment); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chaos_experimentation_v1_experiment_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_chaos_experimentation_v1_experiment_proto_goTypes,
		DependencyIndexes: file_chaos_experimentation_v1_experiment_proto_depIdxs,
		EnumInfos:         file_chaos_experimentation_v1_experiment_proto_enumTypes,
		MessageInfos:      file_chaos_experimentation_v1_experiment_proto_msgTypes,
	}.Build()
	File_chaos_experimentation_v1_experiment_proto = out.File
	file_chaos_experimentation_v1_experiment_proto_rawDesc = nil
	file_chaos_experimentation_v1_experiment_proto_goTypes = nil
	file_chaos_experimentation_v1_experiment_proto_depIdxs = nil
}