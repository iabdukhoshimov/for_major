// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: barcode.proto

package object_builder_service

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

type BarcodeGenerateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TableSlug string `protobuf:"bytes,1,opt,name=table_slug,json=tableSlug,proto3" json:"table_slug,omitempty"`
}

func (x *BarcodeGenerateReq) Reset() {
	*x = BarcodeGenerateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barcode_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BarcodeGenerateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BarcodeGenerateReq) ProtoMessage() {}

func (x *BarcodeGenerateReq) ProtoReflect() protoreflect.Message {
	mi := &file_barcode_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BarcodeGenerateReq.ProtoReflect.Descriptor instead.
func (*BarcodeGenerateReq) Descriptor() ([]byte, []int) {
	return file_barcode_proto_rawDescGZIP(), []int{0}
}

func (x *BarcodeGenerateReq) GetTableSlug() string {
	if x != nil {
		return x.TableSlug
	}
	return ""
}

type BarcodeGenerateRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number string `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *BarcodeGenerateRes) Reset() {
	*x = BarcodeGenerateRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barcode_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BarcodeGenerateRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BarcodeGenerateRes) ProtoMessage() {}

func (x *BarcodeGenerateRes) ProtoReflect() protoreflect.Message {
	mi := &file_barcode_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BarcodeGenerateRes.ProtoReflect.Descriptor instead.
func (*BarcodeGenerateRes) Descriptor() ([]byte, []int) {
	return file_barcode_proto_rawDescGZIP(), []int{1}
}

func (x *BarcodeGenerateRes) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

var File_barcode_proto protoreflect.FileDescriptor

var file_barcode_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x62, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x16, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x33, 0x0a, 0x12, 0x42, 0x61, 0x72, 0x63, 0x6f,
	0x64, 0x65, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1d, 0x0a,
	0x0a, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x6c, 0x75, 0x67, 0x22, 0x2c, 0x0a, 0x12,
	0x42, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x32, 0x76, 0x0a, 0x0e, 0x42, 0x61,
	0x72, 0x63, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x64, 0x0a, 0x08,
	0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x12, 0x2a, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x42, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x1a, 0x2a, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x61,
	0x72, 0x63, 0x6f, 0x64, 0x65, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x22, 0x00, 0x42, 0x21, 0x5a, 0x1f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_barcode_proto_rawDescOnce sync.Once
	file_barcode_proto_rawDescData = file_barcode_proto_rawDesc
)

func file_barcode_proto_rawDescGZIP() []byte {
	file_barcode_proto_rawDescOnce.Do(func() {
		file_barcode_proto_rawDescData = protoimpl.X.CompressGZIP(file_barcode_proto_rawDescData)
	})
	return file_barcode_proto_rawDescData
}

var file_barcode_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_barcode_proto_goTypes = []interface{}{
	(*BarcodeGenerateReq)(nil), // 0: object_builder_service.BarcodeGenerateReq
	(*BarcodeGenerateRes)(nil), // 1: object_builder_service.BarcodeGenerateRes
}
var file_barcode_proto_depIdxs = []int32{
	0, // 0: object_builder_service.BarcodeService.Generate:input_type -> object_builder_service.BarcodeGenerateReq
	1, // 1: object_builder_service.BarcodeService.Generate:output_type -> object_builder_service.BarcodeGenerateRes
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_barcode_proto_init() }
func file_barcode_proto_init() {
	if File_barcode_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_barcode_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BarcodeGenerateReq); i {
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
		file_barcode_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BarcodeGenerateRes); i {
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
			RawDescriptor: file_barcode_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_barcode_proto_goTypes,
		DependencyIndexes: file_barcode_proto_depIdxs,
		MessageInfos:      file_barcode_proto_msgTypes,
	}.Build()
	File_barcode_proto = out.File
	file_barcode_proto_rawDesc = nil
	file_barcode_proto_goTypes = nil
	file_barcode_proto_depIdxs = nil
}
