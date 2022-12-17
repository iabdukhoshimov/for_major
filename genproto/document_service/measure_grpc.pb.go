// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: measure.proto

package document_service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MeasureServiceClient is the client API for MeasureService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MeasureServiceClient interface {
	Create(ctx context.Context, in *MeasureRequest, opts ...grpc.CallOption) (*MeasureId, error)
	GetAll(ctx context.Context, in *GetAllMeasureRequest, opts ...grpc.CallOption) (*GetMeasuresList, error)
	GetById(ctx context.Context, in *MeasureId, opts ...grpc.CallOption) (*Measure, error)
	Update(ctx context.Context, in *MeasureRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Delete(ctx context.Context, in *MeasureId, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateMeasures(ctx context.Context, in *MeasuresList, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type measureServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMeasureServiceClient(cc grpc.ClientConnInterface) MeasureServiceClient {
	return &measureServiceClient{cc}
}

func (c *measureServiceClient) Create(ctx context.Context, in *MeasureRequest, opts ...grpc.CallOption) (*MeasureId, error) {
	out := new(MeasureId)
	err := c.cc.Invoke(ctx, "/document_service.MeasureService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *measureServiceClient) GetAll(ctx context.Context, in *GetAllMeasureRequest, opts ...grpc.CallOption) (*GetMeasuresList, error) {
	out := new(GetMeasuresList)
	err := c.cc.Invoke(ctx, "/document_service.MeasureService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *measureServiceClient) GetById(ctx context.Context, in *MeasureId, opts ...grpc.CallOption) (*Measure, error) {
	out := new(Measure)
	err := c.cc.Invoke(ctx, "/document_service.MeasureService/GetById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *measureServiceClient) Update(ctx context.Context, in *MeasureRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/document_service.MeasureService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *measureServiceClient) Delete(ctx context.Context, in *MeasureId, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/document_service.MeasureService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *measureServiceClient) UpdateMeasures(ctx context.Context, in *MeasuresList, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/document_service.MeasureService/UpdateMeasures", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MeasureServiceServer is the server API for MeasureService service.
// All implementations must embed UnimplementedMeasureServiceServer
// for forward compatibility
type MeasureServiceServer interface {
	Create(context.Context, *MeasureRequest) (*MeasureId, error)
	GetAll(context.Context, *GetAllMeasureRequest) (*GetMeasuresList, error)
	GetById(context.Context, *MeasureId) (*Measure, error)
	Update(context.Context, *MeasureRequest) (*emptypb.Empty, error)
	Delete(context.Context, *MeasureId) (*emptypb.Empty, error)
	UpdateMeasures(context.Context, *MeasuresList) (*emptypb.Empty, error)
	mustEmbedUnimplementedMeasureServiceServer()
}

// UnimplementedMeasureServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMeasureServiceServer struct {
}

func (UnimplementedMeasureServiceServer) Create(context.Context, *MeasureRequest) (*MeasureId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedMeasureServiceServer) GetAll(context.Context, *GetAllMeasureRequest) (*GetMeasuresList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedMeasureServiceServer) GetById(context.Context, *MeasureId) (*Measure, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedMeasureServiceServer) Update(context.Context, *MeasureRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedMeasureServiceServer) Delete(context.Context, *MeasureId) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedMeasureServiceServer) UpdateMeasures(context.Context, *MeasuresList) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMeasures not implemented")
}
func (UnimplementedMeasureServiceServer) mustEmbedUnimplementedMeasureServiceServer() {}

// UnsafeMeasureServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MeasureServiceServer will
// result in compilation errors.
type UnsafeMeasureServiceServer interface {
	mustEmbedUnimplementedMeasureServiceServer()
}

func RegisterMeasureServiceServer(s grpc.ServiceRegistrar, srv MeasureServiceServer) {
	s.RegisterService(&MeasureService_ServiceDesc, srv)
}

func _MeasureService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MeasureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeasureServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/document_service.MeasureService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeasureServiceServer).Create(ctx, req.(*MeasureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeasureService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllMeasureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeasureServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/document_service.MeasureService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeasureServiceServer).GetAll(ctx, req.(*GetAllMeasureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeasureService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MeasureId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeasureServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/document_service.MeasureService/GetById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeasureServiceServer).GetById(ctx, req.(*MeasureId))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeasureService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MeasureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeasureServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/document_service.MeasureService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeasureServiceServer).Update(ctx, req.(*MeasureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeasureService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MeasureId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeasureServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/document_service.MeasureService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeasureServiceServer).Delete(ctx, req.(*MeasureId))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeasureService_UpdateMeasures_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MeasuresList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeasureServiceServer).UpdateMeasures(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/document_service.MeasureService/UpdateMeasures",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeasureServiceServer).UpdateMeasures(ctx, req.(*MeasuresList))
	}
	return interceptor(ctx, in, info, handler)
}

// MeasureService_ServiceDesc is the grpc.ServiceDesc for MeasureService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MeasureService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "document_service.MeasureService",
	HandlerType: (*MeasureServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _MeasureService_Create_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _MeasureService_GetAll_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _MeasureService_GetById_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _MeasureService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _MeasureService_Delete_Handler,
		},
		{
			MethodName: "UpdateMeasures",
			Handler:    _MeasureService_UpdateMeasures_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "measure.proto",
}