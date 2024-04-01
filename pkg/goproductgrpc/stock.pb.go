// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: pkg/goproductgrpc/stock.proto

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

type StockVoid struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StockVoid) Reset() {
	*x = StockVoid{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_goproductgrpc_stock_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StockVoid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StockVoid) ProtoMessage() {}

func (x *StockVoid) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_goproductgrpc_stock_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StockVoid.ProtoReflect.Descriptor instead.
func (*StockVoid) Descriptor() ([]byte, []int) {
	return file_pkg_goproductgrpc_stock_proto_rawDescGZIP(), []int{0}
}

type ReqIncrementStock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId int64 `protobuf:"varint,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
}

func (x *ReqIncrementStock) Reset() {
	*x = ReqIncrementStock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_goproductgrpc_stock_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqIncrementStock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqIncrementStock) ProtoMessage() {}

func (x *ReqIncrementStock) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_goproductgrpc_stock_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqIncrementStock.ProtoReflect.Descriptor instead.
func (*ReqIncrementStock) Descriptor() ([]byte, []int) {
	return file_pkg_goproductgrpc_stock_proto_rawDescGZIP(), []int{1}
}

func (x *ReqIncrementStock) GetProductId() int64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

type ReqDecrementStock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId int64 `protobuf:"varint,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
}

func (x *ReqDecrementStock) Reset() {
	*x = ReqDecrementStock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_goproductgrpc_stock_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqDecrementStock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqDecrementStock) ProtoMessage() {}

func (x *ReqDecrementStock) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_goproductgrpc_stock_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqDecrementStock.ProtoReflect.Descriptor instead.
func (*ReqDecrementStock) Descriptor() ([]byte, []int) {
	return file_pkg_goproductgrpc_stock_proto_rawDescGZIP(), []int{2}
}

func (x *ReqDecrementStock) GetProductId() int64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

var File_pkg_goproductgrpc_stock_proto protoreflect.FileDescriptor

var file_pkg_goproductgrpc_stock_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0d, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x67, 0x72, 0x70, 0x63, 0x22, 0x0b,
	0x0a, 0x09, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x32, 0x0a, 0x11, 0x52,
	0x65, 0x71, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x22,
	0x32, 0x0a, 0x11, 0x52, 0x65, 0x71, 0x44, 0x65, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53,
	0x74, 0x6f, 0x63, 0x6b, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x49, 0x64, 0x32, 0xa7, 0x01, 0x0a, 0x05, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x4e, 0x0a,
	0x0e, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12,
	0x20, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x52, 0x65, 0x71, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x6f, 0x63,
	0x6b, 0x1a, 0x18, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x4e, 0x0a,
	0x0e, 0x44, 0x65, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12,
	0x20, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x52, 0x65, 0x71, 0x44, 0x65, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x6f, 0x63,
	0x6b, 0x1a, 0x18, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x42, 0x32, 0x5a,
	0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x48, 0x69, 0x64, 0x61,
	0x79, 0x61, 0x74, 0x68, 0x61, 0x6d, 0x69, 0x72, 0x2f, 0x67, 0x6f, 0x75, 0x73, 0x65, 0x72, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x67, 0x72, 0x70,
	0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_goproductgrpc_stock_proto_rawDescOnce sync.Once
	file_pkg_goproductgrpc_stock_proto_rawDescData = file_pkg_goproductgrpc_stock_proto_rawDesc
)

func file_pkg_goproductgrpc_stock_proto_rawDescGZIP() []byte {
	file_pkg_goproductgrpc_stock_proto_rawDescOnce.Do(func() {
		file_pkg_goproductgrpc_stock_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_goproductgrpc_stock_proto_rawDescData)
	})
	return file_pkg_goproductgrpc_stock_proto_rawDescData
}

var file_pkg_goproductgrpc_stock_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pkg_goproductgrpc_stock_proto_goTypes = []interface{}{
	(*StockVoid)(nil),         // 0: goproductgrpc.StockVoid
	(*ReqIncrementStock)(nil), // 1: goproductgrpc.ReqIncrementStock
	(*ReqDecrementStock)(nil), // 2: goproductgrpc.ReqDecrementStock
}
var file_pkg_goproductgrpc_stock_proto_depIdxs = []int32{
	1, // 0: goproductgrpc.Stock.IncrementStock:input_type -> goproductgrpc.ReqIncrementStock
	2, // 1: goproductgrpc.Stock.DecrementStock:input_type -> goproductgrpc.ReqDecrementStock
	0, // 2: goproductgrpc.Stock.IncrementStock:output_type -> goproductgrpc.StockVoid
	0, // 3: goproductgrpc.Stock.DecrementStock:output_type -> goproductgrpc.StockVoid
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_goproductgrpc_stock_proto_init() }
func file_pkg_goproductgrpc_stock_proto_init() {
	if File_pkg_goproductgrpc_stock_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_goproductgrpc_stock_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StockVoid); i {
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
		file_pkg_goproductgrpc_stock_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqIncrementStock); i {
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
		file_pkg_goproductgrpc_stock_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqDecrementStock); i {
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
			RawDescriptor: file_pkg_goproductgrpc_stock_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_goproductgrpc_stock_proto_goTypes,
		DependencyIndexes: file_pkg_goproductgrpc_stock_proto_depIdxs,
		MessageInfos:      file_pkg_goproductgrpc_stock_proto_msgTypes,
	}.Build()
	File_pkg_goproductgrpc_stock_proto = out.File
	file_pkg_goproductgrpc_stock_proto_rawDesc = nil
	file_pkg_goproductgrpc_stock_proto_goTypes = nil
	file_pkg_goproductgrpc_stock_proto_depIdxs = nil
}
