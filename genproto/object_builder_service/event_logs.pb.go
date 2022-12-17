// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: event_logs.proto

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

type GetEventLogsListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TableSlug string `protobuf:"bytes,1,opt,name=table_slug,json=tableSlug,proto3" json:"table_slug,omitempty"`
	Offset    int32  `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit     int32  `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *GetEventLogsListRequest) Reset() {
	*x = GetEventLogsListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_logs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEventLogsListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventLogsListRequest) ProtoMessage() {}

func (x *GetEventLogsListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_event_logs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventLogsListRequest.ProtoReflect.Descriptor instead.
func (*GetEventLogsListRequest) Descriptor() ([]byte, []int) {
	return file_event_logs_proto_rawDescGZIP(), []int{0}
}

func (x *GetEventLogsListRequest) GetTableSlug() string {
	if x != nil {
		return x.TableSlug
	}
	return ""
}

func (x *GetEventLogsListRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetEventLogsListRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetEventLogById struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetEventLogById) Reset() {
	*x = GetEventLogById{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_logs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEventLogById) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventLogById) ProtoMessage() {}

func (x *GetEventLogById) ProtoReflect() protoreflect.Message {
	mi := &file_event_logs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventLogById.ProtoReflect.Descriptor instead.
func (*GetEventLogById) Descriptor() ([]byte, []int) {
	return file_event_logs_proto_rawDescGZIP(), []int{1}
}

func (x *GetEventLogById) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type EventLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Guid          string `protobuf:"bytes,1,opt,name=guid,proto3" json:"guid,omitempty"`
	Id            string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Date          string `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
	TableSlug     string `protobuf:"bytes,4,opt,name=table_slug,json=tableSlug,proto3" json:"table_slug,omitempty"`
	EffectedTable string `protobuf:"bytes,5,opt,name=effected_table,json=effectedTable,proto3" json:"effected_table,omitempty"`
}

func (x *EventLog) Reset() {
	*x = EventLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_logs_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventLog) ProtoMessage() {}

func (x *EventLog) ProtoReflect() protoreflect.Message {
	mi := &file_event_logs_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventLog.ProtoReflect.Descriptor instead.
func (*EventLog) Descriptor() ([]byte, []int) {
	return file_event_logs_proto_rawDescGZIP(), []int{2}
}

func (x *EventLog) GetGuid() string {
	if x != nil {
		return x.Guid
	}
	return ""
}

func (x *EventLog) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *EventLog) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *EventLog) GetTableSlug() string {
	if x != nil {
		return x.TableSlug
	}
	return ""
}

func (x *EventLog) GetEffectedTable() string {
	if x != nil {
		return x.EffectedTable
	}
	return ""
}

type GetEventLogListsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventLogs []*EventLog `protobuf:"bytes,1,rep,name=event_logs,json=eventLogs,proto3" json:"event_logs,omitempty"`
	Count     int32       `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GetEventLogListsResponse) Reset() {
	*x = GetEventLogListsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_logs_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEventLogListsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventLogListsResponse) ProtoMessage() {}

func (x *GetEventLogListsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_event_logs_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventLogListsResponse.ProtoReflect.Descriptor instead.
func (*GetEventLogListsResponse) Descriptor() ([]byte, []int) {
	return file_event_logs_proto_rawDescGZIP(), []int{3}
}

func (x *GetEventLogListsResponse) GetEventLogs() []*EventLog {
	if x != nil {
		return x.EventLogs
	}
	return nil
}

func (x *GetEventLogListsResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_event_logs_proto protoreflect.FileDescriptor

var file_event_logs_proto_rawDesc = []byte{
	0x0a, 0x10, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x16, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x66, 0x0a, 0x17, 0x47, 0x65,
	0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73,
	0x6c, 0x75, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x53, 0x6c, 0x75, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x22, 0x21, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x6f,
	0x67, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x88, 0x01, 0x0a, 0x08, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c,
	0x6f, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x67, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x67, 0x75, 0x69, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x5f, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x6c, 0x75, 0x67, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x66, 0x66,
	0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64, 0x54, 0x61, 0x62, 0x6c, 0x65,
	0x22, 0x71, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x4c,
	0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0a,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65,
	0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c,
	0x6f, 0x67, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x32, 0xdc, 0x01, 0x0a, 0x10, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67,
	0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6e, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x2f, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65,
	0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x58, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x53,
	0x69, 0x6e, 0x67, 0x6c, 0x65, 0x12, 0x27, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47,
	0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x20,
	0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67,
	0x22, 0x00, 0x42, 0x21, 0x5a, 0x1f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_event_logs_proto_rawDescOnce sync.Once
	file_event_logs_proto_rawDescData = file_event_logs_proto_rawDesc
)

func file_event_logs_proto_rawDescGZIP() []byte {
	file_event_logs_proto_rawDescOnce.Do(func() {
		file_event_logs_proto_rawDescData = protoimpl.X.CompressGZIP(file_event_logs_proto_rawDescData)
	})
	return file_event_logs_proto_rawDescData
}

var file_event_logs_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_event_logs_proto_goTypes = []interface{}{
	(*GetEventLogsListRequest)(nil),  // 0: object_builder_service.GetEventLogsListRequest
	(*GetEventLogById)(nil),          // 1: object_builder_service.GetEventLogById
	(*EventLog)(nil),                 // 2: object_builder_service.EventLog
	(*GetEventLogListsResponse)(nil), // 3: object_builder_service.GetEventLogListsResponse
}
var file_event_logs_proto_depIdxs = []int32{
	2, // 0: object_builder_service.GetEventLogListsResponse.event_logs:type_name -> object_builder_service.EventLog
	0, // 1: object_builder_service.EventLogsService.GetList:input_type -> object_builder_service.GetEventLogsListRequest
	1, // 2: object_builder_service.EventLogsService.GetSingle:input_type -> object_builder_service.GetEventLogById
	3, // 3: object_builder_service.EventLogsService.GetList:output_type -> object_builder_service.GetEventLogListsResponse
	2, // 4: object_builder_service.EventLogsService.GetSingle:output_type -> object_builder_service.EventLog
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_event_logs_proto_init() }
func file_event_logs_proto_init() {
	if File_event_logs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_event_logs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEventLogsListRequest); i {
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
		file_event_logs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEventLogById); i {
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
		file_event_logs_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventLog); i {
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
		file_event_logs_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEventLogListsResponse); i {
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
			RawDescriptor: file_event_logs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_event_logs_proto_goTypes,
		DependencyIndexes: file_event_logs_proto_depIdxs,
		MessageInfos:      file_event_logs_proto_msgTypes,
	}.Build()
	File_event_logs_proto = out.File
	file_event_logs_proto_rawDesc = nil
	file_event_logs_proto_goTypes = nil
	file_event_logs_proto_depIdxs = nil
}
