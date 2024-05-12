// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: invoices.proto

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

// PosInvoiceServiceClient is the client API for PosInvoiceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PosInvoiceServiceClient interface {
	CreatePosInvoice(ctx context.Context, in *CreatePosInvoiceRequest, opts ...grpc.CallOption) (*CreatePosInvoiceResponse, error)
	ReadPosInvoice(ctx context.Context, in *ReadPosInvoiceRequest, opts ...grpc.CallOption) (*ReadPosInvoiceResponse, error)
	UpdatePosInvoice(ctx context.Context, in *UpdatePosInvoiceRequest, opts ...grpc.CallOption) (*UpdatePosInvoiceResponse, error)
	DeletePosInvoice(ctx context.Context, in *DeletePosInvoiceRequest, opts ...grpc.CallOption) (*DeletePosInvoiceResponse, error)
	ReadAllPosInvoices(ctx context.Context, in *ReadAllPosInvoicesRequest, opts ...grpc.CallOption) (*ReadAllPosInvoicesResponse, error)
}

type posInvoiceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPosInvoiceServiceClient(cc grpc.ClientConnInterface) PosInvoiceServiceClient {
	return &posInvoiceServiceClient{cc}
}

func (c *posInvoiceServiceClient) CreatePosInvoice(ctx context.Context, in *CreatePosInvoiceRequest, opts ...grpc.CallOption) (*CreatePosInvoiceResponse, error) {
	out := new(CreatePosInvoiceResponse)
	err := c.cc.Invoke(ctx, "/pos.PosInvoiceService/CreatePosInvoice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *posInvoiceServiceClient) ReadPosInvoice(ctx context.Context, in *ReadPosInvoiceRequest, opts ...grpc.CallOption) (*ReadPosInvoiceResponse, error) {
	out := new(ReadPosInvoiceResponse)
	err := c.cc.Invoke(ctx, "/pos.PosInvoiceService/ReadPosInvoice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *posInvoiceServiceClient) UpdatePosInvoice(ctx context.Context, in *UpdatePosInvoiceRequest, opts ...grpc.CallOption) (*UpdatePosInvoiceResponse, error) {
	out := new(UpdatePosInvoiceResponse)
	err := c.cc.Invoke(ctx, "/pos.PosInvoiceService/UpdatePosInvoice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *posInvoiceServiceClient) DeletePosInvoice(ctx context.Context, in *DeletePosInvoiceRequest, opts ...grpc.CallOption) (*DeletePosInvoiceResponse, error) {
	out := new(DeletePosInvoiceResponse)
	err := c.cc.Invoke(ctx, "/pos.PosInvoiceService/DeletePosInvoice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *posInvoiceServiceClient) ReadAllPosInvoices(ctx context.Context, in *ReadAllPosInvoicesRequest, opts ...grpc.CallOption) (*ReadAllPosInvoicesResponse, error) {
	out := new(ReadAllPosInvoicesResponse)
	err := c.cc.Invoke(ctx, "/pos.PosInvoiceService/ReadAllPosInvoices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PosInvoiceServiceServer is the server API for PosInvoiceService service.
// All implementations must embed UnimplementedPosInvoiceServiceServer
// for forward compatibility
type PosInvoiceServiceServer interface {
	CreatePosInvoice(context.Context, *CreatePosInvoiceRequest) (*CreatePosInvoiceResponse, error)
	ReadPosInvoice(context.Context, *ReadPosInvoiceRequest) (*ReadPosInvoiceResponse, error)
	UpdatePosInvoice(context.Context, *UpdatePosInvoiceRequest) (*UpdatePosInvoiceResponse, error)
	DeletePosInvoice(context.Context, *DeletePosInvoiceRequest) (*DeletePosInvoiceResponse, error)
	ReadAllPosInvoices(context.Context, *ReadAllPosInvoicesRequest) (*ReadAllPosInvoicesResponse, error)
	mustEmbedUnimplementedPosInvoiceServiceServer()
}

// UnimplementedPosInvoiceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPosInvoiceServiceServer struct {
}

func (UnimplementedPosInvoiceServiceServer) CreatePosInvoice(context.Context, *CreatePosInvoiceRequest) (*CreatePosInvoiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePosInvoice not implemented")
}
func (UnimplementedPosInvoiceServiceServer) ReadPosInvoice(context.Context, *ReadPosInvoiceRequest) (*ReadPosInvoiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadPosInvoice not implemented")
}
func (UnimplementedPosInvoiceServiceServer) UpdatePosInvoice(context.Context, *UpdatePosInvoiceRequest) (*UpdatePosInvoiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePosInvoice not implemented")
}
func (UnimplementedPosInvoiceServiceServer) DeletePosInvoice(context.Context, *DeletePosInvoiceRequest) (*DeletePosInvoiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePosInvoice not implemented")
}
func (UnimplementedPosInvoiceServiceServer) ReadAllPosInvoices(context.Context, *ReadAllPosInvoicesRequest) (*ReadAllPosInvoicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadAllPosInvoices not implemented")
}
func (UnimplementedPosInvoiceServiceServer) mustEmbedUnimplementedPosInvoiceServiceServer() {}

// UnsafePosInvoiceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PosInvoiceServiceServer will
// result in compilation errors.
type UnsafePosInvoiceServiceServer interface {
	mustEmbedUnimplementedPosInvoiceServiceServer()
}

func RegisterPosInvoiceServiceServer(s grpc.ServiceRegistrar, srv PosInvoiceServiceServer) {
	s.RegisterService(&PosInvoiceService_ServiceDesc, srv)
}

func _PosInvoiceService_CreatePosInvoice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePosInvoiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PosInvoiceServiceServer).CreatePosInvoice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pos.PosInvoiceService/CreatePosInvoice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PosInvoiceServiceServer).CreatePosInvoice(ctx, req.(*CreatePosInvoiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PosInvoiceService_ReadPosInvoice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadPosInvoiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PosInvoiceServiceServer).ReadPosInvoice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pos.PosInvoiceService/ReadPosInvoice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PosInvoiceServiceServer).ReadPosInvoice(ctx, req.(*ReadPosInvoiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PosInvoiceService_UpdatePosInvoice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePosInvoiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PosInvoiceServiceServer).UpdatePosInvoice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pos.PosInvoiceService/UpdatePosInvoice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PosInvoiceServiceServer).UpdatePosInvoice(ctx, req.(*UpdatePosInvoiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PosInvoiceService_DeletePosInvoice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePosInvoiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PosInvoiceServiceServer).DeletePosInvoice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pos.PosInvoiceService/DeletePosInvoice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PosInvoiceServiceServer).DeletePosInvoice(ctx, req.(*DeletePosInvoiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PosInvoiceService_ReadAllPosInvoices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadAllPosInvoicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PosInvoiceServiceServer).ReadAllPosInvoices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pos.PosInvoiceService/ReadAllPosInvoices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PosInvoiceServiceServer).ReadAllPosInvoices(ctx, req.(*ReadAllPosInvoicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PosInvoiceService_ServiceDesc is the grpc.ServiceDesc for PosInvoiceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PosInvoiceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pos.PosInvoiceService",
	HandlerType: (*PosInvoiceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePosInvoice",
			Handler:    _PosInvoiceService_CreatePosInvoice_Handler,
		},
		{
			MethodName: "ReadPosInvoice",
			Handler:    _PosInvoiceService_ReadPosInvoice_Handler,
		},
		{
			MethodName: "UpdatePosInvoice",
			Handler:    _PosInvoiceService_UpdatePosInvoice_Handler,
		},
		{
			MethodName: "DeletePosInvoice",
			Handler:    _PosInvoiceService_DeletePosInvoice_Handler,
		},
		{
			MethodName: "ReadAllPosInvoices",
			Handler:    _PosInvoiceService_ReadAllPosInvoices_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "invoices.proto",
}