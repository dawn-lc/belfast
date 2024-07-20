// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: SC_26022.proto

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

type SC_26022 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result          *uint32 `protobuf:"varint,1,req,name=result" json:"result,omitempty"`
	RegisterDate    *uint32 `protobuf:"varint,2,req,name=register_date,json=registerDate" json:"register_date,omitempty"`
	GuildName       *string `protobuf:"bytes,3,req,name=guild_name,json=guildName" json:"guild_name,omitempty"`
	ChapterId       *uint32 `protobuf:"varint,4,req,name=chapter_id,json=chapterId" json:"chapter_id,omitempty"`
	MarryNumber     *uint32 `protobuf:"varint,5,req,name=marry_number,json=marryNumber" json:"marry_number,omitempty"`
	MedalNumber     *uint32 `protobuf:"varint,6,req,name=medal_number,json=medalNumber" json:"medal_number,omitempty"`
	FurnitureNumber *uint32 `protobuf:"varint,7,req,name=furniture_number,json=furnitureNumber" json:"furniture_number,omitempty"`
	FurnitureWorth  *uint32 `protobuf:"varint,8,req,name=furniture_worth,json=furnitureWorth" json:"furniture_worth,omitempty"`
	CharacterId     *uint32 `protobuf:"varint,9,req,name=character_id,json=characterId" json:"character_id,omitempty"`
	FirstLadyId     *uint32 `protobuf:"varint,10,req,name=first_lady_id,json=firstLadyId" json:"first_lady_id,omitempty"`
	FirstLadyName   *string `protobuf:"bytes,11,req,name=first_lady_name,json=firstLadyName" json:"first_lady_name,omitempty"`
	FirstLadyTime   *uint32 `protobuf:"varint,12,req,name=first_lady_time,json=firstLadyTime" json:"first_lady_time,omitempty"`
	FirstOnline     *uint32 `protobuf:"varint,13,req,name=first_online,json=firstOnline" json:"first_online,omitempty"`
	WorldMaxTask    *uint32 `protobuf:"varint,14,req,name=world_max_task,json=worldMaxTask" json:"world_max_task,omitempty"`
	CollectNum      *uint32 `protobuf:"varint,15,req,name=collect_num,json=collectNum" json:"collect_num,omitempty"`
	Combat          *uint32 `protobuf:"varint,16,req,name=combat" json:"combat,omitempty"`
	ShipNumTotal    *uint32 `protobuf:"varint,17,req,name=ship_num_total,json=shipNumTotal" json:"ship_num_total,omitempty"`
	ShipNum_120     *uint32 `protobuf:"varint,18,req,name=ship_num_120,json=shipNum120" json:"ship_num_120,omitempty"`
	ShipNum_125     *uint32 `protobuf:"varint,19,req,name=ship_num_125,json=shipNum125" json:"ship_num_125,omitempty"`
	Love200Num      *uint32 `protobuf:"varint,20,req,name=love200_num,json=love200Num" json:"love200_num,omitempty"`
	SkinNum         *uint32 `protobuf:"varint,21,req,name=skin_num,json=skinNum" json:"skin_num,omitempty"`
	SkinShipNum     *uint32 `protobuf:"varint,22,req,name=skin_ship_num,json=skinShipNum" json:"skin_ship_num,omitempty"`
}

func (x *SC_26022) Reset() {
	*x = SC_26022{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SC_26022_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SC_26022) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SC_26022) ProtoMessage() {}

func (x *SC_26022) ProtoReflect() protoreflect.Message {
	mi := &file_SC_26022_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SC_26022.ProtoReflect.Descriptor instead.
func (*SC_26022) Descriptor() ([]byte, []int) {
	return file_SC_26022_proto_rawDescGZIP(), []int{0}
}

func (x *SC_26022) GetResult() uint32 {
	if x != nil && x.Result != nil {
		return *x.Result
	}
	return 0
}

func (x *SC_26022) GetRegisterDate() uint32 {
	if x != nil && x.RegisterDate != nil {
		return *x.RegisterDate
	}
	return 0
}

func (x *SC_26022) GetGuildName() string {
	if x != nil && x.GuildName != nil {
		return *x.GuildName
	}
	return ""
}

func (x *SC_26022) GetChapterId() uint32 {
	if x != nil && x.ChapterId != nil {
		return *x.ChapterId
	}
	return 0
}

func (x *SC_26022) GetMarryNumber() uint32 {
	if x != nil && x.MarryNumber != nil {
		return *x.MarryNumber
	}
	return 0
}

func (x *SC_26022) GetMedalNumber() uint32 {
	if x != nil && x.MedalNumber != nil {
		return *x.MedalNumber
	}
	return 0
}

func (x *SC_26022) GetFurnitureNumber() uint32 {
	if x != nil && x.FurnitureNumber != nil {
		return *x.FurnitureNumber
	}
	return 0
}

func (x *SC_26022) GetFurnitureWorth() uint32 {
	if x != nil && x.FurnitureWorth != nil {
		return *x.FurnitureWorth
	}
	return 0
}

func (x *SC_26022) GetCharacterId() uint32 {
	if x != nil && x.CharacterId != nil {
		return *x.CharacterId
	}
	return 0
}

func (x *SC_26022) GetFirstLadyId() uint32 {
	if x != nil && x.FirstLadyId != nil {
		return *x.FirstLadyId
	}
	return 0
}

func (x *SC_26022) GetFirstLadyName() string {
	if x != nil && x.FirstLadyName != nil {
		return *x.FirstLadyName
	}
	return ""
}

func (x *SC_26022) GetFirstLadyTime() uint32 {
	if x != nil && x.FirstLadyTime != nil {
		return *x.FirstLadyTime
	}
	return 0
}

func (x *SC_26022) GetFirstOnline() uint32 {
	if x != nil && x.FirstOnline != nil {
		return *x.FirstOnline
	}
	return 0
}

func (x *SC_26022) GetWorldMaxTask() uint32 {
	if x != nil && x.WorldMaxTask != nil {
		return *x.WorldMaxTask
	}
	return 0
}

func (x *SC_26022) GetCollectNum() uint32 {
	if x != nil && x.CollectNum != nil {
		return *x.CollectNum
	}
	return 0
}

func (x *SC_26022) GetCombat() uint32 {
	if x != nil && x.Combat != nil {
		return *x.Combat
	}
	return 0
}

func (x *SC_26022) GetShipNumTotal() uint32 {
	if x != nil && x.ShipNumTotal != nil {
		return *x.ShipNumTotal
	}
	return 0
}

func (x *SC_26022) GetShipNum_120() uint32 {
	if x != nil && x.ShipNum_120 != nil {
		return *x.ShipNum_120
	}
	return 0
}

func (x *SC_26022) GetShipNum_125() uint32 {
	if x != nil && x.ShipNum_125 != nil {
		return *x.ShipNum_125
	}
	return 0
}

func (x *SC_26022) GetLove200Num() uint32 {
	if x != nil && x.Love200Num != nil {
		return *x.Love200Num
	}
	return 0
}

func (x *SC_26022) GetSkinNum() uint32 {
	if x != nil && x.SkinNum != nil {
		return *x.SkinNum
	}
	return 0
}

func (x *SC_26022) GetSkinShipNum() uint32 {
	if x != nil && x.SkinShipNum != nil {
		return *x.SkinShipNum
	}
	return 0
}

var File_SC_26022_proto protoreflect.FileDescriptor

var file_SC_26022_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x53, 0x43, 0x5f, 0x32, 0x36, 0x30, 0x32, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74, 0x22, 0x82, 0x06, 0x0a, 0x08, 0x53, 0x43,
	0x5f, 0x32, 0x36, 0x30, 0x32, 0x32, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x23,
	0x0a, 0x0d, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x02, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x02, 0x28, 0x09, 0x52, 0x09, 0x67, 0x75, 0x69, 0x6c, 0x64, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x09, 0x63, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x61, 0x72, 0x72, 0x79, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x05, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x0b, 0x6d, 0x61, 0x72, 0x72, 0x79, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x64, 0x61, 0x6c, 0x5f, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x06, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x0b, 0x6d, 0x65, 0x64, 0x61,
	0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x10, 0x66, 0x75, 0x72, 0x6e, 0x69,
	0x74, 0x75, 0x72, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x07, 0x20, 0x02, 0x28,
	0x0d, 0x52, 0x0f, 0x66, 0x75, 0x72, 0x6e, 0x69, 0x74, 0x75, 0x72, 0x65, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x27, 0x0a, 0x0f, 0x66, 0x75, 0x72, 0x6e, 0x69, 0x74, 0x75, 0x72, 0x65, 0x5f,
	0x77, 0x6f, 0x72, 0x74, 0x68, 0x18, 0x08, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x0e, 0x66, 0x75, 0x72,
	0x6e, 0x69, 0x74, 0x75, 0x72, 0x65, 0x57, 0x6f, 0x72, 0x74, 0x68, 0x12, 0x21, 0x0a, 0x0c, 0x63,
	0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x02, 0x28,
	0x0d, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x22,
	0x0a, 0x0d, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6c, 0x61, 0x64, 0x79, 0x5f, 0x69, 0x64, 0x18,
	0x0a, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x0b, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4c, 0x61, 0x64, 0x79,
	0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6c, 0x61, 0x64, 0x79,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x02, 0x28, 0x09, 0x52, 0x0d, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x4c, 0x61, 0x64, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x5f, 0x6c, 0x61, 0x64, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20,
	0x02, 0x28, 0x0d, 0x52, 0x0d, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4c, 0x61, 0x64, 0x79, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6f, 0x6e, 0x6c, 0x69,
	0x6e, 0x65, 0x18, 0x0d, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x0b, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4f,
	0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x6d,
	0x61, 0x78, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x0e, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x0c, 0x77,
	0x6f, 0x72, 0x6c, 0x64, 0x4d, 0x61, 0x78, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x1f, 0x0a, 0x0b, 0x63,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x0f, 0x20, 0x02, 0x28, 0x0d,
	0x52, 0x0a, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x4e, 0x75, 0x6d, 0x12, 0x16, 0x0a, 0x06,
	0x63, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x18, 0x10, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x06, 0x63, 0x6f,
	0x6d, 0x62, 0x61, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x73, 0x68, 0x69, 0x70, 0x5f, 0x6e, 0x75, 0x6d,
	0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x11, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x0c, 0x73, 0x68,
	0x69, 0x70, 0x4e, 0x75, 0x6d, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x20, 0x0a, 0x0c, 0x73, 0x68,
	0x69, 0x70, 0x5f, 0x6e, 0x75, 0x6d, 0x5f, 0x31, 0x32, 0x30, 0x18, 0x12, 0x20, 0x02, 0x28, 0x0d,
	0x52, 0x0a, 0x73, 0x68, 0x69, 0x70, 0x4e, 0x75, 0x6d, 0x31, 0x32, 0x30, 0x12, 0x20, 0x0a, 0x0c,
	0x73, 0x68, 0x69, 0x70, 0x5f, 0x6e, 0x75, 0x6d, 0x5f, 0x31, 0x32, 0x35, 0x18, 0x13, 0x20, 0x02,
	0x28, 0x0d, 0x52, 0x0a, 0x73, 0x68, 0x69, 0x70, 0x4e, 0x75, 0x6d, 0x31, 0x32, 0x35, 0x12, 0x1f,
	0x0a, 0x0b, 0x6c, 0x6f, 0x76, 0x65, 0x32, 0x30, 0x30, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x14, 0x20,
	0x02, 0x28, 0x0d, 0x52, 0x0a, 0x6c, 0x6f, 0x76, 0x65, 0x32, 0x30, 0x30, 0x4e, 0x75, 0x6d, 0x12,
	0x19, 0x0a, 0x08, 0x73, 0x6b, 0x69, 0x6e, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x15, 0x20, 0x02, 0x28,
	0x0d, 0x52, 0x07, 0x73, 0x6b, 0x69, 0x6e, 0x4e, 0x75, 0x6d, 0x12, 0x22, 0x0a, 0x0d, 0x73, 0x6b,
	0x69, 0x6e, 0x5f, 0x73, 0x68, 0x69, 0x70, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x16, 0x20, 0x02, 0x28,
	0x0d, 0x52, 0x0b, 0x73, 0x6b, 0x69, 0x6e, 0x53, 0x68, 0x69, 0x70, 0x4e, 0x75, 0x6d, 0x42, 0x0c,
	0x5a, 0x0a, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
}

var (
	file_SC_26022_proto_rawDescOnce sync.Once
	file_SC_26022_proto_rawDescData = file_SC_26022_proto_rawDesc
)

func file_SC_26022_proto_rawDescGZIP() []byte {
	file_SC_26022_proto_rawDescOnce.Do(func() {
		file_SC_26022_proto_rawDescData = protoimpl.X.CompressGZIP(file_SC_26022_proto_rawDescData)
	})
	return file_SC_26022_proto_rawDescData
}

var file_SC_26022_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SC_26022_proto_goTypes = []any{
	(*SC_26022)(nil), // 0: belfast.SC_26022
}
var file_SC_26022_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_SC_26022_proto_init() }
func file_SC_26022_proto_init() {
	if File_SC_26022_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_SC_26022_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*SC_26022); i {
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
			RawDescriptor: file_SC_26022_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SC_26022_proto_goTypes,
		DependencyIndexes: file_SC_26022_proto_depIdxs,
		MessageInfos:      file_SC_26022_proto_msgTypes,
	}.Build()
	File_SC_26022_proto = out.File
	file_SC_26022_proto_rawDesc = nil
	file_SC_26022_proto_goTypes = nil
	file_SC_26022_proto_depIdxs = nil
}
