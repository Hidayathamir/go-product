// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: pkg/goproductgrpc/ping.proto

package goproductgrpc

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

type PingEmpty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PingEmpty) Reset() {
	*x = PingEmpty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_goproductgrpc_ping_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingEmpty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingEmpty) ProtoMessage() {}

func (x *PingEmpty) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_goproductgrpc_ping_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingEmpty.ProtoReflect.Descriptor instead.
func (*PingEmpty) Descriptor() ([]byte, []int) {
	return file_pkg_goproductgrpc_ping_proto_rawDescGZIP(), []int{0}
}

type ResPing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ResPing) Reset() {
	*x = ResPing{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_goproductgrpc_ping_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResPing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResPing) ProtoMessage() {}

func (x *ResPing) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_goproductgrpc_ping_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResPing.ProtoReflect.Descriptor instead.
func (*ResPing) Descriptor() ([]byte, []int) {
	return file_pkg_goproductgrpc_ping_proto_rawDescGZIP(), []int{1}
}

func (x *ResPing) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_pkg_goproductgrpc_ping_proto protoreflect.FileDescriptor

var file_pkg_goproductgrpc_ping_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d,
	0x67, 0x6f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x67, 0x72, 0x70, 0x63, 0x22, 0x0b, 0x0a,
	0x09, 0x50, 0x69, 0x6e, 0x67, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x23, 0x0a, 0x07, 0x52, 0x65,
	0x73, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32,
	0x42, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x3a, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12,
	0x18, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x50, 0x69, 0x6e, 0x67, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x50, 0x69, 0x6e,
	0x67, 0x22, 0x00, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x48, 0x69, 0x64, 0x61, 0x79, 0x61, 0x74, 0x68, 0x61, 0x6d, 0x69, 0x72, 0x2f, 0x67,
	0x6f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_goproductgrpc_ping_proto_rawDescOnce sync.Once
	file_pkg_goproductgrpc_ping_proto_rawDescData = file_pkg_goproductgrpc_ping_proto_rawDesc
)

func file_pkg_goproductgrpc_ping_proto_rawDescGZIP() []byte {
	file_pkg_goproductgrpc_ping_proto_rawDescOnce.Do(func() {
		file_pkg_goproductgrpc_ping_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_goproductgrpc_ping_proto_rawDescData)
	})
	return file_pkg_goproductgrpc_ping_proto_rawDescData
}

var file_pkg_goproductgrpc_ping_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_goproductgrpc_ping_proto_goTypes = []interface{}{
	(*PingEmpty)(nil), // 0: goproductgrpc.PingEmpty
	(*ResPing)(nil),   // 1: goproductgrpc.ResPing
}
var file_pkg_goproductgrpc_ping_proto_depIdxs = []int32{
	0, // 0: goproductgrpc.Ping.Ping:input_type -> goproductgrpc.PingEmpty
	1, // 1: goproductgrpc.Ping.Ping:output_type -> goproductgrpc.ResPing
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_goproductgrpc_ping_proto_init() }
func file_pkg_goproductgrpc_ping_proto_init() {
	if File_pkg_goproductgrpc_ping_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_goproductgrpc_ping_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingEmpty); i {
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
		file_pkg_goproductgrpc_ping_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResPing); i {
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
			RawDescriptor: file_pkg_goproductgrpc_ping_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_goproductgrpc_ping_proto_goTypes,
		DependencyIndexes: file_pkg_goproductgrpc_ping_proto_depIdxs,
		MessageInfos:      file_pkg_goproductgrpc_ping_proto_msgTypes,
	}.Build()
	File_pkg_goproductgrpc_ping_proto = out.File
	file_pkg_goproductgrpc_ping_proto_rawDesc = nil
	file_pkg_goproductgrpc_ping_proto_goTypes = nil
	file_pkg_goproductgrpc_ping_proto_depIdxs = nil
}
