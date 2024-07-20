// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: CELLSINFO.proto

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

type CELLSINFO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Row    *uint32 `protobuf:"varint,1,req,name=row" json:"row,omitempty"`
	Column *uint32 `protobuf:"varint,2,req,name=column" json:"column,omitempty"`
	Color  *uint32 `protobuf:"varint,3,req,name=color" json:"color,omitempty"`
}

func (x *CELLSINFO) Reset() {
	*x = CELLSINFO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CELLSINFO_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CELLSINFO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CELLSINFO) ProtoMessage() {}

func (x *CELLSINFO) ProtoReflect() protoreflect.Message {
	mi := &file_CELLSINFO_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CELLSINFO.ProtoReflect.Descriptor instead.
func (*CELLSINFO) Descriptor() ([]byte, []int) {
	return file_CELLSINFO_proto_rawDescGZIP(), []int{0}
}

func (x *CELLSINFO) GetRow() uint32 {
	if x != nil && x.Row != nil {
		return *x.Row
	}
	return 0
}

func (x *CELLSINFO) GetColumn() uint32 {
	if x != nil && x.Column != nil {
		return *x.Column
	}
	return 0
}

func (x *CELLSINFO) GetColor() uint32 {
	if x != nil && x.Color != nil {
		return *x.Color
	}
	return 0
}

var File_CELLSINFO_proto protoreflect.FileDescriptor

var file_CELLSINFO_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x43, 0x45, 0x4c, 0x4c, 0x53, 0x49, 0x4e, 0x46, 0x4f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x07, 0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74, 0x22, 0x4b, 0x0a, 0x09, 0x43, 0x45,
	0x4c, 0x4c, 0x53, 0x49, 0x4e, 0x46, 0x4f, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x6f, 0x77, 0x18, 0x01,
	0x20, 0x02, 0x28, 0x0d, 0x52, 0x03, 0x72, 0x6f, 0x77, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6c,
	0x75, 0x6d, 0x6e, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d,
	0x6e, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x02, 0x28, 0x0d,
	0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66,
}

var (
	file_CELLSINFO_proto_rawDescOnce sync.Once
	file_CELLSINFO_proto_rawDescData = file_CELLSINFO_proto_rawDesc
)

func file_CELLSINFO_proto_rawDescGZIP() []byte {
	file_CELLSINFO_proto_rawDescOnce.Do(func() {
		file_CELLSINFO_proto_rawDescData = protoimpl.X.CompressGZIP(file_CELLSINFO_proto_rawDescData)
	})
	return file_CELLSINFO_proto_rawDescData
}

var file_CELLSINFO_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_CELLSINFO_proto_goTypes = []any{
	(*CELLSINFO)(nil), // 0: belfast.CELLSINFO
}
var file_CELLSINFO_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_CELLSINFO_proto_init() }
func file_CELLSINFO_proto_init() {
	if File_CELLSINFO_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_CELLSINFO_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CELLSINFO); i {
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
			RawDescriptor: file_CELLSINFO_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CELLSINFO_proto_goTypes,
		DependencyIndexes: file_CELLSINFO_proto_depIdxs,
		MessageInfos:      file_CELLSINFO_proto_msgTypes,
	}.Build()
	File_CELLSINFO_proto = out.File
	file_CELLSINFO_proto_rawDesc = nil
	file_CELLSINFO_proto_goTypes = nil
	file_CELLSINFO_proto_depIdxs = nil
}
