// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: SC_10803_CN_JP_KR_TW.proto

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

type SC_10803_CN_JP_KR_TW struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GatewayIp   *string `protobuf:"bytes,1,req,name=gateway_ip,json=gatewayIp" json:"gateway_ip,omitempty"`
	GatewayPort *uint32 `protobuf:"varint,2,req,name=gateway_port,json=gatewayPort" json:"gateway_port,omitempty"`
	ProxyIp     *string `protobuf:"bytes,3,opt,name=proxy_ip,json=proxyIp" json:"proxy_ip,omitempty"`
	ProxyPort   *uint32 `protobuf:"varint,4,opt,name=proxy_port,json=proxyPort" json:"proxy_port,omitempty"`
}

func (x *SC_10803_CN_JP_KR_TW) Reset() {
	*x = SC_10803_CN_JP_KR_TW{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SC_10803_CN_JP_KR_TW_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SC_10803_CN_JP_KR_TW) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SC_10803_CN_JP_KR_TW) ProtoMessage() {}

func (x *SC_10803_CN_JP_KR_TW) ProtoReflect() protoreflect.Message {
	mi := &file_SC_10803_CN_JP_KR_TW_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SC_10803_CN_JP_KR_TW.ProtoReflect.Descriptor instead.
func (*SC_10803_CN_JP_KR_TW) Descriptor() ([]byte, []int) {
	return file_SC_10803_CN_JP_KR_TW_proto_rawDescGZIP(), []int{0}
}

func (x *SC_10803_CN_JP_KR_TW) GetGatewayIp() string {
	if x != nil && x.GatewayIp != nil {
		return *x.GatewayIp
	}
	return ""
}

func (x *SC_10803_CN_JP_KR_TW) GetGatewayPort() uint32 {
	if x != nil && x.GatewayPort != nil {
		return *x.GatewayPort
	}
	return 0
}

func (x *SC_10803_CN_JP_KR_TW) GetProxyIp() string {
	if x != nil && x.ProxyIp != nil {
		return *x.ProxyIp
	}
	return ""
}

func (x *SC_10803_CN_JP_KR_TW) GetProxyPort() uint32 {
	if x != nil && x.ProxyPort != nil {
		return *x.ProxyPort
	}
	return 0
}

var File_SC_10803_CN_JP_KR_TW_proto protoreflect.FileDescriptor

var file_SC_10803_CN_JP_KR_TW_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x53, 0x43, 0x5f, 0x31, 0x30, 0x38, 0x30, 0x33, 0x5f, 0x43, 0x4e, 0x5f, 0x4a, 0x50,
	0x5f, 0x4b, 0x52, 0x5f, 0x54, 0x57, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x62, 0x65,
	0x6c, 0x66, 0x61, 0x73, 0x74, 0x22, 0x92, 0x01, 0x0a, 0x14, 0x53, 0x43, 0x5f, 0x31, 0x30, 0x38,
	0x30, 0x33, 0x5f, 0x43, 0x4e, 0x5f, 0x4a, 0x50, 0x5f, 0x4b, 0x52, 0x5f, 0x54, 0x57, 0x12, 0x1d,
	0x0a, 0x0a, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x69, 0x70, 0x18, 0x01, 0x20, 0x02,
	0x28, 0x09, 0x52, 0x09, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x49, 0x70, 0x12, 0x21, 0x0a,
	0x0c, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20,
	0x02, 0x28, 0x0d, 0x52, 0x0b, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x50, 0x6f, 0x72, 0x74,
	0x12, 0x19, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f, 0x69, 0x70, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x49, 0x70, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x09, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x50, 0x6f, 0x72, 0x74, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
}

var (
	file_SC_10803_CN_JP_KR_TW_proto_rawDescOnce sync.Once
	file_SC_10803_CN_JP_KR_TW_proto_rawDescData = file_SC_10803_CN_JP_KR_TW_proto_rawDesc
)

func file_SC_10803_CN_JP_KR_TW_proto_rawDescGZIP() []byte {
	file_SC_10803_CN_JP_KR_TW_proto_rawDescOnce.Do(func() {
		file_SC_10803_CN_JP_KR_TW_proto_rawDescData = protoimpl.X.CompressGZIP(file_SC_10803_CN_JP_KR_TW_proto_rawDescData)
	})
	return file_SC_10803_CN_JP_KR_TW_proto_rawDescData
}

var file_SC_10803_CN_JP_KR_TW_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SC_10803_CN_JP_KR_TW_proto_goTypes = []any{
	(*SC_10803_CN_JP_KR_TW)(nil), // 0: belfast.SC_10803_CN_JP_KR_TW
}
var file_SC_10803_CN_JP_KR_TW_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_SC_10803_CN_JP_KR_TW_proto_init() }
func file_SC_10803_CN_JP_KR_TW_proto_init() {
	if File_SC_10803_CN_JP_KR_TW_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_SC_10803_CN_JP_KR_TW_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*SC_10803_CN_JP_KR_TW); i {
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
			RawDescriptor: file_SC_10803_CN_JP_KR_TW_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SC_10803_CN_JP_KR_TW_proto_goTypes,
		DependencyIndexes: file_SC_10803_CN_JP_KR_TW_proto_depIdxs,
		MessageInfos:      file_SC_10803_CN_JP_KR_TW_proto_msgTypes,
	}.Build()
	File_SC_10803_CN_JP_KR_TW_proto = out.File
	file_SC_10803_CN_JP_KR_TW_proto_rawDesc = nil
	file_SC_10803_CN_JP_KR_TW_proto_goTypes = nil
	file_SC_10803_CN_JP_KR_TW_proto_depIdxs = nil
}
