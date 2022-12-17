// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: subcategory.proto

package corporate_service

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateSubcategoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subcategories []*Subcategory `protobuf:"bytes,1,rep,name=subcategories,proto3" json:"subcategories,omitempty"`
	CategoryId    string         `protobuf:"bytes,2,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
}

func (x *CreateSubcategoryRequest) Reset() {
	*x = CreateSubcategoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subcategory_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSubcategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSubcategoryRequest) ProtoMessage() {}

func (x *CreateSubcategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_subcategory_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSubcategoryRequest.ProtoReflect.Descriptor instead.
func (*CreateSubcategoryRequest) Descriptor() ([]byte, []int) {
	return file_subcategory_proto_rawDescGZIP(), []int{0}
}

func (x *CreateSubcategoryRequest) GetSubcategories() []*Subcategory {
	if x != nil {
		return x.Subcategories
	}
	return nil
}

func (x *CreateSubcategoryRequest) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

type GetAllSubcategoriesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count         int32          `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Subcategories []*Subcategory `protobuf:"bytes,2,rep,name=subcategories,proto3" json:"subcategories,omitempty"`
}

func (x *GetAllSubcategoriesResponse) Reset() {
	*x = GetAllSubcategoriesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subcategory_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllSubcategoriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllSubcategoriesResponse) ProtoMessage() {}

func (x *GetAllSubcategoriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_subcategory_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllSubcategoriesResponse.ProtoReflect.Descriptor instead.
func (*GetAllSubcategoriesResponse) Descriptor() ([]byte, []int) {
	return file_subcategory_proto_rawDescGZIP(), []int{1}
}

func (x *GetAllSubcategoriesResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *GetAllSubcategoriesResponse) GetSubcategories() []*Subcategory {
	if x != nil {
		return x.Subcategories
	}
	return nil
}

type SubcategoryPrimaryKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SubcategoryPrimaryKey) Reset() {
	*x = SubcategoryPrimaryKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subcategory_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubcategoryPrimaryKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubcategoryPrimaryKey) ProtoMessage() {}

func (x *SubcategoryPrimaryKey) ProtoReflect() protoreflect.Message {
	mi := &file_subcategory_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubcategoryPrimaryKey.ProtoReflect.Descriptor instead.
func (*SubcategoryPrimaryKey) Descriptor() ([]byte, []int) {
	return file_subcategory_proto_rawDescGZIP(), []int{2}
}

func (x *SubcategoryPrimaryKey) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Subcategory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      *Language `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CreatedAt string    `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Subcategory) Reset() {
	*x = Subcategory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subcategory_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Subcategory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subcategory) ProtoMessage() {}

func (x *Subcategory) ProtoReflect() protoreflect.Message {
	mi := &file_subcategory_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subcategory.ProtoReflect.Descriptor instead.
func (*Subcategory) Descriptor() ([]byte, []int) {
	return file_subcategory_proto_rawDescGZIP(), []int{3}
}

func (x *Subcategory) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Subcategory) GetName() *Language {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *Subcategory) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type GetAllSubcategoriesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CategoryId string `protobuf:"bytes,1,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	Offset     int32  `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit      int32  `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Search     string `protobuf:"bytes,4,opt,name=search,proto3" json:"search,omitempty"`
}

func (x *GetAllSubcategoriesRequest) Reset() {
	*x = GetAllSubcategoriesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subcategory_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllSubcategoriesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllSubcategoriesRequest) ProtoMessage() {}

func (x *GetAllSubcategoriesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_subcategory_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllSubcategoriesRequest.ProtoReflect.Descriptor instead.
func (*GetAllSubcategoriesRequest) Descriptor() ([]byte, []int) {
	return file_subcategory_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllSubcategoriesRequest) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

func (x *GetAllSubcategoriesRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetAllSubcategoriesRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetAllSubcategoriesRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

var File_subcategory_proto protoreflect.FileDescriptor

var file_subcategory_proto_rawDesc = []byte{
	0x0a, 0x11, 0x73, 0x75, 0x62, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x11, 0x63, 0x6f, 0x72, 0x70, 0x6f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x68, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x81, 0x01, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x44,
	0x0a, 0x0d, 0x73, 0x75, 0x62, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x6f, 0x72, 0x70, 0x6f, 0x72, 0x61, 0x74,
	0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x75, 0x62, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x0d, 0x73, 0x75, 0x62, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x69, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x49, 0x64, 0x22, 0x79, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53,
	0x75, 0x62, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x44, 0x0a, 0x0d, 0x73, 0x75,
	0x62, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x6f, 0x72, 0x70, 0x6f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x75, 0x62, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x52, 0x0d, 0x73, 0x75, 0x62, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73,
	0x22, 0x27, 0x0a, 0x15, 0x53, 0x75, 0x62, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x50,
	0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x6d, 0x0a, 0x0b, 0x53, 0x75, 0x62,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2f, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x6f, 0x72, 0x70, 0x6f, 0x72, 0x61,
	0x74, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x83, 0x01, 0x0a, 0x1a, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x53, 0x75, 0x62, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x32, 0xc7,
	0x02, 0x0a, 0x12, 0x53, 0x75, 0x62, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6c, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d,
	0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x12, 0x2b, 0x2e, 0x63, 0x6f, 0x72, 0x70, 0x6f, 0x72,
	0x61, 0x74, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x75, 0x62, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x63, 0x6f, 0x72, 0x70, 0x6f, 0x72, 0x61, 0x74, 0x65,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x75, 0x62, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x00, 0x12, 0x69, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x2d, 0x2e,
	0x63, 0x6f, 0x72, 0x70, 0x6f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x75, 0x62, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x63,
	0x6f, 0x72, 0x70, 0x6f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x75, 0x62, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x58,
	0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x79, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x49, 0x44, 0x12, 0x28, 0x2e, 0x63, 0x6f, 0x72, 0x70, 0x6f, 0x72, 0x61, 0x74, 0x65,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x75, 0x62, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x1c, 0x5a, 0x1a, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x70, 0x6f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_subcategory_proto_rawDescOnce sync.Once
	file_subcategory_proto_rawDescData = file_subcategory_proto_rawDesc
)

func file_subcategory_proto_rawDescGZIP() []byte {
	file_subcategory_proto_rawDescOnce.Do(func() {
		file_subcategory_proto_rawDescData = protoimpl.X.CompressGZIP(file_subcategory_proto_rawDescData)
	})
	return file_subcategory_proto_rawDescData
}

var file_subcategory_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_subcategory_proto_goTypes = []interface{}{
	(*CreateSubcategoryRequest)(nil),    // 0: corporate_service.CreateSubcategoryRequest
	(*GetAllSubcategoriesResponse)(nil), // 1: corporate_service.GetAllSubcategoriesResponse
	(*SubcategoryPrimaryKey)(nil),       // 2: corporate_service.SubcategoryPrimaryKey
	(*Subcategory)(nil),                 // 3: corporate_service.Subcategory
	(*GetAllSubcategoriesRequest)(nil),  // 4: corporate_service.GetAllSubcategoriesRequest
	(*Language)(nil),                    // 5: corporate_service.Language
	(*emptypb.Empty)(nil),               // 6: google.protobuf.Empty
}
var file_subcategory_proto_depIdxs = []int32{
	3, // 0: corporate_service.CreateSubcategoryRequest.subcategories:type_name -> corporate_service.Subcategory
	3, // 1: corporate_service.GetAllSubcategoriesResponse.subcategories:type_name -> corporate_service.Subcategory
	5, // 2: corporate_service.Subcategory.name:type_name -> corporate_service.Language
	0, // 3: corporate_service.SubcategoryService.CreateMultiple:input_type -> corporate_service.CreateSubcategoryRequest
	4, // 4: corporate_service.SubcategoryService.GetAll:input_type -> corporate_service.GetAllSubcategoriesRequest
	2, // 5: corporate_service.SubcategoryService.DeleteByCategoryID:input_type -> corporate_service.SubcategoryPrimaryKey
	0, // 6: corporate_service.SubcategoryService.CreateMultiple:output_type -> corporate_service.CreateSubcategoryRequest
	1, // 7: corporate_service.SubcategoryService.GetAll:output_type -> corporate_service.GetAllSubcategoriesResponse
	6, // 8: corporate_service.SubcategoryService.DeleteByCategoryID:output_type -> google.protobuf.Empty
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_subcategory_proto_init() }
func file_subcategory_proto_init() {
	if File_subcategory_proto != nil {
		return
	}
	file_helper_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_subcategory_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSubcategoryRequest); i {
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
		file_subcategory_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllSubcategoriesResponse); i {
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
		file_subcategory_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubcategoryPrimaryKey); i {
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
		file_subcategory_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Subcategory); i {
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
		file_subcategory_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllSubcategoriesRequest); i {
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
			RawDescriptor: file_subcategory_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_subcategory_proto_goTypes,
		DependencyIndexes: file_subcategory_proto_depIdxs,
		MessageInfos:      file_subcategory_proto_msgTypes,
	}.Build()
	File_subcategory_proto = out.File
	file_subcategory_proto_rawDesc = nil
	file_subcategory_proto_goTypes = nil
	file_subcategory_proto_depIdxs = nil
}