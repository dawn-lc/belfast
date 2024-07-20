// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: BLUPRINTINFO.proto

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

type BLUPRINTINFO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             *uint32 `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	ShipId         *uint32 `protobuf:"varint,2,req,name=ship_id,json=shipId" json:"ship_id,omitempty"`
	StartTime      *uint32 `protobuf:"varint,3,req,name=start_time,json=startTime" json:"start_time,omitempty"`
	BluePrintLevel *uint32 `protobuf:"varint,4,req,name=blue_print_level,json=bluePrintLevel" json:"blue_print_level,omitempty"`
	Exp            *uint32 `protobuf:"varint,5,req,name=exp" json:"exp,omitempty"`
	StartDuration  *uint32 `protobuf:"varint,6,opt,name=start_duration,json=startDuration" json:"start_duration,omitempty"`
}

func (x *BLUPRINTINFO) Reset() {
	*x = BLUPRINTINFO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BLUPRINTINFO_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BLUPRINTINFO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BLUPRINTINFO) ProtoMessage() {}

func (x *BLUPRINTINFO) ProtoReflect() protoreflect.Message {
	mi := &file_BLUPRINTINFO_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BLUPRINTINFO.ProtoReflect.Descriptor instead.
func (*BLUPRINTINFO) Descriptor() ([]byte, []int) {
	return file_BLUPRINTINFO_proto_rawDescGZIP(), []int{0}
}

func (x *BLUPRINTINFO) GetId() uint32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *BLUPRINTINFO) GetShipId() uint32 {
	if x != nil && x.ShipId != nil {
		return *x.ShipId
	}
	return 0
}

func (x *BLUPRINTINFO) GetStartTime() uint32 {
	if x != nil && x.StartTime != nil {
		return *x.StartTime
	}
	return 0
}

func (x *BLUPRINTINFO) GetBluePrintLevel() uint32 {
	if x != nil && x.BluePrintLevel != nil {
		return *x.BluePrintLevel
	}
	return 0
}

func (x *BLUPRINTINFO) GetExp() uint32 {
	if x != nil && x.Exp != nil {
		return *x.Exp
	}
	return 0
}

func (x *BLUPRINTINFO) GetStartDuration() uint32 {
	if x != nil && x.StartDuration != nil {
		return *x.StartDuration
	}
	return 0
}

var File_BLUPRINTINFO_proto protoreflect.FileDescriptor

var file_BLUPRINTINFO_proto_rawDesc = []byte{
	0x0a, 0x12, 0x42, 0x4c, 0x55, 0x50, 0x52, 0x49, 0x4e, 0x54, 0x49, 0x4e, 0x46, 0x4f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74, 0x22, 0xb9, 0x01,
	0x0a, 0x0c, 0x42, 0x4c, 0x55, 0x50, 0x52, 0x49, 0x4e, 0x54, 0x49, 0x4e, 0x46, 0x4f, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17,
	0x0a, 0x07, 0x73, 0x68, 0x69, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0d, 0x52,
	0x06, 0x73, 0x68, 0x69, 0x70, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x09, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x62, 0x6c, 0x75, 0x65, 0x5f, 0x70,
	0x72, 0x69, 0x6e, 0x74, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x02, 0x28, 0x0d,
	0x52, 0x0e, 0x62, 0x6c, 0x75, 0x65, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x12, 0x10, 0x0a, 0x03, 0x65, 0x78, 0x70, 0x18, 0x05, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x03, 0x65,
	0x78, 0x70, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
}

var (
	file_BLUPRINTINFO_proto_rawDescOnce sync.Once
	file_BLUPRINTINFO_proto_rawDescData = file_BLUPRINTINFO_proto_rawDesc
)

func file_BLUPRINTINFO_proto_rawDescGZIP() []byte {
	file_BLUPRINTINFO_proto_rawDescOnce.Do(func() {
		file_BLUPRINTINFO_proto_rawDescData = protoimpl.X.CompressGZIP(file_BLUPRINTINFO_proto_rawDescData)
	})
	return file_BLUPRINTINFO_proto_rawDescData
}

var file_BLUPRINTINFO_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_BLUPRINTINFO_proto_goTypes = []any{
	(*BLUPRINTINFO)(nil), // 0: belfast.BLUPRINTINFO
}
var file_BLUPRINTINFO_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_BLUPRINTINFO_proto_init() }
func file_BLUPRINTINFO_proto_init() {
	if File_BLUPRINTINFO_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_BLUPRINTINFO_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*BLUPRINTINFO); i {
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
			RawDescriptor: file_BLUPRINTINFO_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_BLUPRINTINFO_proto_goTypes,
		DependencyIndexes: file_BLUPRINTINFO_proto_depIdxs,
		MessageInfos:      file_BLUPRINTINFO_proto_msgTypes,
	}.Build()
	File_BLUPRINTINFO_proto = out.File
	file_BLUPRINTINFO_proto_rawDesc = nil
	file_BLUPRINTINFO_proto_goTypes = nil
	file_BLUPRINTINFO_proto_depIdxs = nil
}
