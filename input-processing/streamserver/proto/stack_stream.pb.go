// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: stack_stream.proto

package proto

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Stream message containing text
type TextInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *TextInput) Reset() {
	*x = TextInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stack_stream_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextInput) ProtoMessage() {}

func (x *TextInput) ProtoReflect() protoreflect.Message {
	mi := &file_stack_stream_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextInput.ProtoReflect.Descriptor instead.
func (*TextInput) Descriptor() ([]byte, []int) {
	return file_stack_stream_proto_rawDescGZIP(), []int{0}
}

func (x *TextInput) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// The stream message containing lines that has word error in them
type ErrorWordLines struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ErrorWordLines) Reset() {
	*x = ErrorWordLines{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stack_stream_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorWordLines) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorWordLines) ProtoMessage() {}

func (x *ErrorWordLines) ProtoReflect() protoreflect.Message {
	mi := &file_stack_stream_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorWordLines.ProtoReflect.Descriptor instead.
func (*ErrorWordLines) Descriptor() ([]byte, []int) {
	return file_stack_stream_proto_rawDescGZIP(), []int{1}
}

func (x *ErrorWordLines) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_stack_stream_proto protoreflect.FileDescriptor

var file_stack_stream_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x25, 0x0a, 0x09, 0x54, 0x65, 0x78,
	0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x2a, 0x0a, 0x0e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x57, 0x6f, 0x72, 0x64, 0x4c, 0x69, 0x6e,
	0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x62, 0x0a, 0x0c,
	0x54, 0x65, 0x78, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x12, 0x52, 0x0a, 0x0d,
	0x46, 0x69, 0x6e, 0x64, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x57, 0x6f, 0x72, 0x64, 0x12, 0x10, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a,
	0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x57, 0x6f, 0x72,
	0x64, 0x4c, 0x69, 0x6e, 0x65, 0x73, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x22, 0x09,
	0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x3a, 0x01, 0x2a, 0x28, 0x01, 0x30, 0x01,
	0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stack_stream_proto_rawDescOnce sync.Once
	file_stack_stream_proto_rawDescData = file_stack_stream_proto_rawDesc
)

func file_stack_stream_proto_rawDescGZIP() []byte {
	file_stack_stream_proto_rawDescOnce.Do(func() {
		file_stack_stream_proto_rawDescData = protoimpl.X.CompressGZIP(file_stack_stream_proto_rawDescData)
	})
	return file_stack_stream_proto_rawDescData
}

var file_stack_stream_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_stack_stream_proto_goTypes = []interface{}{
	(*TextInput)(nil),      // 0: proto.TextInput
	(*ErrorWordLines)(nil), // 1: proto.ErrorWordLines
}
var file_stack_stream_proto_depIdxs = []int32{
	0, // 0: proto.TextStreamer.FindErrorWord:input_type -> proto.TextInput
	1, // 1: proto.TextStreamer.FindErrorWord:output_type -> proto.ErrorWordLines
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_stack_stream_proto_init() }
func file_stack_stream_proto_init() {
	if File_stack_stream_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stack_stream_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TextInput); i {
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
		file_stack_stream_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorWordLines); i {
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
			RawDescriptor: file_stack_stream_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_stack_stream_proto_goTypes,
		DependencyIndexes: file_stack_stream_proto_depIdxs,
		MessageInfos:      file_stack_stream_proto_msgTypes,
	}.Build()
	File_stack_stream_proto = out.File
	file_stack_stream_proto_rawDesc = nil
	file_stack_stream_proto_goTypes = nil
	file_stack_stream_proto_depIdxs = nil
}
