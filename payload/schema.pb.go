// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.2
// source: payload/schema.proto

package payload

import (
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

type Flag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Flag) Reset() {
	*x = Flag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payload_schema_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Flag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Flag) ProtoMessage() {}

func (x *Flag) ProtoReflect() protoreflect.Message {
	mi := &file_payload_schema_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Flag.ProtoReflect.Descriptor instead.
func (*Flag) Descriptor() ([]byte, []int) {
	return file_payload_schema_proto_rawDescGZIP(), []int{0}
}

func (x *Flag) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type FlagReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *FlagReply) Reset() {
	*x = FlagReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payload_schema_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlagReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlagReply) ProtoMessage() {}

func (x *FlagReply) ProtoReflect() protoreflect.Message {
	mi := &file_payload_schema_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlagReply.ProtoReflect.Descriptor instead.
func (*FlagReply) Descriptor() ([]byte, []int) {
	return file_payload_schema_proto_rawDescGZIP(), []int{1}
}

func (x *FlagReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_payload_schema_proto protoreflect.FileDescriptor

var file_payload_schema_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22,
	0x16, 0x0a, 0x04, 0x46, 0x6c, 0x61, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x25, 0x0a, 0x09, 0x46, 0x6c, 0x61, 0x67, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x3a,
	0x0a, 0x08, 0x47, 0x72, 0x70, 0x63, 0x53, 0x69, 0x6e, 0x6b, 0x12, 0x2e, 0x0a, 0x07, 0x52, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x12, 0x0d, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e,
	0x46, 0x6c, 0x61, 0x67, 0x1a, 0x12, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x46,
	0x6c, 0x61, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x46, 0x0a, 0x19, 0x63, 0x6f,
	0x6d, 0x2e, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x66, 0x69, 0x72, 0x65, 0x68, 0x6f, 0x73, 0x65, 0x2e,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x69, 0x6b, 0x75, 0x70, 0x69, 0x6e, 0x2f, 0x6c, 0x65, 0x61, 0x72,
	0x6e, 0x2d, 0x66, 0x69, 0x72, 0x65, 0x68, 0x6f, 0x73, 0x65, 0x2f, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_payload_schema_proto_rawDescOnce sync.Once
	file_payload_schema_proto_rawDescData = file_payload_schema_proto_rawDesc
)

func file_payload_schema_proto_rawDescGZIP() []byte {
	file_payload_schema_proto_rawDescOnce.Do(func() {
		file_payload_schema_proto_rawDescData = protoimpl.X.CompressGZIP(file_payload_schema_proto_rawDescData)
	})
	return file_payload_schema_proto_rawDescData
}

var file_payload_schema_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_payload_schema_proto_goTypes = []interface{}{
	(*Flag)(nil),      // 0: payload.Flag
	(*FlagReply)(nil), // 1: payload.FlagReply
}
var file_payload_schema_proto_depIdxs = []int32{
	0, // 0: payload.GrpcSink.Receive:input_type -> payload.Flag
	1, // 1: payload.GrpcSink.Receive:output_type -> payload.FlagReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_payload_schema_proto_init() }
func file_payload_schema_proto_init() {
	if File_payload_schema_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_payload_schema_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Flag); i {
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
		file_payload_schema_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlagReply); i {
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
			RawDescriptor: file_payload_schema_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_payload_schema_proto_goTypes,
		DependencyIndexes: file_payload_schema_proto_depIdxs,
		MessageInfos:      file_payload_schema_proto_msgTypes,
	}.Build()
	File_payload_schema_proto = out.File
	file_payload_schema_proto_rawDesc = nil
	file_payload_schema_proto_goTypes = nil
	file_payload_schema_proto_depIdxs = nil
}
