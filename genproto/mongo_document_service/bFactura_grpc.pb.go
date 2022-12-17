// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: bFactura.proto

package mongo_document_service

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

// BuildingFacturaServiceClient is the client API for BuildingFacturaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BuildingFacturaServiceClient interface {
	Upsert(ctx context.Context, in *BuildingFactura, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Get(ctx context.Context, in *ById, opts ...grpc.CallOption) (*BuildingFactura, error)
	Delete(ctx context.Context, in *ById, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateStatus(ctx context.Context, in *UpdateStatusReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type buildingFacturaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBuildingFacturaServiceClient(cc grpc.ClientConnInterface) BuildingFacturaServiceClient {
	return &buildingFacturaServiceClient{cc}
}

func (c *buildingFacturaServiceClient) Upsert(ctx context.Context, in *BuildingFactura, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/mongo_document_service.BuildingFacturaService/Upsert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buildingFacturaServiceClient) Get(ctx context.Context, in *ById, opts ...grpc.CallOption) (*BuildingFactura, error) {
	out := new(BuildingFactura)
	err := c.cc.Invoke(ctx, "/mongo_document_service.BuildingFacturaService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buildingFacturaServiceClient) Delete(ctx context.Context, in *ById, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/mongo_document_service.BuildingFacturaService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buildingFacturaServiceClient) UpdateStatus(ctx context.Context, in *UpdateStatusReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/mongo_document_service.BuildingFacturaService/UpdateStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BuildingFacturaServiceServer is the server API for BuildingFacturaService service.
// All implementations must embed UnimplementedBuildingFacturaServiceServer
// for forward compatibility
type BuildingFacturaServiceServer interface {
	Upsert(context.Context, *BuildingFactura) (*emptypb.Empty, error)
	Get(context.Context, *ById) (*BuildingFactura, error)
	Delete(context.Context, *ById) (*emptypb.Empty, error)
	UpdateStatus(context.Context, *UpdateStatusReq) (*emptypb.Empty, error)
	mustEmbedUnimplementedBuildingFacturaServiceServer()
}

// UnimplementedBuildingFacturaServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBuildingFacturaServiceServer struct {
}

func (UnimplementedBuildingFacturaServiceServer) Upsert(context.Context, *BuildingFactura) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Upsert not implemented")
}
func (UnimplementedBuildingFacturaServiceServer) Get(context.Context, *ById) (*BuildingFactura, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedBuildingFacturaServiceServer) Delete(context.Context, *ById) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedBuildingFacturaServiceServer) UpdateStatus(context.Context, *UpdateStatusReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStatus not implemented")
}
func (UnimplementedBuildingFacturaServiceServer) mustEmbedUnimplementedBuildingFacturaServiceServer() {
}

// UnsafeBuildingFacturaServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BuildingFacturaServiceServer will
// result in compilation errors.
type UnsafeBuildingFacturaServiceServer interface {
	mustEmbedUnimplementedBuildingFacturaServiceServer()
}

func RegisterBuildingFacturaServiceServer(s grpc.ServiceRegistrar, srv BuildingFacturaServiceServer) {
	s.RegisterService(&BuildingFacturaService_ServiceDesc, srv)
}

func _BuildingFacturaService_Upsert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuildingFactura)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildingFacturaServiceServer).Upsert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mongo_document_service.BuildingFacturaService/Upsert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildingFacturaServiceServer).Upsert(ctx, req.(*BuildingFactura))
	}
	return interceptor(ctx, in, info, handler)
}

func _BuildingFacturaService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildingFacturaServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mongo_document_service.BuildingFacturaService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildingFacturaServiceServer).Get(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

func _BuildingFacturaService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildingFacturaServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mongo_document_service.BuildingFacturaService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildingFacturaServiceServer).Delete(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

func _BuildingFacturaService_UpdateStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStatusReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildingFacturaServiceServer).UpdateStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mongo_document_service.BuildingFacturaService/UpdateStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildingFacturaServiceServer).UpdateStatus(ctx, req.(*UpdateStatusReq))
	}
	return interceptor(ctx, in, info, handler)
}

// BuildingFacturaService_ServiceDesc is the grpc.ServiceDesc for BuildingFacturaService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BuildingFacturaService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mongo_document_service.BuildingFacturaService",
	HandlerType: (*BuildingFacturaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Upsert",
			Handler:    _BuildingFacturaService_Upsert_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _BuildingFacturaService_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _BuildingFacturaService_Delete_Handler,
		},
		{
			MethodName: "UpdateStatus",
			Handler:    _BuildingFacturaService_UpdateStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bFactura.proto",
}
