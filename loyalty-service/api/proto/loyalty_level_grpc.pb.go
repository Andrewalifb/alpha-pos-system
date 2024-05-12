// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: loyalty_level.proto

package loyalty_service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PosLoyaltyLevelServiceClient is the client API for PosLoyaltyLevelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PosLoyaltyLevelServiceClient interface {
	CreatePosLoyaltyLevel(ctx context.Context, in *CreatePosLoyaltyLevelRequest, opts ...grpc.CallOption) (*CreatePosLoyaltyLevelResponse, error)
	ReadPosLoyaltyLevel(ctx context.Context, in *ReadPosLoyaltyLevelRequest, opts ...grpc.CallOption) (*ReadPosLoyaltyLevelResponse, error)
	UpdatePosLoyaltyLevel(ctx context.Context, in *UpdatePosLoyaltyLevelRequest, opts ...grpc.CallOption) (*UpdatePosLoyaltyLevelResponse, error)
	DeletePosLoyaltyLevel(ctx context.Context, in *DeletePosLoyaltyLevelRequest, opts ...grpc.CallOption) (*DeletePosLoyaltyLevelResponse, error)
	ReadAllPosLoyaltyLevels(ctx context.Context, in *ReadAllPosLoyaltyLevelsRequest, opts ...grpc.CallOption) (*ReadAllPosLoyaltyLevelsResponse, error)
}

type posLoyaltyLevelServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPosLoyaltyLevelServiceClient(cc grpc.ClientConnInterface) PosLoyaltyLevelServiceClient {
	return &posLoyaltyLevelServiceClient{cc}
}

func (c *posLoyaltyLevelServiceClient) CreatePosLoyaltyLevel(ctx context.Context, in *CreatePosLoyaltyLevelRequest, opts ...grpc.CallOption) (*CreatePosLoyaltyLevelResponse, error) {
	out := new(CreatePosLoyaltyLevelResponse)
	err := c.cc.Invoke(ctx, "/pos.PosLoyaltyLevelService/CreatePosLoyaltyLevel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *posLoyaltyLevelServiceClient) ReadPosLoyaltyLevel(ctx context.Context, in *ReadPosLoyaltyLevelRequest, opts ...grpc.CallOption) (*ReadPosLoyaltyLevelResponse, error) {
	out := new(ReadPosLoyaltyLevelResponse)
	err := c.cc.Invoke(ctx, "/pos.PosLoyaltyLevelService/ReadPosLoyaltyLevel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *posLoyaltyLevelServiceClient) UpdatePosLoyaltyLevel(ctx context.Context, in *UpdatePosLoyaltyLevelRequest, opts ...grpc.CallOption) (*UpdatePosLoyaltyLevelResponse, error) {
	out := new(UpdatePosLoyaltyLevelResponse)
	err := c.cc.Invoke(ctx, "/pos.PosLoyaltyLevelService/UpdatePosLoyaltyLevel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *posLoyaltyLevelServiceClient) DeletePosLoyaltyLevel(ctx context.Context, in *DeletePosLoyaltyLevelRequest, opts ...grpc.CallOption) (*DeletePosLoyaltyLevelResponse, error) {
	out := new(DeletePosLoyaltyLevelResponse)
	err := c.cc.Invoke(ctx, "/pos.PosLoyaltyLevelService/DeletePosLoyaltyLevel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *posLoyaltyLevelServiceClient) ReadAllPosLoyaltyLevels(ctx context.Context, in *ReadAllPosLoyaltyLevelsRequest, opts ...grpc.CallOption) (*ReadAllPosLoyaltyLevelsResponse, error) {
	out := new(ReadAllPosLoyaltyLevelsResponse)
	err := c.cc.Invoke(ctx, "/pos.PosLoyaltyLevelService/ReadAllPosLoyaltyLevels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PosLoyaltyLevelServiceServer is the server API for PosLoyaltyLevelService service.
// All implementations must embed UnimplementedPosLoyaltyLevelServiceServer
// for forward compatibility
type PosLoyaltyLevelServiceServer interface {
	CreatePosLoyaltyLevel(context.Context, *CreatePosLoyaltyLevelRequest) (*CreatePosLoyaltyLevelResponse, error)
	ReadPosLoyaltyLevel(context.Context, *ReadPosLoyaltyLevelRequest) (*ReadPosLoyaltyLevelResponse, error)
	UpdatePosLoyaltyLevel(context.Context, *UpdatePosLoyaltyLevelRequest) (*UpdatePosLoyaltyLevelResponse, error)
	DeletePosLoyaltyLevel(context.Context, *DeletePosLoyaltyLevelRequest) (*DeletePosLoyaltyLevelResponse, error)
	ReadAllPosLoyaltyLevels(context.Context, *ReadAllPosLoyaltyLevelsRequest) (*ReadAllPosLoyaltyLevelsResponse, error)
	mustEmbedUnimplementedPosLoyaltyLevelServiceServer()
}

// UnimplementedPosLoyaltyLevelServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPosLoyaltyLevelServiceServer struct {
}

func (UnimplementedPosLoyaltyLevelServiceServer) CreatePosLoyaltyLevel(context.Context, *CreatePosLoyaltyLevelRequest) (*CreatePosLoyaltyLevelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePosLoyaltyLevel not implemented")
}
func (UnimplementedPosLoyaltyLevelServiceServer) ReadPosLoyaltyLevel(context.Context, *ReadPosLoyaltyLevelRequest) (*ReadPosLoyaltyLevelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadPosLoyaltyLevel not implemented")
}
func (UnimplementedPosLoyaltyLevelServiceServer) UpdatePosLoyaltyLevel(context.Context, *UpdatePosLoyaltyLevelRequest) (*UpdatePosLoyaltyLevelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePosLoyaltyLevel not implemented")
}
func (UnimplementedPosLoyaltyLevelServiceServer) DeletePosLoyaltyLevel(context.Context, *DeletePosLoyaltyLevelRequest) (*DeletePosLoyaltyLevelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePosLoyaltyLevel not implemented")
}
func (UnimplementedPosLoyaltyLevelServiceServer) ReadAllPosLoyaltyLevels(context.Context, *ReadAllPosLoyaltyLevelsRequest) (*ReadAllPosLoyaltyLevelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadAllPosLoyaltyLevels not implemented")
}
func (UnimplementedPosLoyaltyLevelServiceServer) mustEmbedUnimplementedPosLoyaltyLevelServiceServer() {
}

// UnsafePosLoyaltyLevelServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PosLoyaltyLevelServiceServer will
// result in compilation errors.
type UnsafePosLoyaltyLevelServiceServer interface {
	mustEmbedUnimplementedPosLoyaltyLevelServiceServer()
}

func RegisterPosLoyaltyLevelServiceServer(s grpc.ServiceRegistrar, srv PosLoyaltyLevelServiceServer) {
	s.RegisterService(&PosLoyaltyLevelService_ServiceDesc, srv)
}

func _PosLoyaltyLevelService_CreatePosLoyaltyLevel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePosLoyaltyLevelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PosLoyaltyLevelServiceServer).CreatePosLoyaltyLevel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pos.PosLoyaltyLevelService/CreatePosLoyaltyLevel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PosLoyaltyLevelServiceServer).CreatePosLoyaltyLevel(ctx, req.(*CreatePosLoyaltyLevelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PosLoyaltyLevelService_ReadPosLoyaltyLevel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadPosLoyaltyLevelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PosLoyaltyLevelServiceServer).ReadPosLoyaltyLevel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pos.PosLoyaltyLevelService/ReadPosLoyaltyLevel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PosLoyaltyLevelServiceServer).ReadPosLoyaltyLevel(ctx, req.(*ReadPosLoyaltyLevelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PosLoyaltyLevelService_UpdatePosLoyaltyLevel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePosLoyaltyLevelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PosLoyaltyLevelServiceServer).UpdatePosLoyaltyLevel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pos.PosLoyaltyLevelService/UpdatePosLoyaltyLevel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PosLoyaltyLevelServiceServer).UpdatePosLoyaltyLevel(ctx, req.(*UpdatePosLoyaltyLevelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PosLoyaltyLevelService_DeletePosLoyaltyLevel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePosLoyaltyLevelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PosLoyaltyLevelServiceServer).DeletePosLoyaltyLevel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pos.PosLoyaltyLevelService/DeletePosLoyaltyLevel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PosLoyaltyLevelServiceServer).DeletePosLoyaltyLevel(ctx, req.(*DeletePosLoyaltyLevelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PosLoyaltyLevelService_ReadAllPosLoyaltyLevels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadAllPosLoyaltyLevelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PosLoyaltyLevelServiceServer).ReadAllPosLoyaltyLevels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pos.PosLoyaltyLevelService/ReadAllPosLoyaltyLevels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PosLoyaltyLevelServiceServer).ReadAllPosLoyaltyLevels(ctx, req.(*ReadAllPosLoyaltyLevelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PosLoyaltyLevelService_ServiceDesc is the grpc.ServiceDesc for PosLoyaltyLevelService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PosLoyaltyLevelService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pos.PosLoyaltyLevelService",
	HandlerType: (*PosLoyaltyLevelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePosLoyaltyLevel",
			Handler:    _PosLoyaltyLevelService_CreatePosLoyaltyLevel_Handler,
		},
		{
			MethodName: "ReadPosLoyaltyLevel",
			Handler:    _PosLoyaltyLevelService_ReadPosLoyaltyLevel_Handler,
		},
		{
			MethodName: "UpdatePosLoyaltyLevel",
			Handler:    _PosLoyaltyLevelService_UpdatePosLoyaltyLevel_Handler,
		},
		{
			MethodName: "DeletePosLoyaltyLevel",
			Handler:    _PosLoyaltyLevelService_DeletePosLoyaltyLevel_Handler,
		},
		{
			MethodName: "ReadAllPosLoyaltyLevels",
			Handler:    _PosLoyaltyLevelService_ReadAllPosLoyaltyLevels_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "loyalty_level.proto",
}
