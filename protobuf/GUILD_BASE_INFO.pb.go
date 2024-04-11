// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.25.3
// source: GUILD_BASE_INFO.proto

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

type GUILD_BASE_INFO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              *uint32 `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Policy          *uint32 `protobuf:"varint,2,req,name=policy" json:"policy,omitempty"`
	Faction         *uint32 `protobuf:"varint,3,req,name=faction" json:"faction,omitempty"`
	Name            *string `protobuf:"bytes,4,req,name=name" json:"name,omitempty"`
	Level           *uint32 `protobuf:"varint,5,req,name=level" json:"level,omitempty"`
	Announce        *string `protobuf:"bytes,6,req,name=announce" json:"announce,omitempty"`
	Manifesto       *string `protobuf:"bytes,7,req,name=manifesto" json:"manifesto,omitempty"`
	Exp             *uint32 `protobuf:"varint,8,req,name=exp" json:"exp,omitempty"`
	MemberCount     *uint32 `protobuf:"varint,9,req,name=member_count,json=memberCount" json:"member_count,omitempty"`
	ChangeFactionCd *uint32 `protobuf:"varint,10,req,name=change_faction_cd,json=changeFactionCd" json:"change_faction_cd,omitempty"`
	KickLeaderCd    *uint32 `protobuf:"varint,11,req,name=kick_leader_cd,json=kickLeaderCd" json:"kick_leader_cd,omitempty"`
}

func (x *GUILD_BASE_INFO) Reset() {
	*x = GUILD_BASE_INFO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GUILD_BASE_INFO_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GUILD_BASE_INFO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GUILD_BASE_INFO) ProtoMessage() {}

func (x *GUILD_BASE_INFO) ProtoReflect() protoreflect.Message {
	mi := &file_GUILD_BASE_INFO_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GUILD_BASE_INFO.ProtoReflect.Descriptor instead.
func (*GUILD_BASE_INFO) Descriptor() ([]byte, []int) {
	return file_GUILD_BASE_INFO_proto_rawDescGZIP(), []int{0}
}

func (x *GUILD_BASE_INFO) GetId() uint32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *GUILD_BASE_INFO) GetPolicy() uint32 {
	if x != nil && x.Policy != nil {
		return *x.Policy
	}
	return 0
}

func (x *GUILD_BASE_INFO) GetFaction() uint32 {
	if x != nil && x.Faction != nil {
		return *x.Faction
	}
	return 0
}

func (x *GUILD_BASE_INFO) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *GUILD_BASE_INFO) GetLevel() uint32 {
	if x != nil && x.Level != nil {
		return *x.Level
	}
	return 0
}

func (x *GUILD_BASE_INFO) GetAnnounce() string {
	if x != nil && x.Announce != nil {
		return *x.Announce
	}
	return ""
}

func (x *GUILD_BASE_INFO) GetManifesto() string {
	if x != nil && x.Manifesto != nil {
		return *x.Manifesto
	}
	return ""
}

func (x *GUILD_BASE_INFO) GetExp() uint32 {
	if x != nil && x.Exp != nil {
		return *x.Exp
	}
	return 0
}

func (x *GUILD_BASE_INFO) GetMemberCount() uint32 {
	if x != nil && x.MemberCount != nil {
		return *x.MemberCount
	}
	return 0
}

func (x *GUILD_BASE_INFO) GetChangeFactionCd() uint32 {
	if x != nil && x.ChangeFactionCd != nil {
		return *x.ChangeFactionCd
	}
	return 0
}

func (x *GUILD_BASE_INFO) GetKickLeaderCd() uint32 {
	if x != nil && x.KickLeaderCd != nil {
		return *x.KickLeaderCd
	}
	return 0
}

var File_GUILD_BASE_INFO_proto protoreflect.FileDescriptor

var file_GUILD_BASE_INFO_proto_rawDesc = []byte{
	0x0a, 0x15, 0x47, 0x55, 0x49, 0x4c, 0x44, 0x5f, 0x42, 0x41, 0x53, 0x45, 0x5f, 0x49, 0x4e, 0x46,
	0x4f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74,
	0x22, 0xbe, 0x02, 0x0a, 0x0f, 0x47, 0x55, 0x49, 0x4c, 0x44, 0x5f, 0x42, 0x41, 0x53, 0x45, 0x5f,
	0x49, 0x4e, 0x46, 0x4f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x02,
	0x20, 0x02, 0x28, 0x0d, 0x52, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x18, 0x0a, 0x07,
	0x66, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x07, 0x66,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x02, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x12, 0x1a, 0x0a, 0x08, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x18, 0x06, 0x20, 0x02,
	0x28, 0x09, 0x52, 0x08, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x6d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x6f, 0x18, 0x07, 0x20, 0x02, 0x28, 0x09, 0x52,
	0x09, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x78,
	0x70, 0x18, 0x08, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x03, 0x65, 0x78, 0x70, 0x12, 0x21, 0x0a, 0x0c,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x02,
	0x28, 0x0d, 0x52, 0x0b, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x2a, 0x0a, 0x11, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x66, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x63, 0x64, 0x18, 0x0a, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x0f, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x46, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x6b,
	0x69, 0x63, 0x6b, 0x5f, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x63, 0x64, 0x18, 0x0b, 0x20,
	0x02, 0x28, 0x0d, 0x52, 0x0c, 0x6b, 0x69, 0x63, 0x6b, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x43,
	0x64, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
}

var (
	file_GUILD_BASE_INFO_proto_rawDescOnce sync.Once
	file_GUILD_BASE_INFO_proto_rawDescData = file_GUILD_BASE_INFO_proto_rawDesc
)

func file_GUILD_BASE_INFO_proto_rawDescGZIP() []byte {
	file_GUILD_BASE_INFO_proto_rawDescOnce.Do(func() {
		file_GUILD_BASE_INFO_proto_rawDescData = protoimpl.X.CompressGZIP(file_GUILD_BASE_INFO_proto_rawDescData)
	})
	return file_GUILD_BASE_INFO_proto_rawDescData
}

var file_GUILD_BASE_INFO_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GUILD_BASE_INFO_proto_goTypes = []interface{}{
	(*GUILD_BASE_INFO)(nil), // 0: belfast.GUILD_BASE_INFO
}
var file_GUILD_BASE_INFO_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_GUILD_BASE_INFO_proto_init() }
func file_GUILD_BASE_INFO_proto_init() {
	if File_GUILD_BASE_INFO_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_GUILD_BASE_INFO_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GUILD_BASE_INFO); i {
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
			RawDescriptor: file_GUILD_BASE_INFO_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GUILD_BASE_INFO_proto_goTypes,
		DependencyIndexes: file_GUILD_BASE_INFO_proto_depIdxs,
		MessageInfos:      file_GUILD_BASE_INFO_proto_msgTypes,
	}.Build()
	File_GUILD_BASE_INFO_proto = out.File
	file_GUILD_BASE_INFO_proto_rawDesc = nil
	file_GUILD_BASE_INFO_proto_goTypes = nil
	file_GUILD_BASE_INFO_proto_depIdxs = nil
}