// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.2
// source: msg_id.proto

package pb

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

// 消息ID
type MsgID int32

const (
	MsgID_UNKNOWN_MSGID MsgID = 0 // 未知消息ID
	MsgID_HEARTBEAT     MsgID = 1 // 心跳
	MsgID_LOGIN         MsgID = 2 // 登录
)

// Enum value maps for MsgID.
var (
	MsgID_name = map[int32]string{
		0: "UNKNOWN_MSGID",
		1: "HEARTBEAT",
		2: "LOGIN",
	}
	MsgID_value = map[string]int32{
		"UNKNOWN_MSGID": 0,
		"HEARTBEAT":     1,
		"LOGIN":         2,
	}
)

func (x MsgID) Enum() *MsgID {
	p := new(MsgID)
	*p = x
	return p
}

func (x MsgID) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MsgID) Descriptor() protoreflect.EnumDescriptor {
	return file_msg_id_proto_enumTypes[0].Descriptor()
}

func (MsgID) Type() protoreflect.EnumType {
	return &file_msg_id_proto_enumTypes[0]
}

func (x MsgID) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MsgID.Descriptor instead.
func (MsgID) EnumDescriptor() ([]byte, []int) {
	return file_msg_id_proto_rawDescGZIP(), []int{0}
}

var File_msg_id_proto protoreflect.FileDescriptor

var file_msg_id_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6d, 0x73, 0x67, 0x5f, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x70, 0x62, 0x2a, 0x34, 0x0a, 0x05, 0x4d, 0x73, 0x67, 0x49, 0x44, 0x12, 0x11, 0x0a, 0x0d, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x4d, 0x53, 0x47, 0x49, 0x44, 0x10, 0x00, 0x12, 0x0d,
	0x0a, 0x09, 0x48, 0x45, 0x41, 0x52, 0x54, 0x42, 0x45, 0x41, 0x54, 0x10, 0x01, 0x12, 0x09, 0x0a,
	0x05, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x10, 0x02, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_msg_id_proto_rawDescOnce sync.Once
	file_msg_id_proto_rawDescData = file_msg_id_proto_rawDesc
)

func file_msg_id_proto_rawDescGZIP() []byte {
	file_msg_id_proto_rawDescOnce.Do(func() {
		file_msg_id_proto_rawDescData = protoimpl.X.CompressGZIP(file_msg_id_proto_rawDescData)
	})
	return file_msg_id_proto_rawDescData
}

var file_msg_id_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_msg_id_proto_goTypes = []interface{}{
	(MsgID)(0), // 0: pb.MsgID
}
var file_msg_id_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_msg_id_proto_init() }
func file_msg_id_proto_init() {
	if File_msg_id_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_msg_id_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_msg_id_proto_goTypes,
		DependencyIndexes: file_msg_id_proto_depIdxs,
		EnumInfos:         file_msg_id_proto_enumTypes,
	}.Build()
	File_msg_id_proto = out.File
	file_msg_id_proto_rawDesc = nil
	file_msg_id_proto_goTypes = nil
	file_msg_id_proto_depIdxs = nil
}