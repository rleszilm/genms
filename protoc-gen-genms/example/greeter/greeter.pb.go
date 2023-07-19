// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.1
// source: greeter.proto

package greeter

import (
	_ "github.com/rleszilm/genms/protoc-gen-genms/annotations"
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

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value      string  `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Constraint *string `protobuf:"bytes,2,opt,name=constraint,proto3,oneof" json:"constraint,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greeter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_greeter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_greeter_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Message) GetConstraint() string {
	if x != nil && x.Constraint != nil {
		return *x.Constraint
	}
	return ""
}

var File_greeter_proto protoreflect.FileDescriptor

var file_greeter_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x67, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x67, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67,
	0x65, 0x6e, 0x2d, 0x67, 0x65, 0x6e, 0x6d, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x53, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x23, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x73, 0x74,
	0x72, 0x61, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x63,
	0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b,
	0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x32, 0xba, 0x02, 0x0a, 0x08,
	0x57, 0x69, 0x74, 0x68, 0x48, 0x54, 0x54, 0x50, 0x12, 0x46, 0x0a, 0x0b, 0x48, 0x65, 0x6c, 0x6c,
	0x6f, 0x48, 0x54, 0x54, 0x50, 0x55, 0x55, 0x12, 0x10, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x65,
	0x72, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x10, 0x2e, 0x67, 0x72, 0x65, 0x65,
	0x74, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x13, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0d, 0x12, 0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x75, 0x75,
	0x12, 0x48, 0x0a, 0x0b, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x48, 0x54, 0x54, 0x50, 0x55, 0x53, 0x12,
	0x10, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x1a, 0x10, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x12, 0x0b, 0x2f, 0x76, 0x31,
	0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x75, 0x73, 0x30, 0x01, 0x12, 0x48, 0x0a, 0x0b, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x48, 0x54, 0x54, 0x50, 0x53, 0x55, 0x12, 0x10, 0x2e, 0x67, 0x72, 0x65, 0x65,
	0x74, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x10, 0x2e, 0x67, 0x72,
	0x65, 0x65, 0x74, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x13, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x12, 0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f,
	0x73, 0x75, 0x28, 0x01, 0x12, 0x4a, 0x0a, 0x0b, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x48, 0x54, 0x54,
	0x50, 0x53, 0x53, 0x12, 0x10, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x10, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x12,
	0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x73, 0x73, 0x28, 0x01, 0x30, 0x01,
	0x1a, 0x06, 0xc2, 0xf3, 0x18, 0x02, 0x08, 0x01, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x6c, 0x65, 0x73, 0x7a, 0x69, 0x6c, 0x6d, 0x2f,
	0x67, 0x65, 0x6e, 0x6d, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x6d, 0x73, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2f, 0x67, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_greeter_proto_rawDescOnce sync.Once
	file_greeter_proto_rawDescData = file_greeter_proto_rawDesc
)

func file_greeter_proto_rawDescGZIP() []byte {
	file_greeter_proto_rawDescOnce.Do(func() {
		file_greeter_proto_rawDescData = protoimpl.X.CompressGZIP(file_greeter_proto_rawDescData)
	})
	return file_greeter_proto_rawDescData
}

var file_greeter_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_greeter_proto_goTypes = []interface{}{
	(*Message)(nil), // 0: greeter.Message
}
var file_greeter_proto_depIdxs = []int32{
	0, // 0: greeter.WithHTTP.HelloHTTPUU:input_type -> greeter.Message
	0, // 1: greeter.WithHTTP.HelloHTTPUS:input_type -> greeter.Message
	0, // 2: greeter.WithHTTP.HelloHTTPSU:input_type -> greeter.Message
	0, // 3: greeter.WithHTTP.HelloHTTPSS:input_type -> greeter.Message
	0, // 4: greeter.WithHTTP.HelloHTTPUU:output_type -> greeter.Message
	0, // 5: greeter.WithHTTP.HelloHTTPUS:output_type -> greeter.Message
	0, // 6: greeter.WithHTTP.HelloHTTPSU:output_type -> greeter.Message
	0, // 7: greeter.WithHTTP.HelloHTTPSS:output_type -> greeter.Message
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_greeter_proto_init() }
func file_greeter_proto_init() {
	if File_greeter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_greeter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
	file_greeter_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_greeter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_greeter_proto_goTypes,
		DependencyIndexes: file_greeter_proto_depIdxs,
		MessageInfos:      file_greeter_proto_msgTypes,
	}.Build()
	File_greeter_proto = out.File
	file_greeter_proto_rawDesc = nil
	file_greeter_proto_goTypes = nil
	file_greeter_proto_depIdxs = nil
}