// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: CS_13103.proto

package protobuf

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

type CS_13103 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Act      *uint32 `protobuf:"varint,1,req,name=act" json:"act,omitempty"`
	GroupId  *uint32 `protobuf:"varint,2,req,name=group_id,json=groupId" json:"group_id,omitempty"`
	ActArg_1 *uint32 `protobuf:"varint,3,opt,name=act_arg_1,json=actArg1" json:"act_arg_1,omitempty"`
	ActArg_2 *uint32 `protobuf:"varint,4,opt,name=act_arg_2,json=actArg2" json:"act_arg_2,omitempty"`
	ActArg_3 *uint32 `protobuf:"varint,5,opt,name=act_arg_3,json=actArg3" json:"act_arg_3,omitempty"`
	ActArg_4 *uint32 `protobuf:"varint,6,opt,name=act_arg_4,json=actArg4" json:"act_arg_4,omitempty"`
	ActArg_5 *uint32 `protobuf:"varint,7,opt,name=act_arg_5,json=actArg5" json:"act_arg_5,omitempty"`
}

func (x *CS_13103) Reset() {
	*x = CS_13103{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CS_13103_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CS_13103) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CS_13103) ProtoMessage() {}

func (x *CS_13103) ProtoReflect() protoreflect.Message {
	mi := &file_CS_13103_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CS_13103.ProtoReflect.Descriptor instead.
func (*CS_13103) Descriptor() ([]byte, []int) {
	return file_CS_13103_proto_rawDescGZIP(), []int{0}
}

func (x *CS_13103) GetAct() uint32 {
	if x != nil && x.Act != nil {
		return *x.Act
	}
	return 0
}

func (x *CS_13103) GetGroupId() uint32 {
	if x != nil && x.GroupId != nil {
		return *x.GroupId
	}
	return 0
}

func (x *CS_13103) GetActArg_1() uint32 {
	if x != nil && x.ActArg_1 != nil {
		return *x.ActArg_1
	}
	return 0
}

func (x *CS_13103) GetActArg_2() uint32 {
	if x != nil && x.ActArg_2 != nil {
		return *x.ActArg_2
	}
	return 0
}

func (x *CS_13103) GetActArg_3() uint32 {
	if x != nil && x.ActArg_3 != nil {
		return *x.ActArg_3
	}
	return 0
}

func (x *CS_13103) GetActArg_4() uint32 {
	if x != nil && x.ActArg_4 != nil {
		return *x.ActArg_4
	}
	return 0
}

func (x *CS_13103) GetActArg_5() uint32 {
	if x != nil && x.ActArg_5 != nil {
		return *x.ActArg_5
	}
	return 0
}

var File_CS_13103_proto protoreflect.FileDescriptor

var file_CS_13103_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x43, 0x53, 0x5f, 0x31, 0x33, 0x31, 0x30, 0x33, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74, 0x22, 0xc3, 0x01, 0x0a, 0x08, 0x43, 0x53,
	0x5f, 0x31, 0x33, 0x31, 0x30, 0x33, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x63, 0x74, 0x18, 0x01, 0x20,
	0x02, 0x28, 0x0d, 0x52, 0x03, 0x61, 0x63, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x09, 0x61, 0x63, 0x74, 0x5f, 0x61, 0x72, 0x67, 0x5f, 0x31,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x61, 0x63, 0x74, 0x41, 0x72, 0x67, 0x31, 0x12,
	0x1a, 0x0a, 0x09, 0x61, 0x63, 0x74, 0x5f, 0x61, 0x72, 0x67, 0x5f, 0x32, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x07, 0x61, 0x63, 0x74, 0x41, 0x72, 0x67, 0x32, 0x12, 0x1a, 0x0a, 0x09, 0x61,
	0x63, 0x74, 0x5f, 0x61, 0x72, 0x67, 0x5f, 0x33, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07,
	0x61, 0x63, 0x74, 0x41, 0x72, 0x67, 0x33, 0x12, 0x1a, 0x0a, 0x09, 0x61, 0x63, 0x74, 0x5f, 0x61,
	0x72, 0x67, 0x5f, 0x34, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x61, 0x63, 0x74, 0x41,
	0x72, 0x67, 0x34, 0x12, 0x1a, 0x0a, 0x09, 0x61, 0x63, 0x74, 0x5f, 0x61, 0x72, 0x67, 0x5f, 0x35,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x61, 0x63, 0x74, 0x41, 0x72, 0x67, 0x35, 0x42,
	0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
}

var (
	file_CS_13103_proto_rawDescOnce sync.Once
	file_CS_13103_proto_rawDescData = file_CS_13103_proto_rawDesc
)

func file_CS_13103_proto_rawDescGZIP() []byte {
	file_CS_13103_proto_rawDescOnce.Do(func() {
		file_CS_13103_proto_rawDescData = protoimpl.X.CompressGZIP(file_CS_13103_proto_rawDescData)
	})
	return file_CS_13103_proto_rawDescData
}

var file_CS_13103_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_CS_13103_proto_goTypes = []any{
	(*CS_13103)(nil), // 0: belfast.CS_13103
}
var file_CS_13103_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_CS_13103_proto_init() }
func file_CS_13103_proto_init() {
	if File_CS_13103_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_CS_13103_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CS_13103); i {
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
			RawDescriptor: file_CS_13103_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CS_13103_proto_goTypes,
		DependencyIndexes: file_CS_13103_proto_depIdxs,
		MessageInfos:      file_CS_13103_proto_msgTypes,
	}.Build()
	File_CS_13103_proto = out.File
	file_CS_13103_proto_rawDesc = nil
	file_CS_13103_proto_goTypes = nil
	file_CS_13103_proto_depIdxs = nil
}
