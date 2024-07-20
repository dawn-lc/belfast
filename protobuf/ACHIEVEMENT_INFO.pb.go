// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: ACHIEVEMENT_INFO.proto

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

type ACHIEVEMENT_INFO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        *uint32 `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Progress  *uint32 `protobuf:"varint,2,req,name=progress" json:"progress,omitempty"`
	Timestamp *uint32 `protobuf:"varint,3,req,name=timestamp" json:"timestamp,omitempty"`
}

func (x *ACHIEVEMENT_INFO) Reset() {
	*x = ACHIEVEMENT_INFO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ACHIEVEMENT_INFO_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ACHIEVEMENT_INFO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ACHIEVEMENT_INFO) ProtoMessage() {}

func (x *ACHIEVEMENT_INFO) ProtoReflect() protoreflect.Message {
	mi := &file_ACHIEVEMENT_INFO_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ACHIEVEMENT_INFO.ProtoReflect.Descriptor instead.
func (*ACHIEVEMENT_INFO) Descriptor() ([]byte, []int) {
	return file_ACHIEVEMENT_INFO_proto_rawDescGZIP(), []int{0}
}

func (x *ACHIEVEMENT_INFO) GetId() uint32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *ACHIEVEMENT_INFO) GetProgress() uint32 {
	if x != nil && x.Progress != nil {
		return *x.Progress
	}
	return 0
}

func (x *ACHIEVEMENT_INFO) GetTimestamp() uint32 {
	if x != nil && x.Timestamp != nil {
		return *x.Timestamp
	}
	return 0
}

var File_ACHIEVEMENT_INFO_proto protoreflect.FileDescriptor

var file_ACHIEVEMENT_INFO_proto_rawDesc = []byte{
	0x0a, 0x16, 0x41, 0x43, 0x48, 0x49, 0x45, 0x56, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x49, 0x4e,
	0x46, 0x4f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x62, 0x65, 0x6c, 0x66, 0x61, 0x73,
	0x74, 0x22, 0x5c, 0x0a, 0x10, 0x41, 0x43, 0x48, 0x49, 0x45, 0x56, 0x45, 0x4d, 0x45, 0x4e, 0x54,
	0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28,
	0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03,
	0x20, 0x02, 0x28, 0x0d, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42,
	0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
}

var (
	file_ACHIEVEMENT_INFO_proto_rawDescOnce sync.Once
	file_ACHIEVEMENT_INFO_proto_rawDescData = file_ACHIEVEMENT_INFO_proto_rawDesc
)

func file_ACHIEVEMENT_INFO_proto_rawDescGZIP() []byte {
	file_ACHIEVEMENT_INFO_proto_rawDescOnce.Do(func() {
		file_ACHIEVEMENT_INFO_proto_rawDescData = protoimpl.X.CompressGZIP(file_ACHIEVEMENT_INFO_proto_rawDescData)
	})
	return file_ACHIEVEMENT_INFO_proto_rawDescData
}

var file_ACHIEVEMENT_INFO_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ACHIEVEMENT_INFO_proto_goTypes = []any{
	(*ACHIEVEMENT_INFO)(nil), // 0: belfast.ACHIEVEMENT_INFO
}
var file_ACHIEVEMENT_INFO_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ACHIEVEMENT_INFO_proto_init() }
func file_ACHIEVEMENT_INFO_proto_init() {
	if File_ACHIEVEMENT_INFO_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ACHIEVEMENT_INFO_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ACHIEVEMENT_INFO); i {
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
			RawDescriptor: file_ACHIEVEMENT_INFO_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ACHIEVEMENT_INFO_proto_goTypes,
		DependencyIndexes: file_ACHIEVEMENT_INFO_proto_depIdxs,
		MessageInfos:      file_ACHIEVEMENT_INFO_proto_msgTypes,
	}.Build()
	File_ACHIEVEMENT_INFO_proto = out.File
	file_ACHIEVEMENT_INFO_proto_rawDesc = nil
	file_ACHIEVEMENT_INFO_proto_goTypes = nil
	file_ACHIEVEMENT_INFO_proto_depIdxs = nil
}
