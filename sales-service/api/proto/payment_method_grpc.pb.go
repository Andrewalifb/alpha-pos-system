// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: payment_method.proto

package sales_service

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

// PosPaymentMethodServiceClient is the client API for PosPaymentMethodService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PosPaymentMethodServiceClient interface {
	CreatePosPaymentMethod(ctx context.Context, in *CreatePosPaymentMethodRequest, opts ...grpc.CallOption) (*CreatePosPaymentMethodResponse, error)
	ReadPosPaymentMethod(ctx context.Context, in *ReadPosPaymentMethodRequest, opts ...grpc.CallOption) (*ReadPosPaymentMethodResponse, error)
	UpdatePosPaymentMethod(ctx context.Context, in *UpdatePosPaymentMethodRequest, opts ...grpc.CallOption) (*UpdatePosPaymentMethodResponse, error)
	DeletePosPaymentMethod(ctx context.Context, in *DeletePosPaymentMethodRequest, opts ...grpc.CallOption) (*DeletePosPaymentMethodResponse, error)
	ReadAllPosPaymentMethods(ctx context.Context, in *ReadAllPosPaymentMethodsRequest, opts ...grpc.CallOption) (*ReadAllPosPaymentMethodsResponse, error)
}

type posPaymentMethodServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPosPaymentMethodServiceClient(cc grpc.ClientConnInterface) PosPaymentMethodServiceClient {
	return &posPaymentMethodServiceClient{cc}
}

func (c *posPaymentMethodServiceClient) CreatePosPaymentMethod(ctx context.Context, in *CreatePosPaymentMethodRequest, opts ...grpc.CallOption) (*CreatePosPaymentMethodResponse, error) {
	out := new(CreatePosPaymentMethodResponse)
	err := c.cc.Invoke(ctx, "/pos.PosPaymentMethodService/CreatePosPaymentMethod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *posPaymentMethodServiceClient) ReadPosPaymentMethod(ctx context.Context, in *ReadPosPaymentMethodRequest, opts ...grpc.CallOption) (*ReadPosPaymentMethodResponse, error) {
	out := new(ReadPosPaymentMethodResponse)
	err := c.cc.Invoke(ctx, "/pos.PosPaymentMethodService/ReadPosPaymentMethod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *posPaymentMethodServiceClient) UpdatePosPaymentMethod(ctx context.Context, in *UpdatePosPaymentMethodRequest, opts ...grpc.CallOption) (*UpdatePosPaymentMethodResponse, error) {
	out := new(UpdatePosPaymentMethodResponse)
	err := c.cc.Invoke(ctx, "/pos.PosPaymentMethodService/UpdatePosPaymentMethod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *posPaymentMethodServiceClient) DeletePosPaymentMethod(ctx context.Context, in *DeletePosPaymentMethodRequest, opts ...grpc.CallOption) (*DeletePosPaymentMethodResponse, error) {
	out := new(DeletePosPaymentMethodResponse)
	err := c.cc.Invoke(ctx, "/pos.PosPaymentMethodService/DeletePosPaymentMethod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *posPaymentMethodServiceClient) ReadAllPosPaymentMethods(ctx context.Context, in *ReadAllPosPaymentMethodsRequest, opts ...grpc.CallOption) (*ReadAllPosPaymentMethodsResponse, error) {
	out := new(ReadAllPosPaymentMethodsResponse)
	err := c.cc.Invoke(ctx, "/pos.PosPaymentMethodService/ReadAllPosPaymentMethods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PosPaymentMethodServiceServer is the server API for PosPaymentMethodService service.
// All implementations must embed UnimplementedPosPaymentMethodServiceServer
// for forward compatibility
type PosPaymentMethodServiceServer interface {
	CreatePosPaymentMethod(context.Context, *CreatePosPaymentMethodRequest) (*CreatePosPaymentMethodResponse, error)
	ReadPosPaymentMethod(context.Context, *ReadPosPaymentMethodRequest) (*ReadPosPaymentMethodResponse, error)
	UpdatePosPaymentMethod(context.Context, *UpdatePosPaymentMethodRequest) (*UpdatePosPaymentMethodResponse, error)
	DeletePosPaymentMethod(context.Context, *DeletePosPaymentMethodRequest) (*DeletePosPaymentMethodResponse, error)
	ReadAllPosPaymentMethods(context.Context, *ReadAllPosPaymentMethodsRequest) (*ReadAllPosPaymentMethodsResponse, error)
	mustEmbedUnimplementedPosPaymentMethodServiceServer()
}

// UnimplementedPosPaymentMethodServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPosPaymentMethodServiceServer struct {
}

func (UnimplementedPosPaymentMethodServiceServer) CreatePosPaymentMethod(context.Context, *CreatePosPaymentMethodRequest) (*CreatePosPaymentMethodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePosPaymentMethod not implemented")
}
func (UnimplementedPosPaymentMethodServiceServer) ReadPosPaymentMethod(context.Context, *ReadPosPaymentMethodRequest) (*ReadPosPaymentMethodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadPosPaymentMethod not implemented")
}
func (UnimplementedPosPaymentMethodServiceServer) UpdatePosPaymentMethod(context.Context, *UpdatePosPaymentMethodRequest) (*UpdatePosPaymentMethodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePosPaymentMethod not implemented")
}
func (UnimplementedPosPaymentMethodServiceServer) DeletePosPaymentMethod(context.Context, *DeletePosPaymentMethodRequest) (*DeletePosPaymentMethodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePosPaymentMethod not implemented")
}
func (UnimplementedPosPaymentMethodServiceServer) ReadAllPosPaymentMethods(context.Context, *ReadAllPosPaymentMethodsRequest) (*ReadAllPosPaymentMethodsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadAllPosPaymentMethods not implemented")
}
func (UnimplementedPosPaymentMethodServiceServer) mustEmbedUnimplementedPosPaymentMethodServiceServer() {
}

// UnsafePosPaymentMethodServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PosPaymentMethodServiceServer will
// result in compilation errors.
type UnsafePosPaymentMethodServiceServer interface {
	mustEmbedUnimplementedPosPaymentMethodServiceServer()
}

func RegisterPosPaymentMethodServiceServer(s grpc.ServiceRegistrar, srv PosPaymentMethodServiceServer) {
	s.RegisterService(&PosPaymentMethodService_ServiceDesc, srv)
}

func _PosPaymentMethodService_CreatePosPaymentMethod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePosPaymentMethodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PosPaymentMethodServiceServer).CreatePosPaymentMethod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pos.PosPaymentMethodService/CreatePosPaymentMethod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PosPaymentMethodServiceServer).CreatePosPaymentMethod(ctx, req.(*CreatePosPaymentMethodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PosPaymentMethodService_ReadPosPaymentMethod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadPosPaymentMethodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PosPaymentMethodServiceServer).ReadPosPaymentMethod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pos.PosPaymentMethodService/ReadPosPaymentMethod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PosPaymentMethodServiceServer).ReadPosPaymentMethod(ctx, req.(*ReadPosPaymentMethodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PosPaymentMethodService_UpdatePosPaymentMethod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePosPaymentMethodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PosPaymentMethodServiceServer).UpdatePosPaymentMethod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pos.PosPaymentMethodService/UpdatePosPaymentMethod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PosPaymentMethodServiceServer).UpdatePosPaymentMethod(ctx, req.(*UpdatePosPaymentMethodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PosPaymentMethodService_DeletePosPaymentMethod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePosPaymentMethodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PosPaymentMethodServiceServer).DeletePosPaymentMethod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pos.PosPaymentMethodService/DeletePosPaymentMethod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PosPaymentMethodServiceServer).DeletePosPaymentMethod(ctx, req.(*DeletePosPaymentMethodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PosPaymentMethodService_ReadAllPosPaymentMethods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadAllPosPaymentMethodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PosPaymentMethodServiceServer).ReadAllPosPaymentMethods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pos.PosPaymentMethodService/ReadAllPosPaymentMethods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PosPaymentMethodServiceServer).ReadAllPosPaymentMethods(ctx, req.(*ReadAllPosPaymentMethodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PosPaymentMethodService_ServiceDesc is the grpc.ServiceDesc for PosPaymentMethodService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PosPaymentMethodService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pos.PosPaymentMethodService",
	HandlerType: (*PosPaymentMethodServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePosPaymentMethod",
			Handler:    _PosPaymentMethodService_CreatePosPaymentMethod_Handler,
		},
		{
			MethodName: "ReadPosPaymentMethod",
			Handler:    _PosPaymentMethodService_ReadPosPaymentMethod_Handler,
		},
		{
			MethodName: "UpdatePosPaymentMethod",
			Handler:    _PosPaymentMethodService_UpdatePosPaymentMethod_Handler,
		},
		{
			MethodName: "DeletePosPaymentMethod",
			Handler:    _PosPaymentMethodService_DeletePosPaymentMethod_Handler,
		},
		{
			MethodName: "ReadAllPosPaymentMethods",
			Handler:    _PosPaymentMethodService_ReadAllPosPaymentMethods_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "payment_method.proto",
}
