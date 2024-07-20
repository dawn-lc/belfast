// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: CHILD_PLAN_CELL.proto

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

type CHILD_PLAN_CELL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Day   *uint32           `protobuf:"varint,1,req,name=day" json:"day,omitempty"`
	Index *uint32           `protobuf:"varint,2,req,name=index" json:"index,omitempty"`
	Value []*CHILD_PLAN_VAL `protobuf:"bytes,3,rep,name=value" json:"value,omitempty"`
}

func (x *CHILD_PLAN_CELL) Reset() {
	*x = CHILD_PLAN_CELL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CHILD_PLAN_CELL_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CHILD_PLAN_CELL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CHILD_PLAN_CELL) ProtoMessage() {}

func (x *CHILD_PLAN_CELL) ProtoReflect() protoreflect.Message {
	mi := &file_CHILD_PLAN_CELL_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CHILD_PLAN_CELL.ProtoReflect.Descriptor instead.
func (*CHILD_PLAN_CELL) Descriptor() ([]byte, []int) {
	return file_CHILD_PLAN_CELL_proto_rawDescGZIP(), []int{0}
}

func (x *CHILD_PLAN_CELL) GetDay() uint32 {
	if x != nil && x.Day != nil {
		return *x.Day
	}
	return 0
}

func (x *CHILD_PLAN_CELL) GetIndex() uint32 {
	if x != nil && x.Index != nil {
		return *x.Index
	}
	return 0
}

func (x *CHILD_PLAN_CELL) GetValue() []*CHILD_PLAN_VAL {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_CHILD_PLAN_CELL_proto protoreflect.FileDescriptor

var file_CHILD_PLAN_CELL_proto_rawDesc = []byte{
	0x0a, 0x15, 0x43, 0x48, 0x49, 0x4c, 0x44, 0x5f, 0x50, 0x4c, 0x41, 0x4e, 0x5f, 0x43, 0x45, 0x4c,
	0x4c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74,
	0x1a, 0x14, 0x43, 0x48, 0x49, 0x4c, 0x44, 0x5f, 0x50, 0x4c, 0x41, 0x4e, 0x5f, 0x56, 0x41, 0x4c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x68, 0x0a, 0x0f, 0x43, 0x48, 0x49, 0x4c, 0x44, 0x5f,
	0x50, 0x4c, 0x41, 0x4e, 0x5f, 0x43, 0x45, 0x4c, 0x4c, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x61, 0x79,
	0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x03, 0x64, 0x61, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x12, 0x2d, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74, 0x2e, 0x43, 0x48, 0x49, 0x4c, 0x44,
	0x5f, 0x50, 0x4c, 0x41, 0x4e, 0x5f, 0x56, 0x41, 0x4c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
}

var (
	file_CHILD_PLAN_CELL_proto_rawDescOnce sync.Once
	file_CHILD_PLAN_CELL_proto_rawDescData = file_CHILD_PLAN_CELL_proto_rawDesc
)

func file_CHILD_PLAN_CELL_proto_rawDescGZIP() []byte {
	file_CHILD_PLAN_CELL_proto_rawDescOnce.Do(func() {
		file_CHILD_PLAN_CELL_proto_rawDescData = protoimpl.X.CompressGZIP(file_CHILD_PLAN_CELL_proto_rawDescData)
	})
	return file_CHILD_PLAN_CELL_proto_rawDescData
}

var file_CHILD_PLAN_CELL_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_CHILD_PLAN_CELL_proto_goTypes = []any{
	(*CHILD_PLAN_CELL)(nil), // 0: belfast.CHILD_PLAN_CELL
	(*CHILD_PLAN_VAL)(nil),  // 1: belfast.CHILD_PLAN_VAL
}
var file_CHILD_PLAN_CELL_proto_depIdxs = []int32{
	1, // 0: belfast.CHILD_PLAN_CELL.value:type_name -> belfast.CHILD_PLAN_VAL
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_CHILD_PLAN_CELL_proto_init() }
func file_CHILD_PLAN_CELL_proto_init() {
	if File_CHILD_PLAN_CELL_proto != nil {
		return
	}
	file_CHILD_PLAN_VAL_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_CHILD_PLAN_CELL_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CHILD_PLAN_CELL); i {
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
			RawDescriptor: file_CHILD_PLAN_CELL_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CHILD_PLAN_CELL_proto_goTypes,
		DependencyIndexes: file_CHILD_PLAN_CELL_proto_depIdxs,
		MessageInfos:      file_CHILD_PLAN_CELL_proto_msgTypes,
	}.Build()
	File_CHILD_PLAN_CELL_proto = out.File
	file_CHILD_PLAN_CELL_proto_rawDesc = nil
	file_CHILD_PLAN_CELL_proto_goTypes = nil
	file_CHILD_PLAN_CELL_proto_depIdxs = nil
}
