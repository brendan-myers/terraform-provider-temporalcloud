// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: temporal/api/cloud/operation/v1/message.proto

package operationv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AsyncOperation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The operation id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The current state of this operation
	// Possible values are: PENDING, IN_PROGRESS, FAILED, CANCELLED, FULFILLED
	State string `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	// The recommended duration to check back for an update in the operation's state
	CheckDuration *durationpb.Duration `protobuf:"bytes,3,opt,name=check_duration,json=checkDuration,proto3" json:"check_duration,omitempty"`
	// The type of operation being performed
	OperationType string `protobuf:"bytes,4,opt,name=operation_type,json=operationType,proto3" json:"operation_type,omitempty"`
	// The input to the operation being performed
	OperationInput *anypb.Any `protobuf:"bytes,5,opt,name=operation_input,json=operationInput,proto3" json:"operation_input,omitempty"`
	// If the operation failed, the reason for the failure
	FailureReason string `protobuf:"bytes,6,opt,name=failure_reason,json=failureReason,proto3" json:"failure_reason,omitempty"`
	// The date and time when the operation initiated
	StartedTime *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=started_time,json=startedTime,proto3" json:"started_time,omitempty"`
	// The date and time when the operation completed
	FinishedTime *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=finished_time,json=finishedTime,proto3" json:"finished_time,omitempty"`
}

func (x *AsyncOperation) Reset() {
	*x = AsyncOperation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_api_cloud_operation_v1_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AsyncOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AsyncOperation) ProtoMessage() {}

func (x *AsyncOperation) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_cloud_operation_v1_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AsyncOperation.ProtoReflect.Descriptor instead.
func (*AsyncOperation) Descriptor() ([]byte, []int) {
	return file_temporal_api_cloud_operation_v1_message_proto_rawDescGZIP(), []int{0}
}

func (x *AsyncOperation) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AsyncOperation) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *AsyncOperation) GetCheckDuration() *durationpb.Duration {
	if x != nil {
		return x.CheckDuration
	}
	return nil
}

func (x *AsyncOperation) GetOperationType() string {
	if x != nil {
		return x.OperationType
	}
	return ""
}

func (x *AsyncOperation) GetOperationInput() *anypb.Any {
	if x != nil {
		return x.OperationInput
	}
	return nil
}

func (x *AsyncOperation) GetFailureReason() string {
	if x != nil {
		return x.FailureReason
	}
	return ""
}

func (x *AsyncOperation) GetStartedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartedTime
	}
	return nil
}

func (x *AsyncOperation) GetFinishedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.FinishedTime
	}
	return nil
}

var File_temporal_api_cloud_operation_v1_message_proto protoreflect.FileDescriptor

var file_temporal_api_cloud_operation_v1_message_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76,
	0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1f, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85, 0x03, 0x0a,
	0x0e, 0x41, 0x73, 0x79, 0x6e, 0x63, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x40, 0x0a, 0x0e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3d,
	0x0a, 0x0f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x0e, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x25, 0x0a,
	0x0e, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x52, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x12, 0x3d, 0x0a, 0x0c, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x3f, 0x0a, 0x0d, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x42, 0xc1, 0x02, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x65, 0x6d,
	0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x6b, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61,
	0x6c, 0x69, 0x6f, 0x2f, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x2d, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2d, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x65,
	0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x76, 0x31, 0xa2, 0x02, 0x04, 0x54, 0x41, 0x43, 0x4f,
	0xaa, 0x02, 0x1f, 0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x41, 0x70, 0x69, 0x2e,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x56, 0x31, 0xca, 0x02, 0x1f, 0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x5c, 0x41, 0x70,
	0x69, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x2b, 0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x5c,
	0x41, 0x70, 0x69, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x23, 0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x3a, 0x3a, 0x41,
	0x70, 0x69, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_temporal_api_cloud_operation_v1_message_proto_rawDescOnce sync.Once
	file_temporal_api_cloud_operation_v1_message_proto_rawDescData = file_temporal_api_cloud_operation_v1_message_proto_rawDesc
)

func file_temporal_api_cloud_operation_v1_message_proto_rawDescGZIP() []byte {
	file_temporal_api_cloud_operation_v1_message_proto_rawDescOnce.Do(func() {
		file_temporal_api_cloud_operation_v1_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_temporal_api_cloud_operation_v1_message_proto_rawDescData)
	})
	return file_temporal_api_cloud_operation_v1_message_proto_rawDescData
}

var file_temporal_api_cloud_operation_v1_message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_temporal_api_cloud_operation_v1_message_proto_goTypes = []interface{}{
	(*AsyncOperation)(nil),        // 0: temporal.api.cloud.operation.v1.AsyncOperation
	(*durationpb.Duration)(nil),   // 1: google.protobuf.Duration
	(*anypb.Any)(nil),             // 2: google.protobuf.Any
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_temporal_api_cloud_operation_v1_message_proto_depIdxs = []int32{
	1, // 0: temporal.api.cloud.operation.v1.AsyncOperation.check_duration:type_name -> google.protobuf.Duration
	2, // 1: temporal.api.cloud.operation.v1.AsyncOperation.operation_input:type_name -> google.protobuf.Any
	3, // 2: temporal.api.cloud.operation.v1.AsyncOperation.started_time:type_name -> google.protobuf.Timestamp
	3, // 3: temporal.api.cloud.operation.v1.AsyncOperation.finished_time:type_name -> google.protobuf.Timestamp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_temporal_api_cloud_operation_v1_message_proto_init() }
func file_temporal_api_cloud_operation_v1_message_proto_init() {
	if File_temporal_api_cloud_operation_v1_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_temporal_api_cloud_operation_v1_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AsyncOperation); i {
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
			RawDescriptor: file_temporal_api_cloud_operation_v1_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_temporal_api_cloud_operation_v1_message_proto_goTypes,
		DependencyIndexes: file_temporal_api_cloud_operation_v1_message_proto_depIdxs,
		MessageInfos:      file_temporal_api_cloud_operation_v1_message_proto_msgTypes,
	}.Build()
	File_temporal_api_cloud_operation_v1_message_proto = out.File
	file_temporal_api_cloud_operation_v1_message_proto_rawDesc = nil
	file_temporal_api_cloud_operation_v1_message_proto_goTypes = nil
	file_temporal_api_cloud_operation_v1_message_proto_depIdxs = nil
}
