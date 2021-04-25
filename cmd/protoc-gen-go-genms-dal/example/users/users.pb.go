// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.4
// source: users.proto

package users

import (
	proto "github.com/golang/protobuf/proto"
	_ "github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/annotations"
	types "github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/annotations/types"
	latlng "google.golang.org/genproto/googleapis/type/latlng"
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

type User_Kind int32

const (
	User_Kind_User  User_Kind = 0
	User_Kind_Admin User_Kind = 1
)

// Enum value maps for User_Kind.
var (
	User_Kind_name = map[int32]string{
		0: "Kind_User",
		1: "Kind_Admin",
	}
	User_Kind_value = map[string]int32{
		"Kind_User":  0,
		"Kind_Admin": 1,
	}
)

func (x User_Kind) Enum() *User_Kind {
	p := new(User_Kind)
	*p = x
	return p
}

func (x User_Kind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (User_Kind) Descriptor() protoreflect.EnumDescriptor {
	return file_users_proto_enumTypes[0].Descriptor()
}

func (User_Kind) Type() protoreflect.EnumType {
	return &file_users_proto_enumTypes[0]
}

func (x User_Kind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use User_Kind.Descriptor instead.
func (User_Kind) EnumDescriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{0, 0}
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               int64          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name             string         `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Division         string         `protobuf:"bytes,3,opt,name=division,proto3" json:"division,omitempty"`
	LifetimeScore    float64        `protobuf:"fixed64,4,opt,name=lifetime_score,json=lifetimeScore,proto3" json:"lifetime_score,omitempty"`
	LastScore        float32        `protobuf:"fixed32,5,opt,name=last_score,json=lastScore,proto3" json:"last_score,omitempty"`
	LifetimeWinnings int64          `protobuf:"varint,6,opt,name=lifetime_winnings,json=lifetimeWinnings,proto3" json:"lifetime_winnings,omitempty"`
	LastWinnings     int32          `protobuf:"varint,7,opt,name=last_winnings,json=lastWinnings,proto3" json:"last_winnings,omitempty"`
	Point            *types.Point   `protobuf:"bytes,8,opt,name=point,proto3" json:"point,omitempty"`
	Phone            *types.Phone   `protobuf:"bytes,9,opt,name=phone,proto3" json:"phone,omitempty"`
	Geo              *latlng.LatLng `protobuf:"bytes,10,opt,name=geo,proto3" json:"geo,omitempty"`
	Kind             User_Kind      `protobuf:"varint,11,opt,name=kind,proto3,enum=greeter.User_Kind" json:"kind,omitempty"`
	ByBackend        string         `protobuf:"bytes,12,opt,name=by_backend,json=byBackend,proto3" json:"by_backend,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_users_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetDivision() string {
	if x != nil {
		return x.Division
	}
	return ""
}

func (x *User) GetLifetimeScore() float64 {
	if x != nil {
		return x.LifetimeScore
	}
	return 0
}

func (x *User) GetLastScore() float32 {
	if x != nil {
		return x.LastScore
	}
	return 0
}

func (x *User) GetLifetimeWinnings() int64 {
	if x != nil {
		return x.LifetimeWinnings
	}
	return 0
}

func (x *User) GetLastWinnings() int32 {
	if x != nil {
		return x.LastWinnings
	}
	return 0
}

func (x *User) GetPoint() *types.Point {
	if x != nil {
		return x.Point
	}
	return nil
}

func (x *User) GetPhone() *types.Phone {
	if x != nil {
		return x.Phone
	}
	return nil
}

func (x *User) GetGeo() *latlng.LatLng {
	if x != nil {
		return x.Geo
	}
	return nil
}

func (x *User) GetKind() User_Kind {
	if x != nil {
		return x.Kind
	}
	return User_Kind_User
}

func (x *User) GetByBackend() string {
	if x != nil {
		return x.ByBackend
	}
	return ""
}

var File_users_proto protoreflect.FileDescriptor

var file_users_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x67,
	0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x1a, 0x18, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x2f, 0x6c, 0x61, 0x74, 0x6c, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x35, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2d,
	0x67, 0x65, 0x6e, 0x6d, 0x73, 0x2d, 0x64, 0x61, 0x6c, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x35, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d,
	0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x6d, 0x73, 0x2d, 0x64, 0x61, 0x6c,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xef,
	0x05, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64,
	0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64,
	0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x6c, 0x69, 0x66, 0x65, 0x74,
	0x69, 0x6d, 0x65, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x0d, 0x6c, 0x69, 0x66, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x33, 0x0a,
	0x11, 0x6c, 0x69, 0x66, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x77, 0x69, 0x6e, 0x6e, 0x69, 0x6e,
	0x67, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x42, 0x06, 0xc2, 0xf3, 0x18, 0x02, 0x08, 0x01,
	0x52, 0x10, 0x6c, 0x69, 0x66, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x57, 0x69, 0x6e, 0x6e, 0x69, 0x6e,
	0x67, 0x73, 0x12, 0x31, 0x0a, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x77, 0x69, 0x6e, 0x6e, 0x69,
	0x6e, 0x67, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0c, 0xc2, 0xf3, 0x18, 0x08, 0x12,
	0x06, 0x70, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x57, 0x69, 0x6e,
	0x6e, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x2c, 0x0a, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x65, 0x6e, 0x6d, 0x73, 0x2e, 0x64, 0x61, 0x6c,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x05, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x65, 0x6e, 0x6d, 0x73, 0x2e, 0x64, 0x61, 0x6c, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x12, 0x25, 0x0a, 0x03, 0x67, 0x65, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x4c, 0x61, 0x74,
	0x4c, 0x6e, 0x67, 0x52, 0x03, 0x67, 0x65, 0x6f, 0x12, 0x32, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x2e, 0x4b, 0x69, 0x6e, 0x64, 0x42, 0x0a, 0xc2, 0xf3, 0x18, 0x06,
	0x12, 0x04, 0x74, 0x79, 0x70, 0x65, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x4d, 0x0a, 0x0a,
	0x62, 0x79, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x2e, 0xc2, 0xf3, 0x18, 0x2a, 0x1a, 0x15, 0x12, 0x13, 0x62, 0x79, 0x5f, 0x62, 0x61, 0x63,
	0x6b, 0x65, 0x6e, 0x64, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x22, 0x11, 0x12,
	0x0f, 0x62, 0x79, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x74,
	0x52, 0x09, 0x62, 0x79, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x22, 0x25, 0x0a, 0x04, 0x4b,
	0x69, 0x6e, 0x64, 0x12, 0x0d, 0x0a, 0x09, 0x4b, 0x69, 0x6e, 0x64, 0x5f, 0x55, 0x73, 0x65, 0x72,
	0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x4b, 0x69, 0x6e, 0x64, 0x5f, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x10, 0x01, 0x3a, 0xcb, 0x01, 0xc2, 0xf3, 0x18, 0xc6, 0x01, 0x0a, 0x0d, 0x0a, 0x05, 0x62, 0x79,
	0x20, 0x69, 0x64, 0x12, 0x04, 0x0a, 0x02, 0x69, 0x64, 0x0a, 0x2a, 0x0a, 0x14, 0x62, 0x79, 0x20,
	0x6e, 0x61, 0x6d, 0x65, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x64, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x06, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0a, 0x0a, 0x08, 0x64, 0x69, 0x76,
	0x69, 0x73, 0x69, 0x6f, 0x6e, 0x0a, 0x11, 0x0a, 0x07, 0x62, 0x79, 0x20, 0x6b, 0x69, 0x6e, 0x64,
	0x12, 0x06, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x0a, 0x23, 0x0a, 0x08, 0x62, 0x79, 0x20, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x12, 0x13, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0x0a, 0x0a,
	0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x10, 0x02, 0x22, 0x02, 0x08, 0x01, 0x0a, 0x1c, 0x0a,
	0x08, 0x62, 0x79, 0x20, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x12, 0x05, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x02, 0x0a, 0x16, 0x0a, 0x12, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x20, 0x73, 0x74, 0x75, 0x62, 0x20, 0x6f, 0x6e, 0x6c,
	0x79, 0x18, 0x01, 0x0a, 0x17, 0x0a, 0x13, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x20, 0x73, 0x74, 0x75, 0x62, 0x20, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x02, 0x12, 0x02, 0x01, 0x02,
	0x42, 0x45, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72,
	0x6c, 0x65, 0x73, 0x7a, 0x69, 0x6c, 0x6d, 0x2f, 0x67, 0x65, 0x6e, 0x6d, 0x73, 0x2f, 0x63, 0x6d,
	0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2d,
	0x67, 0x65, 0x6e, 0x6d, 0x73, 0x2d, 0x64, 0x61, 0x6c, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_users_proto_rawDescOnce sync.Once
	file_users_proto_rawDescData = file_users_proto_rawDesc
)

func file_users_proto_rawDescGZIP() []byte {
	file_users_proto_rawDescOnce.Do(func() {
		file_users_proto_rawDescData = protoimpl.X.CompressGZIP(file_users_proto_rawDescData)
	})
	return file_users_proto_rawDescData
}

var file_users_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_users_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_users_proto_goTypes = []interface{}{
	(User_Kind)(0),        // 0: greeter.User.Kind
	(*User)(nil),          // 1: greeter.User
	(*types.Point)(nil),   // 2: genms.dal.types.Point
	(*types.Phone)(nil),   // 3: genms.dal.types.Phone
	(*latlng.LatLng)(nil), // 4: google.type.LatLng
}
var file_users_proto_depIdxs = []int32{
	2, // 0: greeter.User.point:type_name -> genms.dal.types.Point
	3, // 1: greeter.User.phone:type_name -> genms.dal.types.Phone
	4, // 2: greeter.User.geo:type_name -> google.type.LatLng
	0, // 3: greeter.User.kind:type_name -> greeter.User.Kind
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_users_proto_init() }
func file_users_proto_init() {
	if File_users_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_users_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
			RawDescriptor: file_users_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_users_proto_goTypes,
		DependencyIndexes: file_users_proto_depIdxs,
		EnumInfos:         file_users_proto_enumTypes,
		MessageInfos:      file_users_proto_msgTypes,
	}.Build()
	File_users_proto = out.File
	file_users_proto_rawDesc = nil
	file_users_proto_goTypes = nil
	file_users_proto_depIdxs = nil
}
