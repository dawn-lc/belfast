// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: CS_60033.proto

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

type CS_60033 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type *uint32 `protobuf:"varint,1,req,name=type" json:"type,omitempty"`
}

func (x *CS_60033) Reset() {
	*x = CS_60033{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CS_60033_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CS_60033) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CS_60033) ProtoMessage() {}

func (x *CS_60033) ProtoReflect() protoreflect.Message {
	mi := &file_CS_60033_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CS_60033.ProtoReflect.Descriptor instead.
func (*CS_60033) Descriptor() ([]byte, []int) {
	return file_CS_60033_proto_rawDescGZIP(), []int{0}
}

func (x *CS_60033) GetType() uint32 {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return 0
}

var File_CS_60033_proto protoreflect.FileDescriptor

var file_CS_60033_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x43, 0x53, 0x5f, 0x36, 0x30, 0x30, 0x33, 0x33, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74, 0x22, 0x1e, 0x0a, 0x08, 0x43, 0x53, 0x5f,
	0x36, 0x30, 0x30, 0x33, 0x33, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x02, 0x28, 0x0d, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
}

var (
	file_CS_60033_proto_rawDescOnce sync.Once
	file_CS_60033_proto_rawDescData = file_CS_60033_proto_rawDesc
)

func file_CS_60033_proto_rawDescGZIP() []byte {
	file_CS_60033_proto_rawDescOnce.Do(func() {
		file_CS_60033_proto_rawDescData = protoimpl.X.CompressGZIP(file_CS_60033_proto_rawDescData)
	})
	return file_CS_60033_proto_rawDescData
}

var file_CS_60033_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_CS_60033_proto_goTypes = []any{
	(*CS_60033)(nil), // 0: belfast.CS_60033
}
var file_CS_60033_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_CS_60033_proto_init() }
func file_CS_60033_proto_init() {
	if File_CS_60033_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_CS_60033_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CS_60033); i {
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
			RawDescriptor: file_CS_60033_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CS_60033_proto_goTypes,
		DependencyIndexes: file_CS_60033_proto_depIdxs,
		MessageInfos:      file_CS_60033_proto_msgTypes,
	}.Build()
	File_CS_60033_proto = out.File
	file_CS_60033_proto_rawDesc = nil
	file_CS_60033_proto_goTypes = nil
	file_CS_60033_proto_depIdxs = nil
}
