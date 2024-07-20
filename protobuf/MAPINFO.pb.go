// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: MAPINFO.proto

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

type MAPINFO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        *WORLDMAPID        `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
	CellList  []*CHAPTERCELLINFO `protobuf:"bytes,2,rep,name=cell_list,json=cellList" json:"cell_list,omitempty"`
	StateFlag []uint32           `protobuf:"varint,3,rep,name=state_flag,json=stateFlag" json:"state_flag,omitempty"`
	LandList  []*LANDINFO        `protobuf:"bytes,4,rep,name=land_list,json=landList" json:"land_list,omitempty"`
	PosList   []*WORLDPOSINFO    `protobuf:"bytes,5,rep,name=pos_list,json=posList" json:"pos_list,omitempty"`
}

func (x *MAPINFO) Reset() {
	*x = MAPINFO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MAPINFO_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MAPINFO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MAPINFO) ProtoMessage() {}

func (x *MAPINFO) ProtoReflect() protoreflect.Message {
	mi := &file_MAPINFO_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MAPINFO.ProtoReflect.Descriptor instead.
func (*MAPINFO) Descriptor() ([]byte, []int) {
	return file_MAPINFO_proto_rawDescGZIP(), []int{0}
}

func (x *MAPINFO) GetId() *WORLDMAPID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *MAPINFO) GetCellList() []*CHAPTERCELLINFO {
	if x != nil {
		return x.CellList
	}
	return nil
}

func (x *MAPINFO) GetStateFlag() []uint32 {
	if x != nil {
		return x.StateFlag
	}
	return nil
}

func (x *MAPINFO) GetLandList() []*LANDINFO {
	if x != nil {
		return x.LandList
	}
	return nil
}

func (x *MAPINFO) GetPosList() []*WORLDPOSINFO {
	if x != nil {
		return x.PosList
	}
	return nil
}

var File_MAPINFO_proto protoreflect.FileDescriptor

var file_MAPINFO_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x4d, 0x41, 0x50, 0x49, 0x4e, 0x46, 0x4f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74, 0x1a, 0x10, 0x57, 0x4f, 0x52, 0x4c, 0x44, 0x4d,
	0x41, 0x50, 0x49, 0x44, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x43, 0x48, 0x41, 0x50,
	0x54, 0x45, 0x52, 0x43, 0x45, 0x4c, 0x4c, 0x49, 0x4e, 0x46, 0x4f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0e, 0x4c, 0x41, 0x4e, 0x44, 0x49, 0x4e, 0x46, 0x4f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x12, 0x57, 0x4f, 0x52, 0x4c, 0x44, 0x50, 0x4f, 0x53, 0x49, 0x4e, 0x46, 0x4f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe6, 0x01, 0x0a, 0x07, 0x4d, 0x41, 0x50, 0x49, 0x4e, 0x46,
	0x4f, 0x12, 0x23, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74, 0x2e, 0x57, 0x4f, 0x52, 0x4c, 0x44, 0x4d, 0x41, 0x50,
	0x49, 0x44, 0x52, 0x02, 0x69, 0x64, 0x12, 0x35, 0x0a, 0x09, 0x63, 0x65, 0x6c, 0x6c, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x62, 0x65, 0x6c, 0x66,
	0x61, 0x73, 0x74, 0x2e, 0x43, 0x48, 0x41, 0x50, 0x54, 0x45, 0x52, 0x43, 0x45, 0x4c, 0x4c, 0x49,
	0x4e, 0x46, 0x4f, 0x52, 0x08, 0x63, 0x65, 0x6c, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0d, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x65, 0x46, 0x6c, 0x61, 0x67, 0x12, 0x2e, 0x0a, 0x09,
	0x6c, 0x61, 0x6e, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74, 0x2e, 0x4c, 0x41, 0x4e, 0x44, 0x49, 0x4e,
	0x46, 0x4f, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x08,
	0x70, 0x6f, 0x73, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74, 0x2e, 0x57, 0x4f, 0x52, 0x4c, 0x44, 0x50, 0x4f,
	0x53, 0x49, 0x4e, 0x46, 0x4f, 0x52, 0x07, 0x70, 0x6f, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x0c,
	0x5a, 0x0a, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
}

var (
	file_MAPINFO_proto_rawDescOnce sync.Once
	file_MAPINFO_proto_rawDescData = file_MAPINFO_proto_rawDesc
)

func file_MAPINFO_proto_rawDescGZIP() []byte {
	file_MAPINFO_proto_rawDescOnce.Do(func() {
		file_MAPINFO_proto_rawDescData = protoimpl.X.CompressGZIP(file_MAPINFO_proto_rawDescData)
	})
	return file_MAPINFO_proto_rawDescData
}

var file_MAPINFO_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_MAPINFO_proto_goTypes = []any{
	(*MAPINFO)(nil),         // 0: belfast.MAPINFO
	(*WORLDMAPID)(nil),      // 1: belfast.WORLDMAPID
	(*CHAPTERCELLINFO)(nil), // 2: belfast.CHAPTERCELLINFO
	(*LANDINFO)(nil),        // 3: belfast.LANDINFO
	(*WORLDPOSINFO)(nil),    // 4: belfast.WORLDPOSINFO
}
var file_MAPINFO_proto_depIdxs = []int32{
	1, // 0: belfast.MAPINFO.id:type_name -> belfast.WORLDMAPID
	2, // 1: belfast.MAPINFO.cell_list:type_name -> belfast.CHAPTERCELLINFO
	3, // 2: belfast.MAPINFO.land_list:type_name -> belfast.LANDINFO
	4, // 3: belfast.MAPINFO.pos_list:type_name -> belfast.WORLDPOSINFO
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_MAPINFO_proto_init() }
func file_MAPINFO_proto_init() {
	if File_MAPINFO_proto != nil {
		return
	}
	file_WORLDMAPID_proto_init()
	file_CHAPTERCELLINFO_proto_init()
	file_LANDINFO_proto_init()
	file_WORLDPOSINFO_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_MAPINFO_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*MAPINFO); i {
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
			RawDescriptor: file_MAPINFO_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MAPINFO_proto_goTypes,
		DependencyIndexes: file_MAPINFO_proto_depIdxs,
		MessageInfos:      file_MAPINFO_proto_msgTypes,
	}.Build()
	File_MAPINFO_proto = out.File
	file_MAPINFO_proto_rawDesc = nil
	file_MAPINFO_proto_goTypes = nil
	file_MAPINFO_proto_depIdxs = nil
}
