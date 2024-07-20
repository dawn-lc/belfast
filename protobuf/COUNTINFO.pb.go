// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: COUNTINFO.proto

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

type COUNTINFO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StepCount      *uint32  `protobuf:"varint,1,req,name=step_count,json=stepCount" json:"step_count,omitempty"`
	TreasureCount  *uint32  `protobuf:"varint,2,req,name=treasure_count,json=treasureCount" json:"treasure_count,omitempty"`
	TaskProgress   *uint32  `protobuf:"varint,3,req,name=task_progress,json=taskProgress" json:"task_progress,omitempty"`
	ActivateCount  *uint32  `protobuf:"varint,4,req,name=activate_count,json=activateCount" json:"activate_count,omitempty"`
	CollectionList []uint32 `protobuf:"varint,5,rep,name=collection_list,json=collectionList" json:"collection_list,omitempty"`
}

func (x *COUNTINFO) Reset() {
	*x = COUNTINFO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_COUNTINFO_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *COUNTINFO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*COUNTINFO) ProtoMessage() {}

func (x *COUNTINFO) ProtoReflect() protoreflect.Message {
	mi := &file_COUNTINFO_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use COUNTINFO.ProtoReflect.Descriptor instead.
func (*COUNTINFO) Descriptor() ([]byte, []int) {
	return file_COUNTINFO_proto_rawDescGZIP(), []int{0}
}

func (x *COUNTINFO) GetStepCount() uint32 {
	if x != nil && x.StepCount != nil {
		return *x.StepCount
	}
	return 0
}

func (x *COUNTINFO) GetTreasureCount() uint32 {
	if x != nil && x.TreasureCount != nil {
		return *x.TreasureCount
	}
	return 0
}

func (x *COUNTINFO) GetTaskProgress() uint32 {
	if x != nil && x.TaskProgress != nil {
		return *x.TaskProgress
	}
	return 0
}

func (x *COUNTINFO) GetActivateCount() uint32 {
	if x != nil && x.ActivateCount != nil {
		return *x.ActivateCount
	}
	return 0
}

func (x *COUNTINFO) GetCollectionList() []uint32 {
	if x != nil {
		return x.CollectionList
	}
	return nil
}

var File_COUNTINFO_proto protoreflect.FileDescriptor

var file_COUNTINFO_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x49, 0x4e, 0x46, 0x4f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x07, 0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74, 0x22, 0xc6, 0x01, 0x0a, 0x09, 0x43,
	0x4f, 0x55, 0x4e, 0x54, 0x49, 0x4e, 0x46, 0x4f, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x65, 0x70,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x09, 0x73, 0x74,
	0x65, 0x70, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72, 0x65, 0x61, 0x73,
	0x75, 0x72, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0d, 0x52,
	0x0d, 0x74, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23,
	0x0a, 0x0d, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x03, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x0c, 0x74, 0x61, 0x73, 0x6b, 0x50, 0x72, 0x6f, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x0d, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x0d, 0x52, 0x0e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c,
	0x69, 0x73, 0x74, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66,
}

var (
	file_COUNTINFO_proto_rawDescOnce sync.Once
	file_COUNTINFO_proto_rawDescData = file_COUNTINFO_proto_rawDesc
)

func file_COUNTINFO_proto_rawDescGZIP() []byte {
	file_COUNTINFO_proto_rawDescOnce.Do(func() {
		file_COUNTINFO_proto_rawDescData = protoimpl.X.CompressGZIP(file_COUNTINFO_proto_rawDescData)
	})
	return file_COUNTINFO_proto_rawDescData
}

var file_COUNTINFO_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_COUNTINFO_proto_goTypes = []any{
	(*COUNTINFO)(nil), // 0: belfast.COUNTINFO
}
var file_COUNTINFO_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_COUNTINFO_proto_init() }
func file_COUNTINFO_proto_init() {
	if File_COUNTINFO_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_COUNTINFO_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*COUNTINFO); i {
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
			RawDescriptor: file_COUNTINFO_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_COUNTINFO_proto_goTypes,
		DependencyIndexes: file_COUNTINFO_proto_depIdxs,
		MessageInfos:      file_COUNTINFO_proto_msgTypes,
	}.Build()
	File_COUNTINFO_proto = out.File
	file_COUNTINFO_proto_rawDesc = nil
	file_COUNTINFO_proto_goTypes = nil
	file_COUNTINFO_proto_depIdxs = nil
}
