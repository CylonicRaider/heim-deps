// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: testing/testvalidate/v1/test_validate.proto

package testvalidatev1

import (
	reflect "reflect"
	sync "sync"

	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SendRequest) Reset() {
	*x = SendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testing_testvalidate_v1_test_validate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendRequest) ProtoMessage() {}

func (x *SendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_testing_testvalidate_v1_test_validate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendRequest.ProtoReflect.Descriptor instead.
func (*SendRequest) Descriptor() ([]byte, []int) {
	return file_testing_testvalidate_v1_test_validate_proto_rawDescGZIP(), []int{0}
}

func (x *SendRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type SendResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendResponse) Reset() {
	*x = SendResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testing_testvalidate_v1_test_validate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendResponse) ProtoMessage() {}

func (x *SendResponse) ProtoReflect() protoreflect.Message {
	mi := &file_testing_testvalidate_v1_test_validate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendResponse.ProtoReflect.Descriptor instead.
func (*SendResponse) Descriptor() ([]byte, []int) {
	return file_testing_testvalidate_v1_test_validate_proto_rawDescGZIP(), []int{1}
}

type SendStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SendStreamRequest) Reset() {
	*x = SendStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testing_testvalidate_v1_test_validate_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendStreamRequest) ProtoMessage() {}

func (x *SendStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_testing_testvalidate_v1_test_validate_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendStreamRequest.ProtoReflect.Descriptor instead.
func (*SendStreamRequest) Descriptor() ([]byte, []int) {
	return file_testing_testvalidate_v1_test_validate_proto_rawDescGZIP(), []int{2}
}

func (x *SendStreamRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type SendStreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendStreamResponse) Reset() {
	*x = SendStreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testing_testvalidate_v1_test_validate_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendStreamResponse) ProtoMessage() {}

func (x *SendStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_testing_testvalidate_v1_test_validate_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendStreamResponse.ProtoReflect.Descriptor instead.
func (*SendStreamResponse) Descriptor() ([]byte, []int) {
	return file_testing_testvalidate_v1_test_validate_proto_rawDescGZIP(), []int{3}
}

var File_testing_testvalidate_v1_test_validate_proto protoreflect.FileDescriptor

var file_testing_testvalidate_v1_test_validate_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x74,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x30, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x21, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x60, 0x01, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x0e, 0x0a, 0x0c, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x36, 0x0a, 0x11, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48, 0x04,
	0x72, 0x02, 0x60, 0x01, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x14, 0x0a,
	0x12, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x32, 0xd7, 0x01, 0x0a, 0x13, 0x54, 0x65, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x55, 0x0a, 0x04, 0x53,
	0x65, 0x6e, 0x64, 0x12, 0x24, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65,
	0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x67, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x69, 0x0a, 0x0a, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x12, 0x2a, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x55, 0x5a,
	0x53, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2d, 0x65, 0x63, 0x6f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x72,
	0x70, 0x63, 0x2d, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2f, 0x74, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x65, 0x73, 0x74, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_testing_testvalidate_v1_test_validate_proto_rawDescOnce sync.Once
	file_testing_testvalidate_v1_test_validate_proto_rawDescData = file_testing_testvalidate_v1_test_validate_proto_rawDesc
)

func file_testing_testvalidate_v1_test_validate_proto_rawDescGZIP() []byte {
	file_testing_testvalidate_v1_test_validate_proto_rawDescOnce.Do(func() {
		file_testing_testvalidate_v1_test_validate_proto_rawDescData = protoimpl.X.CompressGZIP(file_testing_testvalidate_v1_test_validate_proto_rawDescData)
	})
	return file_testing_testvalidate_v1_test_validate_proto_rawDescData
}

var file_testing_testvalidate_v1_test_validate_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_testing_testvalidate_v1_test_validate_proto_goTypes = []interface{}{
	(*SendRequest)(nil),        // 0: testing.testvalidate.v1.SendRequest
	(*SendResponse)(nil),       // 1: testing.testvalidate.v1.SendResponse
	(*SendStreamRequest)(nil),  // 2: testing.testvalidate.v1.SendStreamRequest
	(*SendStreamResponse)(nil), // 3: testing.testvalidate.v1.SendStreamResponse
}
var file_testing_testvalidate_v1_test_validate_proto_depIdxs = []int32{
	0, // 0: testing.testvalidate.v1.TestValidateService.Send:input_type -> testing.testvalidate.v1.SendRequest
	2, // 1: testing.testvalidate.v1.TestValidateService.SendStream:input_type -> testing.testvalidate.v1.SendStreamRequest
	1, // 2: testing.testvalidate.v1.TestValidateService.Send:output_type -> testing.testvalidate.v1.SendResponse
	3, // 3: testing.testvalidate.v1.TestValidateService.SendStream:output_type -> testing.testvalidate.v1.SendStreamResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_testing_testvalidate_v1_test_validate_proto_init() }
func file_testing_testvalidate_v1_test_validate_proto_init() {
	if File_testing_testvalidate_v1_test_validate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_testing_testvalidate_v1_test_validate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendRequest); i {
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
		file_testing_testvalidate_v1_test_validate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendResponse); i {
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
		file_testing_testvalidate_v1_test_validate_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendStreamRequest); i {
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
		file_testing_testvalidate_v1_test_validate_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendStreamResponse); i {
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
			RawDescriptor: file_testing_testvalidate_v1_test_validate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_testing_testvalidate_v1_test_validate_proto_goTypes,
		DependencyIndexes: file_testing_testvalidate_v1_test_validate_proto_depIdxs,
		MessageInfos:      file_testing_testvalidate_v1_test_validate_proto_msgTypes,
	}.Build()
	File_testing_testvalidate_v1_test_validate_proto = out.File
	file_testing_testvalidate_v1_test_validate_proto_rawDesc = nil
	file_testing_testvalidate_v1_test_validate_proto_goTypes = nil
	file_testing_testvalidate_v1_test_validate_proto_depIdxs = nil
}
