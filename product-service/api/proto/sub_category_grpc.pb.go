// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: sub_category.proto

package product_service

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

// PosProductSubCategoryServiceClient is the client API for PosProductSubCategoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PosProductSubCategoryServiceClient interface {
	CreatePosProductSubCategory(ctx context.Context, in *CreatePosProductSubCategoryRequest, opts ...grpc.CallOption) (*CreatePosProductSubCategoryResponse, error)
	ReadPosProductSubCategory(ctx context.Context, in *ReadPosProductSubCategoryRequest, opts ...grpc.CallOption) (*ReadPosProductSubCategoryResponse, error)
	UpdatePosProductSubCategory(ctx context.Context, in *UpdatePosProductSubCategoryRequest, opts ...grpc.CallOption) (*UpdatePosProductSubCategoryResponse, error)
	DeletePosProductSubCategory(ctx context.Context, in *DeletePosProductSubCategoryRequest, opts ...grpc.CallOption) (*DeletePosProductSubCategoryResponse, error)
	ReadAllPosProductSubCategories(ctx context.Context, in *ReadAllPosProductSubCategoriesRequest, opts ...grpc.CallOption) (*ReadAllPosProductSubCategoriesResponse, error)
}

type posProductSubCategoryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPosProductSubCategoryServiceClient(cc grpc.ClientConnInterface) PosProductSubCategoryServiceClient {
	return &posProductSubCategoryServiceClient{cc}
}

func (c *posProductSubCategoryServiceClient) CreatePosProductSubCategory(ctx context.Context, in *CreatePosProductSubCategoryRequest, opts ...grpc.CallOption) (*CreatePosProductSubCategoryResponse, error) {
	out := new(CreatePosProductSubCategoryResponse)
	err := c.cc.Invoke(ctx, "/pos.PosProductSubCategoryService/CreatePosProductSubCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *posProductSubCategoryServiceClient) ReadPosProductSubCategory(ctx context.Context, in *ReadPosProductSubCategoryRequest, opts ...grpc.CallOption) (*ReadPosProductSubCategoryResponse, error) {
	out := new(ReadPosProductSubCategoryResponse)
	err := c.cc.Invoke(ctx, "/pos.PosProductSubCategoryService/ReadPosProductSubCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *posProductSubCategoryServiceClient) UpdatePosProductSubCategory(ctx context.Context, in *UpdatePosProductSubCategoryRequest, opts ...grpc.CallOption) (*UpdatePosProductSubCategoryResponse, error) {
	out := new(UpdatePosProductSubCategoryResponse)
	err := c.cc.Invoke(ctx, "/pos.PosProductSubCategoryService/UpdatePosProductSubCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *posProductSubCategoryServiceClient) DeletePosProductSubCategory(ctx context.Context, in *DeletePosProductSubCategoryRequest, opts ...grpc.CallOption) (*DeletePosProductSubCategoryResponse, error) {
	out := new(DeletePosProductSubCategoryResponse)
	err := c.cc.Invoke(ctx, "/pos.PosProductSubCategoryService/DeletePosProductSubCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *posProductSubCategoryServiceClient) ReadAllPosProductSubCategories(ctx context.Context, in *ReadAllPosProductSubCategoriesRequest, opts ...grpc.CallOption) (*ReadAllPosProductSubCategoriesResponse, error) {
	out := new(ReadAllPosProductSubCategoriesResponse)
	err := c.cc.Invoke(ctx, "/pos.PosProductSubCategoryService/ReadAllPosProductSubCategories", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PosProductSubCategoryServiceServer is the server API for PosProductSubCategoryService service.
// All implementations must embed UnimplementedPosProductSubCategoryServiceServer
// for forward compatibility
type PosProductSubCategoryServiceServer interface {
	CreatePosProductSubCategory(context.Context, *CreatePosProductSubCategoryRequest) (*CreatePosProductSubCategoryResponse, error)
	ReadPosProductSubCategory(context.Context, *ReadPosProductSubCategoryRequest) (*ReadPosProductSubCategoryResponse, error)
	UpdatePosProductSubCategory(context.Context, *UpdatePosProductSubCategoryRequest) (*UpdatePosProductSubCategoryResponse, error)
	DeletePosProductSubCategory(context.Context, *DeletePosProductSubCategoryRequest) (*DeletePosProductSubCategoryResponse, error)
	ReadAllPosProductSubCategories(context.Context, *ReadAllPosProductSubCategoriesRequest) (*ReadAllPosProductSubCategoriesResponse, error)
	mustEmbedUnimplementedPosProductSubCategoryServiceServer()
}

// UnimplementedPosProductSubCategoryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPosProductSubCategoryServiceServer struct {
}

func (UnimplementedPosProductSubCategoryServiceServer) CreatePosProductSubCategory(context.Context, *CreatePosProductSubCategoryRequest) (*CreatePosProductSubCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePosProductSubCategory not implemented")
}
func (UnimplementedPosProductSubCategoryServiceServer) ReadPosProductSubCategory(context.Context, *ReadPosProductSubCategoryRequest) (*ReadPosProductSubCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadPosProductSubCategory not implemented")
}
func (UnimplementedPosProductSubCategoryServiceServer) UpdatePosProductSubCategory(context.Context, *UpdatePosProductSubCategoryRequest) (*UpdatePosProductSubCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePosProductSubCategory not implemented")
}
func (UnimplementedPosProductSubCategoryServiceServer) DeletePosProductSubCategory(context.Context, *DeletePosProductSubCategoryRequest) (*DeletePosProductSubCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePosProductSubCategory not implemented")
}
func (UnimplementedPosProductSubCategoryServiceServer) ReadAllPosProductSubCategories(context.Context, *ReadAllPosProductSubCategoriesRequest) (*ReadAllPosProductSubCategoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadAllPosProductSubCategories not implemented")
}
func (UnimplementedPosProductSubCategoryServiceServer) mustEmbedUnimplementedPosProductSubCategoryServiceServer() {
}

// UnsafePosProductSubCategoryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PosProductSubCategoryServiceServer will
// result in compilation errors.
type UnsafePosProductSubCategoryServiceServer interface {
	mustEmbedUnimplementedPosProductSubCategoryServiceServer()
}

func RegisterPosProductSubCategoryServiceServer(s grpc.ServiceRegistrar, srv PosProductSubCategoryServiceServer) {
	s.RegisterService(&PosProductSubCategoryService_ServiceDesc, srv)
}

func _PosProductSubCategoryService_CreatePosProductSubCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePosProductSubCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PosProductSubCategoryServiceServer).CreatePosProductSubCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pos.PosProductSubCategoryService/CreatePosProductSubCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PosProductSubCategoryServiceServer).CreatePosProductSubCategory(ctx, req.(*CreatePosProductSubCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PosProductSubCategoryService_ReadPosProductSubCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadPosProductSubCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PosProductSubCategoryServiceServer).ReadPosProductSubCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pos.PosProductSubCategoryService/ReadPosProductSubCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PosProductSubCategoryServiceServer).ReadPosProductSubCategory(ctx, req.(*ReadPosProductSubCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PosProductSubCategoryService_UpdatePosProductSubCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePosProductSubCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PosProductSubCategoryServiceServer).UpdatePosProductSubCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pos.PosProductSubCategoryService/UpdatePosProductSubCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PosProductSubCategoryServiceServer).UpdatePosProductSubCategory(ctx, req.(*UpdatePosProductSubCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PosProductSubCategoryService_DeletePosProductSubCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePosProductSubCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PosProductSubCategoryServiceServer).DeletePosProductSubCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pos.PosProductSubCategoryService/DeletePosProductSubCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PosProductSubCategoryServiceServer).DeletePosProductSubCategory(ctx, req.(*DeletePosProductSubCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PosProductSubCategoryService_ReadAllPosProductSubCategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadAllPosProductSubCategoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PosProductSubCategoryServiceServer).ReadAllPosProductSubCategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pos.PosProductSubCategoryService/ReadAllPosProductSubCategories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PosProductSubCategoryServiceServer).ReadAllPosProductSubCategories(ctx, req.(*ReadAllPosProductSubCategoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PosProductSubCategoryService_ServiceDesc is the grpc.ServiceDesc for PosProductSubCategoryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PosProductSubCategoryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pos.PosProductSubCategoryService",
	HandlerType: (*PosProductSubCategoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePosProductSubCategory",
			Handler:    _PosProductSubCategoryService_CreatePosProductSubCategory_Handler,
		},
		{
			MethodName: "ReadPosProductSubCategory",
			Handler:    _PosProductSubCategoryService_ReadPosProductSubCategory_Handler,
		},
		{
			MethodName: "UpdatePosProductSubCategory",
			Handler:    _PosProductSubCategoryService_UpdatePosProductSubCategory_Handler,
		},
		{
			MethodName: "DeletePosProductSubCategory",
			Handler:    _PosProductSubCategoryService_DeletePosProductSubCategory_Handler,
		},
		{
			MethodName: "ReadAllPosProductSubCategories",
			Handler:    _PosProductSubCategoryService_ReadAllPosProductSubCategories_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sub_category.proto",
}